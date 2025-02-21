package main

import (
	"fcm/common/env"
	"fcm/common/log"
	"fcm/internals/mongodb"
)

func init() {
	initMongoDb()
	initFcm()
}

func initMongoDb() {
	mongodbConfig := mongodb.MongoDBConfig{
		Username: env.GetStringENV("FCM_MONGODB_USERNAME", "root"),
		Password: env.GetStringENV("FCM_MONGODB_PASSWORD", "anhle@!*2025"),
		Host:     env.GetStringENV("FCM_MONGODB_HOST", "localhost"),
		Port:     env.GetIntENV("FCM_MONGODB_PORT", 27017),
		Database: env.GetStringENV("FCM_MONGODB_DATABASE", "fcm"),
	}

	var err error
	// var db mongodb.IMongoDBClient
	_, err = mongodb.NewMongoDBClient(mongodbConfig)
	if err != nil {
		log.Errorf("mongodb connect error: %v", err)
		panic(err)
	}
}

func initFcm() {
	// fcm
}

func main() {

}
