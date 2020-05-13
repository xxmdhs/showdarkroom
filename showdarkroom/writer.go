package showdarkroom

import (
	"fmt"
	"io"
	"os"
)

func witer(cid string, json *string) {
	f, err := os.OpenFile(cid+".json", os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		witer(cid, json)
	}
	_, err1 := io.WriteString(f, *json)
	if err1 != nil {
		fmt.Println(err1)
		witer(cid, json)
	}
	defer f.Close()
}
