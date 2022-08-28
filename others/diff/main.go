package main

import (
	"fmt"
	"strings"

	"github.com/kylelemons/godebug/diff"
)

func main() {

	a := `
 	[Ahmed] [18/10/2022 - 18:52:23]
	Q.25 x 5 DH = 125 DH
	Q.2 x 2 DH = 4 DH
	Q.3 x 23.5 DH = 70.5 DH
	`
	b := `
	[Brahim] [18/10/2022 - 20:01:08]
	Q.25 x 5 DH = 125 DH
	Q.2 x 2 DH = 4 DH
	Q.2 x 6 DH = 12 DH
	Q.1 x 10 DH = 10 DH
	`

	changes := diff.Diff(a, b)
	fmt.Println(changes)

	data := colorize(changes)

	for _, v := range data {
		fmt.Println(v)
	}
}

const (
	ADD    = 1
	DEL    = -1
	NORMAL = 0
)

func css(s string, style int) string {
	switch style {
	case NORMAL:
		return "<div style='background-color:#f2f2f2; color:#292825; border-radius:5px; padding:10px'>" + s + "</div>"
	case ADD:
		return "<div style='background-color:#a1de64; color:white; border-radius:5px; padding:10px'>" + s + "</div>"
	case DEL:
		return "<div style='background-color:#db2143; color:white; border-radius:5px; padding:10px'>" + s + "</div>"

	}
	return s
}

func colorize(changes string) []string {
	slices := strings.Split(strings.TrimSpace(changes), "\n")

	for i, v := range slices {
		switch v[0] {
		case '+':
			slices[i] = css(v, ADD)

		case '-':
			slices[i] = css(v, DEL)

		default:
			slices[i] = css(v, NORMAL)
		}
	}
	return slices
}
