package wInterface

import "github.com/njmdk/okx_v5sdk_go/ws/wImpl"

// 请求数据
type WSParam interface {
	EventType() wImpl.Event
	ToMap() *map[string]string
}
