package yobit

import "github.com/thrasher-/gocryptotrader/currency/symbol"

// Response is a generic struct used for exchange API request result
type Response struct {
	Return  interface{} `json:"return"`
	Success int         `json:"success"`
	Error   string      `json:"error"`
}

// Info holds server time and pair information
type Info struct {
	ServerTime int64           `json:"server_time"`
	Pairs      map[string]Pair `json:"pairs"`
}

// Ticker stores the ticker information
type Ticker struct {
	High          float64 // maximal price
	Low           float64 // minimal price
	Avg           float64 // average price
	Vol           float64 // traded volume
	VolumeCurrent float64 `json:"vol_cur"` // traded volume in currency
	Last          float64 // last transaction price
	Buy           float64 // buying price
	Sell          float64 // selling price
	Updated       int64   // last cache upgrade
}

// Orderbook stores the asks and bids orderbook information
type Orderbook struct {
	Asks [][]float64 `json:"asks"` // selling orders
	Bids [][]float64 `json:"bids"` // buying orders
}

// Trades stores trade information
type Trades struct {
	Type      string  `json:"type"`
	Price     float64 `json:"bid"`
	Amount    float64 `json:"amount"`
	TID       int64   `json:"tid"`
	Timestamp int64   `json:"timestamp"`
}

// ActiveOrders stores active order information
type ActiveOrders struct {
	Pair             string  `json:"pair"`
	Type             string  `json:"type"`
	Amount           float64 `json:"amount"`
	Rate             float64 `json:"rate"`
	TimestampCreated float64 `json:"timestamp_created"`
	Status           int     `json:"status"`
}

// Pair holds pair information
type Pair struct {
	DecimalPlaces int     `json:"decimal_places"` // Quantity of permitted numbers after decimal point
	MinPrice      float64 `json:"min_price"`      // Minimal permitted price
	MaxPrice      float64 `json:"max_price"`      // Maximal permitted price
	MinAmount     float64 `json:"min_amount"`     // Minimal permitted buy or sell amount
	Hidden        int     `json:"hidden"`         // Pair is hidden (0 or 1)
	Fee           float64 `json:"fee"`            // Pair commission
}

// AccountInfo stores the account information for a user
type AccountInfo struct {
	Funds           map[string]float64 `json:"funds"`
	FundsInclOrders map[string]float64 `json:"funds_incl_orders"`
	Rights          struct {
		Info     int `json:"info"`
		Trade    int `json:"trade"`
		Withdraw int `json:"withdraw"`
	} `json:"rights"`
	TransactionCount int     `json:"transaction_count"`
	OpenOrders       int     `json:"open_orders"`
	ServerTime       float64 `json:"server_time"`
	Error            string  `json:"error"`
}

// OrderInfo stores order information
type OrderInfo struct {
	Pair             string  `json:"pair"`
	Type             string  `json:"type"`
	StartAmount      float64 `json:"start_amount"`
	Amount           float64 `json:"amount"`
	Rate             float64 `json:"rate"`
	TimestampCreated float64 `json:"timestamp_created"`
	Status           int     `json:"status"`
}

// CancelOrder is used for the CancelOrder API request response
type CancelOrder struct {
	OrderID float64            `json:"order_id"`
	Funds   map[string]float64 `json:"funds"`
	Error   string             `json:"error"`
}

// Trade stores the trade information
type Trade struct {
	Received float64            `json:"received"`
	Remains  float64            `json:"remains"`
	OrderID  float64            `json:"order_id"`
	Funds    map[string]float64 `json:"funds"`
	Error    string             `json:"error"`
}

// TradeHistory stores trade history
type TradeHistory struct {
	Pair      string  `json:"pair"`
	Type      string  `json:"type"`
	Amount    float64 `json:"amount"`
	Rate      float64 `json:"rate"`
	OrderID   float64 `json:"order_id"`
	MyOrder   int     `json:"is_your_order"`
	Timestamp float64 `json:"timestamp"`
}

// DepositAddress stores a curency deposit address
type DepositAddress struct {
	Address         string  `json:"address"`
	ProcessedAmount float64 `json:"processed_amount"`
	ServerTime      int64   `json:"server_time"`
	Error           string  `json:"error"`
}

// WithdrawCoinsToAddress stores information for a withdrawcoins request
type WithdrawCoinsToAddress struct {
	ServerTime int64  `json:"server_time"`
	Error      string `json:"error"`
}

// CreateCoupon stores information coupon information
type CreateCoupon struct {
	Coupon  string             `json:"coupon"`
	TransID int64              `json:"transID"`
	Funds   map[string]float64 `json:"funds"`
	Error   string             `json:"error"`
}

// RedeemCoupon stores redeem coupon information
type RedeemCoupon struct {
	CouponAmount   float64            `json:"couponAmount,string"`
	CouponCurrency string             `json:"couponCurrency"`
	TransID        int64              `json:"transID"`
	Funds          map[string]float64 `json:"funds"`
	Error          string             `json:"error"`
}

// WithdrawalFees the large list of predefined withdrawal fees
// Prone to change, using highest value
var WithdrawalFees = map[string]float64{
	symbol.ZERO07:     0.002,
	symbol.BIT16:      0.002,
	symbol.TWO015:     0.002,
	symbol.TWO56:      0.0002,
	symbol.TWOBACCO:   0.002,
	symbol.TWOGIVE:    0.01,
	symbol.THIRTY2BIT: 0.002,
	symbol.THREE65:    0.01,
	symbol.FOUR04:     0.01,
	symbol.SEVEN00:    0.01,
	symbol.EIGHTBIT:   0.002,
	symbol.ACLR:       0.002,
	symbol.ACES:       0.01,
	symbol.ACPR:       0.01,
	symbol.ACID:       0.01,
	symbol.ACOIN:      0.01,
	symbol.ACRN:       0.01,
	symbol.ADAM:       0.01,
	symbol.ADT:        0.05,
	symbol.AIB:        0.01,
	symbol.ADZ:        0.002,
	symbol.AECC:       0.002,
	symbol.AM:         0.002,
	symbol.AE:         10,
	symbol.DLT:        0.05,
	symbol.AGRI:       0.01,
	symbol.AGT:        0.01,
	symbol.AST:        0.002,
	symbol.AIR:        0.01,
	symbol.ALEX:       0.01,
	symbol.AUM:        0.002,
	symbol.ALIEN:      0.01,
	symbol.ALIS:       0.05,
	symbol.ALL:        0.01,
	symbol.ASAFE:      0.01,
	symbol.AMBER:      0.002,
	symbol.AMS:        0.002,
	symbol.ANAL:       0.002,
	symbol.ACP:        0.002,
	symbol.ANI:        0.01,
	symbol.ANTI:       0.002,
	symbol.ALTC:       0.01,
	symbol.APT:        0.01,
	symbol.ARCO:       0.002,
	symbol.ALC:        0.01,
	symbol.ANT:        0.01,
	symbol.ARB:        0.002,
	symbol.ARCT:       10,
	symbol.ARCX:       0.01,
	symbol.ARGUS:      0.01,
	symbol.ARH:        0.01,
	symbol.ARM:        0.01,
	symbol.ARNA:       10,
	symbol.ARPA:       0.002,
	symbol.ARTA:       0.01,
	symbol.ABY:        0.01,
	symbol.ARTC:       0.01,
	symbol.AL:         0.01,
	symbol.ASN:        0.002,
	symbol.ADCN:       0.01,
	symbol.ATB:        0.01,
	symbol.ATL:        0.1,
	symbol.ATM:        0.002,
	symbol.ATMCHA:     0.05,
	symbol.ATOM:       0.01,
	symbol.ADC:        0.002,
	symbol.REP:        0.002,
	symbol.ARE:        0.002,
	symbol.AUR:        0.01,
	symbol.AV:         0.002,
	symbol.AXIOM:      0.002,
	symbol.B2B:        10,
	symbol.B2:         0.01,
	symbol.B3:         0.1,
	symbol.BAB:        0.01,
	symbol.BAN:        0.002,
	symbol.BamitCoin:  0.002,
	symbol.NANAS:      0.002,
	symbol.BNT:        0.05,
	symbol.BBCC:       0.002,
	symbol.BAT:        0.05,
	symbol.BTA:        0.002,
	symbol.BSTK:       0.002,
	symbol.BATL:       0.01,
	symbol.BBH:        0.01,
	symbol.BITB:       0.002,
	symbol.BRDD:       0.002,
	symbol.XBTS:       0.01,
	symbol.BVC:        0.01,
	symbol.CHATX:      10,
	symbol.BEEP:       0.01,
	symbol.BEEZ:       0.002,
	symbol.BENJI:      0.01,
	symbol.BERN:       0.002,
	symbol.PROFIT:     0.01,
	symbol.BEST:       0.01,
	symbol.BGF:        0.01,
	symbol.BIGUP:      0.002,
	symbol.BLRY:       0.01,
	symbol.BILL:       0.01,
	symbol.BNB:        0.002,
	symbol.BIOB:       0.01,
	symbol.BIO:        0.1,
	symbol.BIOS:       0.002,
	symbol.BPTN:       10,
	symbol.BTCA:       10,
	symbol.BA:         0.002,
	symbol.BAC:        0.002,
	symbol.BBT:        10,
	symbol.BOSS:       0.01,
	symbol.BRONZ:      0.002,
	symbol.CAT:        0.01,
	symbol.BTD:        0.01,
	symbol.BTC:        0.0012,
	symbol.XBTC21:     0.01,
	symbol.BCA:        0.01,
	symbol.BCH:        0.01,
	symbol.BCP:        0.01,
	symbol.BCD:        0.01,
	symbol.BTDOLL:     0.01,
	symbol.GOD:        0.01,
	symbol.BTG:        0.01,
	symbol.LIZA:       0.01,
	symbol.BTCRED:     10,
	symbol.BTCS:       0.01,
	symbol.BTU:        0.01,
	symbol.BUM:        0.01,
	symbol.LITE:       0.01,
	symbol.BCM:        0.01,
	symbol.BCS:        0.01,
	symbol.BTCU:       0.002,
	symbol.BM:         10,
	symbol.BTCRY:      0.002,
	symbol.BTCR:       0.002,
	symbol.HIRE:       0.002,
	symbol.STU:        10,
	symbol.BITOK:      0.0001,
	symbol.BITON:      0.002,
	symbol.BPC:        0.01,
	symbol.BPOK:       0.01,
	symbol.BTP:        0.002,
	symbol.RNTB:       10,
	symbol.BSH:        0.002,
	symbol.BTS:        5,
	symbol.XBS:        0.002,
	symbol.BITS:       0.01,
	symbol.BST:        0.002,
	symbol.BXT:        0.01,
	symbol.VEG:        0.002,
	symbol.VOLT:       0.01,
	symbol.BTV:        0.01,
	symbol.BITZ:       0.002,
	symbol.BTZ:        0.002,
	symbol.BHC:        0.01,
	symbol.BDC:        0.002,
	symbol.JACK:       0.01,
	symbol.BS:         0.01,
	symbol.BSTAR:      0.01,
	symbol.BLAZR:      0.01,
	symbol.BOD:        0.002,
	symbol.BLUE:       10,
	symbol.BLU:        0.002,
	symbol.BLUS:       0.002,
	symbol.BMT:        10,
	symbol.BOT:        0.002,
	symbol.BOLI:       0.002,
	symbol.BOMB:       0.01,
	symbol.BON:        0.01,
	symbol.BOOM:       0.002,
	symbol.BOSON:      0.01,
	symbol.BSC:        0.002,
	symbol.BRH:        10,
	symbol.BRAIN:      0.01,
	symbol.BRE:        0.002,
	symbol.BTCM:       0.1,
	symbol.BTCO:       0.01,
	symbol.TALK:       0.01,
	symbol.BUB:        0.002,
	symbol.BUY:        0.01,
	symbol.BUZZ:       0.002,
	symbol.BTH:        0.1,
	symbol.C0C0:       0.002,
	symbol.CAB:        0.01,
	symbol.CF:         0.002,
	symbol.CLO:        10,
	symbol.CAM:        0.2,
	symbol.CD:         0.002,
	symbol.CANN:       0.2,
	symbol.CNNC:       0.01,
	symbol.CPC:        0.002,
	symbol.CST:        0.01,
	symbol.CAPT:       0.002,
	symbol.CARBON:     0.01,
	symbol.CME:        0.002,
	symbol.CTK:        0.002,
	symbol.CBD:        0.01,
	symbol.CCC:        0.01,
	symbol.CNT:        0.01,
	symbol.XCE:        0.002,
	symbol.CAG:        1,
	symbol.CHRG:       0.01,
	symbol.CHAT:       0.01,
	symbol.CHEMX:      0.01,
	symbol.CHESS:      0.01,
	symbol.CKS:        0.01,
	symbol.CHILL:      0.01,
	symbol.CHIP:       0.002,
	symbol.CHOOF:      0.01,
	symbol.TIME:       0.05,
	symbol.CRX:        0.01,
	symbol.CIN:        0.01,
	symbol.CLAM:       0.002,
	symbol.POLL:       10,
	symbol.CLICK:      0.002,
	symbol.CLINT:      0.01,
	symbol.CLOAK:      0.002,
	symbol.CLUB:       0.002,
	symbol.CLUD:       0.01,
	symbol.COX:        0.01,
	symbol.COXST:      0.01,
	symbol.CFC:        0.002,
	symbol.CTIC2:      0.01,
	symbol.COIN:       0.01,
	symbol.BTTF:       0.002,
	symbol.C2:         0.01,
	symbol.CAID:       0.002,
	symbol.CL:         10,
	symbol.CTIC:       0.01,
	symbol.CXT:        0.01,
	symbol.CHP:        10,
	symbol.CV2:        0.002,
	symbol.CMT:        0.01,
	symbol.COC:        0.01,
	symbol.COMP:       0.01,
	symbol.CMS:        10,
	symbol.CONX:       0.01,
	symbol.CCX:        0.01,
	symbol.CLR:        10,
	symbol.CORAL:      0.01,
	symbol.CORG:       0.01,
	symbol.CSMIC:      0.01,
	symbol.CMC:        0.01,
	symbol.COV:        0.002,
	symbol.COVX:       10,
	symbol.CRAB:       0.01,
	symbol.CRAFT:      0.01,
	symbol.CRNK:       0.01,
	symbol.CRAVE:      0.002,
	symbol.CRM:        0.01,
	symbol.XCRE:       0.01,
	symbol.CREDIT:     0.002,
	symbol.CREVA:      0.002,
	symbol.CRIME:      0.002,
	symbol.CROC:       0.01,
	symbol.CRC:        10,
	symbol.CRW:        0.002,
	symbol.CRY:        0.002,
	symbol.CBX:        0.002,
	symbol.TKTX:       10,
	symbol.CB:         0.02,
	symbol.CIRC:       0.002,
	symbol.CCB:        0.002,
	symbol.CDO:        0.01,
	symbol.CG:         0.01,
	symbol.CJ:         0.01,
	symbol.CJC:        0.01,
	symbol.CYT:        0.002,
	symbol.CNX:        0.01,
	symbol.CRPS:       0.002,
	symbol.PING:       0.05,
	symbol.CS:         0.002,
	symbol.CWXT:       0.01,
	symbol.CCT:        0.05,
	symbol.CTL:        0.01,
	symbol.CURVES:     0.002,
	symbol.CC:         0.002,
	symbol.CYC:        0.002,
	symbol.CYG:        0.002,
	symbol.CYP:        0.002,
	symbol.FUNK:       0.01,
	symbol.CZECO:      0.01,
	symbol.DALC:       0.1,
	symbol.DLISK:      0.2,
	symbol.MOOND:      0.002,
	symbol.DB:         0.002,
	symbol.DCC:        0.002,
	symbol.DCYP:       0.002,
	symbol.DETH:       0.002,
	symbol.DKC:        0.01,
	symbol.DISK:       0.01,
	symbol.DRKT:       0.002,
	symbol.DTT:        0.002,
	symbol.DASH:       0.002,
	symbol.DASHS:      0.01,
	symbol.DBTC:       0.01,
	symbol.DCT:        0.002,
	symbol.DBET:       10,
	symbol.DEC:        0.002,
	symbol.DCR:        0.05,
	symbol.DECR:       0.002,
	symbol.DEA:        0.01,
	symbol.DPAY:       0.01,
	symbol.DCRE:       0.002,
	symbol.DC:         0.002,
	symbol.DES:        0.01,
	symbol.DEM:        0.002,
	symbol.DXC:        0.01,
	symbol.DCK:        0.01,
	symbol.DGB:        0.002,
	symbol.CUBE:       0.002,
	symbol.DGMS:       0.002,
	symbol.DBG:        0.01,
	symbol.DGCS:       0.002,
	symbol.DBLK:       0.002,
	symbol.DGD:        0.002,
	symbol.DIME:       0.002,
	symbol.DIRT:       0.002,
	symbol.DVD:        10,
	symbol.DMT:        10,
	symbol.NOTE:       0.002,
	symbol.DOGE:       100,
	symbol.DGORE:      0.002,
	symbol.DLC:        0.01,
	symbol.DRT:        0.1,
	symbol.DOTA:       0.01,
	symbol.DOX:        0.002,
	symbol.DRA:        0.002,
	symbol.DFT:        0.002,
	symbol.XDB:        0.002,
	symbol.DRM:        0.002,
	symbol.DRZ:        0.002,
	symbol.DRACO:      0.002,
	symbol.DBIC:       0.002,
	symbol.DUB:        0.002,
	symbol.GUM:        0.002,
	symbol.DUR:        0.01,
	symbol.DUST:       0.002,
	symbol.DUX:        0.01,
	symbol.DXO:        0.01,
	symbol.ECN:        0.01,
	symbol.EDR2:       0.002,
	symbol.EA:         0.002,
	symbol.EAGS:       0.002,
	symbol.EMT:        10,
	symbol.EBONUS:     0.001,
	symbol.ECCHI:      0.01,
	symbol.EKO:        0.01,
	symbol.ECLI:       0.002,
	symbol.ECOB:       3,
	symbol.ECO:        0.01,
	symbol.EDIT:       0.01,
	symbol.EDRC:       0.01,
	symbol.EDC:        0.01,
	symbol.EGAME:      0.01,
	symbol.EGG:        0.002,
	symbol.EGO:        0.01,
	symbol.ELC:        0.01,
	symbol.ELCO:       0.01,
	symbol.ECA:        0.01,
	symbol.EPC:        0.002,
	symbol.ELE:        0.005,
	symbol.ONE337:     0.002,
	symbol.EMB:        0.01,
	symbol.EMC:        0.02,
	symbol.EPY:        0.002,
	symbol.EMPC:       0.01,
	symbol.EMP:        0.002,
	symbol.ENE:        0.01,
	symbol.EET:        10,
	symbol.XNG:        0.01,
	symbol.EGMA:       0.002,
	symbol.ENTER:      0.01,
	symbol.ETRUST:     0.002,
	symbol.EOS:        10,
	symbol.EQL:        0.01,
	symbol.EQM:        0.002,
	symbol.EQT:        0.01,
	symbol.ERR:        0.002,
	symbol.ESC:        0.002,
	symbol.ESP:        0.01,
	symbol.ENT:        0.01,
	symbol.ETCO:       0.2,
	symbol.DOGETH:     0.002,
	symbol.ETH:        0.005,
	symbol.ECASH:      0.1,
	symbol.ETC:        0.005,
	symbol.ELITE:      0.05,
	symbol.ETHS:       0.01,
	symbol.ETL:        1,
	symbol.ETZ:        10,
	symbol.EUC:        0.01,
	symbol.EURC:       0.002,
	symbol.EUROPE:     0.01,
	symbol.EVA:        0.01,
	symbol.EGC:        0.002,
	symbol.EOC:        0.002,
	symbol.EVIL:       0.002,
	symbol.EVO:        0.002,
	symbol.EXB:        0.002,
	symbol.EXIT:       0.01,
	symbol.EXP:        0.01,
	symbol.XT:         0.01,
	symbol.F16:        0.01,
	symbol.FADE:       0.002,
	symbol.DROP:       0.002,
	symbol.FAZZ:       0.01,
	symbol.FX:         0.01,
	symbol.FIDEL:      0.01,
	symbol.FIDGT:      0.01,
	symbol.FIND:       0.002,
	symbol.FPC:        0.01,
	symbol.FIRE:       0.002,
	symbol.FFC:        0.002,
	symbol.FRST:       0.01,
	symbol.FIST:       0.002,
	symbol.FIT:        0.05,
	symbol.FLX:        0.01,
	symbol.FLVR:       0.01,
	symbol.FLY:        0.002,
	symbol.FONZ:       0.002,
	symbol.XFCX:       0.002,
	symbol.FOREX:      0.01,
	symbol.FRN:        0.002,
	symbol.FRK:        0.002,
	symbol.FRWC:       0.01,
	symbol.FGZ:        0.01,
	symbol.FRE:        0.002,
	symbol.FRDC:       0.002,
	symbol.FJC:        0.01,
	symbol.FURY:       0.002,
	symbol.FSN:        0.002,
	symbol.FCASH:      0.002,
	symbol.FTO:        0.01,
	symbol.FUZZ:       0.002,
	symbol.GAKH:       0.01,
	symbol.GBT:        0.01,
	symbol.GAME:       0.01,
	symbol.GML:        0.2,
	symbol.UNITS:      0.01,
	symbol.FOUR20G:    0.01,
	symbol.GSY:        0.002,
	symbol.GENIUS:     0.002,
	symbol.GEN:        0.002,
	symbol.GEO:        0.002,
	symbol.GER:        0.01,
	symbol.GSR:        0.01,
	symbol.SPKTR:      0.002,
	symbol.GIFT:       0.002,
	symbol.WTT:        10,
	symbol.GHS:        0.01,
	symbol.GIG:        0.002,
	symbol.GOT:        0.01,
	symbol.XGTC:       0.01,
	symbol.GIZ:        0.002,
	symbol.GLO:        0.01,
	symbol.GCR:        0.002,
	symbol.BSTY:       0.002,
	symbol.GLC:        0.01,
	symbol.GSX:        0.02,
	symbol.GNO:        0.05,
	symbol.GOAT:       0.002,
	symbol.GO:         0.01,
	symbol.GB:         0.01,
	symbol.GFL:        0.01,
	symbol.MNTP:       10,
	symbol.GP:         0.002,
	symbol.GNT:        0.05,
	symbol.GLUCK:      0.002,
	symbol.GOON:       0.01,
	symbol.GTFO:       0.002,
	symbol.GOTX:       0.01,
	symbol.GPU:        0.01,
	symbol.GRF:        0.002,
	symbol.GRAM:       0.002,
	symbol.GRAV:       0.002,
	symbol.GBIT:       0.002,
	symbol.GREED:      0.002,
	symbol.GE:         0.002,
	symbol.GREENF:     0.01,
	symbol.GRE:        0.01,
	symbol.GREXIT:     0.002,
	symbol.GMCX:       0.002,
	symbol.GROW:       0.01,
	symbol.GSM:        0.002,
	symbol.GT:         0.01,
	symbol.NLG:        0.01,
	symbol.HKN:        10,
	symbol.HAC:        0.05,
	symbol.HALLO:      0.01,
	symbol.HAMS:       0.01,
	symbol.HCC:        0.01,
	symbol.HPC:        0.01,
	symbol.HMC:        0.01,
	symbol.HAWK:       0.01,
	symbol.HAZE:       0.002,
	symbol.HZT:        0.002,
	symbol.HDG:        0.1,
	symbol.HEDG:       0.002,
	symbol.HEEL:       0.002,
	symbol.HMP:        0.01,
	symbol.PLAY:       0.01,
	symbol.HXX:        0.002,
	symbol.XHI:        0.01,
	symbol.HVCO:       0.01,
	symbol.HTC:        0.01,
	symbol.MINH:       0.01,
	symbol.HODL:       0.01,
	symbol.HON:        0.01,
	symbol.HOPE:       0.01,
	symbol.HQX:        10,
	symbol.HSP:        0.002,
	symbol.HTML5:      0.002,
	symbol.HMQ:        0.01,
	symbol.HYPERX:     0.01,
	symbol.HPS:        10,
	symbol.IOC:        0.002,
	symbol.IBANK:      0.01,
	symbol.IBITS:      0.002,
	symbol.ICASH:      0.002,
	symbol.ICOB:       0.01,
	symbol.ICN:        0.002,
	symbol.ICON:       0.01,
	symbol.IETH:       0.1,
	symbol.ILM:        0.002,
	symbol.IMPS:       0.01,
	symbol.NKA:        0.002,
	symbol.INCP:       0.01,
	symbol.IN:         0.01,
	symbol.INC:        0.002,
	symbol.IMS:        0.01,
	symbol.IND:        0.01,
	symbol.XIN:        0.01,
	symbol.IFLT:       0.01,
	symbol.INFX:       0.002,
	symbol.INGT:       0.01,
	symbol.INPAY:      0.01,
	symbol.INSANE:     0.01,
	symbol.INXT:       0.01,
	symbol.IFT:        0.05,
	symbol.INV:        0.01,
	symbol.IVZ:        0.002,
	symbol.ILT:        0.002,
	symbol.IONX:       0.01,
	symbol.ISL:        0.002,
	symbol.ITI:        0.01,
	symbol.ING:        10,
	symbol.IEC:        0.002,
	symbol.IW:         0.01,
	symbol.IXC:        0.01,
	symbol.IXT:        0.05,
	symbol.JPC:        0.002,
	symbol.JANE:       0.01,
	symbol.JWL:        0.01,
	symbol.JNT:        0.01,
	symbol.JIF:        0.002,
	symbol.JOBS:       0.01,
	symbol.JOCKER:     0.01,
	symbol.JW:         0.01,
	symbol.JOK:        0.01,
	symbol.XJO:        0.002,
	symbol.KGB:        0.01,
	symbol.KARMC:      0.01,
	symbol.KARMA:      0.002,
	symbol.KASHH:      0.01,
	symbol.KAT:        0.002,
	symbol.KC:         0.002,
	symbol.KICK:       0.05,
	symbol.KIDS:       0.01,
	symbol.KIN:        10,
	symbol.KNC:        0.01,
	symbol.KISS:       0.01,
	symbol.KOBO:       0.002,
	symbol.TP1:        0.002,
	symbol.KRAK:       0.002,
	symbol.KGC:        0.002,
	symbol.KTK:        0.002,
	symbol.KR:         0.005,
	symbol.KUBO:       0.01,
	symbol.KURT:       0.01,
	symbol.KUSH:       0.01,
	symbol.LANA:       0.01,
	symbol.LTH:        0.01,
	symbol.LAZ:        0.2,
	symbol.LEA:        0.002,
	symbol.LEAF:       0.002,
	symbol.LENIN:      0.01,
	symbol.LEPEN:      0.01,
	symbol.LIR:        0.01,
	symbol.LVG:        0.002,
	symbol.LGBTQ:      0.002,
	symbol.LHC:        10,
	symbol.EXT:        0.002,
	symbol.LBTC:       0.1,
	symbol.LSD:        0.01,
	symbol.LIMX:       0.002,
	symbol.LTD:        0.0000002,
	symbol.LINDA:      0.01,
	symbol.LKC:        0.002,
	symbol.LSK:        0.2,
	symbol.LBTCX:      0.01,
	symbol.LTC:        0.002,
	symbol.LCC:        1,
	symbol.LTCU:       0.01,
	symbol.LTCR:       0.002,
	symbol.LDOGE:      0.002,
	symbol.LTS:        0.002,
	symbol.LIV:        0.01,
	symbol.LIZI:       0.01,
	symbol.LOC:        0.01,
	symbol.LOCX:       10,
	symbol.LOOK:       0.01,
	symbol.LRC:        0.05,
	symbol.LOOT:       0.01,
	symbol.XLTCG:      0.002,
	symbol.BASH:       0.01,
	symbol.LUCKY:      0.002,
	symbol.L7S:        0.002,
	symbol.LDM:        0.05,
	symbol.LUMI:       0.01,
	symbol.LUNA:       0.01,
	symbol.LUN:        0.002,
	symbol.LC:         0.01,
	symbol.LUX:        0.002,
	symbol.MCRN:       0.01,
	symbol.XMG:        0.01,
	symbol.MMXIV:      0.02,
	symbol.MAT:        0.01,
	symbol.MAO:        0.01,
	symbol.MAPC:       0.002,
	symbol.MRB:        0.002,
	symbol.MXT:        0.01,
	symbol.MARV:       0.01,
	symbol.MARX:       0.01,
	symbol.MCAR:       0.002,
	symbol.MM:         0.002,
	symbol.GUP:        0.05,
	symbol.MVC:        0.05,
	symbol.MAVRO:      0.01,
	symbol.MAX:        0.01,
	symbol.MAZE:       0.002,
	symbol.MBIT:       0.01,
	symbol.MCOIN:      0.01,
	symbol.MPRO:       0.01,
	symbol.XMS:        0.002,
	symbol.MLITE:      0.01,
	symbol.MLNC:       0.01,
	symbol.MENTAL:     0.01,
	symbol.MERGEC:     0.01,
	symbol.MTLMC3:     0.002,
	symbol.METAL:      0.002,
	symbol.AMM:        0.01,
	symbol.MDT:        0.002,
	symbol.MUU:        0.01,
	symbol.MILO:       0.01,
	symbol.MND:        0.002,
	symbol.XMINE:      0.002,
	symbol.MNM:        0.01,
	symbol.XNM:        0.01,
	symbol.MIRO:       10,
	symbol.MIS:        0.002,
	symbol.MMXVI:      0.01,
	symbol.MGO:        0.01,
	symbol.MOIN:       0.002,
	symbol.MOJO:       0.01,
	symbol.TAB:        0.002,
	symbol.MCO:        0.005,
	symbol.MONETA:     0.002,
	symbol.MUE:        0.002,
	symbol.MONEY:      0.01,
	symbol.MRP:        0.002,
	symbol.MOTO:       0.002,
	symbol.MULTI:      0.01,
	symbol.MST:        0.01,
	symbol.MVR:        0.01,
	symbol.MYSTIC:     0.002,
	symbol.WISH:       10,
	symbol.NKT:        0.002,
	symbol.NMC:        0.002,
	symbol.NAT:        0.002,
	symbol.ENAU:       10,
	symbol.NAV:        0.002,
	symbol.NEBU:       0.002,
	symbol.NEF:        0.01,
	symbol.XEM:        20,
	symbol.NBIT:       0.01,
	symbol.NETKO:      0.01,
	symbol.NTM:        0.01,
	symbol.NETC:       0.002,
	symbol.NEU:        10,
	symbol.NRC:        1,
	symbol.NTK:        10,
	symbol.NTRN:       0.002,
	symbol.NEVA:       0.01,
	symbol.NIC:        0.01,
	symbol.NKC:        0.002,
	symbol.NYC:        0.01,
	symbol.NZC:        0.01,
	symbol.NICE:       0.002,
	symbol.NET:        0.01,
	symbol.NDOGE:      0.002,
	symbol.XTR:        0.01,
	symbol.N2O:        0.002,
	symbol.NIXON:      0.01,
	symbol.NOC:        0.002,
	symbol.NODC:       0.01,
	symbol.NODES:      0.002,
	symbol.NODX:       0.002,
	symbol.NLC:        0.01,
	symbol.NLC2:       0.01,
	symbol.NOO:        0.002,
	symbol.NVC:        0.002,
	symbol.NPC:        0.002,
	symbol.NUBIS:      0.002,
	symbol.NUKE:       0.002,
	symbol.N7:         0.01,
	symbol.NUM:        0.01,
	symbol.NMR:        0.05,
	symbol.NXE:        0.002,
	symbol.OBS:        0.002,
	symbol.OCEAN:      0.01,
	symbol.OCOW:       0.01,
	symbol.EIGHT88:    0.02,
	symbol.OCC:        0.02,
	symbol.OK:         0.002,
	symbol.ODNT:       0.002,
	symbol.FLAV:       0.002,
	symbol.OLIT:       0.01,
	symbol.OLYMP:      0.01,
	symbol.OMA:        0.002,
	symbol.OMC:        0.01,
	symbol.OMG:        0.01,
	symbol.ONEK:       0.05,
	symbol.ONX:        0.01,
	symbol.XPO:        0.01,
	symbol.OPAL:       0.2,
	symbol.OTN:        0.1,
	symbol.OP:         0.01,
	symbol.OPES:       0.002,
	symbol.OPTION:     0.002,
	symbol.ORLY:       0.01,
	symbol.OS76:       0.002,
	symbol.OZC:        0.002,
	symbol.P7C:        0.002,
	symbol.PAC:        1,
	symbol.PAK:        0.002,
	symbol.PAL:        0.01,
	symbol.PND:        0.002,
	symbol.PINKX:      0.01,
	symbol.POPPY:      0.01,
	symbol.DUO:        0.002,
	symbol.PARA:       0.002,
	symbol.PKB:        0.002,
	symbol.GENE:       0.002,
	symbol.PARTY:      0.01,
	symbol.PYN:        10,
	symbol.XPY:        0.002,
	symbol.CON:        0.002,
	symbol.PAYP:       0.01,
	symbol.PPC:        0.2,
	symbol.GUESS:      10,
	symbol.PEN:        0.002,
	symbol.PTA:        0.002,
	symbol.PEO:        0.002,
	symbol.PSB:        0.01,
	symbol.XPD:        0.01,
	symbol.PXL:        0.002,
	symbol.PHR:        0.002,
	symbol.PIE:        0.01,
	symbol.PIO:        0.01,
	symbol.PIPR:       0.01,
	symbol.SKULL:      0.01,
	symbol.PIVX:       0.002,
	symbol.PLANET:     0.002,
	symbol.PNC:        0.002,
	symbol.XPTX:       0.01,
	symbol.PLNC:       0.002,
	symbol.PLU:        0.01,
	symbol.XPS:        0.01,
	symbol.POKE:       0.01,
	symbol.PLBT:       0.01,
	symbol.POLY:       0.002,
	symbol.POM:        0.001,
	symbol.PONZ2:      0.01,
	symbol.PONZI:      0.01,
	symbol.XSP:        0.002,
	symbol.PPT:        10,
	symbol.XPC:        0.002,
	symbol.PEX:        0.002,
	symbol.TRON:       0.002,
	symbol.POST:       0.01,
	symbol.POSW:       0.01,
	symbol.PWR:        0.01,
	symbol.POWER:      0.002,
	symbol.PRE:        0.002,
	symbol.PRS:        10,
	symbol.PXI:        0.002,
	symbol.PEXT:       10,
	symbol.PRIMU:      0.01,
	symbol.PRX:        0.01,
	symbol.PRM:        0.01,
	symbol.PRIX:       10,
	symbol.XPRO:       0.002,
	symbol.PCM:        0.01,
	symbol.PROC:       0.01,
	symbol.NANOX:      0.01,
	symbol.VRP:        0.01,
	symbol.PTY:        0.002,
	symbol.PSI:        0.002,
	symbol.PSY:        0.002,
	symbol.PULSE:      0.01,
	symbol.PUPA:       0.01,
	symbol.PURE:       0.002,
	symbol.VIDZ:       0.01,
	symbol.PUTIN:      0.01,
	symbol.PX:         0.1,
	symbol.QTM:        0.01,
	symbol.QTZ:        0.002,
	symbol.QBC:        0.01,
	symbol.XQN:        0.02,
	symbol.RBBT:       0.01,
	symbol.RAC:        10,
	symbol.RADI:       0.01,
	symbol.RAD:        0.002,
	symbol.RAI:        0.01,
	symbol.XRA:        0.002,
	symbol.RATIO:      0.002,
	symbol.RCN:        0.01,
	symbol.REA:        10,
	symbol.RCX:        0.01,
	symbol.RDD:        0.002,
	symbol.REE:        0.01,
	symbol.REC:        0.01,
	symbol.REQ:        10,
	symbol.RMS:        0.002,
	symbol.RBIT:       0.01,
	symbol.RNC:        0.002,
	symbol.R:          1,
	symbol.REV:        0.01,
	symbol.RH:         0.01,
	symbol.XRL:        1,
	symbol.RICE:       0.002,
	symbol.RICHX:      0.002,
	symbol.RID:        0.01,
	symbol.RIDE:       0.01,
	symbol.RBT:        0.002,
	symbol.RING:       0.01,
	symbol.RIO:        0.01,
	symbol.RISE:       0.2,
	symbol.ROCKET:     0.01,
	symbol.RPC:        0.01,
	symbol.ROS:        0.002,
	symbol.ROYAL:      0.01,
	symbol.RSGP:       0.01,
	symbol.RBIES:      0.002,
	symbol.RUBIT:      0.002,
	symbol.RBY:        0.01,
	symbol.RUC:        0.01,
	symbol.RUPX:       0.01,
	symbol.RUP:        0.01,
	symbol.RUST:       0.01,
	symbol.SFE:        0.01,
	symbol.SLS:        0.002,
	symbol.SMSR:       0.002,
	symbol.RONIN:      0.002,
	symbol.SAN:        0.05,
	symbol.STV:        0.002,
	symbol.HIFUN:      0.002,
	symbol.MAD:        0.002,
	symbol.SANDG:      0.002,
	symbol.STO:        0.01,
	symbol.SCAN:       0.01,
	symbol.SCITW:      0.002,
	symbol.SCRPT:      0.01,
	symbol.SCRT:       0.002,
	symbol.SED:        0.002,
	symbol.SEEDS:      0.002,
	symbol.B2X:        0.1,
	symbol.SEL:        0.01,
	symbol.SLFI:       0.002,
	symbol.SSC:        0.002,
	symbol.SMBR:       0.2,
	symbol.SEN:        0.002,
	symbol.SENT:       10,
	symbol.SRNT:       10,
	symbol.SEV:        0.01,
	symbol.SP:         0.01,
	symbol.SXC:        0.002,
	symbol.GELD:       0.05,
	symbol.SHDW:       0.01,
	symbol.SDC:        0.02,
	symbol.SAK:        0.002,
	symbol.SHRP:       0.01,
	symbol.SHELL:      0.002,
	symbol.SH:         0.01,
	symbol.SHORTY:     0.01,
	symbol.SHREK:      0.01,
	symbol.SHRM:       0.002,
	symbol.SIB:        0.002,
	symbol.SIGT:       0.01,
	symbol.SLCO:       0.01,
	symbol.SIGU:       0.002,
	symbol.SIX:        0.002,
	symbol.SJW:        0.002,
	symbol.SKB:        0.002,
	symbol.SW:         0.01,
	symbol.SLEEP:      0.01,
	symbol.SLING:      0.01,
	symbol.SMART:      0.01,
	symbol.SMC:        0.002,
	symbol.SMT:        10,
	symbol.SMF:        0.01,
	symbol.SOCC:       0.01,
	symbol.SCL:        0.05,
	symbol.SDAO:       10,
	symbol.SOLAR:      0.01,
	symbol.SOLO:       0.002,
	symbol.SCT:        0.01,
	symbol.SONG:       0.01,
	symbol.SNM:        0.05,
	symbol.ALTCOM:     0.01,
	symbol.SPHTX:      10,
	symbol.SOUL:       0.01,
	symbol.SPC:        0.002,
	symbol.SPACE:      0.002,
	symbol.SBT:        0.01,
	symbol.SPEC:       0.002,
	symbol.SPX:        0.002,
	symbol.SCS:        0.01,
	symbol.SPORT:      0.01,
	symbol.SPT:        0.002,
	symbol.SPR:        0.002,
	symbol.SPEX:       0.002,
	symbol.SQL:        0.002,
	symbol.SBIT:       0.002,
	symbol.STHR:       0.002,
	symbol.STALIN:     0.01,
	symbol.STAR:       0.01,
	symbol.STA:        0.01,
	symbol.START:      0.02,
	symbol.STP:        0.002,
	symbol.SNT:        10,
	symbol.PNK:        0.002,
	symbol.STEPS:      0.002,
	symbol.STK:        0.002,
	symbol.STONK:      0.01,
	symbol.STORJ:      0.05,
	symbol.STORM:      10,
	symbol.STS:        0.002,
	symbol.STRP:       0.002,
	symbol.STY:        10,
	symbol.SUB:        0.01,
	symbol.XMT:        0.002,
	symbol.SNC:        1,
	symbol.SSTC:       0.01,
	symbol.SBTC:       0.1,
	symbol.SUPER:      0.002,
	symbol.SRND:       0.002,
	symbol.STRB:       0.02,
	symbol.M1:         0.002,
	symbol.SPM:        0.01,
	symbol.BUCKS:      0.002,
	symbol.TOKEN:      0.01,
	symbol.SWT:        0.05,
	symbol.SWEET:      0.002,
	symbol.SWING:      0.002,
	symbol.CHSB:       10,
	symbol.SIC:        0.002,
	symbol.SDP:        0.002,
	symbol.XSY:        0.002,
	symbol.SYNX:       0.01,
	symbol.SNRG:       0.002,
	symbol.SYS:        0.002,
	symbol.TAG:        0.01,
	symbol.TAGR:       0.002,
	symbol.TAJ:        0.01,
	symbol.TAK:        0.01,
	symbol.TAKE:       0.01,
	symbol.TAM:        0.002,
	symbol.XTO:        0.01,
	symbol.TAP:        0.01,
	symbol.TLE:        0.01,
	symbol.TSE:        0.01,
	symbol.TLEX:       0.01,
	symbol.TAXI:       10,
	symbol.TCN:        0.01,
	symbol.TDFB:       0.002,
	symbol.TEAM:       0.01,
	symbol.TECH:       0.01,
	symbol.TEC:        0.002,
	symbol.TEK:        0.002,
	symbol.TB:         0.002,
	symbol.TLX:        10,
	symbol.TELL:       0.01,
	symbol.TENNET:     0.002,
	symbol.PAY:        0.002,
	symbol.TES:        0.002,
	symbol.TRA:        0.01,
	symbol.TGS:        10,
	symbol.XVE:        0.01,
	symbol.TCR:        0.01,
	symbol.GCC:        0.002,
	symbol.MAY:        0.01,
	symbol.THOM:       0.01,
	symbol.TIA:        0.002,
	symbol.TIDE:       0.01,
	symbol.TNT:        0.05,
	symbol.TIE:        10,
	symbol.TIT:        0.002,
	symbol.TTC:        0.002,
	symbol.TODAY:      0.01,
	symbol.TBX:        10,
	symbol.TKN:        0.01,
	symbol.TDS:        10,
	symbol.TLOSH:      0.01,
	symbol.TOKC:       0.01,
	symbol.TMRW:       0.01,
	symbol.TOOL:       0.01,
	symbol.TCX:        0.002,
	symbol.TOT:        0.01,
	symbol.TX:         0.002,
	symbol.TRANSF:     0.002,
	symbol.TRAP:       0.01,
	symbol.TBCX:       0.01,
	symbol.TRICK:      0.002,
	symbol.TPG:        0.01,
	symbol.TRX:        300,
	symbol.TFL:        10,
	symbol.TRUMP:      0.002,
	symbol.TNG:        0.002,
	symbol.TUR:        0.002,
	symbol.TWERK:      0.002,
	symbol.TWIST:      0.002,
	symbol.TWO:        0.002,
	symbol.UCASH:      10,
	symbol.UAE:        0.01,
	symbol.XBU:        0.01,
	symbol.UBQ:        0.01,
	symbol.U:          0.002,
	symbol.UDOWN:      0.01,
	symbol.GAIN:       0.01,
	symbol.USC:        0.01,
	symbol.UMC:        0.1,
	symbol.UNF:        0.01,
	symbol.UNIFY:      0.01,
	symbol.UKG:        10,
	symbol.USDE:       0.01,
	symbol.UBTC:       0.1,
	symbol.UIS:        0.01,
	symbol.UNIT:       0.002,
	symbol.UNI:        0.01,
	symbol.UXC:        0.01,
	symbol.URC:        0.002,
	symbol.XUP:        0.002,
	symbol.UFR:        10,
	symbol.URO:        0.002,
	symbol.UTLE:       0.002,
	symbol.VAL:        0.02,
	symbol.VPRC:       0.01,
	symbol.VAPOR:      0.002,
	symbol.VCOIN:      0.002,
	symbol.VEC:        0.002,
	symbol.VEC2:       0.01,
	symbol.VLT:        0.01,
	symbol.VENE:       0.01,
	symbol.VNTX:       0.01,
	symbol.VTN:        0.002,
	symbol.XVG:        0.002,
	symbol.CRED:       10,
	symbol.VERS:       0.01,
	symbol.VTC:        0.2,
	symbol.VTX:        0.002,
	symbol.VIA:        0.002,
	symbol.VTY:        0.01,
	symbol.VIP:        0.002,
	symbol.VISIO:      0.01,
	symbol.VK:         2,
	symbol.VOL:        0.002,
	symbol.VOYA:       0.002,
	symbol.VPN:        0.002,
	symbol.VSL:        0.01,
	symbol.XVS:        0.01,
	symbol.VTL:        0.01,
	symbol.VULC:       0.01,
	symbol.VVI:        10,
	symbol.WGR:        0.01,
	symbol.WAM:        0.01,
	symbol.WARP:       0.002,
	symbol.WASH:       0.01,
	symbol.WAVES:      0.002,
	symbol.WGO:        0.01,
	symbol.WAY:        0.01,
	symbol.WCASH:      0.01,
	symbol.WEALTH:     0.002,
	symbol.WEEK:       0.01,
	symbol.WHO:        0.5,
	symbol.WIC:        0.05,
	symbol.WBB:        0.002,
	symbol.WINE:       0.01,
	symbol.WINK:       0.01,
	symbol.WISC:       0.01,
	symbol.WITCH:      0.01,
	symbol.WMC:        0.01,
	symbol.WOMEN:      0.01,
	symbol.WOK:        0.01,
	symbol.WRC:        10,
	symbol.WRT:        0.01,
	symbol.XCO:        0.002,
	symbol.X2:         0.002,
	symbol.XNX:        0.002,
	symbol.XAU:        0.002,
	symbol.XAV:        0.01,
	symbol.XDE2:       0.002,
	symbol.XDE:        0.002,
	symbol.XIOS:       0.01,
	symbol.XOC:        0.01,
	symbol.XSSX:       0.002,
	symbol.XBY:        0.01,
	symbol.YAC:        0.01,
	symbol.YMC:        0.01,
	symbol.YAY:        0.01,
	symbol.YBC:        0.002,
	symbol.YES:        0.01,
	symbol.YOB2X:      0.01,
	symbol.YOVI:       0.002,
	symbol.ZYD:        0.01,
	symbol.ZEC:        0.02,
	symbol.ZECD:       0.01,
	symbol.ZEIT:       0.002,
	symbol.ZENI:       0.01,
	symbol.ZET2:       0.002,
	symbol.ZET:        0.002,
	symbol.ZMC:        0.002,
	symbol.ZIRK:       0.002,
	symbol.ZLQ:        0.01,
	symbol.ZNE:        0.01,
	symbol.ZONTO:      0.05,
	symbol.ZOOM:       0.002,
	symbol.ZRC:        0.01,
	symbol.ZUR:        0.002,
}
