# YouTube CLI
A simple and easy to use CLI program that allows Gemini to "watch" YouTube videos and to get nuanced information about the YouTube video using Gemini API with custom instructions.

Designed for people, machines, and agents.

Written in golang, allowing for cross-platform support.

If you want to integrate YouTube Data API for search and video metadata extraction in a similar CLI fashion, check out [README for `youtube-search-cli`](./README_yt_search.md). It includes the same installation steps and skills.

## Whom is this for
- For people - You can use this to ask questions, extract content,  and generate transcriptions about the YouTube video. It also saves bandwidth by simply gives you textual representation of the video.
- For machines - You can use this to build on top of this tool, such as GUI frontends, CI/CD, automation, and scripts.
- For agents - Coding and general purpose agents with `bash` tool can use this to understand and ask questions about the YouTube video content without scraping transcripts.  A minimal `SKILL.md` copy is provided for agents.

## Use-cases
- Accurate and multilingual subtitle generation including the visual actions of the video content
- Content extraction such as code or text from the video
- Moderation and safety classification, allowing to provide preambles of YouTube videos before the user can watch it  
- Summarization of the video
- Searching for timestamp of a particular subject, text, or excerpts.

## Implementation Status
It is currently minimalistic program that takes prompt and video ID as input, and text output grounded from YouTube videos.

- [X] Basic functionality
- [ ] Full Linux support and `Makefile` builds
- [ ] Full scripting support such as pipelines (piping commands as prompt), file descriptors like stderr for errors
- [x] Gemini model picker that supports video input (For now it supports 2.5 Flash, 3 Flash, and 3.1 Flash-Lite)
- [ ] Gemini Enterprise Agent Platform (aka Vertex AI) endpoint support and ADC auth
- [ ] Flex and Priority inference for budget tuning
- [ ] Nano Banana 2 based frame extraction
- [ ] Optional Gemma 4-based guardrails for both input and output
- [x] Video processing controls such as media resolution and start/end offset parameter*

*As of 06/23/26, only media resolution parameter is implemented.

# Usage and installation
Download the binary through the [releases](https://github.com/zavocc/youtube-watcher-cli/releases) page and must have the filename `youtube-watcher-cli`.

After the binary is placed onto the `PATH` environment variable, you must then set `GEMINI_API_KEY` environment variable or this program will not work.

You can either set `GEMINI_API_KEY` in `~/.youtube.env` or directly setting into the terminal. For coding agents, it's recommended to set the former so you don't have to directly invoke the API key to the prompt.

Use: 
```
.\youtube-watcher-cli  --video [YOUTUBE_VIDEO_ID_OR_URL] Write your prompt here
```
Note that the prompt must be at the end of the argument, either quoted or unquoted.

## Installing the agent skill
The `youtube-watcher-cli` agent skill allows the agents of your choice effectively use this tool and understand YouTube videos. To install, use the `npx skills` command:

```shell
npx skills add zavocc/youtube-watcher-cli --global     
```

You will be asked where to install the skill based on the agent tools you use.

This will install both `youtube-watcher-cli` and `youtube-search-cli` skills.

## Parameters
- `--video [YOUTUBE_VIDEO_ID_OR_URL]` - Either the public YouTube video URL or ID of the video itself.
- `--model [MODEL_ID]` - An optional parameter of model ID to set to analyze videos, please see [models list](./internal/config/models.go) for list of supported model and defaults.
- `--media-resolution [low|high]` - An optional parameter to set the visual quality when processing video. Use low where speed and cost matters over fine-detail, high if fine detail matters.

`prompt` is placed at the end after named arguments, any arguments placed after `prompt` will be treated as part of the prompt as is. So passing `--model gemini-3-flash-preview` after `prompt` would be treated as prompt.

# Building
You will need the latest version of Golang. I used Go version 1.26.4, `go` is set to PATH and `GOROOT` set to your environment variables pointing to Golang SDK directory. 

Go to this project's root and run
```bash
mkdir outputs
go mod tidy
go build -o .\outputs\youtube-watcher-cli.exe
```

# FAQ
### What is the default model used?
Gemini 2.5 Flash with 1024 thinking budgets.

### Are other non Gemini models will be supported in the future for analyzing YouTube video as a subagent?
No, there are no plans for it. Video understanding capabilities with YouTube videos is only exclusive with Gemini models. However, this utility is designed for other agent harnesses with non-Google models to ask questions about the video.

It is possible for other multimodal models, but it involves more time-consuming process such as downloading the video through yt-dlp, sample frames using FFmpeg and get audio content using speech models, and reason over it. \
But it's not worth a complexity for now, speed and nuances can be compromised, but can be considered.

### Is it free to use?
Yes, as long you get [Gemini API key](https://aistudio.google.com/api-keys) and set `GEMINI_API_KEY` environment variable

Refer to https://ai.google.dev/gemini-api/docs/rate-limits under "Free tier" for more information.

You can also use paid API keys if you wish and benefit from higher rate limits.

### What's the difference between using the Gemini API through REST/SDK vs using this standalone utility to add YouTube multimodal capabilities to my app.
While you can use the [Gemini API](https://ai.google.dev/gemini-api/docs/video-understanding#youtube) directly integrate and to pass YouTube URLs when calling Gemini API to your app

This CLI program is designed to bring YouTube video understanding capabilities as an executable subagent so existing agents like Codex and non-Gemini models can use this to understand the contents of the video without writing additional code which includes:
- Using scrapers or tools like yt-dlp to get the subtitles or calling Gemini API on the fly
- Needing to manually integrate Gemini API as a dependency to the agentic harness
- Having to maintain Gemini API as a dependency to your code in order to support this capability

### Does it rely on subtitles to understand video content?
No, it uses native multimodal video understanding capabilities to get nuances not just speech but also the visual content of the video. 

This what sets apart from other tools like [yt-dlp](https://github.com/yt-dlp/yt-dlp) and [youtube-transcript-api](https://github.com/jdepoix/youtube-transcript-api)

This is similar to Web Fetch tool but for videos, having to understand the whole video nuances instead of relying on transcribed speech.