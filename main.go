package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/zavocc/youtube-watcher-cli/internal/gemini"
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
	// Check for .youtube.env file in home directory and load it if it exists
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Fprintln(os.Stderr, "An error has occured while getting user home directory:", err)
		os.Exit(1)
	}

	// Load .youtube.env file if it exists otherwise we  ignore and still proceed to check existing environment variables
	envFilePath := homeDir + "/.youtube.env"
	_ = godotenv.Load(envFilePath)

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
		fmt.Fprintln(os.Stderr, "Gemini API key environment variable not set. Please set it using GEMINI_API_KEY variable in `~/.youtube.env` or directly in the terminal.")
		os.Exit(1)
	}

	// get the leftover positional arguments as a prompt after parsing command line named arguments
	args := flag.Args()
	if len(args) == 0 {
		fmt.Fprintln(os.Stderr, "A prompt is required along with --id")
		showHelp()
		os.Exit(1)
	}
	prompt := strings.Join(args, " ")

	// check if --id is set
	if *videoID == "" {
		fmt.Fprintln(os.Stderr, "--id is required before the prompt")
		os.Exit(1)
	}

	//  dereference videoID so it can be passed as a string normally
	fmt.Println(gemini.GApiClient(prompt, *videoID))
	os.Exit(0)
}
