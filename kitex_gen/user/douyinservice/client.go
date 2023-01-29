// Code generated by Kitex v0.4.4. DO NOT EDIT.

package douyinservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	user "tiktok-gateway/kitex_gen/user"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	DouyinUserRegisterMethod(ctx context.Context, req *user.DouyinUserRegisterRequest, callOptions ...callopt.Option) (r *user.DouyinUserRegisterResponse, err error)
	DouyinUserLoginMethod(ctx context.Context, req *user.DouyinUserLoginRequest, callOptions ...callopt.Option) (r *user.DouyinUserLoginResponse, err error)
	DouyinUserMethod(ctx context.Context, req *user.DouyinUserRequest, callOptions ...callopt.Option) (r *user.DouyinUserResponse, err error)
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

func (p *kDouyinServiceClient) DouyinUserRegisterMethod(ctx context.Context, req *user.DouyinUserRegisterRequest, callOptions ...callopt.Option) (r *user.DouyinUserRegisterResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DouyinUserRegisterMethod(ctx, req)
}

func (p *kDouyinServiceClient) DouyinUserLoginMethod(ctx context.Context, req *user.DouyinUserLoginRequest, callOptions ...callopt.Option) (r *user.DouyinUserLoginResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DouyinUserLoginMethod(ctx, req)
}

func (p *kDouyinServiceClient) DouyinUserMethod(ctx context.Context, req *user.DouyinUserRequest, callOptions ...callopt.Option) (r *user.DouyinUserResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DouyinUserMethod(ctx, req)
}