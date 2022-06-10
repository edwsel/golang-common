package node

import (
	clientv3 "go.etcd.io/etcd/client/v3"
)

type EtcdConfigs struct {
	Hosts []string
}

type Etcd struct {
	*clientv3.Client
	Config EtcdConfigs
}

func NewEtcd(config EtcdConfigs) *Etcd {
	return &Etcd{
		Config: config,
	}
}

func (e *Etcd) Name() string {
	return "etcd"
}

func (e *Etcd) Run() error {
	client, err := clientv3.New(clientv3.Config{
		Endpoints: e.Config.Hosts,
	})

	if err != nil {
		return err
	}

	e.Client = client

	return nil
}

func (e *Etcd) Close() error {
	err := e.Client.Close()

	if err != nil {
		return err
	}

	return nil
}
