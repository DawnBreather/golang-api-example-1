package dto

import (
  "context"
  "fmt"
  "go.mongodb.org/mongo-driver/bson"
  "go.mongodb.org/mongo-driver/mongo"
  "src/utils"
)

func AggregateDocuments(collectionClient *mongo.Collection, pipeline []bson.D) []map[string]interface{}{
  var rawResults []bson.D

  cursor, err := collectionClient.Aggregate(context.TODO(), pipeline)
  if err != nil {
    fmt.Println(err)
  }

  err = cursor.All(context.TODO(), &rawResults)
  if err != nil {
    fmt.Println(err)
  }

  return utils.ConvertArrayOfBsonDToArrayOfMap(rawResults)
}