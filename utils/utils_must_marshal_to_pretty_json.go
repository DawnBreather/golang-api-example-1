package utils

import (
  "encoding/json"
  "log"
)

func MustMarshalToPrettyJson(obj interface{}) []byte{
  res, err := json.MarshalIndent(obj, "  ", "  ")
  if err != nil {
    log.Printf("Error must marshalling object {%v} to pretty JSON: %v", obj, err)
  }
  return res
}