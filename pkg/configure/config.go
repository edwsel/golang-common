package configure

import (
	conf "go.uber.org/config"
	"os"
)

func LoadConfig(configPath string, config any) error {
	configProvider, err := conf.NewYAML(
		conf.File(configPath),
		conf.Expand(os.LookupEnv),
	)

	if err != nil {
		return err
	}

	err = configProvider.Get("").Populate(config)

	if err != nil {
		return err
	}

	return nil

}
