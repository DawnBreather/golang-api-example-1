package dto

import (
  "go.mongodb.org/mongo-driver/bson"
  "go.mongodb.org/mongo-driver/mongo"
  "go.mongodb.org/mongo-driver/mongo/options"
)

//// ADD RATING
//// https://stackoverflow.com/a/68219291/4265419
//// https://stackoverflow.com/a/71270009/4265419
//// https://www.statology.org/mongodb-add-field/
/////////////////////////////////////////////////
func AppendRatingToSongById(idHexString string, rating uint8) (updateResult *mongo.UpdateResult, ok bool){
  if filter, ok := compileFilterForAppendRatingToSongByIdFunction(idHexString); ok {
    return UpdateOneDocument(
      GetSongsCollectionClient(),
      filter,
      compilePushForAppendRatingToSongByIdFunction(rating),
      options.Update().SetUpsert(true),
    ), ok
  }
  return
}

func compileFilterForAppendRatingToSongByIdFunction(objectIdHexString string) (filter bson.M, ok bool){
  if songId, ok := parseHexStringToObjectId(objectIdHexString); ok {
    return bson.M{
      "_id": songId,
    }, ok
  }
  return
}

func compilePushForAppendRatingToSongByIdFunction(rating uint8) bson.M{
  return bson.M{
    "$push": bson.M{
      "ratings": rating,
    },
  }
}