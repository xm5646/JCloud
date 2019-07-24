/**
 * 功能描述: 添加集群
 * @Date: 2019-07-24
 * @author: lixiaoming
 */
package cluster

import (
	. "apiserver/controllers"
	"apiserver/util"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/lexkong/log/lager"
)

// @Summary 添加集群
// @Description 通过k8s master api token 添加集群
// @Tags cluster
// @Accept  json
// @Produce  json
// @Param Authorization header string true "用户登录状态的token"
// @Param cluster body cluster.CreateRequest true "{}"
// @Success 200 {object} controllers.Response "{}"
// @Router /v1/cluster [post]
func Create(c *gin.Context) {
	log.Info("Create function called.", lager.Data{"X-Request-Id": util.GetReqID(c)})
	SendResponse(c, nil, "ok")
}
