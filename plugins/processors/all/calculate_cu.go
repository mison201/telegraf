//go:build !custom || processors || processors.calculate_cu

package all

import _ "github.com/influxdata/telegraf/plugins/processors/calculate_cu" // register plugin
