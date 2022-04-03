package main

import (
  "go.mongodb.org/mongo-driver/mongo"
  "go.mongodb.org/mongo-driver/mongo/options"
  "log"
)

const(
  DTO_TIMEOUT_SECONDS = 10
)

var dtoClient *mongo.Client = nil


func GetDtoClient() *mongo.Client{
  initializeDtoClient()
  return dtoClient
}

func initializeDtoClient(){
  var err error
  if dtoClient == nil {
    ctx, cancel := getContextWithTimeout(DTO_TIMEOUT_SECONDS)
    defer cancel()
    dtoClient, err = mongo.Connect(ctx, getMongoDbClientOptions())
    handleUnsuccessfulInitializeDtoClient(err)
  }
}

func handleUnsuccessfulInitializeDtoClient(err error){
  logUnsuccessfulGetDtoClient(err)
}

func getMongoDbClientOptions() *options.ClientOptions{
  return options.Client().
    ApplyURI(GetMongoDbConnectionString()).
    SetServerAPIOptions(options.ServerAPI(options.ServerAPIVersion1))
}

func logUnsuccessfulGetDtoClient(err error){
  if err != nil {
    log.Fatalf("Unable to initialize MongoDB connection: %v", err)
  }
}