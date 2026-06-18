param(
    [string]$Version = "dev"
)

$arch = go env GOARCH
$platform = go env GOOS

New-Item -Path "outputs" -ItemType Directory -Force

go build `
  -ldflags "-X main.version=$Version" `
  -o .\outputs\youtube-watcher-cli-$platform-$arch-$Version.exe .