package commands

import (
	"testing"
)

func TestGet_HelpNoResponseToHelpEndpoin(t *testing.T) {
	msg, str := Get_Help("http://localhost:3001/api/BAD_ENDPOINT")
	substring := "Help Endpoint Not found Either! Report This to An Admin"
	if substring != str {
		t.Errorf("\n\nError: Str Should Be Supplied With A String Value If Status Code Error For Help Endpoint Greater Than 399:\nStr Expected To Be: %q, Instead got %q\n\n", substring, str)
	}
	if len(msg) > 0 {
		t.Errorf("\n\nError: Msg Body Should Be Nil If Request To Endpoint Help Could Not Be Reached\nMsg Expected To Be Empty, Instead got %q\n\n", msg)
	}
}
func TestGet_HelpWithResponseToHelpEndpoint(t *testing.T) {
	msg, str := Get_Help("http://localhost:3001/api/help")
	substring := ""
	if substring != str {
		t.Errorf("\n\nError: Str Should Only Be Filled If HTTP Status Code Less Than 400:\nStr Expected To Be: %q, Instead got %q\n\n", substring, str)
	}
	if len(msg) < 1 {
		t.Errorf("\n\nError: Msg Body Should Be Contain The Help Response Data For A Microservice\nMsg Expected To Be ???, Instead got %q\n\n", msg)
	}
}
