// Code generated by mdatagen. DO NOT EDIT.

package metadata

import "go.opentelemetry.io/collector/confmap"

// MetricConfig provides common config for a particular metric.
type MetricConfig struct {
	Enabled bool `mapstructure:"enabled"`

	enabledSetByUser bool
}

func (ms *MetricConfig) Unmarshal(parser *confmap.Conf) error {
	if parser == nil {
		return nil
	}
	err := parser.Unmarshal(ms, confmap.WithErrorUnused())
	if err != nil {
		return err
	}
	ms.enabledSetByUser = parser.IsSet("enabled")
	return nil
}

// MetricsConfig provides config for hostmetricsreceiver/network metrics.
type MetricsConfig struct {
	SystemNetworkConnections    MetricConfig `mapstructure:"system.network.connections"`
	SystemNetworkConntrackCount MetricConfig `mapstructure:"system.network.conntrack.count"`
	SystemNetworkConntrackMax   MetricConfig `mapstructure:"system.network.conntrack.max"`
	SystemNetworkDropped        MetricConfig `mapstructure:"system.network.dropped"`
	SystemNetworkErrors         MetricConfig `mapstructure:"system.network.errors"`
	SystemNetworkIo             MetricConfig `mapstructure:"system.network.io"`
	SystemNetworkPackets        MetricConfig `mapstructure:"system.network.packets"`
	SystemNetworkUDPBufErrors   MetricConfig `mapstructure:"system.network.udp.buf_errors"`
	SystemNetworkUDPDatagrams   MetricConfig `mapstructure:"system.network.udp.datagrams"`
	SystemNetworkUDPErrors      MetricConfig `mapstructure:"system.network.udp.errors"`
	SystemNetworkUDPNoPorts     MetricConfig `mapstructure:"system.network.udp.no_ports"`
}

func DefaultMetricsConfig() MetricsConfig {
	return MetricsConfig{
		SystemNetworkConnections: MetricConfig{
			Enabled: true,
		},
		SystemNetworkConntrackCount: MetricConfig{
			Enabled: false,
		},
		SystemNetworkConntrackMax: MetricConfig{
			Enabled: false,
		},
		SystemNetworkDropped: MetricConfig{
			Enabled: true,
		},
		SystemNetworkErrors: MetricConfig{
			Enabled: true,
		},
		SystemNetworkIo: MetricConfig{
			Enabled: true,
		},
		SystemNetworkPackets: MetricConfig{
			Enabled: true,
		},
		SystemNetworkUDPBufErrors: MetricConfig{
			Enabled: false,
		},
		SystemNetworkUDPDatagrams: MetricConfig{
			Enabled: false,
		},
		SystemNetworkUDPErrors: MetricConfig{
			Enabled: false,
		},
		SystemNetworkUDPNoPorts: MetricConfig{
			Enabled: false,
		},
	}
}

// MetricsBuilderConfig is a configuration for hostmetricsreceiver/network metrics builder.
type MetricsBuilderConfig struct {
	Metrics MetricsConfig `mapstructure:"metrics"`
}

func DefaultMetricsBuilderConfig() MetricsBuilderConfig {
	return MetricsBuilderConfig{
		Metrics: DefaultMetricsConfig(),
	}
}
