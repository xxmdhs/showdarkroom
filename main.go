package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/xxmdhs/showdarkroom/get"
)

func main() {
	data := map[string][]get.BanData{}
	ch := make(chan *get.Baninfo, 20)
	go tosave(&data, ch)
	var cid, oldcid string

	b, err := ioutil.ReadFile("data.json")
	if err == nil {
		var d jsonData
		err = json.Unmarshal(b, &d)
		if err == nil {
			oldcid = d.Cid
		}
	}

	var i int64 = 0
	for {
		log.Println(i)
		b, err := get.GetBanData(int(i))
		if err != nil {
			log.Println(err)
			if !errors.Is(err, get.Errjson) {
				time.Sleep(3 * time.Second)
				continue
			}
		}
		if i == 0 {
			cid = b.Message.Cid
		}
		ch <- b

		if b.Message.Dataexist == "1" {
			i, err = strconv.ParseInt(b.Message.Cid, 10, 64)
			must(err)
			if oldcid != "" {
				ocid, err := strconv.ParseInt(oldcid, 10, 64)
				must(err)
				if i < ocid {
					break
				}
			}
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
	jw.Encode(jsonData{
		Cid:  cid,
		Date: strconv.FormatInt(time.Now().Unix(), 10),
		Data: data,
	})
}

type jsonData struct {
	Cid  string                   `json:"cid"`
	Date string                   `json:"date"`
	Data map[string][]get.BanData `json:"data"`
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
