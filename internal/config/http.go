package config

type Http struct {
	Name         string `yaml:"name"`
	Host         string `yaml:"host"`
	Port         uint   `yaml:"port"`
	GlobalPrefix string `yaml:"global_prefix"`
}
