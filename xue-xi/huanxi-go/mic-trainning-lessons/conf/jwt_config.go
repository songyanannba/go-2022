package conf

type JWTConfig struct {
	SingingKey string `mspstructure:"signing_key"`
}
