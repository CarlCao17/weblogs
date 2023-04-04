package log

import (
	"context"

	"github.com/cloudwego/kitex/pkg/klog"
)

func Fatalf(format string, v ...any) {
	klog.Fatalf(format, v...)
}

func Errorf(format string, v ...any) {
	klog.Errorf(format, v...)
}

func Warnf(format string, v ...any) {
	klog.Warnf(format, v...)
}

func Noticef(format string, v ...any) {
	klog.Noticef(format, v...)
}

func Infof(format string, v ...any) {
	klog.Infof(format, v...)
}

func Debugf(format string, v ...any) {
	klog.Debugf(format, v...)
}

func Tracef(format string, v ...any) {
	klog.Tracef(format, v...)
}

func CtxFatalf(ctx context.Context, format string, v ...any) {
	klog.CtxFatalf(ctx, format, v...)
}

func CtxErrorf(ctx context.Context, format string, v ...any) {
	klog.CtxErrorf(ctx, format, v...)
}

func CtxWarnf(ctx context.Context, format string, v ...any) {
	klog.CtxWarnf(ctx, format, v...)
}

func CtxNoticef(ctx context.Context, format string, v ...any) {
	klog.CtxNoticef(ctx, format, v...)
}

func CtxInfof(ctx context.Context, format string, v ...any) {
	klog.CtxInfof(ctx, format, v...)
}

func CtxDebugf(ctx context.Context, format string, v ...any) {
	klog.CtxDebugf(ctx, format, v...)
}

func CtxTracef(ctx context.Context, format string, v ...any) {
	klog.CtxTracef(ctx, format, v...)
}
