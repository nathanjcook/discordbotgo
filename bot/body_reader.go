package bot

import (
	"encoding/json"
	"fmt"
	"log"
)

func Body_Reader(body []byte) string {
	var outer string
	var inner string
	var objmap []map[string]interface{}
	if err := json.Unmarshal(body, &objmap); err != nil {
		log.Fatal(err)
	}
	for i := range objmap {
		for j := range objmap[i] {
			inner += fmt.Sprintf("%v", objmap[i][j]) + "\n\n"
		}
		outer += inner + "\n\n\n"
		inner = ""
	}
	message := outer
	return message
}
