package main

import (
  "flag"
  "log"
)

const (
  Version = "0.0.1"
)

var (
  verbose = flag.Bool("verbose", false, "Debugging mode")
)

func main(){
  flag.Parse()

  log.Println("proxy-ng", Version)
  log.Println("Usage: proxy-ng")
  flag.PrintDefaults()
}
