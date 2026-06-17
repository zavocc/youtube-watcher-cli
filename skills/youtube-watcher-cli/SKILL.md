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
youtube-watcher-cli --id [YOUTUBE_VIDEO_ID] [prompt]
```

- `--id [YOUTUBE_VIDEO_ID]`: Required. Pass only the video ID, such as `dQw4w9WgXcQ` from `https://www.youtube.com/watch?v=dQw4w9WgXcQ`.
- `prompt`: Required. Place the prompt after all named arguments. The binary joins all remaining positional arguments into the prompt.
- `--version`: Print the binary version and exit.
- `--help`: Show usage help.

Do not place named options after the prompt. Anything after the prompt is treated as prompt text.

## Workflow

1. Extract the video ID from the user's URL when needed.
2. Confirm `GEMINI_API_KEY` is available in the command environment.
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

## Failure Handling

- If `GEMINI_API_KEY` is missing, ask the user to set it before retrying.
- If the user provides a full YouTube URL, extract the `v` parameter or short URL ID instead of passing the full URL.
- If the binary is missing from `PATH`, ask the user for the executable path or to install the release binary.
- If the prompt is absent, ask for the question or extraction task to run against the video.
