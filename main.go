package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"temporal-example/helloworkflow"

	"go.temporal.io/sdk/client"
)

func main() {
	http.HandleFunc("/", helloTemporal)
	http.ListenAndServe(":8083", nil)
}

func helloTemporal(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "Hello World!")
	starter()
}

func starter() {
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create client ", err)
	}
	defer c.Close()

	workflowOptions := client.StartWorkflowOptions{
		ID:        "workflow-options-id-starter-main.go",
		TaskQueue: "hello-world-queue-1",
	}

	we, err := c.ExecuteWorkflow(context.Background(), workflowOptions, helloworkflow.Workflow, "string-name-to-workflow")
	if err != nil {
		log.Fatalln("Unable to execute workflow.", err)
	}
	var result string
	err = we.Get(context.Background(), &result)
	if err != nil {
		log.Fatalln("Unable to get workflow result.")
	}
	log.Println("Workflow result:", result)

}
