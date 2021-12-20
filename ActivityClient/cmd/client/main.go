package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/adamgordonbell/cloudservices/activityclient/internal/client"
)

func main() {
	add := flag.Bool("add", false, "Add activity")
	get := flag.Bool("get", false, "Get activity")

	flag.Parse()

	activitiesClient := &client.Activities{}

	switch {
	case *get:
		id, err := strconv.Atoi(os.Args[1])
		if err != nil {
			println("Invalid Offset: Not an integer")
			os.Exit(1)
		}
		fmt.Printf("Getting %d", id)
		a, err := activitiesClient.Retrieve(id)
		if err != nil {
			println("Error:", err)
			os.Exit(1)
		}
		fmt.Printf("%+v\n", a)
	case *add:
		if len(os.Args) != 3 {
			println(`Usage: --add "message"`)
			os.Exit(1)
		}
		a := client.Activity{Time: time.Now(), Description: os.Args[2]}
		id, err := activitiesClient.Insert(a)
		if err != nil {
			println("Error:", err)
			os.Exit(1)
		}
		fmt.Printf("Added: %+v as %d\n", a, id)
	default:
		flag.Usage()
		os.Exit(1)
	}
}
