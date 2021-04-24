package get

import (
	"fmt"
	"strconv"

	"github.com/yosuke-furukawa/json5/encoding/json5"
)

func GetBanData(cid int) (*Baninfo, error) {
	data, err := getData("https://www.mcbbs.net/forum.php?mod=misc&action=showdarkroom&cid=" + strconv.Itoa(cid) + "&ajaxdata=json")
	if err != nil {
		return nil, fmt.Errorf("GetBanData: %w", err)
	}
	var b Baninfo
	err = json5.Unmarshal(data, &b)
	if err != nil {
		return nil, fmt.Errorf("GetBanData: %w", err)
	}
	return &b, err
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
