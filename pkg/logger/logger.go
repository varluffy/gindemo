package logger

import (
	"context"
	"fmt"
	"go.uber.org/zap"
)

var log *Logger

const (
	TraceIDKey      = "trace_id"
	UserIDKey       = "user_id"
	SpanTitleKey    = "span_title"
	SpanFunctionKey = "span_function"
	VersionKey      = "version"
)

// TraceIDFunc 定义获取跟踪ID的函数
type TraceIDFunc func() string

var (
	version     string
	traceIDFunc TraceIDFunc
)

type Logger struct {
	_log *zap.Logger
}

type Config struct {
	Level      string // debug/info/error
	Format     string // text/json
	Output     string // stdout/stderr/file
	OutputFile string
	MaxSize    int  // 单个文件最大 MB
	MaxBackup  int  // 最多备份几个
	MaxAge     int  // 保留多长时间
	Compress   bool // 是否压缩
}

func NewLogger(cfg *Config) *Logger {
	log = &Logger{_log: newZapLogger(cfg)}
	return log
}

func (l Logger) Debug(args ...interface{}) {
	l._log.Debug(fmt.Sprint(args...))
}

func (l Logger) Debugf(format string, args ...interface{}) {
	l._log.Debug(fmt.Sprintf(format, args...))
}

func (l Logger) Info(args ...interface{}) {
	l._log.Info(fmt.Sprint(args...))
}

func (l Logger) Infof(format string, args ...interface{}) {
	l._log.Info(fmt.Sprintf(format, args...))
}

func (l Logger) Warn(args ...interface{}) {
	l._log.Warn(fmt.Sprint(args...))
}

func (l Logger) Warnf(format string, args ...interface{}) {
	l._log.Warn(fmt.Sprintf(format, args...))
}

func (l Logger) Error(args ...interface{}) {
	l._log.Error(fmt.Sprint(args...))
}

func (l Logger) Errorf(format string, args ...interface{}) {
	l._log.Error(fmt.Sprintf(format, args...))
}

func (l Logger) Fatal(args ...interface{}) {
	l._log.Fatal(fmt.Sprint(args...))
}

func (l Logger) Fatalf(format string, args ...interface{}) {
	l._log.Fatal(fmt.Sprintf(format, args...))
}

func (l *Logger) With(key string, value interface{}) *Logger {
	return &Logger{l._log.With(zap.Any(key, value))}
}

func (l *Logger) WithField(key string, value interface{}) *Logger {
	return &Logger{l._log.With(zap.Any(key, value))}
}

func (l Logger) Sync() {
	_ = l._log.Sync()
}

func (l *Logger) WithFields(fields map[string]interface{}) *Logger {
	i := 0
	var log *Logger
	for k, v := range fields {
		if i == 0 {
			log = l.WithField(k, v)
		} else {
			log = log.WithField(k, v)
		}
		i++
	}
	return log
}

func Sync() {
	log.Sync()
}

// SetVersion 设定版本
func SetVersion(v string) {
	version = v
}

// SetTraceIDFunc 设定追踪ID的处理函数
func SetTraceIDFunc(fn TraceIDFunc) {
	traceIDFunc = fn
}

func getTraceID() string {
	if traceIDFunc != nil {
		return traceIDFunc()
	}
	return ""
}

type (
	traceIDContextKey struct{}
	userIDContextKey  struct{}
)

// NewTraceIDContext 创建跟踪ID上下文
func NewTraceIDContext(ctx context.Context, traceID string) context.Context {
	return context.WithValue(ctx, traceIDContextKey{}, traceID)
}

// FromTraceIDContext 从上下文中获取跟踪ID
func FromTraceIDContext(ctx context.Context) string {
	v := ctx.Value(traceIDContextKey{})
	if v != nil {
		if s, ok := v.(string); ok {
			return s
		}
	}
	return getTraceID()
}

// NewUserIDContext 创建用户ID上下文
func NewUserIDContext(ctx context.Context, userID string) context.Context {
	return context.WithValue(ctx, userIDContextKey{}, userID)
}

// FromUserIDContext 从上下文中获取用户ID
func FromUserIDContext(ctx context.Context) string {
	v := ctx.Value(userIDContextKey{})
	if v != nil {
		if s, ok := v.(string); ok {
			return s
		}
	}
	return ""
}

type spanOptions struct {
	Title    string
	FuncName string
}

// SpanOption 定义跟踪单元的数据项
type SpanOption func(*spanOptions)

// SetSpanTitle 设置跟踪单元的标题
func SetSpanTitle(title string) SpanOption {
	return func(o *spanOptions) {
		o.Title = title
	}
}

// SetSpanFuncName 设置跟踪单元的函数名
func SetSpanFuncName(funcName string) SpanOption {
	return func(o *spanOptions) {
		o.FuncName = funcName
	}
}

// StartSpan 开始一个追踪单元
func StartSpan(ctx context.Context, opts ...SpanOption) *Logger {
	if ctx == nil {
		ctx = context.Background()
	}

	var o spanOptions
	for _, opt := range opts {
		opt(&o)
	}

	fields := map[string]interface{}{
		VersionKey: version,
	}
	if v := FromTraceIDContext(ctx); v != "" {
		fields[TraceIDKey] = v
	}
	if v := FromUserIDContext(ctx); v != "" {
		fields[UserIDKey] = v
	}
	if v := o.Title; v != "" {
		fields[SpanTitleKey] = v
	}
	if v := o.FuncName; v != "" {
		fields[SpanFunctionKey] = v
	}

	return log.WithFields(fields)
}

func Debug(ctx context.Context, args ...interface{}) {
	StartSpan(ctx).Debug(args...)
}

func Debugf(ctx context.Context, format string, args ...interface{}) {
	StartSpan(ctx).Debugf(format, args...)
}

func Info(ctx context.Context, args ...interface{}) {
	StartSpan(ctx).Info(args...)
}

func Infof(ctx context.Context, format string, args ...interface{}) {
	StartSpan(ctx).Infof(format, args...)
}

func Warn(ctx context.Context, args ...interface{}) {
	StartSpan(ctx).Warn(args...)
}

func Warnf(ctx context.Context, format string, args ...interface{}) {
	StartSpan(ctx).Warnf(format, args...)
}

func Error(ctx context.Context, args ...interface{}) {
	StartSpan(ctx).Error(args)
}

func Errorf(ctx context.Context, format string, args ...interface{}) {
	StartSpan(ctx).Errorf(format, args...)
}

func Fatal(ctx context.Context, args ...interface{}) {
	StartSpan(ctx).Fatal(args...)
}

func Fatalf(ctx context.Context, format string, args ...interface{}) {
	StartSpan(ctx).Fatalf(format, args...)
}
