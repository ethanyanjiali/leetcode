package main

import (
	"fmt"
	"strings"
)

func getDestination(src string) string {
	dst := strings.Replace(src, ".", "", -1)
	dst = strings.Split(dst, "+")[0]
	return dst
}

func numUniqueEmails(emails []string) int {
	dstMap := map[string]bool{}
	for _, email := range emails {
		parts := strings.Split(email, "@")
		domain := parts[1]
		local := parts[0]
		dst := getDestination(local) + "@" + domain
		dstMap[dst] = true
	}
	fmt.Printf("%v", dstMap)
	return len(dstMap)
}

func main() {
	emails := []string{"test.email+alex@leetcode.com", "test.e.mail+bob.cathy@leetcode.com", "testemail+david@lee.tcode.com"}
	fmt.Println(numUniqueEmails(emails))
}
