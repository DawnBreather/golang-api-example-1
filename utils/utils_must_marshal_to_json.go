package utils

import (
  "encoding/json"
  "log"
)

func MustMarshalToJson(object interface{}) []byte{
  res, err := json.Marshal(object)
  logErrorOnMustMarshalToJson(res, err)

  return res
}

func logErrorOnMustMarshalToJson(object interface{}, err error){
  if err != nil {
    log.Printf("Error marshalling to JSON {%v}: %v", object, err)
  }
}