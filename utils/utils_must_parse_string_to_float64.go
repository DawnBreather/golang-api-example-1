package utils

import (
  "log"
  "strconv"
)

func MustParseStringToFloat64(raw string) float64{
  result, err := strconv.ParseFloat(raw, 64)
  logErrorOnMustParseStringToFloat64(raw, err)

  return result
}

func logErrorOnMustParseStringToFloat64(raw string, err error)(isError bool){
  if err != nil {
    log.Printf("Error parsing {%s} to float64: %v", raw, err)
    isError = true
  }

  return
}