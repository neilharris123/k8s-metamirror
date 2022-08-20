package config

import (
  "log"
  "github.com/kelseyhightower/envconfig"
)

var Metadata struct {
    Annotation  string `envconfig:"annotation"`
    Label       string `envconfig:"label"`
}

func init() {
  err := envconfig.Process("mm", &Metadata)
  if err != nil {
    log.Fatal(err)
  }
}
