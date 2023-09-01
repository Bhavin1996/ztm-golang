package main

import "unicode"

func main() {
	data := []rune{'a', 'b', 'c', 'd'}
	var capitilezd []rune
	capIt := func(r rune) {
		capitilezd = append(capitilezd, unicode.ToUpper(r))
	}

}
