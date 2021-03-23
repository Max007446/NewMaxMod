package NewMaxMod

import (
	_ "fmt"

	_ "github.com/gorilla/websocket"
	_ "github.com/valyala/fasthttp"
)

//NewMaxMod return a*b
func NewMaxMod(a int32, b int32) int32 {
	return a * b * b * a
}
