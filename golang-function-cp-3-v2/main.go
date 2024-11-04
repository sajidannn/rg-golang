package main

import (
	"fmt"
	"strings"
)

func FindShortestName(names string) string {
	var namesSlice []string
	
	if strings.Contains(names, ";") {
		namesSlice = strings.Split(names, ";")
	} else if strings.Contains(names, ",") {
		namesSlice = strings.Split(names, ",")
	} else {
		namesSlice = strings.Split(names, " ")
	}

	shortestName := namesSlice[0]
	for _, name := range namesSlice {
		if len(name) < len(shortestName) {
			shortestName = name
		}
		if len(name) == len(shortestName) && name < shortestName {
			shortestName = name
		}
	}

	return shortestName 
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(FindShortestName("Hanif Joko Tio Andi Budi Caca Hamdan")) // "Tio"
	fmt.Println(FindShortestName("Budi;Tia;Tio"))                         // "Tia"
}
