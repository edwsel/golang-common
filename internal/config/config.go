package config

type Config struct {
	Http   Http   `yaml:"http"`
	Logger Logger `yaml:"logger"`
}

func NewConfig() Config {
	return Config{}
}
