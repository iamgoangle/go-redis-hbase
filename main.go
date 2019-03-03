package main

import (
	"context"
	"log"
	"os"

	"github.com/tsuna/gohbase"
	"github.com/tsuna/gohbase/hrpc"
)

func main() {
	client := gohbase.NewClient(os.Getenv("HBASE_HOST"))
	values := map[string]map[string][]byte{"f": map[string][]byte{"gggg": []byte(`{"Name":"Alice","Body":"Hello","Time":1294706395881547000}`)}}
	putRequest, err := hrpc.NewPutStr(context.Background(), "test", "key", values)
	rsp, err := client.Put(putRequest)
	if err != nil {
		panic(err)
	}

	log.Println(rsp)
}
