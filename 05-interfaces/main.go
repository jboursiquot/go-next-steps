package main

import (
	"fmt"

	"github.com/jboursiquot/go-next-steps/05-interfaces/stringutils"
)

type uppercaser interface {
	Upper(string) string
}

type lowercaser interface {
	Lower(string) string
}

func main() {
	c := stringutils.Caser{}
	fmt.Println(upcase(c, "hello"))
	fmt.Println(lowcase(c, "hello"))
}

func upcase(u uppercaser, str string) string {
	return u.Upper(str)
}

func lowcase(l lowercaser, str string) string {
	return l.Lower(str)
}
