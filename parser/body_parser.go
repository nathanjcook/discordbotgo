package parser

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
)

func Body_Parse() {
	var body_interface []interface{}
	input := "'text_without_key_test' -input_key 'text_with_key_test' 'array_test1, array_test2, array_test3' -array_key 'array_test4, array_test5, array_test6'"

	r := regexp.MustCompile(`('[^']+'|\S+)`)
	inputs := r.FindAllString(input, -1)

	for i := 0; i < len(inputs); i++ {

		if strings.HasPrefix(inputs[i], "-") && i+1 < len(inputs) {

			key := strings.TrimPrefix(inputs[i], "-")
			value := inputs[i+1]

			if strings.Contains(value, ",") && strings.Contains(value, "'") {

				body_list := strings.Split(strings.Trim(value, "'"), ", ")
				body_interface = append(body_interface, map[string]interface{}{key: body_list})

			} else {

				body_interface = append(body_interface, map[string]interface{}{key: strings.Trim(value, "', ")})

			}
			i++
		} else {

			if strings.Contains(inputs[i], "'") && strings.Contains(inputs[i], ",") {

				body_list := strings.Split(strings.Trim(inputs[i], "'"), ", ")
				body_interface = append(body_interface, body_list)

			} else {

				body_interface = append(body_interface, inputs[i])

			}
		}
	}

	full_body := map[string]interface{}{"body": body_interface}
	full_body_json, err := json.Marshal(full_body)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(string(full_body_json))
}
