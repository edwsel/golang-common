package provider

import "github.com/sirupsen/logrus"

type Provider struct {
	Log *logrus.Logger
}

func New() *Provider {
	return &Provider{}
}
