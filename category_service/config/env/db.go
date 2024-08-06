package env

type DBConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Name     string
}

func (cfg *DBConfig) Load() {
	cfg.Host = getEnv("DB_HOST")
	cfg.Port = getIntEnv("DB_PORT")
	cfg.User = getEnv("DB_USER")
	cfg.Password = getEnv("DB_PASSWORD")
	cfg.Name = getEnv("DB_NAME")
}
