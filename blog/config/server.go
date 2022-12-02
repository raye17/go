package config

type serverConfig struct {
	SiteName string `json:"siteName"`
	Env      string `json:"env"`
	Port     string `json:"port"`
	LogLevel string `json:"log_level"`
}
