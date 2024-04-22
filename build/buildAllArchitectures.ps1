param(
    [string]$ProjectDirectory,
    [string]$Version
)

# Check if project directory and version number are provided
if (-not $ProjectDirectory -or -not $Version) {
    Write-Host "Usage: .\build.ps1 -ProjectDirectory <project_directory> -Version <version>"
    exit 1
}

# Check if project directory exists
if (-not (Test-Path $ProjectDirectory -PathType Container)) {
    Write-Host "Project directory not found: $ProjectDirectory"
    exit 1
}

# Build for darwin/amd64
$env:GOOS="darwin"
$env:GOARCH="amd64"
go build -o "out\urled_darwin_amd64_v$Version" $ProjectDirectory

# Build for darwin/arm64
$env:GOOS="darwin"
$env:GOARCH="arm64"
go build -o "out\urled_darwin_arm64_v$Version" $ProjectDirectory

# Build for freebsd/386
$env:GOOS="freebsd"
$env:GOARCH="386"
go build -o "out\urled_freebsd_386_v$Version" $ProjectDirectory

# Build for freebsd/amd64
$env:GOOS="freebsd"
$env:GOARCH="amd64"
go build -o "out\urled_freebsd_amd64_v$Version" $ProjectDirectory

# Build for freebsd/arm
$env:GOOS="freebsd"
$env:GOARCH="arm"
go build -o "out\urled_freebsd_arm_v$Version" $ProjectDirectory

# Build for linux/386
$env:GOOS="linux"
$env:GOARCH="386"
go build -o "out\urled_linux_386_v$Version" $ProjectDirectory

# Build for linux/amd64
$env:GOOS="linux"
$env:GOARCH="amd64"
go build -o "out\urled_linux_amd64_v$Version" $ProjectDirectory

# Build for linux/armv5
$env:GOOS="linux"
$env:GOARCH="arm"
$env:GOARM="5"
go build -o "out\urled_linux_armv5_v$Version" $ProjectDirectory

# Build for linux/armv6
$env:GOOS="linux"
$env:GOARCH="arm"
$env:GOARM="6"
go build -o "out\urled_linux_armv6_v$Version" $ProjectDirectory

# Build for linux/armv7
$env:GOOS="linux"
$env:GOARCH="arm"
$env:GOARM="7"
go build -o "out\urled_linux_armv7_v$Version" $ProjectDirectory

# Build for linux/arm64
$env:GOOS="linux"
$env:GOARCH="arm64"
go build -o "out\urled_linux_arm64_v$Version" $ProjectDirectory

set GOARM=

# Build for windows/amd64
$env:GOOS="windows"
$env:GOARCH="amd64"
go build -o "out\urled_windows_amd64_v$Version.exe" $ProjectDirectory

set GOOS=
set GOARCH=


Write-Host "Builds completed."
