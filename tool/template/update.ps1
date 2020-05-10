Set-Location {{PROJECT_PATH}}
if(Test-Path "output"){
  Write-Verbose "Create Static File"
  New-Item -ItemType "directory" -Force output -ErrorAction 'silentlycontinue'
  Copy-Item asset output/ -ErrorAction 'silentlycontinue' -Recurse
  Copy-Item index.html output  -ErrorAction 'silentlycontinue'
}else{
  Write-Verbose "Static Exist, Compile Only"
}
$Env:GOARCH="wasm"
$Env:GOOS="js"
go get github.com/TobiasYin/goi@latest
go build -o output/main.wasm main.go
