package main

import (
	"context"
	"fmt"
	"os"

	"github.com/zavocc/youtube-watcher-cli/internal/shared"
)

func showHelp() {
	helpString := "YouTube Video Search version " + shared.Version + "." +
		"\n\nUsage: " + os.Args[0] + " [subcommands]\n" +
		"\nSubcommands:\n" +
		" search                Search YouTube for videos and playlists\n" +
		" playlist              List videos in a playlist from its playlist ID\n" +
		" video                 Get metadata for a specific video by its ID\n" +
		" channel               List videos from a channel by its ID, username, or handle\n" +
		"\nTo get help for a specific subcommand, run:\n" +
		" " + os.Args[0] + " <subcommand> --help" +
		"\n\n" +
		"Supplemental options:\n" +
		" --help     Show this help\n" +
		" --version  Print version"

	fmt.Println(helpString)
}

func printVersion() {
	fmt.Println(shared.Version)
	os.Exit(0)
}

func main() {
	// check if osArgs is only 1
	if len(os.Args) < 2 {
		showHelp()
		os.Exit(1)
	}

	// share context
	ctx := context.Background()

	switch os.Args[1] {
	case "--help", "-h", "help":
		showHelp()
		os.Exit(0)
	case "--version", "-v", "version":
		printVersion()
		os.Exit(0)
	case "search":
		runSearch(ctx, os.Args[2:])
	case "playlist":
		runPlaylistQuery(ctx, os.Args[2:])
	case "video":
		runVideoQuery(ctx, os.Args[2:])
	case "channel":
		runChanQuery(ctx, os.Args[2:])
	default:
		fmt.Fprintln(os.Stderr, "Unknown subcommand:", os.Args[1])
		showHelp()
		os.Exit(1)
	}
}
