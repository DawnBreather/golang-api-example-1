package dto

import (
  "context"
  "go.mongodb.org/mongo-driver/mongo"
  "go.mongodb.org/mongo-driver/mongo/options"
  "log"
)

func UpdateOneDocument(collectionClient *mongo.Collection, filter interface{}, update interface{}, options *options.UpdateOptions) *mongo.UpdateResult{
  updateResult, err := GetSongsCollectionClient().UpdateOne(context.TODO(), filter, update, options)
  if err != nil {
    log.Printf("Error on update one document: %v", err)
  }

  return updateResult
}