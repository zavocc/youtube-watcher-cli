---
name: youtube-watcher-cli
description: Use the local YouTube Watcher CLI binary to ask Gemini grounded questions about YouTube videos, including summarization, transcription-style extraction, visual-content inspection, moderation preambles, code/text extraction, and multilingual video understanding. Trigger when a user provides a YouTube video URL or video ID and wants information from the video through the standalone binary.
---

# YouTube Watcher CLI

Use the `youtube-watcher-cli` binary when a task needs grounded understanding of a YouTube video through Gemini's native video input support.

## Requirements

- Treat this as a standalone binary tool, not a library.
- Expect `youtube-watcher-cli` or `youtube-watcher-cli.exe` to be available on `PATH`, unless the user provides an explicit executable path.
- Require `GEMINI_API_KEY` in the environment before invoking the binary.
- Use the YouTube video ID, not the full URL, with `--id`.

## Arguments

```bash
youtube-watcher-cli --id [YOUTUBE_VIDEO_ID] --model [MODEL_ID] [prompt]
```

- `--id [YOUTUBE_VIDEO_ID]`: Required. Pass only the video ID, such as `dQw4w9WgXcQ` from `https://www.youtube.com/watch?v=dQw4w9WgXcQ`.
- `--model [MODEL_ID]`: Optional. Specify a model to use to process the video, defaults to `gemini-2.5-flash` if not specified. See the supported models section below for choosing model.
- `--media-resolution [RESOLUTION]`: Optional. Specify the media resolution for the video, such as `low` or `high`. Defaults to `low` if not provided. Use `low` to prioritize speed and cost over extreme fine-detail, and `high` for better visual fidelity and fine details over cost of speed and budget.

`prompt` must be placed after all named arguments. The tool joins all remaining positional arguments into a prompt.

### Utility arguments

These will be prioritized if provided, overrides other parameters and only prints the help and version info, then quits the program.
- `--version`: Print the binary version and exit.
- `--help`: Show usage help.

Do not place named options after the prompt. Anything after the prompt is treated as prompt text.

## Supported models

- `gemini-2.5-flash` - Best balance for speed, cost, and intelligence. It is the default with `1024` thinking budgets.
- `gemini-3-flash-preview` - Inherits it's larger Pro-grade intelligence at fraction of the cost with improved vision understanding, but it is priced higher than 2.5 Flash. To ensure it meets cost and latency, this model has minimal near-zero reasoning effort set. Best at vision, world factual knowledge, and multilingual understanding.
- `gemini-3.1-flash-lite` - Google's latest Flash-lite line of model that outperforms 2.5 Flash model and is cheaper than both 2.5 Flash and 3 Flash Preview. Useful for quick video overviews and long videos for time and budget constrained scenarios.

## Workflow

1. Extract the video ID from the user's URL when needed.
2. Confirm if non-empty `~/.youtube.env` file or `GEMINI_API_KEY` is available in the command environment.
3. Run the binary with `--id` before the prompt.
4. Read the answer from stdout and report the relevant result to the user.

## Prompting

Write direct prompts for the video task. Good prompt shapes include:

```bash
youtube-watcher-cli --id dQw4w9WgXcQ summarize the video with key timestamps
youtube-watcher-cli --id dQw4w9WgXcQ extract any visible code or terminal commands
youtube-watcher-cli --id dQw4w9WgXcQ describe visual actions and spoken content in detail
youtube-watcher-cli --id dQw4w9WgXcQ classify whether this video is safe to show before playback
```

Quote the prompt if the shell or command runner requires it, but keep it as the final positional argument.

## Model capabilities

### What it can do:

See visual frames and hear audio of the video, text prompt, and small system instruction always appended to define it's role. It can also understand timestamps of the video associated with the frame but it can be inaccurate (use with caution, must be treated as approximate indicators and not exact).

### What it can't do

It cannot see YouTube video ID, title, or other metadata. Use `yt-dlp` or YouTube Data API to get that information separately if needed.

It may also struggle with very long videos due to context limit. Such as videos exceeding more than 1.5 hours with audio.

## Failure Handling

- If `GEMINI_API_KEY` is missing, ask the user to set `GEMINI_API_KEY` in `~/.youtube.env` or directly in the terminal before retrying. The binary will search for this file and read it without the need of sourcing, otherwise it will fall back based on existing environment variables set without it.
- If the user provides a full YouTube URL, extract the `v` parameter or short URL ID instead of passing the full URL.
- If the binary is missing from `PATH`, ask the user for the executable path or to install the release binary.
- If the prompt is absent, ask for the question or extraction task to run against the video.
