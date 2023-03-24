package handler

import (
	"net/http"

	"cicd.robot/robot/internal/logic"
	"cicd.robot/robot/internal/svc"
	"cicd.robot/robot/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func RobotHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FlyPigeonRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewRobotLogic(r.Context(), svcCtx)
		resp, err := l.Robot(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
