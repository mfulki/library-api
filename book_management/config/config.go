package config

import (
	"sync"

	"book-service/config/env"
)

type Config interface {
	Load()
}

var (
	App  = new(env.AppConfig)
	DB   = new(env.DBConfig)
	Hash = new(env.HashConfig)
)

var configs = []Config{
	App,
	DB,
	Hash,
}

func Load() {
	wg := new(sync.WaitGroup)
	wgFunc := func(wg *sync.WaitGroup, cfg Config) {
		cfg.Load()
		wg.Done()
	}

	wg.Add(len(configs))
	for _, cfg := range configs {
		go wgFunc(wg, cfg)
	}

	wg.Wait()
}
