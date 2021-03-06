package main

import (
	"context"
	"log"

	"github.com/itdevsamurai/gke/simplejob/job/handler"
)

func main() {
	log.Println("Job Started.")

	ctx := context.Background()
	handler.SimpleJobHandler{}.Run(ctx)

	log.Println("Job Finished.")
}
