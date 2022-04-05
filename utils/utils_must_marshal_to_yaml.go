package utils

import (
  "gopkg.in/yaml.v2"
  "log"
)

func MustMarshalToYaml(object interface{}) []byte{
  res, err := yaml.Marshal(object)
  logErrorOnMustMarshalToYaml(res, err)

  return res
}

func logErrorOnMustMarshalToYaml(object interface{}, err error){
  if err != nil {
    log.Printf("Error must marshalling to YAML {%v}: %v", object, err)
  }
}