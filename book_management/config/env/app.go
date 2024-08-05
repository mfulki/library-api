package env

type AppConfig struct {
	Env  string
	Port int
}

func (cfg *AppConfig) Load() {
	cfg.Env = getEnv("APP_ENV")
	cfg.Port = getIntEnv("APP_PORT")
}
