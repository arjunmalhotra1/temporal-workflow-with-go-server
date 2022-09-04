package main

import (
	"log"
	"temporal-example/helloworkflow"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func main() {
	c, err := client.NewClient(client.Options{})
	if err != nil {
		log.Fatal("Unable to make cleint ", err)
	}

	defer c.Close()

	w := worker.New(c, "hello-world-queue-1", worker.Options{})
	w.RegisterWorkflow(helloworkflow.Workflow)
	w.RegisterActivity(helloworkflow.Activity)
	w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln(" Unable to start wotkflow ")
	}
}
