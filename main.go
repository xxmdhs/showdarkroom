package main

import (
	"fmt"
	http "showdarkroom/showdarkroom"
	"strconv"
	"sync"
)

var ww sync.WaitGroup

func main() {
	i := http.GetStartCid()
	ii, _ := strconv.Atoi(i)
	a := http.Get{
		STRATCID: ii * 10,
		ENDCID:   ii/3*2 + 2,
	}
	b := http.Get{
		STRATCID: ii/3*2 + 2,
		ENDCID:   ii/3 + 2,
	}
	c := http.Get{
		STRATCID: ii/3 + 2,
		ENDCID:   0,
	}
	ww.Add(1)
	go a.Toget()
	ww.Add(1)
	go b.Toget()
	ww.Add(1)
	go c.Toget()
	ww.Wait()
	fmt.Println("完成")
}
