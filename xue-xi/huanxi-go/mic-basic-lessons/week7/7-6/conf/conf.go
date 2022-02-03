package conf

type GinConfig struct {
	Host string `mapstructure:"host"`
	Port int32  `mapstructure:"port"`
}

type ServerConfig struct {
	ServerName string    `mapstructure:"serverName"`
	GinConfig  GinConfig `mapstructure:"ginConfig"`
}
