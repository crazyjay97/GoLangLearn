package main

import (
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	args := os.Args //获取输入的参数
	/*	遍历参数
		for i,s := range args{
			print("--",i,"--",s,"\n")
		}
	*/
	//初始化端口以及检验用户输入
	port := "8080"
	argsLenth := len(args)
	if argsLenth != 2 && argsLenth != 1 {
		print("please enter the port\n")
		return
	}
	if argsLenth == 2 {
		port = args[1]
	}
	portLenth := len(port)
	if portLenth != 4 {
		print("error port\n")
		return
	}
	//加载本地文件
	loadPage("", "")
	var handle http.Handler
	//绑定host 开始服务
	print("the server is starting...the port:", port, "\n")
	err := http.ListenAndServe("0.0.0.0:"+port, handle)
	print(err, "\n")
	print("please review your port")
}

func loadPage(path string, urlPath string) {
	var files []os.FileInfo
	if path == "" {
		path = "./"
	}
	if urlPath == "" {
		urlPath = "/"
	}
	//获取当前文件夹下所有文件包括文件夹
	files, _ = ioutil.ReadDir(path)
	for _, f := range files {
		if f.IsDir() { //如果当前是一个文件夹
			loadPage(appendPath(path, f.Name()), appendUrlPath(urlPath, f.Name()))
		} else {
			//绑定文件名
			if path == "./" {
				loadHandle(path+f.Name(), urlPath+f.Name())
			} else {
				loadHandle(path+"/"+f.Name(), urlPath+"/"+f.Name())
			}
		}
	}
}

//文件名和文件绑定
func loadHandle(path string, urlPath string) {
	bytes, _ := ioutil.ReadFile(path)
	//print("\n", urlPath, "\n")
	http.HandleFunc(urlPath, func(writer http.ResponseWriter, request *http.Request) {
		writer.Write(bytes)
	})
}

//拼接路径
func appendPath(prefixPath string, suffixPath string) string {
	if prefixPath == "./" {
		return prefixPath + suffixPath
	}
	return prefixPath + "/" + suffixPath
}

//拼接url->端口号后部分
func appendUrlPath(prefixUrlPath string, suffixUrlPath string) string {
	if prefixUrlPath == "/" {
		return prefixUrlPath + suffixUrlPath
	}
	return prefixUrlPath + "/" + suffixUrlPath
}
