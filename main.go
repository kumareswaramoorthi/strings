package main

import (
	"bytes"
	"fmt"
	"strings"
)

var (
	Repeated string
	Previous string
	Pre      string
	First    bool = true
	Output   string
	i        int
)

func main() {
	// Scanning the Input
	fmt.Print("Enter text: ")
	var input string
	fmt.Scanln(&input)
	// Remove duplicates from input and store it in the 'noDupes' variable
	noDupes := removeDups(input)
	//repeatedChar function stores the repeated characters in 'Repeated' variable
	repeatedChar(input)
	// Mapper function maps the output
	mapper(noDupes, Repeated)

}

// Remove duplicates from input and store it in the 'noDupes' variable
func removeDups(s string) string {
	var buf bytes.Buffer
	var last rune
	for i, r := range s {
		if r != last || i == 0 {
			buf.WriteRune(r)
			last = r
		}
	}
	return buf.String()
}

//repeatedChar function stores the repeated characters in 'Repeated' variable
func repeatedChar(noDupes string) {
	for _, rune := range noDupes {
		v := fmt.Sprintf("%c", rune)
		if Pre == v {
			if strings.Count(noDupes, v) > 1 {
				Repeated = Repeated + fmt.Sprintf("%c", rune)
			}
		}
		// storing the v in Pre inorder to avoid the character present in correct sequence getting stored to 'Repeated' variable
		Pre = v
	}
}

// Mapper function maps the output
func mapper(noDupes, Repeated string) {
	// ranging noDupes
	for id, rune := range noDupes {
		v := fmt.Sprintf("%c", rune)
		Previous = v
		if Previous == Repeated[0:1] {
			for _, r := range Repeated {
				v := fmt.Sprintf("%c", r)
				// for first time iteration the append will be of different format
				if First {
					i = id + 2
					Output = noDupes[0:i] + v
				} else {
					// for the seconf time iteration the append will be of different format
					l := i + 1
					if len(noDupes) >= l {
						// appending the output in required format
						Output = Output + noDupes[i:l] + v
						i = i + 1
					}
				}
				// Setting the flag false for the non-first time iteration
				First = false
			}
			break
		}
	}
	// The lenght of string beyond the length of 'noDupes' needs to be appended to output
	parts := strings.Split(noDupes, Output[len(Output)-2:len(Output)-1])
	// Printing the Output
	fmt.Println("Output", Output+parts[1])
}
