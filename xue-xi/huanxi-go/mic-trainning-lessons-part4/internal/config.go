package internal

type CartOrderSrvConfig struct {
	SrvName string   `mapstructure:"srvName" json:"srvName"`
	Host    string   `mapstructure:"host" json:"host"`
	Port    int      `mapstructure:"port" json:"port"`
	Tags    []string `mapstructure:"tags" json:"tags"`
}

type CartOrderWebConfig struct {
	SrvName string   `mapstructure:"srvName" json:"srvName"`
	Host    string   `mapstructure:"host" json:"host"`
	Port    int      `mapstructure:"port" json:"port"`
	Tags    []string `mapstructure:"tags" json:"tags"`
}

type AppConfig struct {
	DBConfig           DBConfig           `mapstructure:"db" json:"db"`
	RedisConfig        RedisConfig        `mapstructure:"redis" json:"redis"`
	ConsulConfig       ConsulConfig       `mapstructure:"consul" json:"consul"`
	CartOrderSrvConfig CartOrderSrvConfig `mapstructure:"cart_order_srv" json:"cart_order_srv"`
	CartOrderWebConfig CartOrderWebConfig `mapstructure:"cart_order_web" json:"cart_order_web"`
	JWTConfig          JWTConfig          `mapstructure:"jwt" json:"jwt"`
	Debug              bool               `mapstructure:"debug" json:"debug"`
}
