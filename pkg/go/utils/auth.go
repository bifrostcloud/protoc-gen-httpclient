package utils

import "encoding/base64"

// BasicAuth - returns
func BasicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}
