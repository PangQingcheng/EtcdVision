package etcd

import (
	"etcd-vision/controller/response"
	"etcd-vision/entity"
	"etcd-vision/service"
	"github.com/gin-gonic/gin"
)

func NewETCD(c *gin.Context) {
	db := &entity.ETCDDataSource{}
	err := c.Bind(db)
	if err != nil {
		response.BadRequestErr(c, err)
		return
	}

	err = service.CreateEtcdDatasource(db)
	if err != nil {
		response.BadRequestErr(c, err)
		return
	}

	response.Success(c, "ok", nil)
}

func ListETCD(c *gin.Context) {
	list, err := service.GetAllEtcdDatasource()
	if err != nil {
		response.BadRequestErr(c, err)
		return
	}

	response.Success(c, "ok", list)
}

func DeleteETCD(c *gin.Context) {
	name := c.Param("name")

	err := service.DeleteEtcdDatasource(name)
	if err != nil {
		response.BadRequestErr(c, err)
		return
	}

	response.Success(c, "ok", nil)
}

func ConnectETCD(c *gin.Context) {
	name := c.Param("name")

	err := service.ConnectETCD(name)
	if err != nil {
		response.BadRequestErr(c, err)
		return
	}

	response.Success(c, "ok", nil)
}

func DisConnectETCD(c *gin.Context) {
	name := c.Param("name")

	err := service.DisConnectETCD(name)
	if err != nil {
		response.BadRequestErr(c, err)
		return
	}

	response.Success(c, "ok", nil)
}

func GetETCDConnectStatus(c *gin.Context) {

}

func GetETCDKeys(c *gin.Context) {
	name := c.Param("name")
	prefix := c.Query("prefix")

	keys, err := service.ListKeys(name, prefix)
	if err != nil {
		response.BadRequestErr(c, err)
		return
	}

	response.Success(c, "ok", keys)
}

func GetETCDValue(c *gin.Context) {
	name := c.Param("name")
	key := c.Query("key")

	value, err := service.GetValue(name, key)
	if err != nil {
		response.BadRequestErr(c, err)
		return
	}

	response.Success(c, "ok", value)
}

func SetETCDValue(c *gin.Context) {
	name := c.Param("name")

	kv := &entity.KeyValue{}
	err := c.BindJSON(kv)
	if err != nil {
		response.BadRequestErr(c, err)
		return
	}

	revision, err := service.SetValue(name, kv)
	if err != nil {
		response.BadRequestErr(c, err)
		return
	}

	response.Success(c, "ok", map[string]interface{}{
		"revision": revision,
		"key":      kv.Key,
		"value":    kv.Value,
	})
}
