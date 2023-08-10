package service

import (
	"encoding/json"
	"etcd-vision/entity"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"os"
	"path/filepath"
	"time"

	log "github.com/sirupsen/logrus"
)

type ETCD struct {
	DataSource *entity.ETCDDataSource
	Client     *clientv3.Client
	Status     string
}

var etcds map[string]*ETCD

func init() {
	etcds = make(map[string]*ETCD)
	filepath.Walk("config/datasources", func(path string, info os.FileInfo, err error) error {
		// 判断是否为文件
		if !info.IsDir() {
			data, err := os.ReadFile(path)
			if err != nil {
				log.Errorf("init etcd data sources error: read file %s error: %v", path, err)
				return nil
			}
			source := &entity.ETCDDataSource{}
			err = json.Unmarshal(data, source)
			if err != nil {
				log.Errorf("init etcd data sources error: parse file %s error: %v", path, err)
				return nil
			}
			if source.Name != "" {
				etcds[source.Name] = &ETCD{
					DataSource: source,
				}
			}
		}
		return nil
	})
}

func GetAllEtcdDatasource() ([]*entity.ETCDDataSource, error) {
	var list []*entity.ETCDDataSource
	for _, etcd := range etcds {
		list = append(list, etcd.DataSource)
	}
	return list, nil
}

func CreateEtcdDatasource(source *entity.ETCDDataSource) error {
	// 和缓存比对是否名称冲突
	for name, _ := range etcds {
		if name == source.Name {
			return fmt.Errorf("name %s existed", source.Name)
		}
	}
	// 持久化数据
	data, err := json.MarshalIndent(source, "", "  ")
	if err != nil {
		return err
	}
	err = os.WriteFile("config/datasources/"+source.Name, data, os.ModePerm)
	if err != nil {
		return err
	}

	// 缓存更新
	etcds[source.Name] = &ETCD{
		DataSource: source,
	}
	return nil
}

func DeleteEtcdDatasource(name string) error {
	_, ok := etcds[name]
	if !ok {
		return fmt.Errorf("not Found etcd: %s", name)
	}

	os.Remove("config/datasources/" + name)

	_ = DisConnectETCD(name)

	delete(etcds, name)

	return nil
}

func ConnectETCD(name string) error {
	etcd, ok := etcds[name]
	if !ok {
		return fmt.Errorf("not Found etcd: %s", name)
	}
	if etcd.Client != nil {
		return nil
	}

	// 创建一个 etcd 客户端连接
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   etcd.DataSource.Endpoints, // etcd 服务器地址
		DialTimeout: 5 * time.Second,           // 连接超时时间
	})
	if err != nil {
		return err
	}
	etcd.Client = cli
	etcd.Status = "Connected"

	return nil
}

func DisConnectETCD(name string) error {
	etcd, ok := etcds[name]
	if !ok {
		return fmt.Errorf("not Found etcd: %s", name)
	}

	if etcd.Client == nil {
		return nil
	}

	err := etcd.Client.Close()
	if err != nil {
		return err
	}

	etcd.Client = nil
	etcd.Status = "DisConnected"

	return nil
}
