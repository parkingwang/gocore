package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-irain/logger"
)

type after struct {
	hooks []CustomHookFunc
}

func (mw *after) MiddleWare(egn *gin.Engine) {
	egn.Use(func(ctx *gin.Context) {
		func(ctx *gin.Context) {
			ctx.Next()
		}(ctx)

		// 执行自定义钩子方法
		for _, hook := range mw.hooks {
			hook(ctx)
		}
		logResponse(ctx)  // 记录响应信息
		markResponse(ctx) // 标记响应信息
	})
}

// NewAfterMW 后置中间件
func NewAfterMW(hooks []CustomHookFunc) MiddleWarer {
	return &after{hooks: hooks}
}

// logResponse 响应信息记录到日志
func logResponse(ctx *gin.Context) {
	RequestID, _ := ctx.Get("request_id")
	data, _ := ctx.Get("response_data")
	logger.Info(RequestID, "Reponse:", data)
}

// markResponse 标记响应
func markResponse(ctx *gin.Context) {
	RequestID, _ := ctx.Get("request_id")
	sTime, _ := ctx.Get("http_stime")
	eTime := time.Now()
	duration := eTime.Sub(sTime.(time.Time))
	logger.Info(RequestID, "******* Duration:", duration.String(), "<", sTime, " ~ ", eTime, ">", "*******")
}
