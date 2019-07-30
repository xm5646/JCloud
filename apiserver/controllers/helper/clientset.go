/**
 * 功能描述: k8s clientset工具
 * @Date: 2019-07-25
 * @author: lixiaoming
 */
package helper

import (
	"apiserver/model"
	"apiserver/pkg/errno"
	"apiserver/service/service_k8s/client"
	"github.com/jinzhu/gorm"
	"github.com/lexkong/log"
	"k8s.io/client-go/kubernetes"
)

// 根据集群ID获取clientset
func GetK8sClientSetByClusterID(clusterID string) (*kubernetes.Clientset, error) {
	// 根据集群ID获取集群
	clusterModel := model.ClusterModel{}
	err := model.DB.Instance.Where("id = ?", clusterID).First(&clusterModel).Error
	if err == gorm.ErrRecordNotFound {
		log.Errorf(err, "Can't find cluster by clusterID %s.", clusterID)
		return nil, err
	}
	if err != nil {
		log.Errorf(err, errno.ErrDatabase.Message)
		return nil, err
	}

	// 获取clientSet
	clientSet, err := client.NewK8sClientSet(clusterModel.Name, clusterModel.APIProtocol, clusterModel.APIHost, clusterModel.APIToken)
	if err != nil {
		log.Errorf(err, "Create kubernetes clientset failed, config:%#v.", clusterModel)
		return nil, err
	}
	return clientSet, nil
}
