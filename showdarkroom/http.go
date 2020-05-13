package showdarkroom

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func getjson(cid string) *string {
	client := http.Client{
		Timeout: 5 * time.Second,
	}
	rep, err := client.Get("https://www.mcbbs.net/forum.php?mod=misc&action=showdarkroom&cid=" + cid + "&ajaxdata=json")
	rep.Header.Add("Accept", "*/*")
	rep.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.163 Safari/537.36")
	b, err := ioutil.ReadAll(rep.Body)
	if err != nil || rep.StatusCode != http.StatusOK {
		fmt.Println(err)
		time.Sleep(1 * time.Second)
		getjson(cid)
	}
	txt := string(b)
	defer rep.Body.Close()
	return &txt
}

func GetStartCid() string {
	client := http.Client{
		Timeout: 5 * time.Second,
	}
	rep, err := client.Get("https://www.mcbbs.net/forum.php?mod=misc&action=showdarkroom")
	rep.Header.Add("Accept", "*/*")
	rep.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.163 Safari/537.36")
	if err != nil || rep.StatusCode != http.StatusOK {
		fmt.Println(err)
		time.Sleep(1 * time.Second)
		GetStartCid()
	}
	t := bufio.NewScanner(rep.Body)
	var txt string
	var a int
	var b int
	for t.Scan() {
		txt = t.Text()
		if strings.Contains(txt, "id=\"darkroommore\" cid=\"") {
			a = strings.Index(txt, "id=\"darkroommore\" cid=\"")
			b = strings.Index(txt, "\">更多</a></div></span>")
			return txt[a+23 : b]
		}
	}
	defer rep.Body.Close()
	return ""
}
