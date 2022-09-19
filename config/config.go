package config

import (
  "log"
  "github.com/kelseyhightower/envconfig"
)

var Metadata struct {
    Annotations  string `envconfig:"annotations"`
    Labels       string `envconfig:"labels"`
}

func init() {
  err := envconfig.Process("mm", &Metadata)
  if err != nil {
    log.Fatal(err)
  }
}
