package utils

import (
  "log"
  "strconv"
)

func MustParseStringToInt32(raw string) int{
  result, err := strconv.Atoi(raw)
  logErrorOnMustParseStringToInt32(raw, err)

  return result
}

func logErrorOnMustParseStringToInt32(raw string, err error){
  if err != nil {
    log.Printf("Error parsing {%s} to int32: %v", raw, err)
  }
}