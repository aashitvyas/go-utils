package main

import (
	"log"
	"os/exec"

	ps "github.com/mitchellh/go-ps"
)

func main() {
	StopAmazonEcsAgent()
	StartAmazonEcsAgent()
}

func StopAmazonEcsAgent() {
	p, err := ps.Processes()
	if err != nil {
		log.Println(err)
	}
	for _, p1 := range p {
		if p1.Executable() == "amazon-ecs-agent.exe" {
			log.Println("Amazon ECS Agent process has been found")
			log.Println("Stopping Amazon ECS Agent process")
			cmd := exec.Command("net", "stop", "amazonECS")
			err := cmd.Run()
			if err != nil {
				log.Println(err)
			}
			log.Println("Amazon ECS Agent has successfully stopped")
			break

		}
	}
}

// Start Amazon ECS Agent process to start accpeting new LR once SA update is done
func StartAmazonEcsAgent() {
	p, err := ps.Processes()
	if err != nil {
		log.Println(err)
	}
	for _, p1 := range p {
		if p1.Executable() != "amazon-ecs-agent.exe" {
			log.Println("Amazon ECS Agent process missing")
			log.Println("Starting Amazon ECS Agent process")
			cmd := exec.Command("net", "start", "amazonECS")
			err := cmd.Run()
			if err != nil {
				log.Println(err)
			}
			log.Println("Amazon ECS Agent has succesfully started ")
			break
		}
	}
}
