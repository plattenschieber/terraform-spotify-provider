$global:Version = "0.1.0"
$global:Organization = "plattenschieber"
$global:Provider = "spotify"
$global:WorkDir = Get-Location

function Build {
  param (
    [Parameter()]
    [string]$Os,
    [Parameter()]
    [string]$Arch
  )

  if ($Os -eq "windows") { $fileEnding = ".exe" } else { $fileEnding = "" }

  $filePath = "$env:APPDATA/terraform.d/plugins/local/$Organization/$Provider/$Version/$($Os)_$($Arch)/terraform-provider-$($Provider)_v$($Version)$($fileEnding)"
  New-Item -Path "$WorkDir/bin/" -ItemType Directory -ErrorAction SilentlyContinue
  $zipPath = "$WorkDir/bin/terraform-provider-$($Provider)_$($Version)_$($Os)_$($Arch).zip"

  $env:CGO_ENABLED = 0
  $env:GOOS = $Os
  $env:GOARCH = $Arch

  go build -ldflags "-s -w -X main.version=$Version" -o $filePath
  Compress-Archive -DestinationPath $zipPath -LiteralPath $filePath

  $hash = Get-FileHash $zipPath -Algorithm SHA256
  Add-Content -Path "$WorkDir/bin/terraform-provider-$($Provider)_$($Version)_SHA256SUMS" -Value "$($hash.Hash.ToLower())  $([System.IO.Path]::GetFileName(($zipPath)))"
}

# Clean up old release
Remove-Item -Path "$WorkDir/bin/*.zip" -ErrorAction SilentlyContinue
Remove-Item -Path "$WorkDir/bin/terraform-provider-$($Provider)_$($Version)_SHA256SUMS" -ErrorAction SilentlyContinue
Remove-Item -Path "$WorkDir/bin/terraform-provider-$($Provider)_$($Version)_SHA256SUMS.sig" -ErrorAction SilentlyContinue

# Build binaries to local provider path and pack them into /bin
Build "linux" "amd64"
Build "windows" "amd64"
Build "windows" "386"

# Create a sig file in Git Bash using this
# gpg --local-user "Michael Hinni" --output /c/Users/Cxxxxx/Repos/terraform-provider-aadae/bin/terraform-provider-aadae_0.9.0_SHA256SUMS.sig --detach-sign /c/Users/C340685/Repos/terraform-provider-aadae/bin/terraform-provider-aadae_0.9.0_SHA256SUMS