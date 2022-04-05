package api

import (
  "fmt"
  "github.com/gin-gonic/gin"
  "net/http"
  "src/dto"
  "src/utils"
)

// TODO: add headers
// GET /songs/avg/difficulty
// with pagination
// /songs/avg/difficulty?level=13
func GetSongsAverageDifficulty(ctx *gin.Context){
  if ! handleSuccessfulGetSongsAverageDifficulty(ctx){
    handleUnsuccessfulGetSongsAverageDifficulty(ctx)
  }
}

func handleSuccessfulGetSongsAverageDifficulty(ctx *gin.Context) (ok bool){
  if songsWithPaginationData, ok := dto.GetSongsAverageDifficultyByLevel(extractUriParametersForGetSongsAverageDifficulty(ctx)); ok {
    ctx.String(http.StatusOK, fmt.Sprintf("%.2f", songsWithPaginationData))
  }
  return
}

func extractUriParametersForGetSongsAverageDifficulty(ctx *gin.Context) (level int){
  return utils.MustParseStringToInt32(ctx.DefaultQuery("level", "-1"))
}

func handleUnsuccessfulGetSongsAverageDifficulty(ctx *gin.Context){
  ctx.Status(http.StatusInternalServerError)
}
