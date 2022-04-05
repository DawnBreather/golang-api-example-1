package utils

import (
  "log"
  "strconv"
)

func MustParseStringToInt64(raw string) int64{
  result, err := strconv.ParseInt(raw, 10, 64)
  logErrorOnMustParseStringToInt64(raw, err)

  return result
}

func logErrorOnMustParseStringToInt64(raw string, err error){
  if err != nil {
    log.Printf("Error parsing {%s} to int64: %v", raw, err)
  }
}