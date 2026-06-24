package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/zavocc/youtube-watcher-cli/internal/dataapi"
	"github.com/zavocc/youtube-watcher-cli/internal/shared"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

func showHelp() {
	helpString := "YouTube Video Search version " + shared.Version + "." +
		"\n\nUsage: " + os.Args[0] + " 'search_query'\n" +
		" --max-results         Maximum number of results to return [DEFAULT: 10]\n" +
		" --next-page-token     Token for the next page of results, can be obtained from previous search results\n" +
		" query                	Search query [REQUIRED]" +
		"\n\n" +
		"Supplemental options:\n" +
		" --help     Show help\n" +
		" --version  Print version"

	fmt.Println(helpString)
}

func printVersion() {
	fmt.Println(shared.Version)
	os.Exit(0)
}

func main() {
	if err := shared.LoadEnvironment(); err != nil {
		fmt.Fprintln(os.Stderr, "An error has occurred while loading the environment:", err)
		os.Exit(1)
	}

	// Check if environment variable exists
	apikey, exists := os.LookupEnv("YOUTUBE_DATA_API_KEY")

	if !exists {
		fmt.Fprintln(os.Stderr, "YouTube Data API key environment variable not set. Please set it using YOUTUBE_DATA_API_KEY variable in `~/.youtube.env` or directly in the terminal.")
		os.Exit(1)
	}

	// Authenticate
	ctx := context.Background()
	service, err := youtube.NewService(ctx, option.WithAPIKey(apikey))
	if err != nil {
		fmt.Fprintln(os.Stderr, "An error has occurred while creating the YouTube service:", err)
		os.Exit(1)
	}

	// Check for args and parse it and use flag.Parse instead of os.Args to ensure positional accuracy
	flag.Usage = showHelp
	maxResults := flag.String("max-results", "10", "Maximum number of results to return")
	nextPageToken := flag.String("next-page-token", "", "Token for the next page of results")
	invokeVersion := flag.Bool("version", false, "Print version")
	flag.Parse()

	// This will print version and exit, regardless of any conditions e.g. API key is set, id is set, prompt is specified
	if *invokeVersion {
		printVersion()
	}

	// get the leftover positional arguments as a prompt after parsing command line named arguments
	args := flag.Args()
	if len(args) == 0 {
		fmt.Fprintln(os.Stderr, "A query is required, please provide a search query.")
		showHelp()
		os.Exit(1)
	}
	query := strings.Join(args, " ")
	maxResultsValue, err := strconv.ParseInt(*maxResults, 10, 64)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Invalid value for --max-results:", err)
		os.Exit(1)
	}

	searchResponse, err := dataapi.Search(service, query, maxResultsValue, *nextPageToken)
	if err != nil {
		fmt.Fprintln(os.Stderr, "An error has occurred while searching for videos:", err)
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
