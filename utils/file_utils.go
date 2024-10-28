package utils

import (
	"os"

	"github.com/360EntSecGroup-Skylar/excelize"
)

// 读取文件
func ReadFile(path string) ([]byte, error) {
	data, err := os.ReadFile(path)
	return data, err
}

// 读取字符串文件
func ReadStringFile(path string) (string, error) {
	data, err := ReadFile(path)
	if err == nil {
		return string(data), nil
	} else {
		return "", err
	}
}

// 读取Excel
func ReadExcel(excelPath string) ([]map[string]string, error) {
	var result []map[string]string
	var err error
	xlsxFile, openErr := excelize.OpenFile(excelPath)
	if openErr == nil {
		rows := xlsxFile.GetRows(xlsxFile.GetSheetName(xlsxFile.GetActiveSheetIndex()))
		var keyArray []string
		hasKey := false
		for _, row := range rows {
			if !hasKey {
				for _, colCell := range row {
					keyArray = append(keyArray, colCell)
				}
				hasKey = true
			} else {
				for index, colCell := range row {
					key := keyArray[index]
					dataMap := make(map[string]string)
					dataMap[key] = colCell
					result = append(result, dataMap)
				}
			}
		}
	} else {
		err = openErr
	}

	return result, err
}
