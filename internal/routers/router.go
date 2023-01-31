// Code generated by hertz generator.

package routers

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	handler "tiktok-gateway/internal/handler"
)

// customizeRegister registers customize routers.
func customizedRegister(r *server.Hertz) {
	r.GET("/ping", handler.Ping)
	group := r.Group("/douyin")
	group.GET("/feed", handler.DouyinFeedMethod)
	group.POST("/publish/action", handler.DouyinPublishActionMethod)
	group.GET("/publish/list", handler.DouyinPublishListMethod)
	group.POST("/favorite/action", handler.DouyinFavoriteActionMethod)
	group.GET("/favorite/list", handler.DouyinFavoriteListMethod)
	group.POST("/comment/action", handler.DouyinCommentActionMethod)
	group.GET("/comment/list", handler.DouyinCommentListMethod)
	
	// your code ...
}
