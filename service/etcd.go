package service

import (
	"encoding/json"
	"etcd-vision/entity"
	"fmt"
	"os"
	"path/filepath"
	"time"

	log "github.com/sirupsen/logrus"
	"go.etcd.io/etcd/clientv3"
)

var etcds map[string]*entity.ETCDDataSource

func init() {
	etcds = make(map[string]*entity.ETCDDataSource)
	filepath.Walk("config/datasources", func(path string, info os.FileInfo, err error) error {
		// 判断是否为文件
		if !info.IsDir() {
			data, err := os.ReadFile(path)
			if err != nil {
				log.Errorf("init etcd data sources error: read file %s error: %v", path, err)
				return nil
			}
			db := &entity.ETCDDataSource{}
			err = json.Unmarshal(data, db)
			if err != nil {
				log.Errorf("init etcd data sources error: parse file %s error: %v", path, err)
				return nil
			}
			if db.Name != "" {
				etcds[db.Name] = db
			}
		}
		return nil
	})
}

func CreateEtcdDatasource(db *entity.ETCDDataSource) error {
	// 和缓存比对是否名称冲突
	for name, _ := range etcds {
		if name == db.Name {
			return fmt.Errorf("name %s existed", db.Name)
		}
	}
	// 持久化数据
	data, err := json.MarshalIndent(db, "", "  ")
	if err != nil {
		return err
	}
	err = os.WriteFile("config/datasources/"+db.Name, data, os.ModePerm)
	if err != nil {
		return err
	}

	// 缓存更新
	etcds[db.Name] = db
	return nil
}

func ConnectETCD(name string) error {
	db := etcds[name]

	// 创建一个 etcd 客户端连接
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"http://localhost:2379"}, // etcd 服务器地址
		DialTimeout: 5 * time.Second,                   // 连接超时时间
	})
}
