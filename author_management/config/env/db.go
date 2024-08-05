package env

type DBConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Name     string
}

func (c *DBConfig) Load() {
	c.Host = getEnv("DB_HOST")
	c.Port = getInitEnv("DB_PORT")
	c.User = getEnv("DB_USER")
	c.Password = getEnv("DB_PASSWORD")
	c.Name = getEnv("DB_NAME")
}
