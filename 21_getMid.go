package main

func GetMiddle(s string) string {
	var mid = int(len(s) / 2)
	if len(s)%2 > 0 {
		return s[mid : mid+1]
	}
	return s[mid-1 : mid+1]
}
