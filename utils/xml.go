package utils

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

//加载xml对象
func LoadXmlObject(path string, xmlPointer interface{}) bool {
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("error: %v", err)
		return false
	}
	defer func() {
		_ = file.Close()
	}()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("error: %v", err)
		return false
	}
	err = xml.Unmarshal(data, xmlPointer)
	if err != nil {
		fmt.Printf("error: %v", err)
		return false
	}
	return true
}
