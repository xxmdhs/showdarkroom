package showdarkroom

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"
)

type Get struct {
	cid      int
	STRATCID int
	ENDCID   int
}

var W sync.WaitGroup

func (g *Get) Toget() {
	var cid int
	g.cid = g.STRATCID
	for g.cid > g.ENDCID {
		cid = g.cid
		time.Sleep(500 * time.Millisecond)
		txt := getjson(strconv.Itoa(g.cid))
		witer(strconv.Itoa(g.cid), txt)
		g.cid = getcid(*txt)
		fmt.Println(g.cid)
		if cid == g.cid {
			W.Done()
		}
	}
	W.Done()
}

func getcid(json string) int {
	a := strings.Index(json, "\"cid\":\"")
	b := strings.Index(json, "\"},\"data\"")
	c := json[a+7 : b]
	d, err := strconv.Atoi(c)
	if err != nil {
		fmt.Println(err)
	}
	return d
}
