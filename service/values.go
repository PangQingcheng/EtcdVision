package service

import (
	"context"
	"etcd-vision/entity"
	"fmt"
)

func GetValue(source, key string) (value string, err error) {
	etcd, ok := etcds[source]
	if !ok {
		return "", fmt.Errorf("not Found etcd: %s", source)
	}

	if etcd.Client == nil {
		return "", fmt.Errorf("etcd %s not connected", source)
	}

	resp, err := etcd.Client.Get(context.Background(), key)
	if err != nil {
		return
	}

	for _, kv := range resp.Kvs {
		value = string(kv.Value)
		break
	}
	return
}

func SetValue(source string, kv *entity.KeyValue) (revision int64, err error) {
	etcd, ok := etcds[source]
	if !ok {
		return 0, fmt.Errorf("not Found etcd: %s", source)
	}

	if etcd.Client == nil {
		return 0, fmt.Errorf("etcd %s not connected", source)
	}

	resp, err := etcd.Client.Put(context.Background(), kv.Key, kv.Value)
	if err != nil {
		return
	}

	return resp.Header.Revision, nil
}
