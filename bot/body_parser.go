package bot

import (
	"encoding/json"
	"regexp"
	"strings"

	"go.uber.org/zap"
)

var Key string
var Sub_key string
var Num int

func Body_Parser(input string) ([]byte, string) {
	var str = ""
	r := regexp.MustCompile(`(\[[^\]]+\]|'[^']+'|\S+)`)
	inputs := r.FindAllString(input, -1)

	body_map := make(map[string]interface{})
	sub_map := make(map[string]interface{})

	for i := 0; i < len(inputs); i++ {
		if strings.HasPrefix(inputs[i], "-") {
			if i+1 >= len(inputs) {
				str = "Invalid JSON: JSON Cannot End With The Key And Only The Key"
				var full_body_json []byte
				return full_body_json, str
			}
			Key = strings.TrimPrefix(inputs[i], "-")
			Num = i + 1

			if Num < len(inputs) && strings.HasSuffix(inputs[Num], ":") {
				sub_map = make(map[string]interface{})
			} else if Num < len(inputs) {
				if strings.HasPrefix(inputs[Num], "[") && strings.HasSuffix(inputs[Num], "]") {
					body_map[Key] = strings.Trim(inputs[Num], "[]")
				} else if strings.Contains(inputs[Num], ",") && strings.Contains(inputs[Num], "'") {
					body_list := strings.Split(strings.Trim(inputs[Num], "'"), ", ")
					body_map[Key] = body_list
				} else {
					body_map[Key] = strings.Trim(inputs[Num], "', ")
				}
				i++
			}
		}

		if strings.HasSuffix(inputs[i], ":") {
			if i+1 >= len(inputs) {
				str = "Invalid JSON: JSON Cannot End With The SubKey And Only The SubKey"
				var full_body_json []byte
				return full_body_json, str
			}
			Sub_key = strings.TrimSuffix(inputs[i], ":")
			Num = i + 1

			if Num < len(inputs) {
				if strings.Contains(inputs[Num], "-") {
					continue
				} else if strings.HasPrefix(inputs[Num], "[") && strings.HasSuffix(inputs[Num], "]") {
					sub_map[Sub_key] = strings.Trim(inputs[Num], "[]")
				} else if strings.Contains(inputs[Num], ",") && strings.Contains(inputs[Num], "'") {
					body_list := strings.Split(strings.Trim(inputs[Num], "'"), ", ")
					sub_map[Sub_key] = body_list
				} else {
					sub_map[Sub_key] = strings.Trim(inputs[Num], "', ")
				}
				i++
			}
			body_map[Key] = sub_map
		}
	}
	full_body_json, err := json.Marshal(body_map)

	if err != nil {
		zap.L().Error("Error:", zap.Error(err))
	}

	return full_body_json, str
}
