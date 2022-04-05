package utils

import "go.mongodb.org/mongo-driver/bson"

func ConvertArrayOfBsonDToArrayOfMap(raw []bson.D) (res []map[string]interface{}){
  res = []map[string]interface{}{}
  for _, bsonDocument := range raw{
    res = append(res, bsonDocument.Map())
  }
  return
}