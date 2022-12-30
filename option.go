package zlog

import "path/filepath"

type Options struct {
	Development  bool
	LogFileDir   string
	AppName      string
	MaxSize      int //文件多大开始切分
	MaxBackups   int //保留文件个数
	MaxAge       int //文件保留最大实际
	Level        string
	CtxKey       string //通过 ctx 传递 zlog 对象
	WriteFile    bool
	WriteConsole bool
}

type ZLogOptions func(*Options)

func newOptions(opts ...ZLogOptions) *Options {
	opt := &Options{
		Development:  true,
		AppName:      "zlog-app",
		MaxSize:      100,
		MaxBackups:   60,
		MaxAge:       30,
		Level:        "debug",
		CtxKey:       "zlog-ctx",
		WriteFile:    false,
		WriteConsole: true,
	}
	opt.LogFileDir, _ = filepath.Abs(filepath.Dir(filepath.Join(".")))
	opt.LogFileDir += "/logs/"
	for _, o := range opts {
		o(opt)
	}
	return opt
}

func SetDevelopment(development bool) ZLogOptions {
	return func(options *Options) {
		options.Development = development
	}
}

func SetLogFileDir(logFileDir string) ZLogOptions {
	return func(options *Options) {
		options.LogFileDir = logFileDir
	}
}

func SetAppName(appName string) ZLogOptions {
	return func(options *Options) {
		options.AppName = appName
	}
}

func SetMaxSize(maxSize int) ZLogOptions {
	return func(options *Options) {
		options.MaxSize = maxSize
	}
}
func SetMaxBackups(maxBackups int) ZLogOptions {
	return func(options *Options) {
		options.MaxBackups = maxBackups
	}
}
func SetMaxAge(maxAge int) ZLogOptions {
	return func(options *Options) {
		options.MaxAge = maxAge
	}
}

func SetLevel(level string) ZLogOptions {
	return func(options *Options) {
		options.Level = level
	}
}

func SetCtxKey(ctxKey string) ZLogOptions {
	return func(options *Options) {
		options.CtxKey = ctxKey
	}
}

func SetWriteFile(writeFile bool) ZLogOptions {
	return func(options *Options) {
		options.WriteFile = writeFile
	}
}

func SetWriteConsole(writeConsole bool) ZLogOptions {
	return func(options *Options) {
		options.WriteConsole = writeConsole
	}
}
