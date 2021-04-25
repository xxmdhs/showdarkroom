package get

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/yosuke-furukawa/json5/encoding/json5"
)

func GetBanData(cid int) (*Baninfo, error) {
	data, err := getData("https://www.mcbbs.net/forum.php?mod=misc&action=showdarkroom&cid=" + strconv.Itoa(cid) + "&ajaxdata=json")
	if err != nil {
		return nil, fmt.Errorf("GetBanData: %w", err)
	}
	var t test
	err = json5.Unmarshal(data, &t)
	if err != nil {
		return nil, fmt.Errorf("GetBanData: %w", err)
	}
	var d map[string]BanData
	err = json5.Unmarshal(t.Data, &d)
	if err != nil {
		return &Baninfo{Data: d, Message: t.Message}, fmt.Errorf("GetBanData: %v : %w", string(t.Data), Errjson)
	}
	b := Baninfo{
		Data:    d,
		Message: t.Message,
	}
	return &b, err
}

var Errjson = errors.New("json err")

type test struct {
	Data    json5.RawMessage `json:"data"`
	Message banMessage       `json:"message"`
}

type Baninfo struct {
	Data    map[string]BanData `json:"data"`
	Message banMessage         `json:"message"`
}

type BanData struct {
	Action      string `json:"action"`
	Cid         string `json:"cid"`
	Dateline    string `json:"dateline"`
	Groupexpiry string `json:"groupexpiry"`
	Operator    string `json:"operator"`
	Operatorid  string `json:"operatorid"`
	Reason      string `json:"reason"`
	UID         string `json:"uid"`
	Username    string `json:"username"`
}

type banMessage struct {
	Cid       string `json:"cid"`
	Dataexist string `json:"dataexist"`
}
