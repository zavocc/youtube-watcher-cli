package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/zavocc/youtube-watcher-cli/internal/dataapi"
)

func showHelpChan() {
	helpString := "List videos from a specified channel name, ID, or username handle" +
		"\n\nUsage: " + os.Args[0] + " channel [options] [channel ID]\n" +
		"\nChannel options:\n" +
		" --query-type          " + helpChanHelpQueryTypeString + "\n" +
		" --max-results         " + helpMaxResultsString + "\n" +
		" --next-page-token     " + helpNextPageTokenString + "\n" +
		" query                	Channel ID/username/handle [REQUIRED]" +
		"\n\n" +
		"Supplemental options:\n" +
		" --help     Show this subcommand help"

	fmt.Println(helpString)
}

func runChanQuery(ctx context.Context, args []string) {
	flagSet := flag.NewFlagSet("channel", flag.ExitOnError)
	flagSet.Usage = showHelpChan

	// args
	chanQueryType := flagSet.String("query-type", "handle", helpChanHelpQueryTypeString)
	maxResults := flagSet.Int64("max-results", 10, helpMaxResultsString)
	nextPageToken := flagSet.String("next-page-token", "", helpNextPageTokenString)
	showHelp := flagSet.Bool("help", false, helpShowHelpString)

	// get the leftover positional arguments as a prompt after parsing command line named arguments
	flagSet.Parse(args)
	idArgs := flagSet.Args()

	// Show help and exit regardless of  other arguments if --help is provided or if no channel ID is supplied
	if len(idArgs) == 0 || *showHelp {
		showHelpChan()
		os.Exit(1)
	}
	channelID := strings.Join(idArgs, " ")

	service, err := newYouTubeService(ctx)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	channelResponse, err := dataapi.Channel(service, channelID, *chanQueryType, *maxResults, *nextPageToken)
	if err != nil {
		fmt.Fprintln(os.Stderr, "An error has occurred while fetching channel:", err)
		os.Exit(1)
	}

	// Serialize and print the channel results as JSON
	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(channelResponse); err != nil {
		fmt.Fprintln(os.Stderr, "An error has occurred while serializing the channel results:", err)
		os.Exit(1)
	}

	os.Exit(0)
}
