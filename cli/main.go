package main

import (
	"flag"
	log "github.com/sirupsen/logrus"
	"os"
)

func main() {
	flag.Parse()
	k8sResourceYaml, err := os.ReadFile(flag.Args()[0])

	if err != nil {
		log.Error("frowny face\n", err)
	}

	log.Infof("Ya did great \n %s", k8sResourceYaml)
}
