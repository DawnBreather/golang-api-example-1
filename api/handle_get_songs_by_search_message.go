package api

import (
  "github.com/gin-gonic/gin"
  "io/ioutil"
  "net/http"
  "src/dto"
  "src/utils"
)

// TODO: add headers
// GET /songs/search
// /songs/search?relevanceScoreBaseline=0.5
// data="artist title"
func GetSongsBySearchMessage(ctx *gin.Context){
  if ! handleSuccessfulGetSongsBySearchMessage(ctx){
    handleUnsuccessfulGetSongsBySearchMessage(ctx)
  }
}

func handleSuccessfulGetSongsBySearchMessage(ctx *gin.Context) (isOk bool) {
  message, relevanceScoreBaseline, ok, err := extractParametersFromBodyOfGetSongsBySearchMessageRequest(ctx)
  if ok {
    songs := dto.GetSongsByMessage(message, relevanceScoreBaseline)
    ctx.String(http.StatusOK, string(utils.MustMarshalToJson(songs)))
    isOk = true
  } else {
    ctx.String(http.StatusMethodNotAllowed, err.Error())
  }
  return
}

func extractParametersFromBodyOfGetSongsBySearchMessageRequest(ctx *gin.Context) (message string, relevanceScoreBaseline float64, ok bool, err error){
  var messageBytes []byte

  relevanceScoreBaseline = utils.MustParseStringToFloat64(ctx.DefaultQuery("relevanceScoreBaseline", "0.5"))
  messageBytes, err = ioutil.ReadAll(ctx.Request.Body)
  if err != nil {
    return "", 0.0, false, err
  }

  return string(messageBytes), relevanceScoreBaseline, true, nil
}

func handleUnsuccessfulGetSongsBySearchMessage(ctx *gin.Context){
  ctx.Status(http.StatusInternalServerError)
}
