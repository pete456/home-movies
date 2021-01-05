package shelf

import (
	"home-movies/internal/movie"
	"home-movies/config"
	"github.com/BurntSushi/toml"
	"os"
	"strings"
	"log"
)

type Shelf struct {
	Name string
	Title string
	Movies []movie.MovieInfo
}

func extractname(path string) string {
	v := strings.Split(path,"/")
	log.Printf(string(len(v)))
	return v[len(v)-1]
}

func LoadShelf(path string) Shelf {
	s := Shelf{Name:extractname(path)}
	f, err := os.Open(config.Shelfpath+path+"shelf.toml")
	if err != nil {
		log.Panic(err)
	}
	if _, err := toml.DecodeReader(f,&s); err != nil {
		log.Panic(err)
	}
	return s
}
