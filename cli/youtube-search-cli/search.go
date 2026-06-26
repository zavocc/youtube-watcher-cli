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

func showHelpSearch() {
	helpString := "Search YouTube videos by query." +
		"\n\nUsage: " + os.Args[0] + " search [options] [search query]\n" +
		"\nSearch options:\n" +
		" --filter              " + helpSearchFilterString + "\n" +
		" --max-results         " + helpMaxResultsString + "\n" +
		" --next-page-token     " + helpNextPageTokenString + "\n" +
		" query                	Search query [REQUIRED]" +
		"\n\n" +
		"Supplemental options:\n" +
		" --help     " + helpShowHelpString

	fmt.Println(helpString)
}

func runSearch(ctx context.Context, args []string) {
	flagSet := flag.NewFlagSet("search", flag.ExitOnError)
	flagSet.Usage = showHelpSearch

	// args
	filter := flagSet.String("filter", "mixed", helpSearchFilterString)
	maxResults := flagSet.Int64("max-results", 10, helpMaxResultsString)
	nextPageToken := flagSet.String("next-page-token", "", helpNextPageTokenString)
	showHelp := flagSet.Bool("help", false, helpShowHelpString)

	// get the leftover positional arguments as a prompt after parsing command line named arguments
	flagSet.Parse(args)
	queryArgs := flagSet.Args()

	// Show help and exit regardless of  other arguments if --help is provided or if no search query is supplied
	if len(queryArgs) == 0 || *showHelp {
		showHelpSearch()
		os.Exit(1)
	}
	query := strings.Join(queryArgs, " ")

	service, err := newYouTubeService(ctx)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	searchResponse, err := dataapi.Search(service, query, *filter, *maxResults, *nextPageToken)
	if err != nil {
		fmt.Fprintln(os.Stderr, "An error has occurred while searching YouTube:", err)
		os.Exit(1)
	}

	// Serialize and print the search results as JSON
	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(searchResponse); err != nil {
		fmt.Fprintln(os.Stderr, "An error has occurred while serializing the search results:", err)
		os.Exit(1)
	}

	os.Exit(0)
}
