package ws

import (
	"context"
	"github.com/njmdk/okx_v5sdk_go/ws/wImpl"
	"log"
	"time"
)

/*
websocket交易 通用请求
参数说明：

	evtId：封装的事件类型
	id: 请求ID
	op: 请求参数op
	params: 请求参数
	timeOut: 超时时间
*/
func (a *WsClient) jrpcReq(evtId wImpl.Event, op string, id string, params []map[string]interface{}, timeOut ...int) (res bool, detail *ProcessDetail, err error) {
	res = true
	tm := 5000
	if len(timeOut) != 0 {
		tm = timeOut[0]
	}

	req := &wImpl.JRPCReq{
		Id:   id,
		Op:   op,
		Args: params,
	}

	detail = &ProcessDetail{
		EndPoint: a.WsEndPoint,
	}

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Duration(tm)*time.Millisecond)
	defer cancel()
	ctx = context.WithValue(ctx, "detail", detail)

	msg, err := a.process(ctx, evtId, req)
	if err != nil {
		res = false
		log.Println("处理请求失败!", req, err)
		return
	}
	detail.Data = msg

	res, err = checkResult(req, msg)
	if err != nil {
		res = false
		return
	}

	return
}

/*
单个下单
参数说明：

	id: 请求ID
	params: 请求参数
	timeOut: 超时时间
*/
func (a *WsClient) PlaceOrder(id string, param map[string]interface{}, timeOut ...int) (res bool, detail *ProcessDetail, err error) {
	op := "order"
	evtId := wImpl.EVENT_PLACE_ORDER

	var args []map[string]interface{}
	args = append(args, param)

	return a.jrpcReq(evtId, op, id, args, timeOut...)

}

/*
批量下单
参数说明：

	id: 请求ID
	params: 请求参数
	timeOut: 超时时间
*/
func (a *WsClient) BatchPlaceOrders(id string, params []map[string]interface{}, timeOut ...int) (res bool, detail *ProcessDetail, err error) {

	op := "batch-orders"
	evtId := wImpl.EVENT_PLACE_BATCH_ORDERS
	return a.jrpcReq(evtId, op, id, params, timeOut...)

}

/*
单个撤单
参数说明：

	id: 请求ID
	params: 请求参数
	timeOut: 超时时间
*/
func (a *WsClient) CancelOrder(id string, param map[string]interface{}, timeOut ...int) (res bool, detail *ProcessDetail, err error) {

	op := "cancel-order"
	evtId := wImpl.EVENT_CANCEL_ORDER

	var args []map[string]interface{}
	args = append(args, param)

	return a.jrpcReq(evtId, op, id, args, timeOut...)

}

/*
批量撤单
参数说明：

	id: 请求ID
	params: 请求参数
	timeOut: 超时时间
*/
func (a *WsClient) BatchCancelOrders(id string, params []map[string]interface{}, timeOut ...int) (res bool, detail *ProcessDetail, err error) {

	op := "batch-cancel-orders"
	evtId := wImpl.EVENT_CANCEL_BATCH_ORDERS
	return a.jrpcReq(evtId, op, id, params, timeOut...)

}

/*
单个改单
参数说明：

	id: 请求ID
	params: 请求参数
	timeOut: 超时时间
*/
func (a *WsClient) AmendOrder(id string, param map[string]interface{}, timeOut ...int) (res bool, detail *ProcessDetail, err error) {

	op := "amend-order"
	evtId := wImpl.EVENT_AMEND_ORDER

	var args []map[string]interface{}
	args = append(args, param)

	return a.jrpcReq(evtId, op, id, args, timeOut...)

}

/*
批量改单
参数说明：

	id: 请求ID
	params: 请求参数
	timeOut: 超时时间
*/
func (a *WsClient) BatchAmendOrders(id string, params []map[string]interface{}, timeOut ...int) (res bool, detail *ProcessDetail, err error) {

	op := "batch-amend-orders"
	evtId := wImpl.EVENT_AMEND_BATCH_ORDERS
	return a.jrpcReq(evtId, op, id, params, timeOut...)

}
