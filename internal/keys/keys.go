package keys

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"net/http"
	"os"
)

var publicKey any
var privateKey *rsa.PrivateKey

func init() {
	publicKeyPEM, err := os.ReadFile("public.pem")
	if err != nil {
		panic(err)
	}

	publicKeyBlock, _ := pem.Decode(publicKeyPEM)
	publicKey, err = x509.ParsePKIXPublicKey(publicKeyBlock.Bytes)
	if err != nil {
		panic(err)
	}
	

	privateKeyPEM, err := os.ReadFile("private.pem")
	if err != nil {
		panic(err)
	}

	privateKeyBlock, _ := pem.Decode(privateKeyPEM)
	privateKey, err = x509.ParsePKCS1PrivateKey(privateKeyBlock.Bytes)
	if err != nil {
		panic(err)
	}
	
}


func GetPublicKey(w http.ResponseWriter, r *http.Request) {
	publicKeyASN1, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		http.Error(w, "failed to encode key", http.StatusInternalServerError)
		return
	}


	publicKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type: "PUBLIC KEY",
		Bytes: publicKeyASN1,
	})

	resp := map[string]string{
		"publicKey": string(publicKeyPEM),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func DecryptMessage(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Encrypted string `json:"encrypted"`
	}
	json.NewDecoder(r.Body).Decode(&req)

	cipherText, err := base64.StdEncoding.DecodeString(req.Encrypted)
	if err != nil {
		http.Error(w, "invalid ciphertext", http.StatusBadRequest)
		return
	}

	plaintext, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, cipherText)
	if err != nil {
		http.Error(w, "decryption failed", http.StatusInternalServerError)
		return
	}

	fmt.Println(string(plaintext))
}
