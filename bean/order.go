package bean

import (
	"bytes"
	"encoding/json"
)

type Response struct {
	Status string
	Msg    string
	Data   interface{}
}

func (r Response) ToString() string {
	js, _ := json.Marshal(r)
	var buff bytes.Buffer
	buff.Write(js)
	return buff.String()
}
