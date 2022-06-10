package application

import (
	"github.com/sirupsen/logrus"
	"gitlab.com/a-system/golang-common/internal/config"
	"gitlab.com/a-system/golang-common/internal/provider"
	"gitlab.com/a-system/golang-common/pkg/configure"
	"gitlab.com/a-system/golang-common/pkg/runner"
	"os"
)

var Log = logrus.New()

func init() {
	Log.SetLevel(logrus.InfoLevel)
	Log.SetOutput(os.Stdout)
}

type App struct {
	configPath string
	provider   *provider.Provider
	runner     *runner.Runner
	config     config.Config
}

func New(configPath string) *App {
	return &App{
		configPath: configPath,
		provider:   provider.New(),
		runner:     runner.NewRunner(Log),
		config:     config.NewConfig(),
	}
}

func (a *App) Bootstrap() error {
	boot := []func() error{
		a.boostrapConfig,
		a.boostrapLogger,
		a.bootstrapHttp,
		//a.boostrapMetrics,
		//a.boostrapEtcd,
		//a.boostrapNats,
		//a.boostrapGrpc,
		//a.bootstrapValidator,
	}

	return bootExecute(boot)
}

func (a *App) Start() error {
	return a.runner.Run()
}

func (a *App) GracefulShutdown() error {
	return a.runner.GraceFulShutdown()
}

func (a *App) boostrapConfig() error {
	return configure.LoadConfig(a.configPath, &a.config)
}

func (a *App) boostrapLogger() error {
	level, err := a.config.Logger.LevelType()

	if err != nil {
		return err
	}

	Log.SetLevel(level)

	format, err := a.config.Logger.FormatType()

	if err != nil {
		return err
	}

	Log.SetFormatter(format)

	a.provider.Log = Log

	return nil
}

func bootExecute(boot []func() error) error {
	for _, b := range boot {
		if err := b(); err != nil {
			return err
		}
	}

	return nil
}
