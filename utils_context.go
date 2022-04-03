package main

import (
  "context"
  "time"
)

func getContextWithTimeout(seconds time.Duration) (context.Context, context.CancelFunc){
  return context.WithTimeout(context.Background(), seconds*time.Second)
}