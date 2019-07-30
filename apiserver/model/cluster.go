/**
 * 功能描述: 集群实体
 * @Date: 2019-07-24
 * @author: lixiaoming
 */
package model

import "time"

//create table tenx_clusters
//(
//id varchar(45) not null primary key,
//name varchar(45) not null,
//api_protocol varchar(45) null comment 'http or https',
//api_host varchar(500) null,
//api_token varchar(2048) null,
//api_tokenid varchar(45) null,
//api_version varchar(8) null,
//description text null,
//creation_time timestamp null,
//public_ips varchar(1000) null,
//binding_ips varchar(45) null,
//binding_domain varchar(1000) null,
//monitor varchar(1000) null,
//web_terminal_domain varchar(500) null,
//storage varchar(45) null,
//is_default tinyint(1) default '0' not null,
//resource_price varchar(200) null,
//config_detail varchar(200) default '{}' not null,
//extention varchar(1024) null comment '服务出口扩展',
//type tinyint(1) not null,
//edas_config blob null,
//`csb-broker_config` blob null,
//yc_project varchar(200) default '' not null,
//yc_user varchar(200) default '' not null,
//yc_id varchar(200) default '' not null,
//tiller varchar(56) default '' not null,
//constraint cluster_yc_project_name_UNIQUE
//unique (name, yc_project)
//)
//charset=utf8
//;

type ClusterModel struct {
	ID           string    `json:"id" gorm:"column:id;not null;primary_key;"`
	Name         string    `json:"name" gorm:"column:name;not null"`
	APIProtocol  string    `json:"api_protocol" gorm:"column:api_protocol"`
	APIHost      string    `json:"api_host" gorm:"column:api_host"`
	APIToken     string    `json:"api_token" gorm:"column:api_token;type:varchar(2048)"`
	APIVersion   string    `json:"api_version" gorm:"column:api_version"`
	Description  string    `json:"description" gorm:"column:description"`
	CreationTime time.Time `json:"creation_time" gorm:"column:create_time"`
	ISDefault    int8      `json:"is_default" gorm:"column:is_default"`
	ConfigDetail string    `json:"config_detail" gorm:"column:config_detail"`
	Type         int8      `json:"type" gorm:"column:type"`
}

// 表名
func (c *ClusterModel) TableName() string {
	return "jc_clusters"
}

// 插入集群
func (c *ClusterModel) Create() error {
	return DB.Instance.Save(&c).Error
}
