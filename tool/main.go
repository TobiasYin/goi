package main

import (
	"flag"
	"fmt"
	"github.com/TobiasYin/go_web_ui/tool/inline"
	"github.com/TobiasYin/go_web_ui/tool/server"
	"log"
	"os"
	"os/exec"
	"strings"
)

var repo = flag.String("repo", "", "Input You Repo URL, eg: github.com/TobiasYin/go_web_ui")
var run = flag.Bool("run", true, "Run A Project, if use this flag we will run a server in :port")
var _new = flag.Bool("new", false, "New Project Mode")
var port = flag.Int("port", 8080, "Run server port. only in run mode.")
var project = flag.String("project Name", "male", "input your project name, only in _new mode.")
var goPath string

func getProjectPath() string {
	return goPath + "/src/" + *repo
}

func runMode() {
	projectPath := getProjectPath()
	if !server.Exists(projectPath) {
		log.Fatalln("project not exist")
	}
	fmt.Println("cd " + projectPath +" && " + projectPath+"/build.sh")
	cmd := exec.Command("/bin/bash", "-c", "cd " + projectPath +" && " + projectPath+"/build.sh")
	cmd.Stderr = os.Stdout
	err := cmd.Run()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(*port)
	server.Serve(*port, projectPath+"/output/")
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
	cmd := exec.Command("/bin/bash", "-c", "chmod +x "+projectPath+"/build.sh")
	err := cmd.Run()
	if err != nil {
		log.Printf("error in give run perm on build.sh , user chmod to add perm. %e", err)
	}
}
func create(base string, f inline.File) {
	name := fmt.Sprintf("%s/%s", base, f.Name)
	if f.IsDir {
		err := os.Mkdir(name, os.ModePerm)
		if err != nil {
			log.Fatalf("error in create : %s\n", name)
		}
		for _, c := range f.Children {
			create(name, c)
		}
	} else if f.IsFile {
		c := f.Content
		c = strings.Replace(c, "github.com/TobiasYin/go_web_ui/tool/template", *repo, -1)
		c = strings.Replace(c, "{{PROJECT_NAME}}", *project, -1)
		c = strings.Replace(c, "{{PROJECT_PATH}}", getProjectPath(), -1)
		file, err := os.Create(name)
		if err != nil {
			log.Fatalf("error in create : %s\n", name)
		}
		_, err = file.WriteString(c)
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
		newMode()
	} else if !*run {
		log.Fatalln("run or _new mode require.")
	} else {
		runMode()
	}
}
