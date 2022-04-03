package main

import "go.mongodb.org/mongo-driver/mongo"

func GetCollectionClient(databaseName, collectionName string) *mongo.Collection{
  return GetDtoClient().Database(databaseName).Collection(collectionName)
}