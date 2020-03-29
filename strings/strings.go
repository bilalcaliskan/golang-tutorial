package strings

import (
	"fmt"
	"unicode/utf8"
)

func printBytes(s string) {
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
	fmt.Printf("\n")
}

func printChars(s string) {
	for i:= 0; i < len(s); i++ {
		fmt.Printf("%c ",s[i])
	}
}

func printCharsWithRune(s string) {
	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		fmt.Printf("%c ", runes[i])
	}
}

func printCharsAndBytes(s string) {
	fmt.Println("string", s, "has length", len(s), "bytes")
	for index, rune := range s {
		fmt.Printf("%c starts at byte %d\n", rune, index)
	}
}

func getStringLength(s string) {
	fmt.Printf("length of %s is %d\n", s, utf8.RuneCountInString(s))
}

func mutate(s []rune) string {
	s[0] = 'a'
	return string(s)
}

func RunStrings() {
	fmt.Printf("\nBeginning of introduction to strings...\n")
	/*
	A string in Go is a slice of bytes.
	Strings in Go are Unicode compliant and are UTF-8 Encoded.
	 */

	fmt.Printf("\nBeginning of accessing individual bytes of a string...\n")
	/*
	Since a string is a slice of bytes, it's possible to access each byte of a string.
	 */
	name := "Hello World"
	printBytes(name)
	/*
	On above program, len(s) on function printBytes() returns the number of bytes in the string and we use a for loop to
	print those bytes in hexadecimal notation. %x is the format specifier for hexadecimal.
	These are the Unicode UT8-encoded values of "Hello World". A basic understanding of Unicode and UTF-8 is needed to
	understand strings better.
	*/
	name = "Señor"
	printBytes(name)
	fmt.Printf("\n")
	printChars(name)
	/*
	If we try to print the characters of Señor, it outputs S e Ã ± o r which is wrong. Why does this program break for
	Señor when it's perfectly alright with Hello World. The reason is that the Unicode code point of ñ is U+00F1 and its
	UTF-8 encoding occupies 2 bytes c3 and b1. We are trying to print characters assuming that each code point will be
	one byte long which is wrong. In UTF-8 encoding a code point can occupy more than 1 byte. So how do we solve this.
	This is where rune saves us.
	 */

	fmt.Printf("\nBeginning of rune...\n")
	/*
	A rune is a builtin type in Go and it's the alias of int32. rune represents a Unicode code point in Go. It does not
	matter how many bytes the code point occupies, it can be represented by a rune.
	 */
	name = "Señor"
	printBytes(name)
	fmt.Printf("\n")
	printCharsWithRune(name)
	fmt.Printf("\n")
	/*
	In the printCharsWithRune function, the string is converted to a slice of runes.
	 */

	fmt.Printf("\nBeginning of for range loop on a string...\n")
	name = "Señor"
	printCharsAndBytes(name)
	/*
	In the above program, the string is iterated using for range loop. The loop returns the position of the byte where
	the rune starts along with the rune.
	From the above output it's clear that ñ occupies 2 bytes :).
	*/

	fmt.Printf("\nBeginning of constructing string from slice of bytes...\n")
	byteSlice := []byte{0x43, 0x61, 0x66, 0xC3, 0xA9}
	str := string(byteSlice)
	fmt.Println(str)
	/*
	byteSlice in the program above contains the UTF-8 Encoded hex bytes of the string "Café". The program outputs Café.
	What if we have the decimal equivalent of hex values. Will the above program work? Lets check it out.
	 */
	byteSlice = []byte{67, 97, 102, 195, 169} // decimal equivalent of {'\x43', '\x61', '\x66', '\xC3', '\xA9'}
	str = string(byteSlice)
	fmt.Println(str)

	fmt.Printf("\nBeginning of constructing a string from slice of runes...\n")
	runeSlice := []rune{0x0053, 0x0065, 0x00f1, 0x006f, 0x0072}
	str = string(runeSlice)
	fmt.Println(str)
	/*
	In the above program runeSlice contains the Unicode code points of the string Señor in hexadecimal. The program
	outputs Señor.
	 */

	fmt.Printf("\nBeginning of length of the string...\n")
	/*
	The func RuneCountInString(s string) (n int) function of the utf8 package is used to find the length of the string.
	This method takes a string as argument and returns the number of runes in it.
	*/
	word1 := "Señor"
	getStringLength(word1)
	word2 := "Pets"
	getStringLength(word2)

	fmt.Printf("\nBeginning of strings are immutable...\n")
	/*
	Strings are immutable in Go. Once a string is created it's not possible to change it.
	To workaround this string immutability, strings are converted to a slice of runes. Then that slice is mutated
	with whatever changes needed and converted back to a new string.
	*/
	h := "hello"
	fmt.Println(mutate([]rune(h)))
	/*
	In the above program, the mutate function accepts a rune slice as argument. It then changes the first element of the
	slice to 'a', converts the rune back to string and returns it. h is converted to a slice of runes and passed to
	mutate. This program outputs aello
	 */
}