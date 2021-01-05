package net

import (
	//"home-movies/config"
	"home-movies/internal/movie"
	//"bufio"
	//"fmt"
	//"io"
	"net/http"
	//"strings"
)

// Handles retrieving data from disk and sends it to user
func HandleData(w http.ResponseWriter, r *http.Request, ur *Request) {
	http.ServeFile(w,r,movie.GetDataPath(ur.DataName))
}

func HandleShelf(w http.ResponseWriter, r *http.Request, ur *Request) {
}
