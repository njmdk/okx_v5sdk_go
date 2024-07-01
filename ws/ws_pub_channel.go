package ws

import (
	"errors"
	"github.com/njmdk/okx_v5sdk_go/ws/wImpl"
)

/*
产品频道
*/
func (a *WsClient) PubInstruemnts(op string, params []map[string]string, timeOut ...int) (res bool, msg []*Msg, err error) {

	return a.PubChannel(wImpl.EVENT_BOOK_INSTRUMENTS, op, params, timeOut...)
}

func (a *WsClient) PubStatus(op string, timeOut ...int) (res bool, msg []*Msg, err error) {
	return a.PubChannel(wImpl.EVENT_STATUS, op, nil, timeOut...)
}

/*
行情频道
*/
func (a *WsClient) PubTickers(op string, params []map[string]string, timeOut ...int) (res bool, msg []*Msg, err error) {

	return a.PubChannel(wImpl.EVENT_BOOK_TICKERS, op, params, timeOut...)
}

/*
持仓总量频道
*/
func (a *WsClient) PubOpenInsterest(op string, params []map[string]string, timeOut ...int) (res bool, msg []*Msg, err error) {
	return a.PubChannel(wImpl.EVENT_BOOK_OPEN_INTEREST, op, params, timeOut...)
}

/*
K线频道
*/
func (a *WsClient) PubKLine(op string, params []map[string]string, timeOut ...int) (res bool, msg []*Msg, err error) {

	return a.PubChannel(wImpl.EVENT_BOOK_KLINE, op, params, timeOut...)
}

/*
交易频道
*/
func (a *WsClient) PubTrade(op string, params []map[string]string, timeOut ...int) (res bool, msg []*Msg, err error) {

	return a.PubChannel(wImpl.EVENT_BOOK_TRADE, op, params, timeOut...)
}

/*
预估交割/行权价格频道
*/
func (a *WsClient) PubEstDePrice(op string, params []map[string]string, timeOut ...int) (res bool, msg []*Msg, err error) {

	return a.PubChannel(wImpl.EVENT_BOOK_ESTIMATE_PRICE, op, params, timeOut...)

}

/*
标记价格频道
*/
func (a *WsClient) PubMarkPrice(op string, params []map[string]string, timeOut ...int) (res bool, msg []*Msg, err error) {

	return a.PubChannel(wImpl.EVENT_BOOK_MARK_PRICE, op, params, timeOut...)
}

/*
标记价格K线频道
*/
func (a *WsClient) PubMarkPriceCandle(op string, pd wImpl.Period, params []map[string]string, timeOut ...int) (res bool, msg []*Msg, err error) {

	return a.PubChannel(wImpl.EVENT_BOOK_MARK_PRICE_CANDLE_CHART, op, params, timeOut...)
}

/*
限价频道
*/
func (a *WsClient) PubLimitPrice(op string, params []map[string]string, timeOut ...int) (res bool, msg []*Msg, err error) {

	return a.PubChannel(wImpl.EVENT_BOOK_LIMIT_PRICE, op, params, timeOut...)
}

/*
深度频道
*/
func (a *WsClient) PubOrderBooks(op string, channel string, params []map[string]string, timeOut ...int) (res bool, msg []*Msg, err error) {

	switch channel {
	// 400档快照
	case "books":
		return a.PubChannel(wImpl.EVENT_BOOK_ORDER_BOOK, op, params, timeOut...)
	// 5档快照
	case "books5":
		return a.PubChannel(wImpl.EVENT_BOOK_ORDER_BOOK5, op, params, timeOut...)
	// 400 tbt
	case "books-l2-tbt":
		return a.PubChannel(wImpl.EVENT_BOOK_ORDER_BOOK_TBT, op, params, timeOut...)
	// 50 tbt
	case "books50-l2-tbt":
		return a.PubChannel(wImpl.EVENT_BOOK_ORDER_BOOK50_TBT, op, params, timeOut...)

	default:
		err = errors.New("未知的channel")
		return
	}

}

/*
期权定价频道
*/
func (a *WsClient) PubOptionSummary(op string, params []map[string]string, timeOut ...int) (res bool, msg []*Msg, err error) {

	return a.PubChannel(wImpl.EVENT_BOOK_OPTION_SUMMARY, op, params, timeOut...)
}

/*
资金费率频道
*/
func (a *WsClient) PubFundRate(op string, params []map[string]string, timeOut ...int) (res bool, msg []*Msg, err error) {

	return a.PubChannel(wImpl.EVENT_BOOK_FUND_RATE, op, params, timeOut...)
}

/*
指数K线频道
*/
func (a *WsClient) PubKLineIndex(op string, params []map[string]string, timeOut ...int) (res bool, msg []*Msg, err error) {

	return a.PubChannel(wImpl.EVENT_BOOK_KLINE_INDEX, op, params, timeOut...)
}

/*
指数行情频道
*/
func (a *WsClient) PubIndexTickers(op string, params []map[string]string, timeOut ...int) (res bool, msg []*Msg, err error) {

	return a.PubChannel(wImpl.EVENT_BOOK_INDEX_TICKERS, op, params, timeOut...)
}
