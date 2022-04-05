package dto

import (
  "go.mongodb.org/mongo-driver/bson"
  "go.mongodb.org/mongo-driver/mongo"
)

const (
  SONGS_BY_MESSAGE_SEARCH_INDEX = "artists_and_title"
)

const (
  FIELD_NAME_FOR_ID = "_id"
  FIELD_NAME_FOR_TITLE = "title"
  FIELD_NAME_FOR_ARTIST = "artist"
  FIELD_NAME_FOR_LEVEL = "level"
  FIELD_NAME_FOR_RELEASED = "released"
  FIELD_NAME_FOR_DIFFICULTY = "difficulty"
  FIELD_NAME_FOR_RATINGS = "ratings"
  FIELD_NAME_FOR_SCORE = "score"
)

// SEARCH SONGS BY MESSAGE
// GET /songs/search
// https://www.mongodb.com/docs/atlas/atlas-search/tutorial/run-query/#std-label-fts-tutorial-run-query
// https://www.mongodb.com/basics/search-index
// https://www.mongodb.com/docs/drivers/go/master/fundamentals/crud/read-operations/text/
// https://www.mongodb.com/docs/manual/tutorial/text-search-in-aggregation/#text-score
///////////////////////////////
func GetSongsByMessage(message string, relevanceScoreBaseline float64) []map[string]interface{}{
  return AggregateDocuments(
    GetSongsCollectionClient(),
    mongo.Pipeline{
      compileQueryForGetSongsByMessageAggregation(message),
      compileProjectionForGetSongsByMessageAggregation(),
      compileMatchForGetSongsByMessageAggregation(relevanceScoreBaseline),
      compileSortForGetSongsByMessageAggregation(),
    },
  )
}

func compileQueryForGetSongsByMessageAggregation(message string) bson.D{
  return bson.D{
    {"$search", bson.D{
      {"index", SONGS_BY_MESSAGE_SEARCH_INDEX},
      {"text", bson.D{
        {"query", message},
        {"path", bson.D{
          {"wildcard", "*"},
        }},
      }},
    }},
  }
}

func compileProjectionForGetSongsByMessageAggregation() bson.D{
  return bson.D{
    {"$project", bson.D{
      {FIELD_NAME_FOR_ID, 0},
      {FIELD_NAME_FOR_TITLE, 1},
      {FIELD_NAME_FOR_ARTIST, 1},
      {FIELD_NAME_FOR_LEVEL, 1},
      {FIELD_NAME_FOR_RELEASED, 1},
      {FIELD_NAME_FOR_DIFFICULTY, 1},
      {FIELD_NAME_FOR_RATINGS, 1},
      {FIELD_NAME_FOR_SCORE, bson.D{
        {"$meta", "searchScore"},
      }},
    }},
  }
}

func compileMatchForGetSongsByMessageAggregation(relevanceScoreBaseline float64) bson.D{
  return bson.D{
    {"$match", bson.D{
      {FIELD_NAME_FOR_SCORE, bson.D{
        {"$gt", relevanceScoreBaseline},
      }},
    }},
  }
}

func compileSortForGetSongsByMessageAggregation() bson.D{
  return bson.D{
    {"$sort", bson.D{
      {FIELD_NAME_FOR_SCORE, -1},
      {FIELD_NAME_FOR_ARTIST, 1},
      {FIELD_NAME_FOR_TITLE, 1},
    },
    },
  }
}