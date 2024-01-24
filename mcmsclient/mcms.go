// Code generated by goctl. DO NOT EDIT.
// Source: mcms.proto

package mcmsclient

import (
	"context"

	"github.com/iot-synergy/synergy-message-center/types/mcms"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	BaseIDResp            = mcms.BaseIDResp
	BaseResp              = mcms.BaseResp
	BaseUUIDResp          = mcms.BaseUUIDResp
	EmailInfo             = mcms.EmailInfo
	EmailLogInfo          = mcms.EmailLogInfo
	EmailLogListReq       = mcms.EmailLogListReq
	EmailLogListResp      = mcms.EmailLogListResp
	EmailProviderInfo     = mcms.EmailProviderInfo
	EmailProviderListReq  = mcms.EmailProviderListReq
	EmailProviderListResp = mcms.EmailProviderListResp
	Empty                 = mcms.Empty
	IDReq                 = mcms.IDReq
	IDsReq                = mcms.IDsReq
	PageInfoReq           = mcms.PageInfoReq
	SmsInfo               = mcms.SmsInfo
	SmsLogInfo            = mcms.SmsLogInfo
	SmsLogListReq         = mcms.SmsLogListReq
	SmsLogListResp        = mcms.SmsLogListResp
	SmsProviderInfo       = mcms.SmsProviderInfo
	SmsProviderListReq    = mcms.SmsProviderListReq
	SmsProviderListResp   = mcms.SmsProviderListResp
	UUIDReq               = mcms.UUIDReq
	UUIDsReq              = mcms.UUIDsReq

	Mcms interface {
		InitDatabase(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*BaseResp, error)
		SendEmail(ctx context.Context, in *EmailInfo, opts ...grpc.CallOption) (*BaseUUIDResp, error)
		// EmailLog management
		CreateEmailLog(ctx context.Context, in *EmailLogInfo, opts ...grpc.CallOption) (*BaseUUIDResp, error)
		UpdateEmailLog(ctx context.Context, in *EmailLogInfo, opts ...grpc.CallOption) (*BaseResp, error)
		GetEmailLogList(ctx context.Context, in *EmailLogListReq, opts ...grpc.CallOption) (*EmailLogListResp, error)
		GetEmailLogById(ctx context.Context, in *UUIDReq, opts ...grpc.CallOption) (*EmailLogInfo, error)
		DeleteEmailLog(ctx context.Context, in *UUIDsReq, opts ...grpc.CallOption) (*BaseResp, error)
		// EmailProvider management
		CreateEmailProvider(ctx context.Context, in *EmailProviderInfo, opts ...grpc.CallOption) (*BaseIDResp, error)
		UpdateEmailProvider(ctx context.Context, in *EmailProviderInfo, opts ...grpc.CallOption) (*BaseResp, error)
		GetEmailProviderList(ctx context.Context, in *EmailProviderListReq, opts ...grpc.CallOption) (*EmailProviderListResp, error)
		GetEmailProviderById(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*EmailProviderInfo, error)
		DeleteEmailProvider(ctx context.Context, in *IDsReq, opts ...grpc.CallOption) (*BaseResp, error)
		SendSms(ctx context.Context, in *SmsInfo, opts ...grpc.CallOption) (*BaseUUIDResp, error)
		// SmsLog management
		CreateSmsLog(ctx context.Context, in *SmsLogInfo, opts ...grpc.CallOption) (*BaseUUIDResp, error)
		UpdateSmsLog(ctx context.Context, in *SmsLogInfo, opts ...grpc.CallOption) (*BaseResp, error)
		GetSmsLogList(ctx context.Context, in *SmsLogListReq, opts ...grpc.CallOption) (*SmsLogListResp, error)
		GetSmsLogById(ctx context.Context, in *UUIDReq, opts ...grpc.CallOption) (*SmsLogInfo, error)
		DeleteSmsLog(ctx context.Context, in *UUIDsReq, opts ...grpc.CallOption) (*BaseResp, error)
		// SmsProvider management
		CreateSmsProvider(ctx context.Context, in *SmsProviderInfo, opts ...grpc.CallOption) (*BaseIDResp, error)
		UpdateSmsProvider(ctx context.Context, in *SmsProviderInfo, opts ...grpc.CallOption) (*BaseResp, error)
		GetSmsProviderList(ctx context.Context, in *SmsProviderListReq, opts ...grpc.CallOption) (*SmsProviderListResp, error)
		GetSmsProviderById(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*SmsProviderInfo, error)
		DeleteSmsProvider(ctx context.Context, in *IDsReq, opts ...grpc.CallOption) (*BaseResp, error)
	}

	defaultMcms struct {
		cli zrpc.Client
	}
)

func NewMcms(cli zrpc.Client) Mcms {
	return &defaultMcms{
		cli: cli,
	}
}

func (m *defaultMcms) InitDatabase(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*BaseResp, error) {
	client := mcms.NewMcmsClient(m.cli.Conn())
	return client.InitDatabase(ctx, in, opts...)
}

func (m *defaultMcms) SendEmail(ctx context.Context, in *EmailInfo, opts ...grpc.CallOption) (*BaseUUIDResp, error) {
	client := mcms.NewMcmsClient(m.cli.Conn())
	return client.SendEmail(ctx, in, opts...)
}

// EmailLog management
func (m *defaultMcms) CreateEmailLog(ctx context.Context, in *EmailLogInfo, opts ...grpc.CallOption) (*BaseUUIDResp, error) {
	client := mcms.NewMcmsClient(m.cli.Conn())
	return client.CreateEmailLog(ctx, in, opts...)
}

func (m *defaultMcms) UpdateEmailLog(ctx context.Context, in *EmailLogInfo, opts ...grpc.CallOption) (*BaseResp, error) {
	client := mcms.NewMcmsClient(m.cli.Conn())
	return client.UpdateEmailLog(ctx, in, opts...)
}

func (m *defaultMcms) GetEmailLogList(ctx context.Context, in *EmailLogListReq, opts ...grpc.CallOption) (*EmailLogListResp, error) {
	client := mcms.NewMcmsClient(m.cli.Conn())
	return client.GetEmailLogList(ctx, in, opts...)
}

func (m *defaultMcms) GetEmailLogById(ctx context.Context, in *UUIDReq, opts ...grpc.CallOption) (*EmailLogInfo, error) {
	client := mcms.NewMcmsClient(m.cli.Conn())
	return client.GetEmailLogById(ctx, in, opts...)
}

func (m *defaultMcms) DeleteEmailLog(ctx context.Context, in *UUIDsReq, opts ...grpc.CallOption) (*BaseResp, error) {
	client := mcms.NewMcmsClient(m.cli.Conn())
	return client.DeleteEmailLog(ctx, in, opts...)
}

// EmailProvider management
func (m *defaultMcms) CreateEmailProvider(ctx context.Context, in *EmailProviderInfo, opts ...grpc.CallOption) (*BaseIDResp, error) {
	client := mcms.NewMcmsClient(m.cli.Conn())
	return client.CreateEmailProvider(ctx, in, opts...)
}

func (m *defaultMcms) UpdateEmailProvider(ctx context.Context, in *EmailProviderInfo, opts ...grpc.CallOption) (*BaseResp, error) {
	client := mcms.NewMcmsClient(m.cli.Conn())
	return client.UpdateEmailProvider(ctx, in, opts...)
}

func (m *defaultMcms) GetEmailProviderList(ctx context.Context, in *EmailProviderListReq, opts ...grpc.CallOption) (*EmailProviderListResp, error) {
	client := mcms.NewMcmsClient(m.cli.Conn())
	return client.GetEmailProviderList(ctx, in, opts...)
}

func (m *defaultMcms) GetEmailProviderById(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*EmailProviderInfo, error) {
	client := mcms.NewMcmsClient(m.cli.Conn())
	return client.GetEmailProviderById(ctx, in, opts...)
}

func (m *defaultMcms) DeleteEmailProvider(ctx context.Context, in *IDsReq, opts ...grpc.CallOption) (*BaseResp, error) {
	client := mcms.NewMcmsClient(m.cli.Conn())
	return client.DeleteEmailProvider(ctx, in, opts...)
}

func (m *defaultMcms) SendSms(ctx context.Context, in *SmsInfo, opts ...grpc.CallOption) (*BaseUUIDResp, error) {
	client := mcms.NewMcmsClient(m.cli.Conn())
	return client.SendSms(ctx, in, opts...)
}

// SmsLog management
func (m *defaultMcms) CreateSmsLog(ctx context.Context, in *SmsLogInfo, opts ...grpc.CallOption) (*BaseUUIDResp, error) {
	client := mcms.NewMcmsClient(m.cli.Conn())
	return client.CreateSmsLog(ctx, in, opts...)
}

func (m *defaultMcms) UpdateSmsLog(ctx context.Context, in *SmsLogInfo, opts ...grpc.CallOption) (*BaseResp, error) {
	client := mcms.NewMcmsClient(m.cli.Conn())
	return client.UpdateSmsLog(ctx, in, opts...)
}

func (m *defaultMcms) GetSmsLogList(ctx context.Context, in *SmsLogListReq, opts ...grpc.CallOption) (*SmsLogListResp, error) {
	client := mcms.NewMcmsClient(m.cli.Conn())
	return client.GetSmsLogList(ctx, in, opts...)
}

func (m *defaultMcms) GetSmsLogById(ctx context.Context, in *UUIDReq, opts ...grpc.CallOption) (*SmsLogInfo, error) {
	client := mcms.NewMcmsClient(m.cli.Conn())
	return client.GetSmsLogById(ctx, in, opts...)
}

func (m *defaultMcms) DeleteSmsLog(ctx context.Context, in *UUIDsReq, opts ...grpc.CallOption) (*BaseResp, error) {
	client := mcms.NewMcmsClient(m.cli.Conn())
	return client.DeleteSmsLog(ctx, in, opts...)
}

// SmsProvider management
func (m *defaultMcms) CreateSmsProvider(ctx context.Context, in *SmsProviderInfo, opts ...grpc.CallOption) (*BaseIDResp, error) {
	client := mcms.NewMcmsClient(m.cli.Conn())
	return client.CreateSmsProvider(ctx, in, opts...)
}

func (m *defaultMcms) UpdateSmsProvider(ctx context.Context, in *SmsProviderInfo, opts ...grpc.CallOption) (*BaseResp, error) {
	client := mcms.NewMcmsClient(m.cli.Conn())
	return client.UpdateSmsProvider(ctx, in, opts...)
}

func (m *defaultMcms) GetSmsProviderList(ctx context.Context, in *SmsProviderListReq, opts ...grpc.CallOption) (*SmsProviderListResp, error) {
	client := mcms.NewMcmsClient(m.cli.Conn())
	return client.GetSmsProviderList(ctx, in, opts...)
}

func (m *defaultMcms) GetSmsProviderById(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*SmsProviderInfo, error) {
	client := mcms.NewMcmsClient(m.cli.Conn())
	return client.GetSmsProviderById(ctx, in, opts...)
}

func (m *defaultMcms) DeleteSmsProvider(ctx context.Context, in *IDsReq, opts ...grpc.CallOption) (*BaseResp, error) {
	client := mcms.NewMcmsClient(m.cli.Conn())
	return client.DeleteSmsProvider(ctx, in, opts...)
}
