package paytm

import (
	"strings"
	"sort"
	"math/rand"
	"time"
	"crypto/sha256"
	"fmt"
	"encoding/base64"
)

func GetChecksumFromArray(paramsMap map[string]string) (checksum string, err error) {
	var keys = make([]string, 0, 0)
	for k, v := range paramsMap {
		if v != "" {
			keys = append(keys, k)
		}
	}
	sort.Strings(keys)
	var arrayList = make([]string, 0, 0)
	for _, key := range keys {
		if value, ok := paramsMap[key]; ok && value != "" {
			arrayList = append(arrayList, value)
		}
	}
	arrayStr := getArray2Str(arrayList)
	salt := generateSalt(4)
	finalString := arrayStr + "|" + salt
	hash := fmt.Sprintf("%x", sha256.Sum256([]byte(finalString)))
	hashString := hash + salt
	crypt, err := Encrypt([]byte(hashString))
	if err != nil {
		return
	}
	checksum = base64.StdEncoding.EncodeToString(crypt)
	return
}

func VerifyCheckum(paramsMap map[string]string, checksum string) (ok bool) {
	delete(paramsMap, "CHECKSUMHASH")
	var keys = make([]string, 0, 0)
	for k, v := range paramsMap {
		if v != "" {
			keys = append(keys, k)
		}
	}
	sort.Strings(keys)
	var arrayList = make([]string, 0, 0)
	for _, key := range keys {
		if value, ok := paramsMap[key]; ok && value != "" {
			arrayList = append(arrayList, value)
		}
	}
	arrayStr := getArray2StrForVerify(arrayList)
	cs, err := base64.StdEncoding.DecodeString(checksum)
	if err != nil {
		fmt.Printf("base64 DecodeString err [%v]\n", err)
		return
	}
	paytmHash, err := Decrypt(cs)
	if err != nil {
		fmt.Printf("Decrypt err [%v]\n", err)
		return
	}
	paytmHashStr := string(paytmHash)
	salt := paytmHashStr[len(paytmHashStr)-4:]
	finalString := arrayStr + "|" + salt
	h := sha256.New()
	h.Write([]byte(finalString))
	finalStringHash := fmt.Sprintf("%x", h.Sum(nil))
	websiteHashStr := finalStringHash + salt
	if websiteHashStr == paytmHashStr {
		return true
	}
	return false
}

func Encrypt(input []byte) (output []byte, err error) {
	iv := "@@@@&&&&####$$$$"
	crypter, _ := NewCrypter([]byte(PaytmMerchantKey), []byte(iv))
	output, err = crypter.Encrypt(input)
	return
}

func Decrypt(input []byte) (output []byte, err error) {
	iv := "@@@@&&&&####$$$$"
	crypter, err := NewCrypter([]byte(PaytmMerchantKey), []byte(iv))
	output, err = crypter.Decrypt(input)
	return
}

func getArray2Str(arrayList []string) (str string) {
	findme := "REFUND"
	findmepipe := "|"
	flag := 1
	for _, v := range arrayList {
		pos := strings.Index(v, findme)
		pospipe := strings.Index(v, findmepipe)
		if pos != -1 || pospipe != -1 {
			continue
		}
		if flag > 0 {
			str += strings.TrimSpace(v)
			flag = 0
		} else {
			str += "|" + strings.TrimSpace(v)
		}
	}
	return
}

func getArray2StrForVerify(arrayList []string) (str string) {
	flag := 1
	for _, v := range arrayList {
		if flag > 0 {
			str += strings.TrimSpace(v)
			flag = 0
		} else {
			str += "|" + strings.TrimSpace(v)
		}
	}
	return
}

func generateSalt(length int) (salt string) {
	rand.Seed(time.Now().UnixNano())
	data := "AbcDE123IJKLMN67QRSTUVWXYZ"
	data += "aBCdefghijklmn123opq45rs67tuv89wxyz"
	data += "0FGH45OP89"
	for i := 0; i < length; i++ {
		salt += string(data[int(rand.Int()%len(data))])
	}
	return
}
