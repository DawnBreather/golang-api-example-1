package dto

import (
  "go.mongodb.org/mongo-driver/bson/primitive"
  "log"
)

func parseHexStringToObjectId(hexString string) (objectId primitive.ObjectID, ok bool){
  var err error
  objectId, err = primitive.ObjectIDFromHex(hexString)
  if err != nil {
    log.Printf("Error parsing hexString {%s} to objectId: %v", hexString, err)
    return primitive.NilObjectID, ok
  }

  return objectId, true
}