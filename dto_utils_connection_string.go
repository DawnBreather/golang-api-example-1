package main

import (
  "github.com/wissance/stringFormatter"
  "golang.org/x/exp/maps"
)

const (
  PRIMARY_DTO_PARAMETERS string = "PRIMARY"
  OPTIONAL_DTO_PARAMETERS string = "OPTIONAL"
)

var (
  dtoMongoDbConnectionParameters = map[string]interface{}{
    PRIMARY_DTO_PARAMETERS: map[string]interface{}{
      "schema":   "mongodb://",
      "username": "app",
      "password": "m8a5rITB8vIdsvnu",
      "shards": []string{
        "cluster0-shard-00-00.rusxw.mongodb.net:27017",
        "cluster0-shard-00-01.rusxw.mongodb.net:27017",
        "cluster0-shard-00-02.rusxw.mongodb.net:27017",
      },
    },
    OPTIONAL_DTO_PARAMETERS: map[string]interface{}{
      "retryWrites": true,
      "w":           "majority",
      "replicaSet":  "atlas-zhqegh-shard-0",
      "tls":         true,
    },
  }
)

func GetMongoDbConnectionString() string{
  return stringFormatter.FormatComplex("{schema}{username}:{password}@{shard0},{shard1},{shard2}/{parametersString}", getMongoDbConnectionParametersFinalMap())
}

func getMongoDbConnectionParametersFinalMap() (res map[string]interface{}){
  maps.Copy(res, dtoMongoDbConnectionParameters[PRIMARY_DTO_PARAMETERS])
  res["parametersString"] = getMongoDbConnectionOptionalParametersString()
  return res
}

func getMongoDbConnectionOptionalParametersString() string{
  return stringFormatter.FormatComplex("?retryWrites={0}&w={1}&replicaSet={2}&tls={3}", dtoMongoDbConnectionParameters[OPTIONAL_DTO_PARAMETERS].(map[string]interface{}))
}