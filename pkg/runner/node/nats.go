package node

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"strings"
)

type NatsConfig struct {
	ClientName   string
	Servers      []string
	Username     string
	Password     string
	ErrorHandler nats.ErrHandler
}

type Nats struct {
	*nats.EncodedConn
	Config NatsConfig
}

func NewNats(config NatsConfig) *Nats {
	return &Nats{
		Config: config,
	}
}

func (n *Nats) Name() string {
	return "nats.encoder"
}

func (n *Nats) Run() error {
	opts := []nats.Option{
		nats.Name(n.Config.ClientName),
		nats.UserInfo(n.Config.Username, n.Config.Password),
		nats.ErrorHandler(n.Config.ErrorHandler),
	}

	nc, err := nats.Connect(strings.Join(n.Config.Servers, ","), opts...)

	if err != nil {
		return err
	}

	conn, err := nats.NewEncodedConn(nc, nats.JSON_ENCODER)

	if err != nil {
		return err
	}

	n.EncodedConn = conn

	return nil
}

func (n *Nats) Close() error {
	n.EncodedConn.Close()

	return nil
}

func (n *Nats) Publish(payload interface{}, subjects ...string) error {
	for _, v := range subjects {
		err := n.EncodedConn.Publish(v, payload)

		if err != nil {
			return fmt.Errorf("%q subject publish error: %v", v, err)
		}
	}

	return nil
}

func (n *Nats) Subscribe(subject string, handler func(msg any)) error {
	_, err := n.EncodedConn.Subscribe(subject, handler)

	if err != nil {
		return err
	}

	return nil
}
