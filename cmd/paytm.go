package main

import (
	"github.com/yiv/go-paytm"
	"fmt"
	"encoding/json"
	"encoding/base64"
)

func main() {

	GetTransactionStatus()

}

func GetChecksumFromArray() {
	var (
		orderId  = "order456"
		customId = "custom001"
		amount   = "1.00"
	)
	paramList := map[string]string{
		"MID":              paytm.MID,
		"INDUSTRY_TYPE_ID": paytm.INDUSTRY_TYPE_ID,
		"CHANNEL_ID":       paytm.CHANNEL_ID,
		"WEBSITE":          paytm.WEBSITE,
		"CALLBACK_URL":     paytm.CALLBACK_URL + orderId,
		"ORDER_ID":         orderId,
		"CUST_ID":          customId,
		"TXN_AMOUNT":       amount,
	}

	checksum, err := paytm.GetChecksumFromArray(paramList)
	if err != nil {
		fmt.Println("err = ", err)
	} else {
		fmt.Println("checksum = ", checksum)
	}
}
func VerifyCheckum() {
	result := `{"ORDERID":"27000364888888", "MID":"Pay85623985963121", "TXNID":"20180710111212800110168868500018912", "TXNAMOUNT":"1.00", "PAYMENTMODE":"DC", "CURRENCY":"INR", "TXNDATE":"2018-07-10 15:04:56.0", "STATUS":"TXN_SUCCESS", "RESPCODE":"01", "RESPMSG":"Txn Success", "GATEWAYNAME":"HDFC", "BANKTXNID":"4036217121962950", "BANKNAME":"HDFC Bank", "CHECKSUMHASH":"TabnoADfqfWjI3twGIsjTRb97iDXlJSjq3S+fWOOtsz608mo+6JsAy600VZR/uimKR/46bdjrwgREQh4uF0L6IBeuhAhabyzUfJ5s2i5wps="}`
	resultList := map[string]string{}
	json.Unmarshal([]byte(result), &resultList)

	fmt.Println("resultList = ", resultList)

	ok := paytm.VerifyCheckum(resultList, resultList["CHECKSUMHASH"])
	fmt.Println("ok = ", ok)
}

func encrypt() {
	crypt, err := paytm.Encrypt([]byte("test"))
	checksum := base64.StdEncoding.EncodeToString(crypt)
	fmt.Println("err = ", err)
	fmt.Println("checksum = ", checksum)
}

func GetTransactionStatus()  {
	res, err := paytm.GetTransactionStatus("27000364888888", "TabnoADfqfWjI3twGIsjTRb97iDXlJSjq3S+fWOOtsz608mo+6JsAy600VZR/uimKR/46bdjrwgREQh4uF0L6IBeuhAhabyzUfJ5s2i5wps=")
	if err != nil {
		fmt.Println("err = ", err.Error())
		return
	}
	fmt.Println("res = ", res)
}
