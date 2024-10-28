package main

import (
	"devil-tools/utils"
	"fmt"
)

func main() {
	dataMap := make(map[string]interface{})
	dataMap["test"] = "value"
	flag := dataMap["flag"]
	if flag == nil {
		flag = 0
	}
	fmt.Printf("......value: %d\r\n", flag)


	//interfaces, err :=  net.Interfaces()
	//if err != nil {
	//	panic("Poor soul, here is what you got: " + err.Error())
	//}

	//for _, inter := range interfaces {
	//	result, successful := utils.XorEncrypt(inter.HardwareAddr.String(), "daiwei@aicyber.com")
	//	if successful{
	//		fmt.Println(inter.Name, inter.HardwareAddr, result)
	//	}
	//}

	//fmt.Println(utils.CheckKey("KgwkQShbCVczNyxUKB93VTU5A1cnMwJU"))

	//fmt.Println(utils.AppendTimeInfo("KgwkQShbCVczNyxUKB93VTU5A1cnMwJU", 1))
	fmt.Println(utils.CheckKey("BjMiXh8pA1YNPCIDAAAPLQUWUgcbEDZMD3UqXQ=="))
	//r, successful := utils.XorEncrypt("hello world", "abc")
	//if successful {
	//	fmt.Println(len(r))
	//	fmt.Println(r)
	//}
	//
	//r, successful = utils.XorDecrypt("ACU1EgBQAgUHU1saAyUyXA==", "abc")
	//fmt.Println(successful)
	//if successful {
	//	fmt.Println(r)
	//}


}
