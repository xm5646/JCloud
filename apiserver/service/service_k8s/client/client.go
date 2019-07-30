/**
 * 功能描述: 原生k8s clientset创建
 * @Date: 2019-07-25
 * @author: lixiaoming
 */
package client

import (
	"fmt"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/clientcmd/api"
	"time"
)

// 定义超时时间
const timeout = 10 * time.Second

// 创建clientset config
func NewConfig(clusterName, apiProtocol, apiHost, apiToken string) (*rest.Config, error) {
	config := api.NewConfig()
	config.Clusters[clusterName] = &api.Cluster{
		Server:                fmt.Sprintf("%s://%s", apiProtocol, apiHost),
		InsecureSkipTLSVerify: true,
	}
	config.AuthInfos[clusterName] = &api.AuthInfo{Token: apiToken}
	config.Contexts[clusterName] = &api.Context{
		Cluster:  clusterName,
		AuthInfo: clusterName,
	}
	config.CurrentContext = clusterName

	clientBuilder := clientcmd.NewNonInteractiveClientConfig(*config, clusterName, &clientcmd.ConfigOverrides{}, nil)
	return clientBuilder.ClientConfig()
}

// 创建clientset
func NewK8sClientSet(clusterName, apiProtocol, apiHost, apiToken string) (*kubernetes.Clientset, error) {
	cfg, err := NewConfig(clusterName, apiProtocol, apiHost, apiToken)
	if err != nil {
		return nil, err
	}

	clientset, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		return nil, err
	}

	return clientset, nil
}
