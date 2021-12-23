package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/adamgordonbell/cloudservices/activity-client/internal/client"
	"github.com/adamgordonbell/cloudservices/activity-log/internal/server"
)

const defaultURL = "http://localhost:8080/"

func main() {
	add := flag.Bool("add", false, "Add activity")
	get := flag.Bool("get", false, "Get activity")

	flag.Parse()

	activitiesClient := &client.Activities{URL: defaultURL}

	switch {
	case *get:
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Fprintln(os.Stderr, "Invalid Offset: Not an integer")
			os.Exit(1)
		}
		a, err := activitiesClient.Retrieve(id)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error:", err.Error())
			os.Exit(1)
		}
		fmt.Println(a.String())
	case *add:
		if len(os.Args) != 3 {
			fmt.Fprintln(os.Stderr, `Usage: --add "message"`)
			os.Exit(1)
		}
		a := server.Activity{Time: time.Now(), Description: os.Args[2]}
		id, err := activitiesClient.Insert(a)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error:", err.Error())
			os.Exit(1)
		}
		fmt.Printf("Added: %s as %d\n", a.String(), id)
	default:
		flag.Usage()
		os.Exit(1)
	}
}
