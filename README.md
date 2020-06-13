## Installation
```shell
github.com/rkritchat/gfana
```

## Usage
Add GFANA_FIELDS to your configuration to `.env` file:

```shell
GFANA_FIELDS="field1,field2"
```
GFANA_FIELDS is fields that you need to show in Grafana

```go
package main

import (
   	"github.com/rkritchat/gfana"
	  "github.com/go-chi/chi"
)

func main(){
	r := chi.NewRouter()
	gfana.New(r, metric)
	//start your server with chi router
}

//the method will generate metric for grafana
func metric() map[string]string{
	m := make(map[string]string)
	m["field1"] = strconv.Itoa(rand.Intn(100)) //change me
	m["field2"] = strconv.Itoa(rand.Intn(100)) //change me
	return m
}

```
