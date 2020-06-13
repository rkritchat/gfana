package pkg

import (
	"log"
	"os"
	"strings"
	"time"
)

type SearchResp struct {
	Type string `json:"type"`
	Target string `json:"target"`
}

//Grafana query response
type QueryResp struct {
	Target     string `json:"target"`
	Datapoints []details  `json:"datapoints"`
}

//store value and time in UnixNano
type details []interface{}

func InitSearch()[]string{
	key := os.Getenv("gfana2.key")
	if key == "" {
		log.Fatalln("Cannot read gfana2.key")
	}
	return strings.Split(key, ",")
}

func Query(values map[string]string) []QueryResp{
	resp := make([]QueryResp, len(values))
	count := 0
	for k,v := range values{
		d := make([]details, 1)
		d[0] = details{v, time.Now().UnixNano() / 1000000}
		resp[count] = QueryResp{Target: k,Datapoints: d}
		count++
	}
	return resp
}