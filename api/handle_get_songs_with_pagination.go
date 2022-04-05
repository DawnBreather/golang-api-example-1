package api

import (
  "github.com/gin-gonic/gin"
  "net/http"
  "src/dto"
  "src/utils"
)

// TODO: add headers
// GET /songs
// with pagination
// /songs?page=1&limit=10
func GetSongsWithPagination(ctx *gin.Context){
  if ! handleSuccessfulGetSongsWithPagination(ctx){
    handleUnsuccessfulGetSongsWithPagination(ctx)
  }
}

func handleSuccessfulGetSongsWithPagination(ctx *gin.Context) (ok bool){
  page, limit := extractUriParametersForGetSongsWithPagination(ctx)
  if songsWithPaginationData, ok := dto.GetAllSongsWithPagination(page, limit); ok {
    ctx.String(http.StatusOK, string(utils.MustMarshalToJson(songsWithPaginationData)))
  }
  return
}

func extractUriParametersForGetSongsWithPagination(ctx *gin.Context) (page, limit int64){
  return utils.MustParseStringToInt64(ctx.DefaultQuery("page", "1")), utils.MustParseStringToInt64(ctx.DefaultQuery("limit", "10"))
}

func handleUnsuccessfulGetSongsWithPagination(ctx *gin.Context){
  ctx.Status(http.StatusInternalServerError)
}
