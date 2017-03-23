package security

import (
	"bytes"
	"crypto/md5"
	"crypto/rand"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"io"
	"time"

	"golang.org/x/crypto/bcrypt"
)

const PW_SALT_BYTES  = 32
const NONCE_BYTES = 16

type SecurityManager struct {
	secret string
	tokenLifetime int64
}

func NewSecurityManager(secret string, tokenLifetime int64) *SecurityManager {
	return &SecurityManager{
		secret: secret,
		tokenLifetime: tokenLifetime,
	}
}

func (m *SecurityManager) GenerateSalt() (string, error) {
	b, err := m.generateRandomBytes(PW_SALT_BYTES)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func (m *SecurityManager) generateRandomBytes(l int) ([]byte, error) {
	b := make([]byte, l)
	_, err := io.ReadFull(rand.Reader, b)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func (m *SecurityManager) GetMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

func (m *SecurityManager) GetPasswordHash(password string, salt string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password + salt + m.secret), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(hash), err
}

func (m *SecurityManager) GenerateNewToken(input string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(input + m.secret + time.Now().Format("2006-01-02T15:04:05Z")), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(hash), nil
}

func (m *SecurityManager) CreateWSSEToken(username string, password string) *WSSEToken {
	nonce, _ := m.generateRandomBytes(NONCE_BYTES)
	created := time.Now()
	return &WSSEToken{
		Username: username,
		Nonce: nonce,
		Created: time.Now(),
		PasswordDigit: m.createPasswordDigest(string(nonce), created.Format("2006-01-02T15:04:05Z"), password),
	}
}

func (m *SecurityManager) createPasswordDigest(nonce, created, password string) []byte {
	digest := sha1.New()
	digest.Write([]byte(nonce))
	digest.Write([]byte(created))
	digest.Write([]byte(password))
	return digest.Sum(nil)
}

func (m *SecurityManager) ValidateWSSEToken(token *WSSEToken, password string) error {
	now := time.Now()
	if token.Created.Add(time.Duration(m.tokenLifetime) * time.Second).Before(now) {
		return ErrWSSETokenExpired
	}
	digit := m.createPasswordDigest(string(token.Nonce), token.Created.Format("2006-01-02T15:04:05Z"), password)
	//fmt.Printf("Token: %s\nDigit: %s\n", token.PasswordDigit, digit)
	if bytes.Compare(token.PasswordDigit, digit) == 0 {
		return nil
	}
	return ErrWSSETokenInvalid
}

func (m *SecurityManager) ValidatePassword(passwordHash string, salt string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password + salt + m.secret))
}