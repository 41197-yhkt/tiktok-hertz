package handler

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"tiktok-gateway/internal/model"
)

// DouyinPublishActionMethod .
// @router /relation/publish/action [POST]
func DouyinPublishActionMethod(ctx context.Context, c *app.RequestContext) {
	var err error
	var req douyin.DouyinPublishActionRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

}

// DouyinPublishListMethod .
// @router /relation/publish/list [GET]
func DouyinPublishListMethod(ctx context.Context, c *app.RequestContext) {
	// var err error
	// var req douyin.DouyinPublishListRequest
	// err = c.BindAndValidate(&req)
	// if err != nil {
	// 	c.String(consts.StatusBadRequest, err.Error())
	// 	return
	// }

	// c.JSON(consts.StatusOK, resp)
}

