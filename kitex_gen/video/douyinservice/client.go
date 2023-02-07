// Code generated by Kitex v0.4.4. DO NOT EDIT.

package douyinservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	video "tiktok-gateway/kitex_gen/video"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	DouyinPublishActionMethod(ctx context.Context, req *video.DouyinPublishActionRequest, callOptions ...callopt.Option) (r *video.DouyinPublishActionResponse, err error)
	DouyinPublishListMethod(ctx context.Context, req *video.DouyinPublishListRequest, callOptions ...callopt.Option) (r *video.DouyinPublishListResponse, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kDouyinServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kDouyinServiceClient struct {
	*kClient
}

func (p *kDouyinServiceClient) DouyinPublishActionMethod(ctx context.Context, req *video.DouyinPublishActionRequest, callOptions ...callopt.Option) (r *video.DouyinPublishActionResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DouyinPublishActionMethod(ctx, req)
}

func (p *kDouyinServiceClient) DouyinPublishListMethod(ctx context.Context, req *video.DouyinPublishListRequest, callOptions ...callopt.Option) (r *video.DouyinPublishListResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DouyinPublishListMethod(ctx, req)
}