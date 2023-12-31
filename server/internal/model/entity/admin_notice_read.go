// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// AdminNoticeRead is the golang structure for table admin_notice_read.
type AdminNoticeRead struct {
	Id        int64       `json:"id"        description:"记录ID"`
	NoticeId  int64       `json:"noticeId"  description:"公告ID"`
	MemberId  int64       `json:"memberId"  description:"会员ID"`
	Clicks    int         `json:"clicks"    description:"已读次数"`
	UpdatedAt *gtime.Time `json:"updatedAt" description:"更新时间"`
	CreatedAt *gtime.Time `json:"createdAt" description:"阅读时间"`
}
