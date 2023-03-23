//go:build !custom || outputs || outputs.mbr_gateway_config

package all

import _ "github.com/influxdata/telegraf/plugins/outputs/mbr_gateway_config" // register plugin
