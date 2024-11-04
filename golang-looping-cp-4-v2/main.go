package main

import "fmt"

func EmailInfo(email string) string {
	domain := ""
	tld := ""
	isDomain := false
	isTld := false
	for _, char := range email {
		if char == '@' {
			isDomain = true
			continue
		}

		if char == '.' && isDomain{
			isTld = true
			isDomain = false
			continue
		}

		if isDomain {
			domain += string(char)
		} else if isTld {
			tld += string(char)
		} 
	}
	return fmt.Sprintf("Domain: %s dan TLD: %s", domain, tld)
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(EmailInfo("admin@yahoo.co.id"))
}
