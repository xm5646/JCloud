/**
 * 功能描述: 集群接口相关的请求相应结构体
 * @Date: 2019-07-24
 * @author: lixiaoming
 */
package cluster

type CreateRequest struct {
	Name     string `json:"name"`
	APIHost  string `json:"api_host" gorm:"column:api_host"`
	APIToken string `json:"api_token" gorm:"column:api_token"`
}
