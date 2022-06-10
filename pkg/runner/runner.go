package runner

import (
	"context"
	"errors"
	"github.com/sirupsen/logrus"
	"gitlab.com/a-system/golang-common/pkg/runner/node"
	"golang.org/x/sync/errgroup"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

// Runner the module that standardizes the procedure for starting and stopping applications
type Runner struct {
	ctx        context.Context
	cancel     context.CancelFunc
	collection []node.Node
	lock       sync.Mutex
	group      *errgroup.Group
	logger     *logrus.Logger
}

func NewRunner(logger *logrus.Logger) *Runner {
	ctx, cancel := context.WithCancel(context.Background())
	group, ctx := errgroup.WithContext(ctx)

	return &Runner{
		ctx:        ctx,
		cancel:     cancel,
		collection: []node.Node{},
		lock:       sync.Mutex{},
		group:      group,
		logger:     logger,
	}
}

func (r *Runner) Add(service node.Node) {
	r.lock.Lock()
	defer r.lock.Unlock()

	r.collection = append(r.collection, service)
}

func (r *Runner) Run() error {
	r.lock.Lock()
	defer r.lock.Unlock()

	defer func() {
		if v := recover(); v != nil {
			r.logger.Errorf("runner has error %q", v)

		}
	}()

	if len(r.collection) == 0 {
		return errors.New("node found nodes for run")
	}

	for _, service := range r.collection {
		r.group.Go(run(service, r.ctx, r.logger))
	}

	return nil
}

func (r *Runner) Stop() error {
	r.cancel()

	return r.group.Wait()
}

func (r *Runner) GraceFulShutdown() error {
	sigl := make(chan os.Signal, 1)

	signal.Notify(sigl, syscall.SIGINT, syscall.SIGTERM)

	<-sigl

	return r.Stop()
}

func run(service node.Node, ctx context.Context, logger *logrus.Logger) func() error {
	return func() error {
		logger.WithField("service", service.Name()).Info("starting service")

		go func() {
			err := service.Run()

			if err != nil {
				panic(err)
			}
		}()

		logger.WithField("service", service.Name()).Info("started service")

		for {
			select {
			case <-ctx.Done():
				logger.WithField("service", service.Name()).Info("stopping service")

				err := service.Close()

				if err != nil {
					return err
				}

				logger.WithField("service", service.Name()).Info("stopped service")

				return nil
			}
		}
	}
}
