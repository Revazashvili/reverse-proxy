package config

type ReverseProxyConfig struct {
	Cluster map[string]ClusterConfig `json:"clusters"`
}

type ClusterConfig struct {
	Destinations map[string]DestinationConfig `json:"destinations"`
	HealthCheck  HealthCheckConfig            `json:"healthCheck"`
}

type DestinationConfig struct {
	Address string `json:"address"`
}

type HealthCheckConfig struct {
	Enabled  bool   `json:"enabled"`
	Interval int64  `json:"interval"`
	Timeout  int64  `json:"timeout"`
	Path     string `json:"path"`
}
