package contentparser

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
)

func Body_Parser(input string) []byte {
	var body_interface []interface{}

	var txt string
	var msg string

	r := regexp.MustCompile(`('[^']+'|\S+)`)
	inputs := r.FindAllString(input, -1)

	for i := 0; i < len(inputs); i++ {

		if strings.HasPrefix(inputs[i], "-") {
			if i+1 >= len(inputs) {
				txt = "Pre Microservice JSON Body Error"
				msg = "Invalid JSON: JSON Cannot End With The Key And Only The Key"
				fmt.Println(txt + "\n" + msg)
			}

			key := strings.TrimPrefix(inputs[i], "-")
			value := inputs[i+1]

			if strings.Contains(value, ",") && strings.Contains(value, "'") {

				body_list := strings.Split(strings.Trim(value, "'"), ", ")
				body_interface = append(body_interface, map[string]interface{}{key: body_list})

			} else {

				body_interface = append(body_interface, map[string]interface{}{key: strings.Trim(value, "', ")})

			}
			i++
		}
	}
	full_body := map[string]interface{}{"body": body_interface}

	full_body_json, err := json.Marshal(full_body)
	if err != nil {
		fmt.Println("Error:", err)
	}
	return full_body_json
}
