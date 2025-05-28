package helper

import "strings"

func SplitIdPriceToday(id string) (string, string, string) {
	Id := strings.Split(id, "-")
	return Id[0], Id[1], Id[2]
}
