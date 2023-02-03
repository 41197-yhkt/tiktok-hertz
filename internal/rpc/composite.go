package rpc

import (
	"context"
	"tiktok-gateway/kitex_gen/composite"
	"tiktok-gateway/kitex_gen/composite/compositeservice"
	"time"

	"pkg/constants"
	"pkg/errno"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"
)

var compClient compositeservice.Client

func initCompRPC(){
	r, err := etcd.NewEtcdResolver([]string{constants.EtcdAddress})
	if err != nil {
		panic(err)
	}

	c, err := compositeservice.NewClient(
		constants.CompServiceName,
		//client.WithMiddleware(middleware.CommonMiddleware),
		//client.WithInstanceMW(middleware.ClientMiddleware),
		client.WithMuxConnection(1),                       // mux
		client.WithRPCTimeout(3*time.Second),              // rpc timeout
		client.WithConnectTimeout(50*time.Millisecond),    // conn timeout
		client.WithFailureRetry(retry.NewFailurePolicy()), // retry
		client.WithSuite(trace.NewDefaultClientSuite()),   // tracer
		client.WithResolver(r),                            // resolver
	)
	if err != nil {
		// panic(err)
	}
	compClient = c
}

func FavoriteAction(ctx context.Context, req *composite.BasicFavoriteActionRequest) error{
	resp, err:= compClient.BasicFavoriteActionMethod(ctx, req)

	if err != nil{
		return err
	}
	if resp.BaseResp.StatusCode != 0{
		errInt := int(resp.BaseResp.StatusCode)
		return errno.NewErrNo(errInt, errno.Codes[errInt])
	}
	return nil
}

func FavoriteList(ctx context.Context, req *composite.BasicFavoriteListRequest) error{
	resp, err:= compClient.BasicFavoriteListMethod(ctx, req)

	if err != nil{
		return err
	}
	if resp.BaseResp.StatusCode != 0{
		errInt := int(resp.BaseResp.StatusCode)
		return errno.NewErrNo(errInt, errno.Codes[errInt])
	}
	return nil
}

func FeedMethod(ctx context.Context, req *composite.BasicFeedRequest) error{
	resp, err:= compClient.BasicFeedMethod(ctx, req)

	if err != nil{
		return err
	}
	if resp.BaseResp.StatusCode != 0{
		errInt := int(resp.BaseResp.StatusCode)
		return errno.NewErrNo(errInt, errno.Codes[errInt])
	}
	return nil
}

func CommentAction(ctx context.Context, req *composite.BasicCommentActionRequest) error{
	resp, err:= compClient.BasicCommentActionMethod(ctx, req)

	if err != nil{
		return err
	}
	if resp.BaseResp.StatusCode != 0{
		errInt := int(resp.BaseResp.StatusCode)
		return errno.NewErrNo(errInt, errno.Codes[errInt])
	}
	return nil
}

func CommentList(ctx context.Context, req *composite.BasicCommentListRequest) error{
	resp, err:= compClient.BasicCommentListMethod(ctx, req)

	if err != nil{
		return err
	}
	if resp.BaseResp.StatusCode != 0{
		errInt := int(resp.BaseResp.StatusCode)
		return errno.NewErrNo(errInt, errno.Codes[errInt])
	}
	return nil
}