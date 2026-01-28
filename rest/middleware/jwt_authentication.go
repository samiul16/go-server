package middleware

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"
)

func Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, request *http.Request) {
fmt.Println("checking request header", request.Header.Get("Authorization"))

	authHeader := strings.Split(request.Header.Get("Authorization"), " ")
	fmt.Println("authHeader", authHeader)

	if len(authHeader) != 2 || authHeader[0] != "Bearer" {
		http.Error(w, "Unauthorized", 401)
		return
	}
	
	token := authHeader[1]
	fmt.Println("token", token)

	tokenParts := strings.Split(token, ".")
	fmt.Println("tokenParts", tokenParts)

	message := tokenParts[0] + "." + tokenParts[1]
	fmt.Println("message", message)

	requestSignature := tokenParts[2]
	fmt.Println("signature", requestSignature)

	byteArrSecret := []byte("my-secret")
	byteArrMessage := []byte(message)

	h := hmac.New(sha256.New, byteArrSecret)
	h.Write(byteArrMessage)

	preparedSignature := h.Sum(nil)
	preparedSignatureB64 := converToBase64(preparedSignature)

	if preparedSignatureB64 != requestSignature {
		http.Error(w, "Unauthorized", 401)
		return
	}
		
	next.ServeHTTP(w, request)
	})
}


func converToBase64(data []byte) string {
	return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(data)
}
