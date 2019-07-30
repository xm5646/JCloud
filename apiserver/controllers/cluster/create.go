/**
 * 功能描述: 添加集群
 * @Date: 2019-07-24
 * @author: lixiaoming
 */
package cluster

import (
	. "apiserver/controllers"
	"apiserver/model"
	"apiserver/pkg/errno"
	"apiserver/service/service_cluster"
	"apiserver/util"
	"apiserver/util/uuid"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/lexkong/log"
	"github.com/lexkong/log/lager"
	"strings"
	"time"
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

	//TODO 判断当前是否为管理员用户

	// 解析请求参数
	var req CreateRequest
	err := c.ShouldBind(&req)
	if err != nil {
		log.Errorf(err, errno.ErrRequestBody.Message)
		SendResponse(c, errno.ErrRequestBody, nil)
		return
	}

	if req.APIHost == "" || req.APIToken == "" || req.Name == "" {
		SendResponse(c, errno.ErrRequestBody, nil)
		return
	}

	// 创建集群模型并赋值
	cluster := model.ClusterModel{}
	if strings.HasPrefix(req.APIHost, "https://") {
		cluster.APIProtocol = "https"
		cluster.APIHost = strings.Trim(req.APIHost, "https://")
	} else if strings.HasPrefix(req.APIHost, "http://") {
		cluster.APIProtocol = "http"
		cluster.APIHost = strings.Trim(req.APIHost, "http://")
	}

	if cluster.APIProtocol != "http" && cluster.APIProtocol != "https" {
		log.Errorf(nil, "The api protocol is invalid.")
		SendResponse(c, errno.ErrRequestBody, nil)
		return
	}
	cluster.Name = req.Name
	cluster.APIToken = req.APIToken
	cluster.APIVersion = "v1"
	cluster.ID = uuid.CombineClusterUUId(req.APIHost, req.APIToken)
	if checkIfAnyDefaultCluster() == false {
		cluster.ISDefault = 1
	}
	cluster.CreationTime = time.Now()

	//TODO 验证token是否有效
	if service_cluster.VerifyClusterToken(cluster.Name, cluster.APIProtocol, cluster.APIHost, cluster.APIToken) == false {
		log.Errorf(nil, "Verify cluster token failed.")
		SendResponse(c, errno.ErrAPIError, nil)
		return
	}

	// 插入数据库
	err = cluster.Create()
	if err != nil {
		log.Errorf(err, errno.ErrDatabase.Message)
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	SendResponse(c, nil, "ok")
}

func checkIfAnyDefaultCluster() bool {
	cluster := &model.ClusterModel{}
	err := model.DB.Instance.Table(cluster.TableName()).Where("is_default = 1").First(&cluster).Error
	if err == gorm.ErrRecordNotFound {
		return false
	}
	return true
}
