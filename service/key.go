package service

import (
	"context"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
)

func ListKeys(source, prefix string) (keys []string, err error) {
	etcd, ok := etcds[source]
	if !ok {
		return nil, fmt.Errorf("Not Found etcd: %s", source)
	}

	if etcd.Client == nil {
		return nil, fmt.Errorf("Etcd %s not connected", source)
	}

	resp, err := etcd.Client.Get(context.Background(), prefix, clientv3.WithPrefix(), clientv3.WithKeysOnly())
	if err != nil {
		return
	}

	for _, kv := range resp.Kvs {
		keys = append(keys, string(kv.Key))
	}
	return
}
