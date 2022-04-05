package api

import (
  "fmt"
  "github.com/gin-gonic/gin"
  "net/http"
  "src/dto"
  "src/utils"
)

// TODO: add headers
//- GET /song/rating/:song_id
//- Takes in parameter a "song_id" and a "rating"
//- This call adds a rating to the song. Ratings should be between 1 and 5
func GetSongRating(ctx *gin.Context){
  if result, ok := dto.GetSongRatingById(extractUriParametersForGetSongsRating(ctx)); ok {
    if len(result) > 0 {
      ctx.String(http.StatusOK, string(utils.MustMarshalToJson(result[0])))
    } else {
      ctx.Status(http.StatusNotFound)
    }
  } else {
    ctx.String(http.StatusMethodNotAllowed, fmt.Sprintf("{song_id} value must be a valid hex"))
  }
}

func extractUriParametersForGetSongsRating(ctx *gin.Context) (songId string){
 return ctx.Param("song_id")
}