package smsprovider

import (
	"context"

	"github.com/iot-synergy/synergy-message-center/internal/svc"
	"github.com/iot-synergy/synergy-message-center/internal/utils/dberrorhandler"
	"github.com/iot-synergy/synergy-message-center/types/mcms"

	"github.com/iot-synergy/synergy-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetSmsProviderByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetSmsProviderByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSmsProviderByIdLogic {
	return &GetSmsProviderByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetSmsProviderByIdLogic) GetSmsProviderById(in *mcms.IDReq) (*mcms.SmsProviderInfo, error) {
	result, err := l.svcCtx.DB.SmsProvider.Get(l.ctx, in.Id)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &mcms.SmsProviderInfo{
		Id:        &result.ID,
		CreatedAt: pointy.GetPointer(result.CreatedAt.UnixMilli()),
		UpdatedAt: pointy.GetPointer(result.UpdatedAt.UnixMilli()),
		Name:      &result.Name,
		SecretId:  &result.SecretID,
		SecretKey: &result.SecretKey,
		Region:    &result.Region,
		IsDefault: &result.IsDefault,
	}, nil
}
