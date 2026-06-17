package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

var (
	version = "dev-unspecified"
)

func showHelp() {
	helpString := "YouTube Video Watcher version " + version + ". " + "For people, for machines, and for agents." +
		"\n\nUsage: " + os.Args[0] + " --id [YOUTUBE_VIDEO_ID] 'prompt'\n" +
		" --id\tYouTube video ID [REQUIRED]\n" +
		" prompt\tPrompt to ask questions about the video [REQUIRED]" +
		"\n\n" +
		"Supplemental options:\n" +
		" --help\tShow help\n" +
		" --version\tPrint version"

	fmt.Println(helpString)
}

func printVersion() {
	fmt.Println(version)
	os.Exit(0)
}

func main() {
	// Check for args and parse it and use flag.Parse instead of os.Args to ensure positional accuracy
	flag.Usage = showHelp
	videoID := flag.String("id", "", "YouTube Video ID")
	invokeVersion := flag.Bool("version", false, "Print version")
	flag.Parse()

	// This will print version and exit, regardless of any conditions e.g. API key is set, id is set, prompt is specified
	if *invokeVersion {
		printVersion()
	}

	// Check if environment variable
	_, exists := os.LookupEnv("GEMINI_API_KEY")

	if !exists {
		fmt.Println("Gemini API key environment variable not set. Please set it using GEMINI_API_KEY variable.")
		os.Exit(1)
	}

	// get the leftover positional arguments as a prompt after parsing command line named arguments
	args := flag.Args()
	if len(args) == 0 {
		fmt.Println("--id and prompt as las")
		showHelp()
		os.Exit(1)
	}
	prompt := strings.Join(args, " ")

	// check if --id is set
	if *videoID == "" {
		fmt.Println("--id is required before the prompt")
		os.Exit(1)
	}

	textualResponse := inference(prompt, *videoID)
	fmt.Println(textualResponse)
	return
}
