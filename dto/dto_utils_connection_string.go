package dto

import (
  "github.com/wissance/stringFormatter"
  "src/config"
)

//const (
//  PRIMARY_DTO_PARAMETERS string = "PRIMARY"
//  OPTIONAL_DTO_PARAMETERS string = "OPTIONAL"
//)
//
//var (
//  dtoMongoDbConnectionParameters = map[string]interface{}{
//    PRIMARY_DTO_PARAMETERS: map[string]interface{}{
//      "schema":   "mongodb://",
//      "username": "app",
//      "password": "m8a5rITB8vIdsvnu",
//      "shard0": "cluster0-shard-00-00.rusxw.mongodb.net:27017",
//      "shard1": "cluster0-shard-00-01.rusxw.mongodb.net:27017",
//      "shard2": "cluster0-shard-00-02.rusxw.mongodb.net:27017",
//
//    },
//    OPTIONAL_DTO_PARAMETERS: map[string]interface{}{
//      "retryWrites": true,
//      "w":           "majority",
//      "replicaSet":  "atlas-zhqegh-shard-0",
//      "tls":         true,
//    },
//  }
//)

func GetMongoDbConnectionString() string{
  //return stringFormatter.FormatComplex("{schema}{username}:{password}@{shard0},{shard1},{shard2}/{parametersString}", getMongoDbConnectionParametersFinalMap())
  return stringFormatter.FormatComplex("{schema}{username}:{password}@{shard0},{shard1},{shard2}/{parametersString}", map[string]interface{}{
    "schema": config.Get().Dto.Primary.Schema,
    "username": config.Get().Dto.Primary.Username,
    "password": config.Get().Dto.Primary.Password,
    "shard0": getNMongodbShardsNames(config.Get().Dto.Primary.ShardsNamesTemplate, 3)[0],
    "shard1": getNMongodbShardsNames(config.Get().Dto.Primary.ShardsNamesTemplate, 3)[1],
    "shard2": getNMongodbShardsNames(config.Get().Dto.Primary.ShardsNamesTemplate, 3)[2],
    "parametersString": stringFormatter.FormatComplex("?retryWrites={retryWrites}&w={w}&replicaSet={replicaSet}&tls={tls}", map[string]interface{}{
      "retryWrites": config.Get().Dto.Optional.RetryWrites,
      "w": config.Get().Dto.Optional.W,
      "replicaSet": config.Get().Dto.Optional.ReplicaSet,
      "tls": config.Get().Dto.Optional.Tls,
    }),
  })
}

//func getMongoDbConnectionParametersFinalMap() (res map[string]interface{}){
//  res = map[string]interface{}{}
//  maps.Copy(res, dtoMongoDbConnectionParameters[PRIMARY_DTO_PARAMETERS].(map[string]interface{}))
//  res["parametersString"] = getMongoDbConnectionOptionalParametersString()
//  return res
//}
//
//func getMongoDbConnectionOptionalParametersString() string{
//  return stringFormatter.FormatComplex("?retryWrites={retryWrites}&w={w}&replicaSet={replicaSet}&tls={tls}", dtoMongoDbConnectionParameters[OPTIONAL_DTO_PARAMETERS].(map[string]interface{}))
//}