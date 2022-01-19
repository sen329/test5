package controller

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"net/http"
)

func ComputeHmac256(message string, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(message))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func BroadcastChat(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	url := "http://159.138.122.103:50050/api/chat/system"

	messageRaw := r.Form.Get("message")

	signature := ComputeHmac256(messageRaw, goDotEnvVariable("BROADCAST_KEY"))

	postBody, _ := json.Marshal(map[string]string{
		"signature": signature,
		"message":   messageRaw,
	})
	responseBody := bytes.NewBuffer(postBody)

	resp, err := http.Post(url, "application/json", responseBody)
	if err != nil {
		panic(err.Error())
	}
	defer resp.Body.Close()

	// req, err := http.NewRequest("POST", url, responseBody)
	// if err != nil {
	// 	panic(err)
	// }

	json.NewEncoder(w).Encode("Success")

}
