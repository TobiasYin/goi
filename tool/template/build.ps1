Set-Location {{PROJECT_PATH}}
Remove-Item output -Recurse -ErrorAction 'silentlycontinue'
New-Item -ItemType "directory" -Force output -ErrorAction 'silentlycontinue'
Copy-Item assert output/ -ErrorAction 'silentlycontinue' -Recurse
Copy-Item index.html output  -ErrorAction 'silentlycontinue'
go get -d github.com/TobiasYin/go_web_ui
$Env:GOARCH="wasm"
$Env:GOOS="js"
go mod tidy
go get -d github.com/TobiasYin/go_web_ui
go build -o output/main.wasm main.go