# go-paytm
Checksum utilities for paytm write by golang

This library provides three primary functions

* generate checksum
* verify checksum
* get transaction status


Before integrate the paytm, you need set up the configuration

```golang
#config.go

const (
	PaytmMerchantKey = `xxxxxxxxx`
	MID              = `xxxxxxxxxxxxxxxxxxxx`
	INDUSTRY_TYPE_ID = `Retail`
	CHANNEL_ID       = `WAP`
	WEBSITE          = `APPSTAGING`
	CALLBACK_URL     = `https://securegw-stage.paytm.in/theia/paytmCallback?ORDER_ID=`
	TransactionStatusAPI = `https://securegw-stage.paytm.in/merchant-status/getTxnStatus`
)
```
