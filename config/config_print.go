package config

import (
  "fmt"
  "src/utils"
)

// TODO: hide secrets
func Print(){
  fmt.Println("")
  fmt.Println("###### CONFIGURATION DETAILS ######\n")
  fmt.Println(string(utils.MustMarshalToYaml(Get())))
  fmt.Println("###################################\n\n")
}