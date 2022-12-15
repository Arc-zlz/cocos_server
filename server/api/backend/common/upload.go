// Package common
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2022 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
//
package common

import (
	"github.com/gogf/gf/v2/frame/g"
	"hotgo/internal/model/input/sysin"
)

// UploadImageReq 上传图片
type UploadImageReq struct {
	g.Meta `path:"/upload/image" tags:"上传" method:"post" summary:"上传图片"`
}

type UploadImageRes *sysin.AttachmentListModel