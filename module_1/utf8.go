package main

import (
	"fmt"
	"math"
)

func decode(a []byte) []rune {
	var s []rune
	for i := 0; i < len(a); i++ {
		if a[i] < byte(math.Pow(2, 7)) {
			s = append(s, rune(a[i]))
		} else {
			if a[i]&0x000000E0 == 192 {
				s = append(s, rune(a[i]-192)*0x40+rune(a[i+1]-128))
				i++
			} else {
				if a[i]&0x000000F0 == 224 {
					s = append(s, rune(a[i]-224)*0x1000+rune(a[i+1]-128)*0x40+rune(a[i+2]-128))
					i++
					i++
				} else {
					s = append(s, rune(a[i]-240)*262144+rune(a[i+1]-128)*0x1000+rune(a[i+2]-128)*0x40+rune(a[i+3]-128))
					i++
					i++
					i++
				}
			}
		}
	}
	return s
}

func encode(a []rune) []byte {
	var s []byte
	for i := 0; i < len(a); i++ {
		if a[i] < rune(math.Pow(2, 7)) {
			s = append(s, byte(a[i]))
		} else {
			if a[i] < rune(math.Pow(2, 11)) {
				s = append(s, 0xC0+byte(a[i]/0x40), 0x80+byte(a[i]%0x40))
			} else {
				if a[i] < rune(math.Pow(2, 16)) {
					s = append(s, 0xE0+byte(a[i]/0x1000), 0x80+byte((a[i]/0x40)%0x40), 0x80+byte(a[i]%0x40))
				} else {
					s = append(s, 0xF0+byte(a[i]/262144), 0x80+byte(a[i]/0x1000%0x40), 0x80+byte(a[i]/0x40%0x40), 0x80+byte(a[i]%0x40))
				}
			}
		}
	}
	return s
}
func main() {
	var primer1 = []rune{'Ф'}
	ans11 := encode(primer1)
	ans21 := decode(ans11)
	fmt.Print(ans11, ans21)

}
