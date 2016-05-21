/*
 * Copyright 2015 JiaoHu. All rights reserved.
 * JiaoHu PROPRIETARY/CONFIDENTIAL. Use is subject to license terms.
 *
 * Created by tony on 15/12/17.
 */
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"path/filepath"
	"strings"
	"github.com/mind-stack-cn/golang-fileserver/handle"
)

var (
	dir string
	port string
	logging bool
	debug bool
)


const VERSION = "1.0"

func main() {
	//fmt.Println(len(os.Args), os.Args)
	if len(os.Args) > 1 && os.Args[1] == "-v" {
		fmt.Println("Version " + VERSION)
		os.Exit(0)
	}

	flag.StringVar(&dir, "dir", ".", "Specify a directory to server files from.")
	flag.StringVar(&port, "port", ":8088", "Port to bind the file server")
	flag.BoolVar(&logging, "log", true, "Enable Log (true/false)")
	flag.BoolVar(&debug, "debug", true, "Make external assets expire every request")
	flag.Parse()

	if logging == false {
		log.SetOutput(ioutil.Discard)
	}
	// If no path is passed to app, normalize to path formath
	if dir == "." {
		dir, _ = filepath.Abs(dir)
		dir += "/data/"
	}

	if _, err := os.Stat(dir); err != nil {
		log.Printf("Directory %s not exist, Create it", dir)
		errPath := os.MkdirAll(dir, 0777)
		if errPath != nil {
			log.Fatalf("Directory %s not exist, Create it Fail", dir)
			return
		}
	}

	// normalize dir, ending with... /
	if strings.HasSuffix(dir, "/") == false {
		dir = dir + "/"
	}

	mux := http.NewServeMux()

	mux.Handle("/", addDefaultHeaders(http.HandlerFunc(handleReq)))

	log.Printf("Listening on port %s .....\n", port)
	if debug {
		log.Print("Serving data dir in debug mode.. .\n")
	}
	http.ListenAndServe(port, mux)
}

func handleReq(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		handle.FileUpload(dir, w, r)
		return
	}

	log.Print("Request: ", r.RequestURI)
	// See bug #9. For some reason, don't arrive index.html, when asked it..
	if r.URL.Path != "/" && r.URL.Path != "/test/" && strings.HasSuffix(r.URL.Path, "/") && r.FormValue("get_file") != "true" {
		log.Printf("Index dir %s", r.URL.Path)
		http.Error(w, "BadRequest", http.StatusBadRequest)
	} else {
		handle.FileDownload(dir, w, r)
	}
}

func addDefaultHeaders(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		fn(w, r)
	}
}



