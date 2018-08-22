package websocket

import (
	"encoding/base64"
	"github.com/pquerna/ffjson/ffjson"
	"time"
)

type AuthToken struct {
	UserId     int64    	`json:"user_id"`
	Uuid 	   string 		`json:"uuid"`
	Salt       string 		`json:"salt"`
	CreationAt time.Time	`json:"creation_at"`
	Sign       string 		`json:"sign"`
	ClientId   string 		`json:"client_id"`
}

func (a AuthToken) String() string {
	str, _ := ffjson.Marshal(a)
	return base64.RawStdEncoding.EncodeToString([]byte(string(str)))
}
