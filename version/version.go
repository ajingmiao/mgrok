package version

import (
	"fmt"
)

const (
	Major = "1"
	Minor = "3"
	Proto = "0"
)

func MajorMinor() string {
	return fmt.Sprintf("%s.%s", Major, Minor)
}

func Full() string {
	return fmt.Sprintf("%s-%s.%s", Proto, Major, Minor)
}

func Compat(client string, server string) bool {
	return client == server
}
