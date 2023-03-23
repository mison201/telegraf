# Printer Processor Plugin

The printer processor plugin simple prints every metric passing through it.

## Global configuration options <!-- @/docs/includes/plugin_config.md -->

In addition to the plugin-specific configuration settings, plugins support
additional global and plugin configuration settings. These settings are used to
modify metrics, tags, and field or create aliases and configure ordering, etc.
See the [CONFIGURATION.md][CONFIGURATION.md] for more details.

[CONFIGURATION.md]: ../../../docs/CONFIGURATION.md#plugins

## Configuration

```toml @sample.conf
# Print all metrics that pass through this filter.
[[outputs.mbr_gateway_config]]
    pg_user="postgres"
    pg_password="postgres"
    pg_database="massbit-user"
    pg_pool_max_conns=10
    pg_host="localhost"
    pg_port="5432"
    nginx_vts_template_path="./sample-config/input-nginx-vts.tmpl"
    nginx_vts_output_path="./sample-config/input-nginx-vts.conf"
```

## Tags

No tags are applied by this processor.
