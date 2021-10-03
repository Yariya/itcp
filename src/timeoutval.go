package main

import (
	"os"
	"strconv"
)

func ValidateTimeout() (int, int, int, error){
	var timeut 	int
	var stimeut int
	var thrs int

	for x, v := range os.Args {
		if v == "-timeout" {
			t, err := strconv.Atoi(os.Args[x+1])
			if err != nil {
				help(false)
			} else {
				stimeut=t
			}
		} else if v == "-ptime" {
			o, err := strconv.Atoi(os.Args[x+1])
			if err != nil {
				help(false)
			} else {
				timeut=o
			}
		} else if v == "-threads" {
			thr, err := strconv.Atoi(os.Args[x+1])
			if err != nil {
				help(false)
			} else {
				thrs=thr
			}
		}
	}
	return timeut, stimeut, thrs, nil
}
