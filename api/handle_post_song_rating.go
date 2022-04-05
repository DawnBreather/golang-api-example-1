package api

import (
  "fmt"
  "github.com/gin-gonic/gin"
  "net/http"
  "src/dto"
  "src/utils"
)

const (
  RATING_MIN_LIMIT = 1
  RATING_MAX_LIMIT = 5
)

// TODO: add headers
//- POST /song/rating?song_id=theid?rating=2
//- Takes in parameter a "song_id" and a "rating"
//- This call adds a rating to the song. Ratings should be between 1 and 5
func PostSongRating(ctx *gin.Context){
  if songId, rating, ok := extractUriParametersForPostSongsRating(ctx); ok {
    if updateResult, ok := dto.AppendRatingToSongById(songId, rating); ok {
      if updateResult.ModifiedCount > 0 {
        ctx.Status(http.StatusOK)
      }
    } else {
      ctx.String(http.StatusMethodNotAllowed, fmt.Sprintf("{song_id} value must be a valid hex"))
    }
    ctx.Status(http.StatusNotFound)
  } else {
    ctx.String(http.StatusMethodNotAllowed, fmt.Sprintf("{rating} value must be >= %d and <= %d", RATING_MIN_LIMIT, RATING_MAX_LIMIT))
  }
}

func extractUriParametersForPostSongsRating(ctx *gin.Context) (songId string, rating uint8, ok bool){
  songId = ctx.DefaultQuery("song_id", "")
  rating  = utils.MustParseStringToUint8(ctx.DefaultQuery("rating", "0"))
  return songId, rating, IsRatingParameterValueOkForPostSongsRating(rating)
}

func IsRatingParameterValueOkForPostSongsRating(rating uint8) (ok bool){
  return rating >= RATING_MIN_LIMIT && rating <= RATING_MAX_LIMIT
}