package security

import (
	"errors"
	"fmt"
	"log"
	"regexp"
	"time"
	"encoding/base64"
)

var (
	ErrWSSEHeaderInvalid  = errors.New("WSSE header has invalid format")
	ErrWSSETokenExpired   = errors.New("WSSE Token expired")
	ErrWSSETokenInvalid   = errors.New("WSSE token is invalid")
)

type WSSEToken struct {
	Username      string
	Nonce         []byte
	Created       time.Time
	PasswordDigit []byte
}


func (t *WSSEToken) ToString() string {
	return fmt.Sprintf(
		`UsernameToken Username="%s", PasswordDigest="%s", Nonce="%s", Created="%s"`,
		t.Username,
		base64.StdEncoding.EncodeToString(t.PasswordDigit),
		base64.StdEncoding.EncodeToString(t.Nonce),
		t.Created.Format("2006-01-02T15:04:05Z"),
	)
}

func ParseToken(input string) (*WSSEToken, error) {
	r, err := regexp.Compile(`UsernameToken Username="([^"]+)", PasswordDigest="([^"]+)", Nonce="([a-zA-Z0-9+\/]+={0,2})", Created="([^"]+)"`)
	if err != nil {
		log.Panic(err)
	}
	res := r.FindStringSubmatch(input)
	if len(res) != 5 {
		return nil, ErrWSSEHeaderInvalid
	}
	t, err := time.Parse("2006-01-02T15:04:05Z", res[4])
	if err != nil {
		return nil, err
	}
	digit, _ := base64.StdEncoding.DecodeString(res[2])
	nonce, _ := base64.StdEncoding.DecodeString(res[3])
	return &WSSEToken{
		Username: res[1],
		PasswordDigit: digit,
		Nonce: nonce,
		Created: t,
	}, nil
}

