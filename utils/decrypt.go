package utils

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"net"
	"strconv"
	"strings"
	"time"
)

//异或加密
func XorEncrypt(data, key string) (string, bool) {
	if len(data) > 0 && len(key) > 0 {
		keyData := []byte(key)
		keyLength := len(keyData)
		byteData := []byte(data)
		encodeString := base64.StdEncoding.EncodeToString(byteData)
		encodingData := []byte(encodeString)
		var buffer bytes.Buffer
		num := 0
		for index := range encodingData {
			if num >= keyLength {
				num = num % keyLength
			}
			buffer.WriteByte(encodingData[index] ^ keyData[num])
			num += 1
		}
		encodeString = base64.StdEncoding.EncodeToString(buffer.Bytes())
		return encodeString, true
	}
	return "", false
}

//异或解密
func XorDecrypt(secret, key string) (string, bool) {
	keyData := []byte(key)
	keyLength := len(keyData)
	if keyLength > 0 {
		byteData, err := base64.StdEncoding.DecodeString(secret)
		if err == nil {
			var buffer bytes.Buffer
			num := 0
			for index := range byteData {
				if num >= keyLength {
					num = num % keyLength
				}
				buffer.WriteByte(byteData[index] ^ keyData[num])
				num += 1
			}
			byteData, err := base64.StdEncoding.DecodeString(buffer.String())
			if err == nil {
				var result bytes.Buffer
				result.Write(byteData)
				return result.String(), true
			} else {
				//fmt.Printf("XorDecrypt Error: %s\r\n", err.Error())
			}
		} else {
			//fmt.Printf("XorDecrypt Error: %s\r\n", err.Error())
		}
	}
	return "", false
}

//创建序列号
func CreateKey() string {
	interfaces, err := net.Interfaces()
	if err != nil {
		//panic("Poor soul, here is what you got: " + err.Error())
		return "error"
	}

	for _, inter := range interfaces {
		result, successful := XorEncrypt(inter.HardwareAddr.String(), "daiwei@aicyber.com")
		if successful {
			return result
		}
	}
	return ""
}

//验证序列号
func CheckKey(key string) bool {
	interfaces, err := net.Interfaces()
	if err != nil {
		//panic("Poor soul, here is what you got: " + err.Error())
		return false
	}

	for _, inter := range interfaces {
		result, successful := XorEncrypt(inter.HardwareAddr.String(), "daiwei@aicyber.com")
		if successful {
			result, successful = XorDecrypt(key, result)
			if successful {
				timeArray := strings.Split(result, ",")
				if len(timeArray) == 2 {
					start, _ := strconv.ParseInt(timeArray[0], 10, 64)
					end, _ := strconv.ParseInt(timeArray[1], 10, 64)
					current := time.Now().Unix()
					if current >= start && current <= end {
						return true
					}
				}
			}
			//if strings.Compare(key, currentKey) == 0 {
			//	return true
			//}
		}
	}
	return false
}

//添加时间信息
func AppendTimeInfo(key string, day int) string {
	day1 := 24 * 60 * 60
	start := time.Now().Unix()
	time := fmt.Sprintf("%d,%d", start, start+int64(day1*day))
	result, successful := XorEncrypt(time, key)
	if successful {
		return result
	}
	return "error"
}
