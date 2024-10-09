package facade

import (
	"github.com/sirupsen/logrus"
	"project-test/configs"
	"project-test/pkg/loggers"
)

type AutoReconcileFacade struct {
	Logger *logrus.Logger
	Config *configs.Config
}

func NewAutoReconcileFacade(config *configs.Config) *AutoReconcileFacade {
	log := loggers.NewLogger()

	f := &AutoReconcileFacade{
		Logger: log,
		Config: config,
	}

	return f
}
