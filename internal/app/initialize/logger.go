package initialize

import (
	"github.com/spf13/viper"
	"github.com/varluffy/gindemo/pkg/logger"
)

// InitLogger 初始化logger
func InitLogger() *logger.Logger {
	cfg := &logger.Config{
		Level:      viper.GetString("Log.Level"),
		Format:     viper.GetString("Log.Format"),
		Output:     viper.GetString("Log.Output"),
		OutputFile: viper.GetString("Log.OutputFile"),
		MaxSize:    viper.GetInt("Log.MaxSize"),
		MaxBackup:  viper.GetInt("Log.MaxBackup"),
		MaxAge:     viper.GetInt("Log.MaxAge"),
		Compress:   false,
	}
	return logger.NewLogger(cfg)
}
