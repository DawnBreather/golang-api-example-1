package dto

import (
  "go.mongodb.org/mongo-driver/bson"
  "go.mongodb.org/mongo-driver/mongo"
  "src/utils"
)

const(
  FIELD_NAME_FOR_DIFFICULTY_AVERAGE = "difficultyAvg"
)

// SEARCH ONLY LEVEL
// GET /songs/avg/difficulty
// https://www.mongodb.com/docs/drivers/go/current/fundamentals/crud/read-operations/project/
///////////////////////////////
func GetSongsAverageDifficultyByLevel(level int) (res float64, ok bool) {
  return handleResultsForGetSongsAverageDifficultyByLevelFunction(
    AggregateDocuments(
      GetSongsCollectionClient(),
      mongo.Pipeline{
        compileMatchBsonForGetSongsAverageDifficultyByLevelFunction(level),
        compileGroupBsonForGetSongsAverageDifficultyByLevelFunction(),
        compileProjectionBsonForGetSongsAverageDifficultyByLevelFunction(),
      },
    ),
  )
}

func handleResultsForGetSongsAverageDifficultyByLevelFunction(results []map[string]interface{}) (result float64, ok bool){
  if len(results) > 0 {
    return utils.MustParseInterfaceToFloat64(results[0][FIELD_NAME_FOR_DIFFICULTY_AVERAGE]), true
  }
  return
}

func compileMatchBsonForGetSongsAverageDifficultyByLevelFunction(level int) bson.D{
    return bson.D{
      {"$match", bson.D{
        {FIELD_NAME_FOR_LEVEL, bson.D{
          getMatchBsonElementByLevel(level),
        }},
      }},
    }
}

func getMatchBsonElementByLevel(level int) (res bson.E){
  if level > 0 {
    res = bson.E{ "$eq", level }
  } else {
    res = bson.E {"$gt", 0}
  }

  return
}

func compileGroupBsonForGetSongsAverageDifficultyByLevelFunction() bson.D {
  return bson.D{
    {"$group", bson.D{
      {FIELD_NAME_FOR_ID, nil},
      {FIELD_NAME_FOR_DIFFICULTY_AVERAGE, bson.D{
        {"$avg", "$" + FIELD_NAME_FOR_DIFFICULTY},
      }},
    }},
  }
}

func compileProjectionBsonForGetSongsAverageDifficultyByLevelFunction() bson.D{
  return bson.D{
    {"$project", bson.D{
      {FIELD_NAME_FOR_ID, 0},
    }},
  }
}