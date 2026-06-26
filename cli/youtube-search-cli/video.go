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

func showHelpVideo() {
	helpString := "Extract video metadata" +
		"\n\nUsage: " + os.Args[0] + " video [video ID]\n" +
		"\nVideo options:\n" +
		" id                	Video ID [REQUIRED]" +
		"\n\n" +
		"Supplemental options:\n" +
		" --help     " + helpShowHelpString

	fmt.Println(helpString)
}

func runVideoQuery(ctx context.Context, args []string) {
	flagSet := flag.NewFlagSet("video", flag.ExitOnError)
	flagSet.Usage = showHelpVideo

	showHelp := flagSet.Bool("help", false, helpShowHelpString)

	// get the leftover positional arguments as a prompt after parsing command line named arguments
	flagSet.Parse(args)
	idArgs := flagSet.Args()

	// Show help and exit regardless of  other arguments if --help is provided or if no video ID is supplied
	if len(idArgs) == 0 || *showHelp {
		showHelpVideo()
		os.Exit(1)
	}

	videoID := strings.Join(idArgs, " ")
	service, err := newYouTubeService(ctx)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	videoResponse, err := dataapi.Video(service, videoID)
	if err != nil {
		fmt.Fprintln(os.Stderr, "An error has occurred while fetching video:", err)
		os.Exit(1)
	}

	// Serialize and print the video results as JSON
	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(videoResponse); err != nil {
		fmt.Fprintln(os.Stderr, "An error has occurred while serializing the video results:", err)
		os.Exit(1)
	}

	os.Exit(0)
}
