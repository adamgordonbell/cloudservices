package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/adamgordonbell/cloudservices/activity-client/internal/client"
	api "github.com/adamgordonbell/cloudservices/activity-log/api/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const defaultURL = "localhost:8080"

func main() {
	add := flag.Bool("add", false, "Add activity")
	get := flag.Bool("get", false, "Get activity")
	list := flag.Bool("list", false, "List activities")

	flag.Parse()

	activitiesClient := client.NewActivities(defaultURL)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	switch {
	case *get:
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Fprintln(os.Stderr, "Invalid Offset: Not an integer")
			os.Exit(1)
		}
		a, err := activitiesClient.Retrieve(ctx, id)
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
		a := api.Activity{Time: timestamppb.New(time.Now()), Description: os.Args[2]}
		id, err := activitiesClient.Insert(ctx, &a)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error:", err.Error())
			os.Exit(1)
		}
		fmt.Printf("Added: %s as %d\n", asString(&a), id)
	case *list:
		as, err := activitiesClient.List(ctx, 0)
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

func asString(a *api.Activity) string {
	return fmt.Sprintf("ID:%d\t\"%s\"\t%d-%d-%d",
		a.Id, a.Description, a.Time.AsTime().Year(), a.Time.AsTime().Month(), a.Time.AsTime().Day())
}
