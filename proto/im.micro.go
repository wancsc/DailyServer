// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: rpc/proto/im.proto

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	rpc/proto/im.proto

It has these top-level messages:
	PingReq
	PingRes
	PushMsgReq
	PushMsgRes
	PushMsgsReq
	PushMsgsRes
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "context"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for IMService service

type IMService interface {
	Ping(ctx context.Context, in *PingReq, opts ...client.CallOption) (*PingRes, error)
	PushMsg(ctx context.Context, in *PushMsgReq, opts ...client.CallOption) (*PushMsgRes, error)
	PushMsgs(ctx context.Context, in *PushMsgsReq, opts ...client.CallOption) (*PushMsgsRes, error)
}

type iMService struct {
	c    client.Client
	name string
}

func NewIMService(name string, c client.Client) IMService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "proto"
	}
	return &iMService{
		c:    c,
		name: name,
	}
}

func (c *iMService) Ping(ctx context.Context, in *PingReq, opts ...client.CallOption) (*PingRes, error) {
	req := c.c.NewRequest(c.name, "IMService.Ping", in)
	out := new(PingRes)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iMService) PushMsg(ctx context.Context, in *PushMsgReq, opts ...client.CallOption) (*PushMsgRes, error) {
	req := c.c.NewRequest(c.name, "IMService.PushMsg", in)
	out := new(PushMsgRes)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iMService) PushMsgs(ctx context.Context, in *PushMsgsReq, opts ...client.CallOption) (*PushMsgsRes, error) {
	req := c.c.NewRequest(c.name, "IMService.PushMsgs", in)
	out := new(PushMsgsRes)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for IMService service

type IMServiceHandler interface {
	Ping(context.Context, *PingReq, *PingRes) error
	PushMsg(context.Context, *PushMsgReq, *PushMsgRes) error
	PushMsgs(context.Context, *PushMsgsReq, *PushMsgsRes) error
}

func RegisterIMServiceHandler(s server.Server, hdlr IMServiceHandler, opts ...server.HandlerOption) error {
	type iMService interface {
		Ping(ctx context.Context, in *PingReq, out *PingRes) error
		PushMsg(ctx context.Context, in *PushMsgReq, out *PushMsgRes) error
		PushMsgs(ctx context.Context, in *PushMsgsReq, out *PushMsgsRes) error
	}
	type IMService struct {
		iMService
	}
	h := &iMServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&IMService{h}, opts...))
}

type iMServiceHandler struct {
	IMServiceHandler
}

func (h *iMServiceHandler) Ping(ctx context.Context, in *PingReq, out *PingRes) error {
	return h.IMServiceHandler.Ping(ctx, in, out)
}

func (h *iMServiceHandler) PushMsg(ctx context.Context, in *PushMsgReq, out *PushMsgRes) error {
	return h.IMServiceHandler.PushMsg(ctx, in, out)
}

func (h *iMServiceHandler) PushMsgs(ctx context.Context, in *PushMsgsReq, out *PushMsgsRes) error {
	return h.IMServiceHandler.PushMsgs(ctx, in, out)
}
