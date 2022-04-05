package utils

import (
  "log"
  "strconv"
)

func MustParseStringToUint8(raw string) uint8{
  result, err := strconv.ParseInt(raw, 10, 8)
  logErrorOnMustParseStringToUint8(raw, err)

  return uint8(result)
}

func logErrorOnMustParseStringToUint8(raw string, err error){
  if err != nil {
    log.Printf("Error parsing {%s} to int64: %v", raw, err)
  }
}