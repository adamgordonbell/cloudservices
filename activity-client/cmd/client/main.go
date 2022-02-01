package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/adamgordonbell/cloudservices/activity-client/internal/client"
	api "github.com/adamgordonbell/cloudservices/activity-log/api/v1"
)

const defaultURL = "http://localhost:8080/"

func main() {
	add := flag.Bool("add", false, "Add activity")
	get := flag.Bool("get", false, "Get activity")
	list := flag.Bool("list", false, "List activities")

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
		fmt.Println(asString(a))
	case *add:
		if len(os.Args) != 3 {
			fmt.Fprintln(os.Stderr, `Usage: --add "message"`)
			os.Exit(1)
		}
		a := api.Activity{Time: time.Now(), Description: os.Args[2]}
		id, err := activitiesClient.Insert(a)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error:", err.Error())
			os.Exit(1)
		}
		fmt.Printf("Added: %s as %d\n", asString(a), id)
	case *list:
		as, err := activitiesClient.List(0)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error:", err.Error())
			os.Exit(1)
		}
		var output string
		for _, a := range as {
			output += asString(a) + "\n"
		}
		fmt.Println(output)

	default:
		flag.Usage()
		os.Exit(1)
	}
}

func asString(a api.Activity) string {
	return fmt.Sprintf("ID:%d\t\"%s\"\t%d-%d-%d",
		a.ID, a.Description, a.Time.Year(), a.Time.Month(), a.Time.Day())
}
