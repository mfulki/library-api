package config

import (
	"library-api/author_management/config/env"
	"sync"
)

type Config interface {
	Load()
}

func Load() {
	var App = new(env.AppConfig)
	var DB = new(env.DBConfig)
	var Hash = new(env.HashConfig)

	var configs = []Config{
		App,
		DB,
		Hash,
	}
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
