package main

import (
	"encoding/json"
	fmt "fmt"
	"syscall/js"

	"github.com/gogo/protobuf/proto"
)

func main() {
	c := make(chan bool)

	msg := &Message{
		Type: 1,
		Data: []byte("sdfsdfsdfsdf"),
	}

	js.Global().Set("c", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		data, _ := proto.Marshal(msg)
		fmt.Println(data)

		s := make([]interface{}, len(data))
		for i, v := range data {
			s[i] = v
		}

		return js.ValueOf(s)
	}))

	js.Global().Set("d", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		message := &Message{}
		sdsd := make([]byte, args[0].Length())
		for i := 0; i < args[0].Length(); i++ {
			sdsd[i] = byte(args[0].Index(i).Int())
		}
		_ = proto.Unmarshal(sdsd, message)

		//fmt.Println(sdsd, message)

		data := make([]interface{}, len(message.GetData()))
		for i, v := range message.GetData() {
			data[i] = v
		}

		b, _ := json.Marshal(struct {
			Type string
			Data []interface{}
		}{
			Type: message.Type.String(),
			Data: data,
		})

		return js.ValueOf(string(b))
	}))

	js.Global().Set("exit", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		c <- true

		return nil
	}))

	<-c
}
