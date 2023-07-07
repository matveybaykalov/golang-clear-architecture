package file

import "clear-arch/internal/config"

type ConfigLoader struct{}

func (ConfigLoader) Load() config.Config {
	return config.Config{}
}
