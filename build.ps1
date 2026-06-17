param(
    [string]$Version = "dev"
)

$arch = $env:PROCESSOR_ARCHITECTURE ?? "AMD64"

New-Item -Path "outputs" -ItemType Directory -Force

go build `
  -ldflags "-X main.version=$Version" `
  -o .\outputs\youtube-watcher-cli-win32-$arch-$Version.exe .