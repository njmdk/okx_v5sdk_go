package ws

import (
	"github.com/njmdk/okx_v5sdk_go/ws/wImpl"
)

/*
订阅账户频道
*/
func (a *WsClient) PrivAccout(op string, params []map[string]string, timeOut ...int) (res bool, msg []*Msg, err error) {
	return a.PubChannel(wImpl.EVENT_BOOK_ACCOUNT, op, params, timeOut...)
}

/*
订阅持仓频道
*/
func (a *WsClient) PrivPostion(op string, params []map[string]string, timeOut ...int) (res bool, msg []*Msg, err error) {
	return a.PubChannel(wImpl.EVENT_BOOK_POSTION, op, params, timeOut...)
}

/*
订阅订单频道
*/
func (a *WsClient) PrivBookOrder(op string, params []map[string]string, timeOut ...int) (res bool, msg []*Msg, err error) {
	return a.PubChannel(wImpl.EVENT_BOOK_ORDER, op, params, timeOut...)
}

/*
订阅策略委托订单频道
*/
func (a *WsClient) PrivBookAlgoOrder(op string, params []map[string]string, timeOut ...int) (res bool, msg []*Msg, err error) {
	return a.PubChannel(wImpl.EVENT_BOOK_ALG_ORDER, op, params, timeOut...)
}

/*
订阅账户余额和持仓频道
*/
func (a *WsClient) PrivBalAndPos(op string, timeOut ...int) (res bool, msg []*Msg, err error) {
	var params []map[string]string
	arg := make(map[string]string)
	arg["channel"] = "balance_and_position"
	params = append(params, arg)
	return a.PubChannel(wImpl.EVENT_BOOK_B_AND_P, op, params, timeOut...)
}
