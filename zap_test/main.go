package main

import (
	"go.uber.org/zap"
)

func msain() {
	logger, _ := zap.NewProduction() //生产环境
	//logger, _ := zap.NewDevelopment()  //开发环境
	defer logger.Sync() // flushes buffer, if any
	url := "https://imooc.com"
	sugar := logger.Sugar()
	sugar.Infow("failed to fetch URL",
		// Structured context as loosely typed key-value pairs.
		"url", url,
		"attempt", 3,
	)
	sugar.Infof("Failed to fetch URL: %s", url)
}
