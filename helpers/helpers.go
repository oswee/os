package helpers

import (
	"net/url"
	"strings"
)

// PrettyLinks is function to make string lower case and replace spaces with dash.
func PrettyLinks(s string) string {
	lower := strings.ToLower(s)
	dash := url.PathEscape(strings.Replace(lower, " ", "-", -1))
	return dash
}

// CheckErr is just an temporary function
func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
