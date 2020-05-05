package server

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
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
	log.Printf("listening on %d...", port)
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
