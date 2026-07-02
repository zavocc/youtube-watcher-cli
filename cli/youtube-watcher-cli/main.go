package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/zavocc/youtube-watcher-cli/internal/gemini"
	"github.com/zavocc/youtube-watcher-cli/internal/shared"
)

func showHelp() {
	helpString := "YouTube Video Watcher version " + shared.Version + ". " + "For people, for machines, and for agents." +
		"\n\nUsage: " + os.Args[0] + " --video [YOUTUBE_VIDEO_URL_OR_ID] 'prompt'\n" +
		" --video             YouTube video URL or ID [REQUIRED]\n" +
		" --model             Model to use for inference, defaults to " + gemini.DefaultModel + "\n" +
		" --media-resolution  Media resolution for the video. Possible values are only low, high. If not set, it will default for low resolution\n" +
		" prompt              Prompt to ask questions about the video [REQUIRED]" +
		"\n\n" +
		"Supplemental options:\n" +
		" --help     Show help\n" +
		" --version  Print version" +
		"\n\n" +
		"Supported models:\n" +
		" - gemini-2.5-flash (with 1024 thinking budget)\n" +
		" - gemini-3-flash-preview (with minimal thinking level)\n" +
		" - gemini-3.1-flash-lite (with low thinking level)"

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

	// Check for args and parse it and use flag.Parse instead of os.Args to ensure positional accuracy
	flag.Usage = showHelp
	videoID := flag.String("video", "", "YouTube Video URL or ID")
	selectedModel := flag.String("model", gemini.DefaultModel, "Model to use")
	mediaRes := flag.String("media-resolution", "low", "Media resolution for the video (low, medium, high)")
	invokeVersion := flag.Bool("version", false, "Print version")
	flag.Parse()

	// This will print version and exit, regardless of any conditions e.g. API key is set, id is set, prompt is specified
	if *invokeVersion {
		printVersion()
	}

	// get the leftover positional arguments as a prompt after parsing command line named arguments
	args := flag.Args()
	if len(args) == 0 {
		fmt.Fprintln(os.Stderr, "A prompt is required along with --video")
		showHelp()
		os.Exit(1)
	}
	prompt := strings.Join(args, " ")

	// check if --id is set
	if *videoID == "" {
		fmt.Fprintln(os.Stderr, "--video is required before the prompt")
		os.Exit(1)
	}

	//  dereference videoID so it can be passed as a string normally
	ctx := context.Background()
	result, err := gemini.GApiClient(ctx, prompt, *videoID, *selectedModel, *mediaRes)
	if err != nil {
		fmt.Fprintln(os.Stderr, "An error has occurred - ", err)
		os.Exit(1)
	}

	fmt.Println(result)
	os.Exit(0)
}
