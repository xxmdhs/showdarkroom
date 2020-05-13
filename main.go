package main

import (
	"fmt"
	http "showdarkroom/showdarkroom"
	"strconv"
)

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
	http.W.Add(3)
	go a.Toget()
	go b.Toget()
	go c.Toget()
	http.W.Wait()
	fmt.Println("完成")
}
