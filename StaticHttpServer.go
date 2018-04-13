package main

import (
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	args := os.Args
	/*	for i,s := range args{
			print("--",i,"--",s,"\n")
		}
	*/
	port := "8080"
	argsLenth := len(args)
	if argsLenth != 2 {
		print("please enter the port\n")
		return
	}
	port = args[1]
	portLenth := len(port)
	if portLenth != 4 {
		print("error port\n")
		return
	}
	loadPage("", "")
	var handle http.Handler
	http.ListenAndServe("0.0.0.0:"+port, handle)
}

func loadPage(path string, urlPath string) {
	var files []os.FileInfo
	if path == "" {
		path = "./"
	}
	if urlPath == "" {
		urlPath = "/"
	}
	files, _ = ioutil.ReadDir(path)
	for _, f := range files {
		if f.IsDir() {
			var (
				pathDir    string
				urlPathDir string
			)
			if path == "./" {
				pathDir = path + f.Name()
			} else {
				pathDir = path + "/" + f.Name()
			}
			if urlPath == "/" {
				urlPathDir = urlPath + f.Name()
			} else {
				urlPathDir = urlPath + "/" + f.Name()
			}
			loadPage(pathDir, urlPathDir)
		} else {
			if path == "./" {
				bytes, _ := ioutil.ReadFile(path + f.Name())
				http.HandleFunc(urlPath+f.Name(), func(writer http.ResponseWriter, request *http.Request) {
					writer.Write(bytes)
				})
				print("\n", urlPath+f.Name(), "\n")
			} else {
				bytes, _ := ioutil.ReadFile(path + "/" + f.Name())
				print("\n", urlPath+f.Name(), "\n")
				http.HandleFunc(urlPath+"/"+f.Name(), func(writer http.ResponseWriter, request *http.Request) {
					writer.Write(bytes)
				})
			}
		}
	}
}
