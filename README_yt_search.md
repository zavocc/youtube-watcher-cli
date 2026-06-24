# YouTube Search CLI
Search YouTube videos using CLI! Built for machines, terminal and cloud agents, using YouTube Data API v3 and complements with YouTube Watcher CLI.

This command line program let's you search YouTube videos and output raw JSON-formatted results from YouTube Data API through command line.

## Comparison to yt-dlp
While `yt-dlp` is also capable of searching videos through web scraping, it is not suitable for agents and automation/CI scripts running on the cloud network, YouTube blocks known datacenter and cloud IP ranges from extracting and searching data.

This command line tool ensures the 

## Whom is this for
For automation, CI, scripts, and terminal agents, especially those that's running on headless or cloud network, while ensuring compliance with the YouTube Terms of Service in order to get programmatic access.

# Limitations
## Missing functionality
the CLI program is already capable at it's core for searching, listing videos from playlist, and obtaining video metadata, there are some features that's not currently implemented:

- [ ] Transcription extraction.
- [ ] Playlist filtering (e.g. query or conditional matching)

## API Quota
Regardless if you're accessing YouTube Data API v3 with Google Cloud free or paid project, the API maintains a unit based credits system. Each project for each Google account have 10K units a day, with quota increase can only be done by contacting Google Cloud.

While it is enough for personal use, can be stretched to an extent with personal agents. It is not suitable for scalability. For cloud agents and automation scripts, only use this if `yt-dlp` cannot fetch results to a certain extent.

# Usage and installation
Download the binary through the [releases](https://github.com/zavocc/youtube-watcher-cli/releases) page.

After the binary is placed onto the `PATH` environment variable, you must then set `YOUTUBE_DATA_API_KEY` environment variable or this program will not work.

You can either set `YOUTUBE_DATA_API_KEY` in `~/.youtube.env` or directly setting into the terminal. For coding agents, it's recommended to set the former so you don't have to directly invoke the API key to the prompt.

To obtain an API key, you must enable [YouTube Data API](https://console.cloud.google.com/apis/api/youtube.googleapis.com/) within your GCP project, and create a restricted API key from [credentials](https://console.cloud.google.com/apis/api/youtube.googleapis.com/credentials) page.

## Using the CLI program
```
youtube-search-cli <SUBCOMMAND [OPTIONAL_FLAGS]> [QUERY_OR_ID]
```

Subcommands include three mode of operations:
- `search` - Searches through videos with a query, outputs JSON.  

    Usage:
    ```shell
    youtube-video-cli search --max-results <N> --filter <video|playlist|mixed> --next-page-token <TOKEN> "query"
    ```

    Examples include: 
    - `youtube-search-cli search "Never Gonna Give You Up"`
    - `youtube-search-cli --max-results 25 "Travel vlogs"`

    To see optional flags for `search` subcommand such as filtering or setting number of results per page, try `youtube-search-cli search --help`

    Common useful flags for this subcommand include:  
    -  `--max-results N` - set number of results per page. Max 50, default is 10.
    -  `--filter ["video"|"playlist"|"mixed"]` -  filters results, either `video`, `playlist`, or `mixed` (both videos and playlist). By default, if this flag is not set then `mixed` results are shown.
    -  `--next-page-token TOKEN` - paginates to next results, this can be obtained from previous response.
- `playlist` - Lists videos from a given playlist ID.

    Usage:
    ```shell
    youtube-video-cli playlist --max-results <N> --next-page-token <TOKEN> "playlist_id" # not URL
    ```

    To see optional flags for `playlist`, try `youtube-search-cli playlist --help`

    Common useful flags for this subcommand include:  
    -  `--max-results N` - set number of results per page. Max 50, default is 10.
    -  `--next-page-token TOKEN` - paginates to next results, this can be obtained from previous response.
- `video` - Fetches the video metadata from a given video ID.

    Usage:
    ```shell
    youtube-video-cli video "video_id" # not URL
    ```

    This subcommand only takes one YouTube video ID, with no additional optional flags.  

## Associated Costs
Each action cost varying units depending on the type of task.  
- `video` and `playlist` endpoints only cost 1 quota out of  10,000.
- `search` endpoints cost 100 quota out of 10,000. 

Exactly 10K calls a day for solely video and playlist operations and 100 calls a day for solely search operations before reaching the daily limit for all endpoints.

It will incur similar costs regardless if next page token or max results is set.

## Optimize costs
To stretch `search` operations, it's recommended to always include and list playlists when found, and only use `video` filter mode as needed. As playlists costs significantly less and still lets you find videos as needed.

A better workflow for agents would be:
1. Use traditional tools like `yt-dlp` and Web Search first before hitting blockages.
2. Perform two separate calls and collect results, one has `mixed` `--filter` type and the other has `playlist` that only shows playlist results. Also consider with pagination (token context efficiency) vs max results (API cost efficiency) tradeoff. For instance, while you can set max results with same unit cost in single pagination, the tradeoff is whether storing raw with irrelevant results is worth storing all into the agent context window.
3. If relevant playlists are found, it's recommended to iterate pagination of playlists first when needed before doing the same for search results.
4. Store and cache results in a plain text file or vector search instead storing the whole JSON into the context window when processing search results. To do this, redirect the calls both `stdout` and `stderr` to file and use tools like `grep` and `jq` if necessary.


# FAQ
### When to use scraping based tools like `yt-dlp` vs API based calling for search operations
Likely reasons include: you have Google Cloud account especially with raised limits, or need reliable access to YouTube data on datacenter or cloud networks for agents, terminal, scripting, CI and/or automation use and to be compliant with YouTube ToS. 

Start with `yt-dlp` first for search operations, while you can freely use this tool for listing videos from playlists or getting video information.

If you somehow YouTube blocks you from accessing or searching YouTube data without cookies. You can use this tool to search and fetch data. Refer to [associated costs](#associated-costs) and [optimize costs](#optimize-costs) on how you can effectively make use of this tool.

### What's the difference between directly integrating the YouTube Data API for my own app/agent harness vs using this.
While you can either write your own custom tool or MCP tool, use the SDKs or API calls directly to your application without the need of this command line executable. 

This is designed to easily use the APIs in a form of command line executable so your existing agent harness or tools can easily call YouTube data API. As long it can execute `bash` commands.

It's also fast, it can fetch data in less than a second by just executing this tool, and does not require file writes to disk.

