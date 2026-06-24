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

func showHelpPlaylist() {
	helpString := "List videos from playlist" +
		"\n\nUsage: " + os.Args[0] + " playlist [options] [playlist ID]\n" +
		"\nPlaylist options:\n" +
		" --max-results         Maximum number of results to return [DEFAULT: 10]\n" +
		" --next-page-token     Token for the next page of results, can be obtained from previous playlist results\n" +
		" id                	Playlist ID [REQUIRED]" +
		"\n\n" +
		"Supplemental options:\n" +
		" --help     Show this subcommand help"

	fmt.Println(helpString)
}

func runPlaylistQuery(ctx context.Context, args []string) {
	flagSet := flag.NewFlagSet("playlist", flag.ExitOnError)
	flagSet.Usage = showHelpPlaylist

	// args
	maxResults := flagSet.Int64("max-results", 10, "Maximum number of results to return")
	nextPageToken := flagSet.String("next-page-token", "", "Token for the next page of results")
	showHelp := flagSet.Bool("help", false, "Show this subcommand help")

	// get the leftover positional arguments as a prompt after parsing command line named arguments
	flagSet.Parse(args)
	idArgs := flagSet.Args()

	// Show help and exit regardless of  other arguments if --help is provided or if no playlist ID is supplied
	if len(idArgs) == 0 || *showHelp {
		showHelpPlaylist()
		os.Exit(1)
	}
	playlistID := strings.Join(idArgs, " ")

	service, err := newYouTubeService(ctx)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	playlistResponse, err := dataapi.Playlist(service, playlistID, *maxResults, *nextPageToken)
	if err != nil {
		fmt.Fprintln(os.Stderr, "An error has occurred while fetching playlist:", err)
		os.Exit(1)
	}

	// Serialize and print the playlist results as JSON
	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(playlistResponse); err != nil {
		fmt.Fprintln(os.Stderr, "An error has occurred while serializing the playlist results:", err)
		os.Exit(1)
	}

	os.Exit(0)
}
