Set-Location {{PROJECT_PATH}}
if(Test-Path "output"){
  Write-Verbose "Create Static File"
  New-Item -ItemType "directory" -Force output -ErrorAction 'silentlycontinue'
  Copy-Item assert output/ -ErrorAction 'silentlycontinue' -Recurse
  Copy-Item index.html output  -ErrorAction 'silentlycontinue'
}else{
  Write-Verbose "Static Exist, Compile Only"
}
go get -d github.com/TobiasYin/go_web_ui
$Env:GOARCH="wasm"
$Env:GOOS="js"
go build -o output/main.wasm main.go
