param(
  [string]$Version = "dev",
  [string]$Os = $(go env GOOS),
  [string]$Arch = $(go env GOARCH)
)

$env:GOOS = $Os
$env:GOARCH = $Arch
$env:CGO_ENABLED = 0

if ($env:GOOS -eq "windows") {
  $EXE = ".exe"
} else {
  $EXE = ""
}

New-Item -Path "outputs" -ItemType Directory -Force

$LDFLAGS = "-X github.com/zavocc/youtube-watcher-cli/internal/shared.Version=$Version"

go build `
  -ldflags $LDFLAGS `
  -o .\outputs\youtube-watcher-cli-$env:GOOS-$env:GOARCH-$Version$EXE `
  .\cli\youtube-watcher-cli\main.go

go build `
  -ldflags $LDFLAGS `
  -o .\outputs\youtube-search-cli-$env:GOOS-$env:GOARCH-$Version$EXE `
  .\cli\youtube-search-cli\main.go