package main

import (
	"net/http"

	"github.com/Lieoxc/zlog"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

func AddTraceId() gin.HandlerFunc {
	return func(g *gin.Context) {
		traceId := g.GetHeader("traceId")
		if traceId == "" {
			traceId = uuid.New().String()
		}
		ctx, log := zlog.GetLogger().AddCtx(g.Request.Context(), zap.Any("traceId", traceId))
		g.Request = g.Request.WithContext(ctx)
		log.Info("AddTraceId success")
		g.Next()
	}
}

// curl http://127.0.0.1:8888/test
func main() {
	zlog.NewLogger(zlog.SetDevelopment(true))
	g := gin.New()
	g.Use(AddTraceId())
	g.GET("/test", func(context *gin.Context) {
		log := zlog.GetLogger().GetCtx(context.Request.Context())
		log.Info("test")
		log.Debug("test")
		context.JSON(200, "success")
	})
	zlog.GetLogger().Info("web server init success")
	http.ListenAndServe(":8888", g)
}
