package tools

import (
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"net"
	"strconv"
	"time"
)

func UserKey(key string) (oneTimeKey string, rest int) {
	obj := sha256.New()
	timeUnix := time.Now().Unix()
	obj.Write([]byte(strconv.Itoa(int(timeUnix / 60))))
	obj.Write([]byte(key))
	temp := obj.Sum(nil)
	oneTimeKeyByte := make([]byte, 6)
	for i := 0; i < 6; i++ {
		oneTimeKeyByte[i] = temp[i]%10 + '0'
	}
	return string(oneTimeKeyByte), int(60 - timeUnix%60)
}

func CreateUserKey(username, password string) (userKey string) {
	obj := sha512.New()
	obj.Write([]byte(Sha256(username))[:20])
	obj.Write([]byte(Sha256(password)))
	obj.Write([]byte(Sha256(strconv.Itoa(int(time.Now().UnixNano())))))
	addr, _ := net.InterfaceAddrs()
	network := "0.0.0.0"
	if len(addr) != 0 {
		network = addr[0].Network()
	}
	obj.Write([]byte(Sha256(network)))
	return hex.EncodeToString(obj.Sum(nil))
}
