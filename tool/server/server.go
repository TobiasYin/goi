package server

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func Exists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

func Serve(port int, dir string) {
	log.Printf("listening on :%d...\n", port)
	url := fmt.Sprintf("http://127.0.0.1:%d", port)
	log.Printf("visit at %s...\n", url)
	go func() {
		var cmd *exec.Cmd
		if runtime.GOOS == "windows" {
			cmd = exec.Command("powershell.exe", "start", url)
		} else if runtime.GOOS == "darwin" {
			cmd = exec.Command("open", url)
		} else {
			cmd = exec.Command("xdg-open", url)
		}
		_ = cmd.Run()
	}()
	server := http.FileServer(http.Dir(dir))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		if strings.HasSuffix(req.URL.Path, ".wasm") {
			resp.Header().Set("content-type", "application/wasm")
		}
		path := dir + req.RequestURI
		if Exists(path) && !IsDir(path) {
			server.ServeHTTP(resp, req)
		} else {
			defer func() { _ = req.Body.Close() }()
			res, err := ioutil.ReadFile(dir + "/index.html")
			if err != nil {
				resp.WriteHeader(http.StatusInternalServerError)
				_, _ = resp.Write([]byte("internal error"))
				return
			}
			resp.WriteHeader(http.StatusOK)
			_, _ = resp.Write(res)
		}
	})))
}
