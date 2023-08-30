package test

import (
	"encoding/json"
	"etcd-vision/db"
	"etcd-vision/entity"
	"testing"
)

func TestDB(t *testing.T) {
	rq := `{
"name":"etcd-1","endpoints":["http://127.0.0.1:2389"]
}
`
	etcd := &entity.ETCDDataSource{}
	err := json.Unmarshal([]byte(rq), etcd)
	if err != nil {
		t.Fatal(err)
	}
	err = db.GetDB().Create(etcd).Error
	if err != nil {
		t.Fatal(err)
	}

	etcd2 := &entity.ETCDDataSource{}
	err = db.GetDB().Find(etcd2).Where("name = ?", "etcd-1").Error
	if err != nil {
		t.Fatal(err)
	}

	if etcd2.Name != "etcd-1" {
		t.Failed()
	}
}
