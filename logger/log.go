package logger

import (
	"context"
	"fmt"
	"log"
	"os"
)

// LogLevel 表示日志的级别，包括 Error、Warn 和 Info 三种级别
type LogLevel int

const (
	LogLevelError LogLevel = iota
	LogLevelWarn
	LogLevelInfo
)

// customLogger 表示自定义的日志对象
type customLogger struct {
	logger *log.Logger // 内部使用的标准库 logger 对象
	level  LogLevel    // 当前的日志级别
}

var (
	// Logger 表示全局的 CustomLogger 对象
	Logger *customLogger
)

// init 是包的初始化函数，会在包被导入时自动调用
func init() {
	Logger = &customLogger{
		logger: log.New(os.Stderr, "", log.Ldate|log.Ltime|log.Lshortfile),
		level:  LogLevelInfo, // 默认为 Info 级别
	}
}

// SetLevel 设置日志级别
func (l *customLogger) SetLevel(level LogLevel) {
	l.level = level
}

// Error 输出 Error 级别的日志
func (l *customLogger) Error(message string) {
	if l.level >= LogLevelError {
		l.logger.Output(2, fmt.Sprintf("[ERROR] %s", message))
	}
}

// Errorc 输出 Error 级别的日志，包括上下文信息
func (l *customLogger) Errorc(ctx context.Context, message string) {
	if l.level >= LogLevelError {
		l.logger.Output(2, fmt.Sprintf("[ERROR] %s: %s", message, ctx.Value("key")))
	}
}

// Errorf 输出 Error 级别的日志，包括上下文信息，支持格式化字符串
func (l *customLogger) Errorf(ctx context.Context, format string, args ...interface{}) {
	if l.level >= LogLevelError {
		message := fmt.Sprintf(format, args...)
		l.logger.Output(2, fmt.Sprintf("[ERROR] %s: %s", message, ctx.Value("key")))
	}
}

// Warn 输出 Warn 级别的日志
func (l *customLogger) Warn(message string) {
	if l.level >= LogLevelWarn {
		l.logger.Output(2, fmt.Sprintf("[WARN] %s", message))
	}
}

// Warnc 输出 Warn 级别的日志，包括上下文信息
func (l *customLogger) Warnc(ctx context.Context, message string) {
	if l.level >= LogLevelWarn {
		l.logger.Output(2, fmt.Sprintf("[WARN] %s: %s", message, ctx.Value("key")))
	}
}

// Warnf 输出 Warn 级别的日志，包括上下文信息，支持格式化字符串
func (l *customLogger) Warnf(ctx context.Context, format string, args ...interface{}) {
	if l.level >= LogLevelError {
		message := fmt.Sprintf(format, args...)
		l.logger.Output(2, fmt.Sprintf("[WARN] %s: %s", message, ctx.Value("key")))
	}
}

// Info 输出 Info 级别的日志
func (l *customLogger) Info(message string) {
	if l.level >= LogLevelInfo {
		l.logger.Output(2, fmt.Sprintf("[INFO] %s", message))
	}
}

// Infof 输出 Info 级别的日志，包括上下文信息，支持格式化字符串
func (l *customLogger) Infof(ctx context.Context, format string, args ...interface{}) {
	if l.level >= LogLevelError {
		message := fmt.Sprintf(format, args...)
		l.logger.Output(2, fmt.Sprintf("[INFO] %s: %s", message, ctx.Value("key")))
	}
}

// Infoc 输出 Info 级别的日志，包括上下文信息
func (l *customLogger) Infoc(ctx context.Context, message string) {
	if l.level >= LogLevelInfo {
		l.logger.Output(2, fmt.Sprintf("[INFO] %s: %s", message, ctx.Value("key")))
	}
}
