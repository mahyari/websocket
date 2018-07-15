package websocket

import (
	"encoding/base64"
	"github.com/pquerna/ffjson/ffjson"
)

type AuthToken struct {
	UserId     int    `json:"user_id"`
	Salt       string `json:"salt"`
	CreationAt int64  `json:"creation_at"`
	Sign       string `json:"sign"`
	ClientId   string `json:"client_id"`
}

func (a AuthToken) String() string {
	str, _ := ffjson.Marshal(a)
	return base64.RawStdEncoding.EncodeToString([]byte(string(str)))
}
