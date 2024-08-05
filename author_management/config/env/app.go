package env

type AppConfig struct {
	Env  string
	Port int
}

func (c *AppConfig) Load() {
	c.Env = getEnv("APP_ENV")
	c.Port= getInitEnv("APP_PORT")

}
