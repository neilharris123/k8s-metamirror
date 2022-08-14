package config

import (
  "log"
  "github.com/kelseyhightower/envconfig"
)

var Metadata struct {
    annotation  string `envconfig:"annotation"`
    label       string `envconfig:"label"`
}

func init() {
  err := envconfig.Process("", &Metadata)
  if err != nil {
    log.Fatal(err)
  }
}
