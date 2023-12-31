// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysGenCodes is the golang structure for table sys_gen_codes.
type SysGenCodes struct {
	Id            int64       `json:"id"            description:"生成ID"`
	GenType       uint        `json:"genType"       description:"生成类型"`
	GenTemplate   int         `json:"genTemplate"   description:"生成模板"`
	VarName       string      `json:"varName"       description:"实体命名"`
	Options       *gjson.Json `json:"options"       description:"配置选项"`
	DbName        string      `json:"dbName"        description:"数据库名称"`
	TableName     string      `json:"tableName"     description:"主表名称"`
	TableComment  string      `json:"tableComment"  description:"主表注释"`
	DaoName       string      `json:"daoName"       description:"主表dao模型"`
	MasterColumns *gjson.Json `json:"masterColumns" description:"主表字段"`
	AddonName     string      `json:"addonName"     description:"插件名称"`
	Status        int         `json:"status"        description:"生成状态"`
	CreatedAt     *gtime.Time `json:"createdAt"     description:"创建时间"`
	UpdatedAt     *gtime.Time `json:"updatedAt"     description:"更新时间"`
}
