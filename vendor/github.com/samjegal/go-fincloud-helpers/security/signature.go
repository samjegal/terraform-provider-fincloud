package security

import (
	"crypto"
	"crypto/hmac"
	"encoding/base64"
	"net/url"
)

type HMACSecret struct {
	secretKey string
	hashFunc  crypto.Hash
}

func NewSignature(secretKey string, hashFunc crypto.Hash) *HMACSecret {
	return &HMACSecret{
		secretKey: secretKey,
		hashFunc:  hashFunc,
	}
}

func (s *HMACSecret) Signature(method, path, accessKey, timestamp string) (string, error) {
	url, err := url.Parse(path)
	if err != nil {
		return "", err
	}

	h := hmac.New(s.HashFunc().New, []byte(s.secretKey))
	h.Write([]byte(method))
	h.Write([]byte(" "))
	h.Write([]byte(url.RequestURI()))
	h.Write([]byte("\n"))
	h.Write([]byte(timestamp))
	h.Write([]byte("\n"))
	h.Write([]byte(accessKey))
	raw := h.Sum(nil)

	base64signature := base64.StdEncoding.EncodeToString(raw)
	return base64signature, nil
}

func (s *HMACSecret) HashFunc() crypto.Hash {
	return s.hashFunc
}
