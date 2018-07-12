package paytm

import (
	"net/http"
	"bytes"
	"fmt"
	"io/ioutil"
	"encoding/json"
)

type TransactionStatus struct {
	TXNID       string `json:"TXNID"`
	BANKTXNID   string `json:"BANKTXNID"`
	ORDERID     string `json:"ORDERID"`
	TXNAMOUNT   string `json:"TXNAMOUNT"`
	STATUS      string `json:"STATUS"`
	TXNTYPE     string `json:"TXNTYPE"`
	GATEWAYNAME string `json:"GATEWAYNAME"`
	RESPCODE    string `json:"RESPCODE"`
	RESPMSG     string `json:"RESPMSG"`
	BANKNAME    string `json:"BANKNAME"`
	MID         string `json:"MID"`
	PAYMENTMODE string `json:"PAYMENTMODE"`
	REFUNDAMT   string `json:"REFUNDAMT"`
	TXNDATE     string `json:"TXNDATE"`
}

func GetTransactionStatus(orderId string, checksum string) (success bool, err error) {
	var (
		req  *http.Request
		resp *http.Response
		body []byte
	)

	jsonStr := fmt.Sprintf(`{"MID":"%s","ORDERID":"%s","CHECKSUMHASH":"%s"}`, MID, orderId, checksum)

	req, err = http.NewRequest("POST", TransactionStatusAPI, bytes.NewBuffer([]byte(jsonStr)))
	if err != nil {
		return false, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Cache-Control", "no-cache")
	client := &http.Client{}
	resp, err = client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	var txnStatus TransactionStatus
	if err = json.Unmarshal(body, &txnStatus); err != nil {
		return false, err
	}
	if txnStatus.STATUS == "TXN_SUCCESS" {
		return true, nil
	}
	return false, err
}
