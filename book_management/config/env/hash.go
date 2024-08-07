package env

type HashConfig struct {
	Cost int
}

func (cfg *HashConfig) Load() {
	cfg.Cost = getIntEnv("HASH_COST")
}
