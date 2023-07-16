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

func DeleteETCD(c *gin.Context) {

}

func ConnectETCD(c *gin.Context) {

}

func DisConnectETCD(c *gin.Context) {

}

func GetETCDConnectStatus(c *gin.Context) {

}

func GetETCDKeys(c *gin.Context) {

}
