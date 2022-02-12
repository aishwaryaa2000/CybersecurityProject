package encrypt

import (
	"crypto"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha512"
	"encoding/hex"
	"io"
)

//const secretkey string=""

func EncryptMail(publicKey rsa.PublicKey, msg string) []byte {
	cipher, _ := rsa.EncryptOAEP(sha512.New(), rand.Reader, &publicKey, []byte(msg), nil)
	return cipher
}

func DecryptMail(privateKey *rsa.PrivateKey, cipher []byte) []byte {
	plainText, _ := privateKey.Decrypt(nil, cipher, &rsa.OAEPOptions{Hash: crypto.SHA512})
	return plainText
}

func CreateHash(pwd string) string {
	hasher := md5.New()
	hasher.Write([]byte(pwd)) //storing key in slice of bytes
	return hex.EncodeToString(hasher.Sum(nil))
}

//Create func EncryptFile and DecryptFile

func EncryptFile(data []byte) []byte {
	key := CreateHash("Siddharth")
	block, _ := aes.NewCipher([]byte(key))
	gcm, _ := cipher.NewGCM(block)
	nonce := make([]byte, gcm.NonceSize())
	io.ReadFull(rand.Reader, nonce)
	cipherText := gcm.Seal(nonce, nonce, data, nil)
	return cipherText
}

func DecryptFile(data []byte) []byte {
	key := CreateHash("Siddharth")
	block, _ := aes.NewCipher([]byte(key))
	gcm, _ := cipher.NewGCM(block)
	nonce := data[:gcm.NonceSize()]
	cipherText := data[gcm.NonceSize():]
	plainText, _ := gcm.Open(nil, nonce, cipherText, nil)
	return plainText
}
