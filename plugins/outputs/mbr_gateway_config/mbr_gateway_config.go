//go:generate ../../../tools/readme_config_includer/generator
package mbr_gateway_config

// mbr_gateway_config.go

import (
	"context"
	_ "embed"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/influxdata/telegraf"
	"github.com/influxdata/telegraf/plugins/outputs"
	"github.com/jackc/pgx/v5/pgxpool"
)

//go:embed sample.conf
var sampleConfig string
var dbpool *pgxpool.Pool
var ticker *time.Ticker

type MbrGatewayConfig struct {
	PgUser               string `toml:"pg_user"`
	PgPassword           string `toml:"pg_password"`
	PgDatabase           string `toml:"pg_database"`
	PgHost               string `toml:"pg_host"`
	PgPort               string `toml:"pg_port"`
	PgPoolMaxConns       int    `toml:"pg_pool_max_conns"`
	NginxVTSTemplatePath string `toml:"nginx_vts_template_path"`
	NginxVTSOutputPath   string `toml:"nginx_vts_output_path"`

	Log telegraf.Logger `toml:"-"`
}

func (*MbrGatewayConfig) SampleConfig() string {
	return sampleConfig
}

// Init is for setup, and validating config.
func (s *MbrGatewayConfig) Init() error {
	var err error

	// ex: postgres://postgres:postgres@localhost:5432/massbit-user?pool_max_conns=10
	dbConn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?pool_max_conns=%d", s.PgUser, s.PgPassword, s.PgHost, s.PgPort, s.PgDatabase, s.PgPoolMaxConns)
	dbpool, err = pgxpool.New(context.Background(), dbConn)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}

	return err
}

func (s *MbrGatewayConfig) Connect() error {
	// Make any connection required here
	fmt.Println("Connect  ")

	ticker = time.NewTicker(15 * time.Second)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				data := Data{}
				rows, _ := dbpool.Query(context.Background(), "select ip from mbr_gateways where status = 'staked'")
				for rows.Next() {
					var ip string
					err := rows.Scan(&ip)
					if err != nil {
						fmt.Println("error : ", err)
					}
					data.GatewayIPs = append(data.GatewayIPs, ip)
				}

				fmt.Println("GatewayIPs :\n=====================================\n", strings.Join(data.GatewayIPs, "\n"), "=====================================")
				processTemplate(s.NginxVTSTemplatePath, s.NginxVTSOutputPath, data)
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()

	return nil
}

func (s *MbrGatewayConfig) Close() error {
	dbpool.Close()
	ticker.Stop()

	// Close any connections here.
	// Write will not be called once Close is called, so there is no need to synchronize.
	return nil
}

// Write should write immediately to the output, and not buffer writes
// (Telegraf manages the buffer for you). Returning an error will fail this
// batch of writes and the entire batch will be retried automatically.
func (s *MbrGatewayConfig) Write(metrics []telegraf.Metric) error {
	return nil
}

func init() {
	outputs.Add("mbr_gateway_config", func() telegraf.Output { return &MbrGatewayConfig{} })
}
