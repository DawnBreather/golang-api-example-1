package dto

import (
  "fmt"
  "strings"
)

func getNMongodbShardsNames(shardBaseTemplate string, n int)(shardsNames []string){
  for i := 0; i < n; i++ {
    shardsNames = append(shardsNames, strings.ReplaceAll(shardBaseTemplate, "{#}", fmt.Sprintf("%02d", i)))
  }
  return
}