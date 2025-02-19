package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/xxmdhs/showdarkroom/get"
	"github.com/yosuke-furukawa/json5/encoding/json5"
)

func main() {
	var oldcid int64

	j := jsonData{
		Date: strconv.FormatInt(time.Now().Unix(), 10),
		Data: map[string]map[string]get.BanData{},
	}

	b, err := ioutil.ReadFile("data.json")
	must(err)

	sha := Sha256(b)
	j.Hash = sha
	fmt.Println(sha)

	var d jsonData
	err = json.Unmarshal(b, &d)
	must(err)
	temp := d.Cid
	oldcid, _ = strconv.ParseInt(temp, 10, 64)
	if d.Data != nil {
		j.Data = d.Data
	}

	var i int64
	for {
		log.Println(i)
		b, err := get.GetBanData(int(i))
		if err != nil {
			log.Println(err)
			serr := &json5.SyntaxError{}
			if errors.As(err, &serr) {
				break
			}
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
	Hash string                            `json:"hash"`
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

func Sha256(date []byte) string {
	h := sha256.New()
	h.Write(date)
	b := h.Sum(nil)
	return hex.EncodeToString(b)
}
