package main

import (
	"fmt"
	"log"

	"k8s.io/apimachinery/pkg/labels"
)

func main() {
	// c, err := client.New(config.GetConfigOrDie(), client.Options{})
	// if err != nil {
	// 	log.Fatalf("Failed to create client")
	// }
	pvcVersion := "blah"
	lable := GetLabel("release_version", pvcVersion)
	fmt.Println(lable)

}
func GetLabel(labelkey string, labelvalue string) labels.Selector {

	label, err := labels.Parse(fmt.Sprintf("%v=%v", labelkey, labelvalue))
	if err != nil {
		log.Fatalf("error getting label %v : %v", labelkey, labelvalue)
	}

	return label
}
