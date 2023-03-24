package logic

import (
	"cicd.robot/publictool"
	"cicd.robot/robot/internal/svc"
	"cicd.robot/robot/internal/types"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

type RobotLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRobotLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RobotLogic {
	return &RobotLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RobotLogic) Robot(req *types.FlyPigeonRequest) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	l.Logger.Infow("req", logx.Field("req", req))
	ChatIDArray := []string{req.ChatID}
	publictool.SendMarkDownContentRequest(req.BotKey, req.MsgContent, ChatIDArray)
	return &types.Response{
		Message: "OK,get fly pigeon request",
	}, nil
}
