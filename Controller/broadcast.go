package controller

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"log"
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

	messageRaw := r.Form.Get("message")

	signature := ComputeHmac256(messageRaw, goDotEnvVariable("BROADCAST_KEY"))

	postBody, _ := json.Marshal(map[string]string{
		"signature": signature,
		"message":   messageRaw,
	})
	responseBody := bytes.NewBuffer(postBody)

	resp, err := http.Post("http://159.138.122.103:50050/api/chat/system", "application/json", responseBody)
	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}
	defer resp.Body.Close()

	json.NewEncoder(w).Encode("Success")

}
