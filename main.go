package main

import (
  "log"
  "src/api"
  "src/config"
)

//TODO: implement unit tests
func main() {
  config.Print()
  err := api.NewApplicationRouter().Run(config.Get().Api.Socket)
  if err != nil {
    log.Printf("Error opening API socket: %v", err)
  }
}