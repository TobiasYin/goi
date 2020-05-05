package main

import (
	"flag"
	"fmt"
	"github.com/TobiasYin/go_web_ui/tool/inline"
	"github.com/TobiasYin/go_web_ui/tool/server"
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
var project = flag.String("project Name", "", "input your project name, only in _new mode.")
var goPath string
var allowPostfix = map[string]bool{"go": true, "py": true, "sh": true, "html": true, "js": true}

func getProjectPath() string {
	path := goPath + "/src/" + *repo
	path = strings.Replace(path, "\\", "/", -1)
	return path
}

func build(projectPath string)  {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("powershell.exe", "-c", fmt.Sprintf("cd %s; %s/build.sh", projectPath, projectPath))
		cmd.Stderr = os.Stdout
		err := cmd.Run()
		if err != nil {
			log.Println(err)
		}
	} else {
		cmd := exec.Command("/bin/bash", "-c",  fmt.Sprintf("cd %s && %s/build.sh", projectPath, projectPath))
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
	build(projectPath)
	fmt.Println(*port)
	go server.Serve(*port, projectPath+"/output/")
	go func() {
		time.Sleep(time.Second * 5)
		build(projectPath)
	}()
}

func newMode() {
	projectPath := getProjectPath()
	if server.Exists(projectPath) {
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
	if runtime.GOOS != "windows" {
		cmd := exec.Command("/bin/bash", "-c", "chmod +x "+projectPath+"/build.sh")
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

		file, err := os.Create(name)
		if err != nil {
			log.Fatalf("error in create : %s\n", name)
		}
		o := f.Content
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
	goPath = os.Getenv("GOPATH")
	if goPath == "" {
		log.Fatalln("please set you go path in env.")
	}
	if *repo == "" {
		log.Fatalln("repo name require.")
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
		log.Fatalln("run or _new mode require.")
	} else {
		runMode()
	}
}
