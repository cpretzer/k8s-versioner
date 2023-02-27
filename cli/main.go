package main

import (
	"flag"
	log "github.com/sirupsen/logrus"
	"os"
	yaml "sigs.k8s.io/yaml"
)

type K8sResource struct {
	kind       string
	apiVersion string
}

func main() {
	flag.Parse()
	k8sResourceYaml, err := os.ReadFile(flag.Args()[0])

	if err != nil {
		log.Error("frowny face\n", err)
	}

	resourceType, err := getResourceType(string(k8sResourceYaml))

	log.Infof("Ya did great \n %s", *resourceType)
}

func getResourceType(k8sResourceYaml string) (*string, error) {

	log.Infof("marshaling string \n %s", k8sResourceYaml)

	var k8sResource K8sResource

	err := yaml.Unmarshal([]byte(k8sResourceYaml), k8sResource)

	if err != nil {
		log.Errorf("Error unmarshaling %s", err)
		return nil, err
	}

	log.Infof("Your resource is %v", k8sResource)

	return &k8sResource.kind, nil
}
