package CryptoUtil

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io"
	"os"
)

type CryptoInterface interface {
	Decrypt()
	Encrypt()
	Sha256()
	Md5Sum()
}
type Crypto struct {
	key      string
	FilePath string
	Text     string
	Result   string
}

func NewKey() *Crypto {
	k := new(Crypto)
	k.key = "AnKoloft@~delNazok!12345" // key parameter must be 16, 24 or 32,
	return k
}

func (k *Crypto) Sha256() {
	h := sha256.New()
	h.Write([]byte(k.Text))
	k.Result = fmt.Sprintf("%x", h.Sum(nil))
}

func (k *Crypto) Md5Sum() {
	file, err := os.Open(k.FilePath)
	if err != nil {
		return
	}
	defer file.Close()
	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		return
	}
	k.Result = hex.EncodeToString(hash.Sum(nil))
}

func (k *Crypto) Encrypt() {
	key := []byte(k.key)
	plaintext := []byte(k.Text)
	c, err := aes.NewCipher(key)
	if err != nil {
		return

	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return

	}
	k.Result = base64.StdEncoding.EncodeToString(gcm.Seal(nonce, nonce, plaintext, nil))
}

func (k *Crypto) Decrypt() {
	key := []byte(k.key)
	bb, _ := base64.StdEncoding.DecodeString(k.Text)
	ciphertext := []byte(bb)
	c, err := aes.NewCipher(key)
	if err != nil {
		return
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return
	}

	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	t, e := gcm.Open(nil, nonce, ciphertext, nil)
	if e != nil {
		return
	}
	k.Result = string(t)

}
