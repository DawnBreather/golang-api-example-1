package api

import (
  "github.com/gin-gonic/gin"
  "github.com/penglongli/gin-metrics/ginmetrics"
)

// TODO: add swagger
// TODO: add health check
func NewApplicationRouter() *gin.Engine{
  router := gin.New()

  router.
    GET("/songs", GetSongsWithPagination).
    GET("/songs/avg/difficulty", GetSongsAverageDifficulty).
    GET("/songs/search", GetSongsBySearchMessage).
    POST("/song/rating", PostSongRating).
    GET("/song/rating/:song_id", GetSongRating)

  //TODO: Move metrics to segregated router and socket
  m := ginmetrics.GetMonitor()
  m.SetMetricPath("/metrics")
  m.SetDuration([]float64{0.1, 0.3, 1.2, 5, 10})
  m.Use(router)

  return router
}