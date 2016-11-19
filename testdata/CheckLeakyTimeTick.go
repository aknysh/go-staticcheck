package pkg

import "time"

func fn1() {
	for range time.Tick(0) {
		println("")
	}
}

func fn2() {
	for range time.Tick(0) { // MATCH /leaks the underlying ticker/
		println("")
		if true {
			break
		}
	}
}

func fn3() {
	for range time.Tick(0) { // MATCH /leaks the underlying ticker/
		println("")
		if true {
			return
		}
	}
}