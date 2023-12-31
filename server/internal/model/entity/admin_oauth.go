// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// AdminOauth is the golang structure for table admin_oauth.
type AdminOauth struct {
	Id           int64       `json:"id"           description:"主键"`
	MemberId     int64       `json:"memberId"     description:"用户ID"`
	Unionid      string      `json:"unionid"      description:"唯一ID"`
	OauthClient  string      `json:"oauthClient"  description:"授权组别"`
	OauthOpenid  string      `json:"oauthOpenid"  description:"授权开放ID"`
	Sex          int         `json:"sex"          description:"性别"`
	Nickname     string      `json:"nickname"     description:"昵称"`
	HeadPortrait string      `json:"headPortrait" description:"头像"`
	Birthday     *gtime.Time `json:"birthday"     description:"生日"`
	Country      string      `json:"country"      description:"国家"`
	Province     string      `json:"province"     description:"省"`
	City         string      `json:"city"         description:"市"`
	Status       int         `json:"status"       description:"状态"`
	CreatedAt    *gtime.Time `json:"createdAt"    description:"创建时间"`
	UpdatedAt    *gtime.Time `json:"updatedAt"    description:"修改时间"`
}
