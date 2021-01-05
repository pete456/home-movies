package net

import (
	"home-movies/config"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type Request struct {
	RequestStructure string
	DataName	 string
	Args		 []string
}

func urltolist(url string) []string {
	if url[len(url)-1] != '/' {
		url += "/"
	}
	surl := strings.Split(url,"/")
	if surl[0] == "" {
		return surl[1:]
	}
	fmt.Print(surl)
	return surl
}

func parseDataRequest(w http.ResponseWriter, r *http.Request) {
	url := urltolist(r.URL.Path)
	if len(url) < 3 {
		fmt.Fprint(w,"Request not long enough")
		return
	}
	request := Request{
		RequestStructure:url[2],
		DataName:url[1],
		Args:url[3:] }

	switch(request.RequestStructure) {
	case "shelf":
		HandleShelf(w,r,&request)
	case "movie":
		HandleData(w,r,&request)
	default:
		fmt.Fprint(w,"Invalid data type")
		return
	}
}

func InitServer() {
	http.HandleFunc("/getdata/",parseDataRequest)
	port := fmt.Sprintf(":%d",config.Port)
	log.Printf("Starting http server on port %s\n",port)
	log.Fatal(http.ListenAndServe(":8080",nil))
}
