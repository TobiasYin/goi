Set-Location {{PROJECT_PATH}}
Remove-Item output -Recurse -ErrorAction 'silentlycontinue'
New-Item -ItemType "directory" -Force output -ErrorAction 'silentlycontinue'
Copy-Item asset output/ -ErrorAction 'silentlycontinue' -Recurse
Copy-Item index.html output  -ErrorAction 'silentlycontinue'
$Env:GOARCH="wasm"
$Env:GOOS="js"
go get github.com/TobiasYin/goi@latest
go build -o output/main.wasm main.go