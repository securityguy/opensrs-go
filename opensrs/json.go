package opensrs

import "encoding/json"

//goland:noinspection GoUnusedFunction
func prettyJSON(v interface{}) string {
	j, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return ""
	}
	return string(j)
}
