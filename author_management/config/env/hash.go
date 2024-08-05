package env

type HashConfig struct {
	Cost int
}

func (c *HashConfig) Load() {
	c.Cost = getInitEnv("HASH_COST")
}