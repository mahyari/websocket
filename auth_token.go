package websocket

import (
	"encoding/base64"
	"github.com/pquerna/ffjson/ffjson"
)

type AuthToken struct {
	Id         int    `json:"id"`
	Salt       string `json:"salt"`
	CreationAt int64  `json:"creation_at"`
	Sign       string `json:"sign"`
}

func (a AuthToken) String() string {
	str, _ := ffjson.Marshal(a)
	return base64.RawStdEncoding.EncodeToString([]byte(string(str)))
}
