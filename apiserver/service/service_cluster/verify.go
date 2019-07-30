/**
 * 功能描述: 验证集群token是否有效
 * @Date: 2019-07-25
 * @author: lixiaoming
 */
package service_cluster

import (
	"apiserver/service/service_k8s/client"
	"github.com/lexkong/log"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/labels"
)

// 验证集群token是否有效
func VerifyClusterToken(clusterName, APIProtocol, APIHost, APIToken string) bool {
	clientset, err := client.NewK8sClientSet(clusterName, APIProtocol, APIHost, APIToken)
	if err != nil {
		return false
	}
	//nodes, err := client.Nodes().List(meta_v1.ListOptions{
	//	LabelSelector: labels.Everything().String(),
	//	FieldSelector: fields.Everything().String(),
	//})
	nodes, err := clientset.CoreV1().Nodes().List(meta_v1.ListOptions{
		LabelSelector: labels.Everything().String(),
		FieldSelector: fields.Everything().String(),
	})
	if err != nil {
		return false
	}
	log.Infof("Cluster node number is :%d", len(nodes.Items))
	return true
}
