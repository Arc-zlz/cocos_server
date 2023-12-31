// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/os/gtime"
)

// AddonHgexampleTable is the golang structure for table addon_hgexample_table.
type AddonHgexampleTable struct {
	Id          int64       `json:"id"          description:"ID"`
	CategoryId  int64       `json:"categoryId"  description:"分类ID"`
	Flag        *gjson.Json `json:"flag"        description:"标签"`
	Title       string      `json:"title"       description:"标题"`
	Description string      `json:"description" description:"描述"`
	Content     string      `json:"content"     description:"内容"`
	Image       string      `json:"image"       description:"单图"`
	Images      *gjson.Json `json:"images"      description:"多图"`
	Attachfile  string      `json:"attachfile"  description:"附件"`
	Attachfiles *gjson.Json `json:"attachfiles" description:"多附件"`
	Map         *gjson.Json `json:"map"         description:"动态键值对"`
	Star        float64     `json:"star"        description:"推荐星"`
	Price       float64     `json:"price"       description:"价格"`
	Views       int64       `json:"views"       description:"浏览次数"`
	ActivityAt  *gtime.Time `json:"activityAt"  description:"活动时间"`
	StartAt     *gtime.Time `json:"startAt"     description:"开启时间"`
	EndAt       *gtime.Time `json:"endAt"       description:"结束时间"`
	Switch      int         `json:"switch"      description:"开关"`
	Sort        int         `json:"sort"        description:"排序"`
	Avatar      string      `json:"avatar"      description:"头像"`
	Sex         int         `json:"sex"         description:"性别"`
	Qq          string      `json:"qq"          description:"qq"`
	Email       string      `json:"email"       description:"邮箱"`
	Mobile      string      `json:"mobile"      description:"手机号码"`
	Hobby       *gjson.Json `json:"hobby"       description:"爱好"`
	Channel     int         `json:"channel"     description:"渠道"`
	CityId      int64       `json:"cityId"      description:"所在城市"`
	Pid         int64       `json:"pid"         description:"上级ID"`
	Level       int         `json:"level"       description:"树等级"`
	Tree        string      `json:"tree"        description:"关系树"`
	Remark      string      `json:"remark"      description:"备注"`
	Status      int         `json:"status"      description:"状态"`
	CreatedBy   int64       `json:"createdBy"   description:"创建者"`
	UpdatedBy   int64       `json:"updatedBy"   description:"更新者"`
	CreatedAt   *gtime.Time `json:"createdAt"   description:"创建时间"`
	UpdatedAt   *gtime.Time `json:"updatedAt"   description:"修改时间"`
	DeletedAt   *gtime.Time `json:"deletedAt"   description:"删除时间"`
}
