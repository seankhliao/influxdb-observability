module github.com/influxdata/influxdb-observability/otel2influx

go 1.15

require (
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/open-telemetry/opentelemetry-proto/gen/go v0.0.0
	github.com/stretchr/testify v1.5.1
)

replace github.com/open-telemetry/opentelemetry-proto/gen/go v0.0.0 => ../otlp
