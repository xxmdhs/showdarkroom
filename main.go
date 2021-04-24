package main

import (
	"encoding/json"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/xmdhs/showdarkroom/get"
)

func main() {
	data := map[string][]get.BanData{}
	ch := make(chan *get.Baninfo, 20)
	go tosave(&data, ch)

	var i int64 = 0
	for {
		log.Println(i)
		b, err := get.GetBanData(int(i))
		if err != nil {
			log.Println(err)
			time.Sleep(3 * time.Second)
			continue
		}
		ch <- b
		if b.Message.Dataexist == "1" {
			i, err = strconv.ParseInt(b.Message.Cid, 10, 64)
			must(err)
		} else {
			break
		}
		time.Sleep(500 * time.Millisecond)
	}

	close(ch)

	f, err := os.Create("data.json")
	must(err)
	defer f.Close()
	jw := json.NewEncoder(f)
	jw.SetIndent("", "    ")
	jw.SetEscapeHTML(false)
	jw.Encode(data)
}

func tosave(data *map[string][]get.BanData, ch <-chan *get.Baninfo) {
	for v := range ch {
		for k, v := range v.Data {
			if l, ok := (*data)[k]; ok {
				l = append(l, v)
				(*data)[k] = l
			} else {
				(*data)[k] = []get.BanData{v}
			}
		}
	}
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
