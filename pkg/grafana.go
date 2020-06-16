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

var(
	key = ""
)

func InitSearch()[]string{
	key = os.Getenv("GFANA_FIELDS")
	if key == "" {
		log.Fatalln("Cannot read GFANA_FIELDS, stop server..")
	}
	return strings.Split(key, ",")
}

func Query(get func() map[string]string) []QueryResp{
	m := get()
	resp := make([]QueryResp, len(m))
	count := 0

	//make it ordered
	for _, key := range strings.Split(key, ","){
		d := make([]details, 1)
		d[0] = details{m[key], time.Now().UnixNano() / 1000000}
		resp[count] = QueryResp{Target: key,Datapoints: d}
		count++
	}
	return resp
}