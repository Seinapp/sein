package server

import (
	"os"
	"path/filepath"

	"github.com/Seinapp/sein/stable/go/desktop-helper/src/config"
	"github.com/Seinapp/sein/stable/go/desktop-helper/src/lgr"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"
)

// newLogger returns a new logger that write on the specified file
func newLogger(logDirectory string, printer lgr.Printer) *zap.Logger {
	w := zapcore.AddSync(&lumberjack.Logger{
		Filename:   filepath.Join(logDirectory, config.HelperServerLogRelPath),
		MaxSize:    20,
		MaxBackups: 3,
		MaxAge:     30,
	})

	var stdoutEncoder zapcore.Encoder
	if printer.Type() == lgr.TypeJSON {
		stdoutEncoder = zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
	} else {
		stdoutEncoder = zapcore.NewConsoleEncoder(zap.NewProductionEncoderConfig())
	}

	core := zapcore.NewTee(
		zapcore.NewCore(stdoutEncoder, os.Stdout, zap.InfoLevel),
		// File logger with lumberjack, we always use JSON for the files
		zapcore.NewCore(
			zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
			w,
			zap.InfoLevel,
		),
	)
	return zap.New(core)
}
