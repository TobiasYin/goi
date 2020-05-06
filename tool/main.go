package main

import (
	"flag"
	"fmt"
	"github.com/TobiasYin/go_web_ui/tool/inline"
	"github.com/TobiasYin/go_web_ui/tool/server"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

var repo = flag.String("repo", "", "Input You Repo URL, eg: github.com/TobiasYin/go_web_ui")
var run = flag.Bool("run", true, "Run A Project, if use this flag we will run a server in :port")
var _new = flag.Bool("new", false, "New Project Mode")
var port = flag.Int("port", 8080, "Run server port. only in run mode.")
var project = flag.String("project Name", "", "Input your project name, only in _new mode.")
var here = flag.Bool("here", false, "If you use this flag, your project will create Hear not in GOPATH")
var goPath string
var goRoot string
var allowPostfix = map[string]bool{"go": true, "py": true, "sh": true, "html": true, "js": true, "ps1": true}
var scriptPostFix = "sh"
var goPathVerbose = "GOPATH"

func init() {
	if runtime.GOOS == "windows" {
		scriptPostFix = "ps1"
	}
}
func getProjectPath() string {
	if goPath[len(goPath)-1] == '/' {
		goPath = goPath[:len(goPath)-1]
	} else if strings.HasSuffix(goPath, "\\") {
		goPath = goPath[:len(goPath)-2]
	}
	path := goPath + "/src/" + *repo
	if *here {
		path = goPath + "/" + *repo
	}
	path = strings.Replace(path, "\\", "/", -1)
	return path
}

func build(projectPath string, file string) {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("powershell.exe", "-c", fmt.Sprintf("cd %s; %s/%s.%s", projectPath, projectPath, file, scriptPostFix))
		cmd.Stderr = os.Stdout
		err := cmd.Run()
		if err != nil {
			log.Println(err)
		}
	} else {
		cmd := exec.Command("/bin/bash", "-c", fmt.Sprintf("cd %s && %s/%s.%s", projectPath, projectPath, file, scriptPostFix))
		cmd.Stderr = os.Stdout
		err := cmd.Run()
		if err != nil {
			log.Println(err)
		}
	}
}

func runMode() {
	projectPath := getProjectPath()
	if !server.Exists(projectPath) {
		log.Fatalln("project not exist")
	}
	log.Printf("Run Project in %s: %s\n", goPathVerbose, goPath)
	log.Printf("Project PATH: %s\n", projectPath)
	fmt.Println("building...")
	build(projectPath, "build")
	go server.Serve(*port, projectPath+"/output/")
	go func() {
		for {
			time.Sleep(time.Second * 10)
			build(projectPath, "update")
			//fmt.Println("rebuild...")
		}
	}()
	fmt.Println("Project Will Update 1 time per 10 second. press <Enter> will update now.")
	fmt.Println("When Update New Code, you should refresh in your browser to see the change.")
	listenInput()
}

func listenInput() {
	projectPath := getProjectPath()
	for {
		var s string
		_, _ = fmt.Scanln(&s)
		fmt.Println("Rebuild!")
		build(projectPath, "update")
	}
}

func newMode() {
	projectPath := getProjectPath()
	if server.Exists(projectPath) {
		log.Printf("Project Exist in %s: %s\n", goPathVerbose, goPath)
		log.Printf("Project PATH: %s\n", projectPath)
		log.Fatalln("project already exist")
	}
	i := strings.LastIndex(projectPath, "/")
	if i < 0 {
		log.Fatalln("project path error")
	}
	base := projectPath[:i]
	name := projectPath[i+1:]
	inline.Root.Name = name
	create(base, inline.Root)
	log.Printf("Project Create Success in %s: %s\n", goPathVerbose, goPath)
	log.Printf("Project PATH: %s\n", projectPath)
	programName := "go_web_ui_tool"
	if len(os.Args) > 0 {
		programName = os.Args[0]
	}
	var verbose strings.Builder
	verbose.WriteString(fmt.Sprintf("Use Command: \"%s -run ", programName))
	if *here {
		verbose.WriteString("-here ")
	}
	verbose.WriteString(fmt.Sprintf("-repo %s\" to Run The Project(no \")", *repo))
	log.Println(verbose.String())

	if runtime.GOOS == "windows" {
		cmd := exec.Command("powershell.exe", "-c", fmt.Sprintf("cd %s; go mod init %s", projectPath, *repo))
		_ = cmd.Run()
	} else {
		cmd := exec.Command("/bin/bash", "-c", fmt.Sprintf("cd %s; go mod init %s", projectPath, *repo))
		_ = cmd.Run()
	}

	if runtime.GOOS != "windows" {
		cmd := exec.Command("/bin/bash", "-c", fmt.Sprintf("chmod +x %s/build.sh %s/update.sh", projectPath, projectPath))
		err := cmd.Run()
		if err != nil {
			log.Printf("error in give run perm on build.sh , user chmod to add perm. %e", err)
		}
	}
}
func create(base string, f inline.File) {
	name := fmt.Sprintf("%s/%s", base, f.Name)
	if f.IsDir {
		err := os.MkdirAll(name, os.ModePerm)
		if err != nil {
			log.Fatalf("error in create : %s\n", name)
		}
		for _, c := range f.Children {
			create(name, c)
		}
	} else if f.IsFile {
		i := strings.LastIndex(f.Name, ".")
		if i < 0 {
			i = -1
		}
		postfix := f.Name[i+1:]
		if postfix == "sh" && runtime.GOOS == "windows" {
			return
		}
		if postfix == "ps1" && runtime.GOOS != "windows" {
			return
		}
		file, err := os.Create(name)
		if err != nil {
			log.Fatalf("error in create : %s\n", name)
		}
		o := f.Content
		if f.Name == "wasm_exec.js" && goRoot != "" {
			b, err := ioutil.ReadFile(fmt.Sprintf("%s/misc/wasm/wasm_exec.js", goRoot))
			if err == nil {
				if len(b) > 0 {
					o = b
				}
			}
		}
		if allowPostfix[postfix] {
			c := string(o)
			c = strings.Replace(c, "github.com/TobiasYin/go_web_ui/tool/template", *repo, -1)
			c = strings.Replace(c, "{{PROJECT_NAME}}", *project, -1)
			c = strings.Replace(c, "{{PROJECT_PATH}}", getProjectPath(), -1)
			_, err = file.WriteString(c)
		} else {
			_, err = file.Write(o)
		}
		if err != nil {
			log.Fatalf("error in create : %s\n", name)
		}
		_ = file.Close()
	}

}

func main() {
	flag.Parse()
	if *repo == "" && *_new {
		log.Fatalln("repo name require. Input You Repo URL, eg: github.com/TobiasYin/go_web_ui")
	}
	if *repo == "" {
		if server.Exists("build."+scriptPostFix) && server.Exists("update."+scriptPostFix) {
			*here = true
		} else {
			log.Fatalln("repo name require. Input You Repo URL, eg: github.com/TobiasYin/go_web_ui")
		}
	}
	if *here {
		goPath = os.Getenv("PWD")
		goPathVerbose = "Current Path"
	} else {
		goPath = os.Getenv("GOPATH")
		if goPath == "" {
			log.Fatalln("please set you go path in env.")
		}
		if *_new {
			log.Println("Project Will Create in GOPATH, if you want to create here, use -here flag.")
		}
	}
	goRoot = os.Getenv("GOROOT")
	if goRoot == "" {
		log.Println("go root not define. wasm assert may not suit for your go version, set GOROOT in your path.")
	}

	if *_new {
		if *project == "" {
			projectPath := getProjectPath()
			i := strings.LastIndex(projectPath, "/")
			if i < 0 {
				log.Fatalln("project path error")
			}
			*project = projectPath[i+1:]
		}
		newMode()
	} else if !*run {
		log.Fatalln("run or new mode require.")
	} else {
		runMode()
	}
}
