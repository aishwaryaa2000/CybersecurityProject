package encrypt

import (
	"crypto"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha512"
	"encoding/hex"
)

func Encrypts(publicKey rsa.PublicKey, msg string) []byte {
	cipher, _ := rsa.EncryptOAEP(sha512.New(), rand.Reader, &publicKey, []byte(msg), nil)
	return cipher

}

func Decrypt(privateKey *rsa.PrivateKey, cipher []byte) []byte {
	plainText, _ := privateKey.Decrypt(nil, cipher, &rsa.OAEPOptions{Hash: crypto.SHA512})
	return plainText
}

func CreateHash(pwd string) string {
	hasher := md5.New()
	hasher.Write([]byte(pwd)) //storing key in slice of bytes
	return hex.EncodeToString(hasher.Sum(nil))
}
