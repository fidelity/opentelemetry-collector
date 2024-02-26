// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"

	"go.opentelemetry.io/collector/component"
)

var (
	Type      = component.MustNewType("memory_ballast")
	scopeName = "go.opentelemetry.io/collector/extension/ballastextension"
)

const (
	ExtensionStability = component.StabilityLevelDeprecated
)

func Meter(settings component.TelemetrySettings) metric.Meter {
	return settings.MeterProvider.Meter(scopeName)
}

func Tracer(settings component.TelemetrySettings) trace.Tracer {
	return settings.TracerProvider.Tracer(scopeName)
}