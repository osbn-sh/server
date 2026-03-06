package errMsgs

import (
	"fmt"
	"strings"
)

func TranslateErrorMessage(errmsg string) string {

	output := errmsg
	data := map[string]string{
		"(SQLSTATE 23505)": "این فیلد تکراری است",
	}

	for k, v := range data {
		fmt.Println(k, v)
		if strings.Contains(errmsg, k) {
			output = v

			break
		}
	}

	return output
}
