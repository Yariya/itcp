package main

import (
	"fmt"
	"os"
	"strings"
)

func CheckFlags() {
	if strings.Contains(fmt.Sprint(os.Args[1:]), "--de") {
		DeactivateFlag = true
	}
	if strings.Contains(fmt.Sprint(os.Args[1:]), "--notcp") {
		NoTCP = true
	}
	if strings.Contains(fmt.Sprint(os.Args[1:]), "--noicmp") {
		NoICMP = true
	}
}
