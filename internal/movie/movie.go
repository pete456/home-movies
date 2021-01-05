package movie

import (
	"home-movies/config"
	"github.com/BurntSushi/toml"
	"os"
	"log"
)

type MovieInfo struct {
	Name string
	Title string
	PlayTime int
}

func GetDataPath(name string) string {
	return config.Datapath+name+"/"+name+"."+config.Movieformat
}

func OpenInfo(item string) *MovieInfo {
	f,err := os.Open(config.Datapath+item+"/index.toml")
	if err != nil {
		log.Panic(err)
	}
        mi := MovieInfo{Name:item}
        if _, err := toml.DecodeReader(f,&mi); err != nil {
                log.Panic(err)
        }
	f.Close()
	return &mi
}

