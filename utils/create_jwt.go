package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
)

type Header struct {
	Alg string `json:"alg"`
	Typ string `json:"typ"`
}

type Payload struct {
	Sub         string `json:"sub"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	IsShopOwner bool `json:"is_shop_owner"`
}

func CreateJwt(secret string, data Payload) (string, error) {

	header := Header{
		Alg: "HS256",
		Typ: "JWT",
	}

	headerByte, err := json.Marshal(header)

	if err != nil {
		return "", err
	}

	headerB64 := converToBase64(headerByte)

	payloadByte, err := json.Marshal(data)

	if err != nil {
		return "", err
	}

	payloadB64 := converToBase64(payloadByte)

	messageB64 := headerB64 + "." + payloadB64

	byteArrSecret := []byte(secret)
	byteArrMessage := []byte(messageB64)

	h := hmac.New(sha256.New, byteArrSecret)
	h.Write(byteArrMessage)

	signature := h.Sum(nil)
	signatureB64 := converToBase64(signature)

	jwt := headerB64 + "." + payloadB64 + "." + signatureB64

	return jwt , nil

}

func converToBase64(data []byte) string {
	return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(data)
}