package zlog

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

func TestGetLogger(t *testing.T) {
	NewLogger(SetDevelopment(true))
	GetLogger().Info("zlog example success")
	// 可以在中间件内赋值
	ctx, zlog := GetLogger().AddCtx(context.Background(), zap.String("traceId", uuid.New().String()))
	zlog.Debug("TestGetLogger", zap.Any("t", "t"))
	FA(ctx)
	FB(ctx)

	// 可以在中间件内赋值
	ctx, zlog = GetLogger().AddCtx(context.Background(), zap.String("traceId", uuid.New().String()))
	zlog.Info("TestGetLogger", zap.Any("t", "t"))
	FA(ctx)
	FB(ctx)
}

func FA(ctx context.Context) {
	zlog := GetLogger().GetCtx(ctx)
	zlog.Info("FA", zap.Any("a", "a"))
}

func FB(ctx context.Context) {
	zlog := GetLogger().GetCtx(ctx)
	zlog.Info("FB", zap.Any("b", "b"))
	FC(ctx)
}
func FC(ctx context.Context) {
	zlog := GetLogger().GetCtx(ctx)
	zlog.Info("FC", zap.Any("c", "c"))
}
