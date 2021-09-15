package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	domains := []string{"900 google.mail.com", "50 yahoo.com", "1 intel.mail.com", "5 wiki.org"}
	subDomainVisits(domains)
	lang := "Golang"
	fmt.Println(&lang)
}

func subDomainVisits(domains []string) {
	countMap := make(map[string]int)
	for _, d := range domains {
		fields := strings.Fields(d)
		count, _ := strconv.Atoi(fields[0])
		domain := fields[1]
		if c, ok:=countMap[domain]; !ok {
			countMap[domain] = count
		} else {
			countMap[domain] = c + count
		}
		domainParts := strings.Split(domain, ".")
		var subDomain1 string
		if len(domainParts) > 2 {
			subDomain1 = domainParts[2]
			subDomain2 := domainParts[1]
			subDomain := subDomain2 + "." + subDomain1
			if c, ok:=countMap[subDomain]; !ok {
				countMap[subDomain] = count
			} else {
				countMap[subDomain] = c + count
			}
		} else {
			subDomain1 = domainParts[1]
		}
		if c, ok:=countMap[subDomain1]; !ok {
			countMap[subDomain1] = count
		} else {
			countMap[subDomain1] = c + count
		}
	}
	fmt.Println(countMap)
}