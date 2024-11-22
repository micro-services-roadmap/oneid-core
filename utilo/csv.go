package utilo

import (
	"bufio"
	"encoding/csv"
	"errors"
	"os"
	"strconv"
)

type Numeric interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64
}

// ReadAsNestedMap 读取CSV文件并返回嵌套的Map
func ReadAsNestedMap[T Numeric](filePath string, separator rune, fn func(v string, dv T) T, dv T) (map[string]map[string]T, error) {
	nestedMap := make(map[string]map[string]T)

	file, err := os.Open(filePath)
	if err != nil {
		return nil, errors.New("Failed to open file: " + err.Error())
	}
	defer file.Close()

	reader := csv.NewReader(bufio.NewReader(file))
	reader.Comma = separator

	headers := make([]string, 0)
	for {
		record, err := reader.Read()
		if err != nil {
			break
		}

		if len(headers) == 0 {
			headers = record
		} else {
			innerMap := make(map[string]T)

			for i := 1; i < len(headers); i++ {
				key := headers[i]

				var value T
				if i < len(record) {
					value = fn(record[i], dv)
				} else {
					value = dv
				}
				innerMap[key] = value
			}

			outerKey := record[0]
			nestedMap[outerKey] = innerMap
		}
	}

	return nestedMap, nil
}

func Convert2Float32(str string, dv float32) float32 {
	float, err := strconv.ParseFloat(str, 32)
	if err != nil {
		return dv
	}
	return float32(float)
}
