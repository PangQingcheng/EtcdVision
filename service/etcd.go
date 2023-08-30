package service

import (
	"etcd-vision/db"
	"etcd-vision/entity"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"gorm.io/gorm"
	"time"
)

type ETCD struct {
	DataSource *entity.ETCDDataSource
	Client     *clientv3.Client
	Status     string
}

var etcds = map[string]*ETCD{}

func GetAllEtcdDatasource() ([]entity.ETCDDataSource, error) {
	var list []entity.ETCDDataSource
	err := db.GetDB().Find(&list).Error
	return list, err
}

func CreateEtcdDatasource(source *entity.ETCDDataSource) error {
	err := db.GetDB().Transaction(func(tx *gorm.DB) error {
		// 查重
		var sr []entity.ETCDDataSource
		err := tx.Where("name = ?", source.Name).Find(&sr).Error
		if err != nil && err != gorm.ErrRecordNotFound {
			return err
		}
		if len(sr) > 0 {
			return fmt.Errorf("source %s existed", source.Name)
		}
		// 持久化数据
		return tx.Create(source).Error
	})
	return err
}

func DeleteEtcdDatasource(name string) error {
	delete(etcds, name)
	return db.GetDB().Delete(&entity.ETCDDataSource{Name: name}).Error
}

func ConnectETCD(name string) error {
	etcd, ok := etcds[name]
	if ok && etcd.Client != nil {
		return nil
	}

	etcd = &ETCD{}

	source := &entity.ETCDDataSource{}
	err := db.GetDB().Where("name = ?", name).First(source).Error
	if err != nil {
		return err
	}
	etcd.DataSource = source

	// 创建一个 etcd 客户端连接
	endpoints, _ := source.GetEndpoints()
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   endpoints,        // etcd 服务器地址
		DialTimeout: 10 * time.Second, // 连接超时时间
	})
	if err != nil {
		return err
	}
	etcd.Client = cli
	etcd.Status = "Connected"

	// 缓存更新 TODO 加锁
	etcds[name] = etcd

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
