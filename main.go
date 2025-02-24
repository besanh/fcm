package main

import (
	"fcm/common/env"
	"fcm/common/util"
	"fcm/pkgs/mongodb"
	"log/slog"
	"strings"

	log "github.com/besanh/logger/logging/slog"

	"github.com/fluent/fluent-logger-golang/fluent"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(".env"); err != nil {
		panic(err)
	}

	initLogger()
	initMongoDb()
	initFcm()
}

func initLogger() {
	logFile := "tmp/console.log"
	logLevel := log.LEVEL_DEBUG
	switch env.GetStringENV("LOG_LEVEL", "error") {
	case "debug":
		logLevel = log.LEVEL_DEBUG
	case "info":
		logLevel = log.LEVEL_INFO
	case "error":
		logLevel = log.LEVEL_ERROR
	case "warn":
		logLevel = log.LEVEL_WARN
	}
	opts := []log.Option{}
	opts = append(opts, log.WithLevel(logLevel),
		log.WithRotateFile(logFile),
		log.WithFileSource(),
		log.WithTraceId(),
		log.WithAttrs(slog.Attr{
			Key: "environment", Value: slog.StringValue(env.GetStringENV("ENVIRONMENT", "local")),
		}),
	)
	if env.GetStringENV("LOG_SERVER", "") != "" {
		// get server and port from env
		arr := strings.Split(env.GetStringENV("LOG_SERVER", ""), ":")
		if len(arr) >= 2 {
			tag := "fcm"
			client, err := fluent.New(fluent.Config{FluentPort: int(util.ParseInt64(arr[1])), FluentHost: arr[0]})
			if err != nil {
				log.Error(err)
			} else {
				opts = append(opts, log.WithFluentd(client, tag))
			}
		}
	}
	log.SetLogger(log.NewSLogger(opts...))
}

func initMongoDb() {
	mongodbConfig := mongodb.MongoDBConfig{
		Username:      env.GetStringENV("MONGODB_USERNAME", "root"),
		Password:      env.GetStringENV("MONGODB_PASSWORD", "anhle@!*2025"),
		Host:          env.GetStringENV("MONGODB_HOST", "localhost"),
		Port:          env.GetIntENV("MONGODB_PORT", 27017),
		Database:      env.GetStringENV("MONGODB_DATABASE", "fcm"),
		DefaultAuthDb: env.GetStringENV("MONGODB_DEFAULT_AUTH_DB", "admin"),
	}

	var err error
	var db mongodb.IMongoDBClient
	db, err = mongodb.NewMongoDBClient(mongodbConfig)
	if err != nil {
		log.Errorf("mongodb connect error: %v", err)
		panic(err)
	}
	initRepositories(db)
}

func initFcm() {
	// fcm
}

func initRepositories(db mongodb.IMongoDBClient) {

}

func main() {
	engine := gin.Default()

	engine.Run(":8000")
}
