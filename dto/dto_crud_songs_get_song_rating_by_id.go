package dto

import (
  "go.mongodb.org/mongo-driver/bson"
  "go.mongodb.org/mongo-driver/mongo"
)

const (
  FIELD_NAME_FOR_RATINGS_AVERAGE = "avg"
  FIELD_NAME_FOR_RATINGS_MINIMUM = "min"
  FIELD_NAME_FOR_RATINGS_MAXIMUM = "max"
)

// GET SONG'S RATING BY ID
// https://www.mongodb.com/docs/manual/reference/operator/aggregation/avg/
// https://www.mongodb.com/community/forums/t/mongodb-filter-an-array-by-id-using-aggregate/114182
//////////////////////////////////////////////////////////////////////////////////////////////////
func GetSongRatingById(idHexString string) (rating []map[string]interface{}, ok bool){
  if match, ok := compileMatchForGetSongRatingByIdAggregation(idHexString); ok {
    return AggregateDocuments(
      GetSongsCollectionClient(),
      mongo.Pipeline{
        match,
        compileProjectionForGetSongRatingByIdAggregation(),
      },
    ), ok
  }
  return
}

func compileMatchForGetSongRatingByIdAggregation(idHexString string) (match bson.D, ok bool) {
  if songId, ok := parseHexStringToObjectId(idHexString); ok {
    return bson.D{
      {"$match", bson.D{
        {"_id", songId},
      }},
    }, ok
  }
  return
}

func compileProjectionForGetSongRatingByIdAggregation() bson.D {
  return bson.D{
    {"$project", bson.D{
      {FIELD_NAME_FOR_ID, 1},
      {FIELD_NAME_FOR_RATINGS_AVERAGE, bson.D{
        {"$avg", "$" + FIELD_NAME_FOR_RATINGS},
      }},
      {FIELD_NAME_FOR_RATINGS_MINIMUM, bson.D{
        {"$min", "$" + FIELD_NAME_FOR_RATINGS},
      }},
      {FIELD_NAME_FOR_RATINGS_MAXIMUM, bson.D{
        {"$max", "$" + FIELD_NAME_FOR_RATINGS},
      }},
    }},
  }
}