// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// AdminDept is the golang structure for table admin_dept.
type AdminDept struct {
	Id        int64       `json:"id"        description:"部门ID"`
	Pid       int64       `json:"pid"       description:"父部门ID"`
	Name      string      `json:"name"      description:"部门名称"`
	Code      string      `json:"code"      description:"部门编码"`
	Type      string      `json:"type"      description:"部门类型"`
	Leader    string      `json:"leader"    description:"负责人"`
	Phone     string      `json:"phone"     description:"联系电话"`
	Email     string      `json:"email"     description:"邮箱"`
	Level     int         `json:"level"     description:"关系树等级"`
	Tree      string      `json:"tree"      description:"关系树"`
	Sort      int         `json:"sort"      description:"排序"`
	Status    int         `json:"status"    description:"部门状态"`
	CreatedAt *gtime.Time `json:"createdAt" description:"创建时间"`
	UpdatedAt *gtime.Time `json:"updatedAt" description:"更新时间"`
}
