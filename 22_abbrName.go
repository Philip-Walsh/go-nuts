package kata

import (
	"strings"
)

func AbbrevName(name string) (res string) {
	for _, i := range strings.Split(name, " ") {
		res += i[:1]
	}
	return strings.Join(strings.Split(strings.ToUpper(res), ""), ".")
}
