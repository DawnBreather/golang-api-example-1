package utils

import (
  "context"
  "time"
)

func GetContextWithTimeout(seconds time.Duration) (context.Context, context.CancelFunc){
  return context.WithTimeout(context.Background(), seconds*time.Second)
}