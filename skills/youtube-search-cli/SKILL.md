---
name: youtube-search-cli
description: Use the local YouTube Search CLI standalone binary to search YouTube videos and playlists, list videos from a playlist, or fetch video metadata through YouTube Data API v3 with JSON output. Trigger for API-based YouTube discovery and metadata tasks, especially from automation, CI, headless systems, or cloud networks where scraping tools such as yt-dlp are blocked or unreliable.
---

# YouTube Search CLI

Use the `youtube-search-cli` binary for structured YouTube search and metadata retrieval through YouTube Data API v3.

## Requirements

- Treat this as a standalone binary tool, not a library or source package.
- Expect `youtube-search-cli` or `youtube-search-cli.exe` to be available on `PATH`, unless the user provides an explicit executable path.
- Require `YOUTUBE_DATA_API_KEY` in the environment or in `~/.youtube.env`.
- Never ask the user to paste the API key into the prompt or include it in a command.
- Expect successful commands to emit indented JSON on stdout.
- Pass a playlist ID or video ID to ID-based commands, not a YouTube URL.

## Arguments

```bash
youtube-search-cli <SUBCOMMAND> [OPTIONS] [QUERY_OR_ID]
```

Place named options before the positional query or ID.

### Search

```bash
youtube-search-cli search [--filter video|playlist|mixed] [--max-results N] [--next-page-token TOKEN] "QUERY"
```

- `QUERY`: Required search text. Quote it when it contains spaces.
- `--filter`: Optional result type. Accept `video`, `playlist`, or `mixed`; default `mixed`.
- `--max-results`: Optional results per page from 1 to 50; default 10.
- `--next-page-token`: Optional token from a previous response for the next result page.

Examples:

```bash
youtube-search-cli search "Never Gonna Give You Up"
youtube-search-cli search --filter video --max-results 25 "TypeScript tutorials"
youtube-search-cli search --max-results 50 --next-page-token TOKEN "travel vlogs"
```

### Playlist

```bash
youtube-search-cli playlist [--max-results N] [--next-page-token TOKEN] PLAYLIST_ID
```

- `PLAYLIST_ID`: Required playlist ID.
- `--max-results`: Optional results per page from 1 to 50; default 10.
- `--next-page-token`: Optional token from a previous response for the next playlist page.

Example:

```bash
youtube-search-cli playlist --max-results 50 PLxxxxxxxxxxxxxxxx
```

### Video

```bash
youtube-search-cli video VIDEO_ID
```

- `VIDEO_ID`: Required YouTube video ID.
- This subcommand has no data options.

Example:

```bash
youtube-search-cli video dQw4w9WgXcQ
```

### Utility arguments

- `--help` or `-h`: Show top-level help.
- `--version` or `-v`: Print the binary version.
- `<subcommand> --help`: Show subcommand help.

## Cost recommendations

Assume the standard YouTube Data API allocation is 10,000 quota units per project per day unless the user states otherwise.

- Treat each `search` request as 100 units. This permits about 100 search requests if no other operations consume quota.
- Treat each `playlist` or `video` request as 1 unit. Prefer these commands when an ID is already known.
- Try `yt-dlp` or web search first for discovery when they work in the current environment. Use this binary when API reliability, structured results, cloud-network access, or YouTube Terms of Service compliance matters.
- Start discovery with `--filter mixed` so one 100-unit request can expose both videos and useful playlists. Run a separate playlist-only search only when its additional 100-unit cost is justified.
- Follow relevant playlists with the 1-unit `playlist` command before buying more 100-unit search pages.
- Choose `--max-results` deliberately. A larger page does not increase the quota cost of that request and can reduce follow-up calls, but it produces more JSON and may waste context on irrelevant results.
- Treat every `--next-page-token` use as a new API request: another 100 units for `search`, or another 1 unit for `playlist`.
- Save large JSON responses to a file and inspect only relevant fields with tools such as `jq` or `rg` instead of loading the entire response into context.

## Workflow

1. Determine whether the task needs discovery, playlist enumeration, or known-video metadata.
2. Confirm that the binary is available and that `YOUTUBE_DATA_API_KEY` is configured without exposing its value.
3. Select the lowest-cost subcommand that can answer the task.
4. Run the command with named options before the positional query or ID.
5. Parse the JSON response and report only the relevant results. Preserve `nextPageToken` when another page may be needed.
6. Stop paging once the task is answered; do not spend quota collecting unused results.

## Failure handling

- If `YOUTUBE_DATA_API_KEY` is missing, ask the user to set it in `~/.youtube.env` or in the command environment before retrying.
- If the binary is missing from `PATH`, ask for its executable path or direct the user to install the release binary.
- If given a YouTube URL for `video`, extract its video ID. If given a playlist URL for `playlist`, extract its `list` parameter.
- If the API reports quota exhaustion, stop retrying and explain which operation consumed quota. Suggest waiting for quota reset or using a non-API discovery method.
- If the API rejects `--max-results`, keep it between 1 and 50.
- If the command returns malformed or mixed output, preserve stderr separately and parse stdout as JSON.
