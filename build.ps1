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

go build `
  -ldflags "-X main.version=$Version" `
  -o .\outputs\youtube-watcher-cli-$env:GOOS-$env:GOARCH-$Version$EXE