package get

import (
	"fmt"
	"testing"
)

func TestGetBanData(t *testing.T) {
	b, err := GetBanData(0)
	if err != nil {
		t.Fatal(err)
		return
	}
	fmt.Println(b)
}
