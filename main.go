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
	data := map[string]map[string]get.BanData{}
	var oldcid int64

	j := jsonData{
		Date: strconv.FormatInt(time.Now().Unix(), 10),
		Data: data,
	}

	b, err := ioutil.ReadFile("data.json")
	if err == nil {
		var d jsonData
		err = json.Unmarshal(b, &d)
		if err == nil {
			temp := d.Cid
			oldcid, _ = strconv.ParseInt(temp, 10, 64)
		}
	}

	var i int64
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
			j.Cid = b.Message.Cid
		}
		if b.Message.Dataexist == "1" {
			i, err = strconv.ParseInt(b.Message.Cid, 10, 64)
			must(err)
			if i <= oldcid {
				break
			}
		} else {
			j.tosave(b)
			break
		}
		j.tosave(b)
		time.Sleep(500 * time.Millisecond)
	}

	f, err := os.Create("data.json")
	must(err)
	defer f.Close()
	jw := json.NewEncoder(f)
	jw.SetIndent("", "    ")
	jw.SetEscapeHTML(false)
	jw.Encode(j)
}

type jsonData struct {
	Cid  string                            `json:"cid"`
	Date string                            `json:"date"`
	Data map[string]map[string]get.BanData `json:"data"`
}

func (j *jsonData) tosave(data *get.Baninfo) {
	for k, v := range data.Data {
		var m map[string]get.BanData
		var ok bool
		if m, ok = j.Data[k]; !ok {
			m = make(map[string]get.BanData)
		}
		m[v.Cid] = v
		j.Data[k] = m
	}
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
