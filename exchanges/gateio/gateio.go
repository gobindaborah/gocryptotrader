package gateio

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/thrasher-/gocryptotrader/common"
	"github.com/thrasher-/gocryptotrader/config"
	exchange "github.com/thrasher-/gocryptotrader/exchanges"
	"github.com/thrasher-/gocryptotrader/exchanges/request"
	"github.com/thrasher-/gocryptotrader/exchanges/ticker"
)

const (
	gateioTradeURL   = "https://api.gateio.io"
	gateioMarketURL  = "https://data.gateio.io"
	gateioAPIVersion = "api2/1"

	gateioSymbol          = "pairs"
	gateioMarketInfo      = "marketinfo"
	gateioKline           = "candlestick2"
	gateioOrder           = "private"
	gateioBalances        = "private/balances"
	gateioCancelOrder     = "private/cancelOrder"
	gateioCancelAllOrders = "private/cancelAllOrders"
	gateioOpenOrders      = "private/openOrders"
	gateioTicker          = "ticker"
	gateioTickers         = "tickers"
	gateioOrderbook       = "orderBook"

	gateioAuthRate   = 100
	gateioUnauthRate = 100
)

// Gateio is the overarching type across this package
type Gateio struct {
	exchange.Base
}

// SetDefaults sets default values for the exchange
func (g *Gateio) SetDefaults() {
	g.Name = "GateIO"
	g.Enabled = false
	g.Verbose = false
	g.RESTPollingDelay = 10
	g.APIWithdrawPermissions = exchange.AutoWithdrawCrypto
	g.RequestCurrencyPairFormat.Delimiter = "_"
	g.RequestCurrencyPairFormat.Uppercase = false
	g.ConfigCurrencyPairFormat.Delimiter = "_"
	g.ConfigCurrencyPairFormat.Uppercase = true
	g.AssetTypes = []string{ticker.Spot}
	g.SupportsAutoPairUpdating = true
	g.SupportsRESTTickerBatching = true
	g.Requester = request.New(g.Name,
		request.NewRateLimit(time.Second*10, gateioAuthRate),
		request.NewRateLimit(time.Second*10, gateioUnauthRate),
		common.NewHTTPClientWithTimeout(exchange.DefaultHTTPTimeout))
	g.APIUrlDefault = gateioTradeURL
	g.APIUrl = g.APIUrlDefault
	g.APIUrlSecondaryDefault = gateioMarketURL
	g.APIUrlSecondary = g.APIUrlSecondaryDefault
	g.WebsocketInit()
}

// Setup sets user configuration
func (g *Gateio) Setup(exch config.ExchangeConfig) {
	if !exch.Enabled {
		g.SetEnabled(false)
	} else {
		g.Enabled = true
		g.AuthenticatedAPISupport = exch.AuthenticatedAPISupport
		g.SetAPIKeys(exch.APIKey, exch.APISecret, "", false)
		g.APIAuthPEMKey = exch.APIAuthPEMKey
		g.SetHTTPClientTimeout(exch.HTTPTimeout)
		g.SetHTTPClientUserAgent(exch.HTTPUserAgent)
		g.RESTPollingDelay = exch.RESTPollingDelay
		g.Verbose = exch.Verbose
		g.BaseCurrencies = common.SplitStrings(exch.BaseCurrencies, ",")
		g.AvailablePairs = common.SplitStrings(exch.AvailablePairs, ",")
		g.EnabledPairs = common.SplitStrings(exch.EnabledPairs, ",")
		err := g.SetCurrencyPairFormat()
		if err != nil {
			log.Fatal(err)
		}
		err = g.SetAssetTypes()
		if err != nil {
			log.Fatal(err)
		}
		err = g.SetAutoPairDefaults()
		if err != nil {
			log.Fatal(err)
		}
		err = g.SetAPIURL(exch)
		if err != nil {
			log.Fatal(err)
		}
		err = g.SetClientProxyAddress(exch.ProxyAddress)
		if err != nil {
			log.Fatal(err)
		}
	}
}

// GetSymbols returns all supported symbols
func (g *Gateio) GetSymbols() ([]string, error) {
	var result []string

	url := fmt.Sprintf("%s/%s/%s", g.APIUrlSecondary, gateioAPIVersion, gateioSymbol)

	err := g.SendHTTPRequest(url, &result)
	if err != nil {
		return nil, nil
	}
	return result, err
}

// GetMarketInfo returns information about all trading pairs, including
// transaction fee, minimum order quantity, price accuracy and so on
func (g *Gateio) GetMarketInfo() (MarketInfoResponse, error) {
	type response struct {
		Result string        `json:"result"`
		Pairs  []interface{} `json:"pairs"`
	}

	url := fmt.Sprintf("%s/%s/%s", g.APIUrlSecondary, gateioAPIVersion, gateioMarketInfo)

	var res response
	var result MarketInfoResponse
	err := g.SendHTTPRequest(url, &res)
	if err != nil {
		return result, err
	}

	result.Result = res.Result
	for _, v := range res.Pairs {
		item := v.(map[string]interface{})
		for itemk, itemv := range item {
			pairv := itemv.(map[string]interface{})
			result.Pairs = append(result.Pairs, MarketInfoPairsResponse{
				Symbol:        itemk,
				DecimalPlaces: pairv["decimal_places"].(float64),
				MinAmount:     pairv["min_amount"].(float64),
				Fee:           pairv["fee"].(float64),
			})
		}
	}
	return result, nil
}

// GetLatestSpotPrice returns latest spot price of symbol
// updated every 10 seconds
//
// symbol: string of currency pair
func (g *Gateio) GetLatestSpotPrice(symbol string) (float64, error) {
	res, err := g.GetTicker(symbol)
	if err != nil {
		return 0, err
	}

	return res.Last, nil
}

// GetTicker returns a ticker for the supplied symbol
// updated every 10 seconds
func (g *Gateio) GetTicker(symbol string) (TickerResponse, error) {
	url := fmt.Sprintf("%s/%s/%s/%s", g.APIUrlSecondary, gateioAPIVersion, gateioTicker, symbol)

	var res TickerResponse
	err := g.SendHTTPRequest(url, &res)
	if err != nil {
		return res, err
	}
	return res, nil
}

// GetTickers returns tickers for all symbols
func (g *Gateio) GetTickers() (map[string]TickerResponse, error) {
	url := fmt.Sprintf("%s/%s/%s", g.APIUrlSecondary, gateioAPIVersion, gateioTickers)

	resp := make(map[string]TickerResponse)
	err := g.SendHTTPRequest(url, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// GetOrderbook returns the orderbook data for a suppled symbol
func (g *Gateio) GetOrderbook(symbol string) (Orderbook, error) {
	url := fmt.Sprintf("%s/%s/%s/%s", g.APIUrlSecondary, gateioAPIVersion, gateioOrderbook, symbol)

	var resp OrderbookResponse
	err := g.SendHTTPRequest(url, &resp)
	if err != nil {
		return Orderbook{}, err
	}

	if resp.Result != "true" {
		return Orderbook{}, errors.New("result was not true")
	}

	var ob Orderbook

	if len(resp.Asks) == 0 {
		return ob, errors.New("asks are empty")
	}

	// Asks are in reverse order
	for x := len(resp.Asks) - 1; x != 0; x-- {
		data := resp.Asks[x]

		price, err := strconv.ParseFloat(data[0], 64)
		if err != nil {
			continue
		}

		amount, err := strconv.ParseFloat(data[1], 64)
		if err != nil {
			continue
		}

		ob.Asks = append(ob.Asks, OrderbookItem{Price: price, Amount: amount})
	}

	if len(resp.Bids) == 0 {
		return ob, errors.New("bids are empty")
	}

	for x := range resp.Bids {
		data := resp.Bids[x]

		price, err := strconv.ParseFloat(data[0], 64)
		if err != nil {
			continue
		}

		amount, err := strconv.ParseFloat(data[1], 64)
		if err != nil {
			continue
		}

		ob.Bids = append(ob.Bids, OrderbookItem{Price: price, Amount: amount})
	}

	ob.Result = resp.Result
	ob.Elapsed = resp.Elapsed
	return ob, nil
}

// GetSpotKline returns kline data for the most recent time period
func (g *Gateio) GetSpotKline(arg KlinesRequestParams) ([]*KLineResponse, error) {
	url := fmt.Sprintf("%s/%s/%s/%s?group_sec=%d&range_hour=%d",
		g.APIUrlSecondary,
		gateioAPIVersion,
		gateioKline,
		arg.Symbol,
		arg.GroupSec,
		arg.HourSize)

	var rawKlines map[string]interface{}
	err := g.SendHTTPRequest(url, &rawKlines)
	if err != nil {
		return nil, err
	}

	var result []*KLineResponse
	if rawKlines == nil || rawKlines["data"] == nil {
		return nil, fmt.Errorf("rawKlines is nil. Err: %s", err)
	}

	rawKlineDatasString, _ := json.Marshal(rawKlines["data"].([]interface{}))
	rawKlineDatas := [][]interface{}{}
	if err := json.Unmarshal(rawKlineDatasString, &rawKlineDatas); err != nil {
		return nil, fmt.Errorf("rawKlines unmarshal failed. Err: %s", err)
	}

	for _, k := range rawKlineDatas {
		otString, _ := strconv.ParseFloat(k[0].(string), 64)
		ot, err := common.TimeFromUnixTimestampFloat(otString)
		if err != nil {
			return nil, fmt.Errorf("cannot parse Kline.OpenTime. Err: %s", err)
		}
		_vol, err := common.FloatFromString(k[1])
		if err != nil {
			return nil, fmt.Errorf("cannot parse Kline.Volume. Err: %s", err)
		}
		_id, err := common.FloatFromString(k[0])
		if err != nil {
			return nil, fmt.Errorf("cannot parse Kline.Id. Err: %s", err)
		}
		_close, err := common.FloatFromString(k[2])
		if err != nil {
			return nil, fmt.Errorf("cannot parse Kline.Close. Err: %s", err)
		}
		_high, err := common.FloatFromString(k[3])
		if err != nil {
			return nil, fmt.Errorf("cannot parse Kline.High. Err: %s", err)
		}
		_low, err := common.FloatFromString(k[4])
		if err != nil {
			return nil, fmt.Errorf("cannot parse Kline.Low. Err: %s", err)
		}
		_open, err := common.FloatFromString(k[5])
		if err != nil {
			return nil, fmt.Errorf("cannot parse Kline.Open. Err: %s", err)
		}
		result = append(result, &KLineResponse{
			ID:        _id,
			KlineTime: ot,
			Volume:    _vol,
			Close:     _close,
			High:      _high,
			Low:       _low,
			Open:      _open,
		})
	}
	return result, nil
}

// GetBalances obtains the users account balance
func (g *Gateio) GetBalances() (BalancesResponse, error) {
	var result BalancesResponse

	return result,
		g.SendAuthenticatedHTTPRequest("POST", gateioBalances, "", &result)
}

// SpotNewOrder places a new order
func (g *Gateio) SpotNewOrder(arg SpotNewOrderRequestParams) (SpotNewOrderResponse, error) {
	var result SpotNewOrderResponse

	// Be sure to use the correct price precision before calling this
	params := fmt.Sprintf("currencyPair=%s&rate=%s&amount=%s",
		arg.Symbol,
		strconv.FormatFloat(arg.Price, 'f', -1, 64),
		strconv.FormatFloat(arg.Amount, 'f', -1, 64),
	)

	strRequestURL := fmt.Sprintf("%s/%s", gateioOrder, arg.Type)

	err := g.SendAuthenticatedHTTPRequest("POST", strRequestURL, params, &result)
	if err != nil {
		return result, err
	}

	return result, nil
}

// CancelExistingOrder cancels an order given the supplied orderID and symbol
// orderID order ID number
// symbol trade pair (ltc_btc)
func (g *Gateio) CancelExistingOrder(orderID int64, symbol string) (bool, error) {
	type response struct {
		Result  bool   `json:"result"`
		Code    int    `json:"code"`
		Message string `json:"message"`
	}

	var result response
	// Be sure to use the correct price precision before calling this
	params := fmt.Sprintf("orderNumber=%d&currencyPair=%s",
		orderID,
		symbol,
	)
	err := g.SendAuthenticatedHTTPRequest("POST", gateioCancelOrder, params, &result)
	if err != nil {
		return false, err
	}
	if !result.Result {
		return false, fmt.Errorf("code:%d message:%s", result.Code, result.Message)
	}

	return true, nil
}

// SendHTTPRequest sends an unauthenticated HTTP request
func (g *Gateio) SendHTTPRequest(path string, result interface{}) error {
	return g.SendPayload("GET", path, nil, nil, result, false, g.Verbose)
}

// CancelAllExistingOrders all orders for a given symbol and side
// orderType (0: sell,1: buy,-1: unlimited)
func (g *Gateio) CancelAllExistingOrders(orderType int64, symbol string) error {
	type response struct {
		Result  bool   `json:"result"`
		Code    int    `json:"code"`
		Message string `json:"message"`
	}

	var result response
	params := fmt.Sprintf("type=%d&currencyPair=%s",
		orderType,
		symbol,
	)
	err := g.SendAuthenticatedHTTPRequest("POST", gateioCancelAllOrders, params, &result)
	if err != nil {
		return err
	}

	if !result.Result {
		return fmt.Errorf("code:%d message:%s", result.Code, result.Message)
	}

	return nil
}

//
// GetOpenOrders retrieves all orders with an optional symbol filter
func (g *Gateio) GetOpenOrders(symbol string) (OpenOrdersResponse, error) {
	var params string
	var result OpenOrdersResponse

	if symbol != "" {
		params = fmt.Sprintf("currencyPair=%s", symbol)
	}

	err := g.SendAuthenticatedHTTPRequest("POST", gateioOpenOrders, params, &result)
	if err != nil {
		return result, err
	}

	if result.Code > 0 {
		return result, fmt.Errorf("code:%d message:%s", result.Code, result.Message)
	}

	return result, nil
}

// SendAuthenticatedHTTPRequest sends authenticated requests to the Gateio API
// To use this you must setup an APIKey and APISecret from the exchange
func (g *Gateio) SendAuthenticatedHTTPRequest(method, endpoint, param string, result interface{}) error {
	if !g.AuthenticatedAPISupport {
		return fmt.Errorf(exchange.WarningAuthenticatedRequestWithoutCredentialsSet, g.Name)
	}

	headers := make(map[string]string)
	headers["Content-Type"] = "application/x-www-form-urlencoded"
	headers["key"] = g.APIKey

	hmac := common.GetHMAC(common.HashSHA512, []byte(param), []byte(g.APISecret))
	headers["sign"] = common.HexEncodeToString(hmac)

	url := fmt.Sprintf("%s/%s/%s", g.APIUrl, gateioAPIVersion, endpoint)

	var intermidiary json.RawMessage

	err := g.SendPayload(method, url, headers, strings.NewReader(param), &intermidiary, true, g.Verbose)
	if err != nil {
		return err
	}

	errCap := struct {
		Result  bool   `json:"result,string"`
		Code    int    `json:"code"`
		Message string `json:"message"`
	}{}

	if err := common.JSONDecode(intermidiary, &errCap); err == nil {
		if !errCap.Result {
			return fmt.Errorf("GateIO auth request error, code: %d message: %s",
				errCap.Code,
				errCap.Message)
		}
	}

	return common.JSONDecode(intermidiary, result)
}

// GetFee returns an estimate of fee based on type of transaction
func (g *Gateio) GetFee(feeBuilder exchange.FeeBuilder) (fee float64, err error) {
	switch feeBuilder.FeeType {
	case exchange.CryptocurrencyTradeFee:
		feePairs, err := g.GetMarketInfo()
		if err != nil {
			return 0, err
		}
		currencyPair := feeBuilder.FirstCurrency + feeBuilder.Delimiter + feeBuilder.SecondCurrency
		var feeForPair float64
		for _, i := range feePairs.Pairs {
			if strings.EqualFold(currencyPair, i.Symbol) {
				feeForPair = i.Fee
			}
		}
		if feeForPair == 0 {
			return 0, fmt.Errorf("Currency: '%s' failed to find fee data", currencyPair)
		}
		fee = calculateTradingFee(feeForPair, feeBuilder.PurchasePrice, feeBuilder.Amount)
	case exchange.CryptocurrencyWithdrawalFee:
		fee = getCryptocurrencyWithdrawalFee(feeBuilder.FirstCurrency)
	}

	if fee < 0 {
		fee = 0
	}

	return fee, nil
}

func calculateTradingFee(feeForPair, purchasePrice, amount float64) float64 {
	return (feeForPair / 100) * purchasePrice * amount
}

func getCryptocurrencyWithdrawalFee(currency string) float64 {
	return WithdrawalFees[currency]
}
