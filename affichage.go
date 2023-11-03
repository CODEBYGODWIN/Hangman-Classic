package pendu

import (
	"fmt"
	"os"
)

func Display(s []string) {
	final := ""
	for _, i := range s {
		final += i
	}
	fmt.Println(final)
}

func DisplayHangman(hp int) {
	b,_ := os.ReadFile("hangman.txt")
	word := ""
	line := 0
	result := []string{}
	for _, i := range b {
		if line == 8 {
			line = 0
			result = append(result, word)
			word = ""
		}
		if string(i) == "\n" {
			line++
		}
		word += string(i)
	}
	
	print(append(result, word), hp)
}

func DisplayGuillo(hp int) {
	b,_ := os.ReadFile("guillotine.txt")
	word := ""
	ligne := 0
	result := []string{}
	for _, i := range b {
		if ligne == 9 {
			ligne = 0
			result = append(result, word)
			word = ""
		}
		if string(i) == "\n" {
			ligne++
		}
		word += string(i)
	}
	
	print(append(result, word), hp)	
}

func DisplayLetterUsed(s []string) {
	lettre := ""
	for _, i := range s {
		lettre += " " + i
	}
	fmt.Println(lettre)
}

func print(s []string, hp int) {
	if hp < 10 {
		fmt.Println(s[9-hp])
	}
}

func List_word(lWordFull []string, letters []string) []string {
	lWordEmpty := make([]string, len(lWordFull))
	for iEmpty := 0; iEmpty < len(lWordEmpty); iEmpty++ {
		lWordEmpty[iEmpty] = "_"
	}
	for letter := 0; letter < len(letters); letter++ {
		for index := 0; index < len(lWordFull); index++ {
			if letters[letter] == lWordFull[index] {
				lWordEmpty[index] = letters[letter]
			}
		}
	}
	return lWordEmpty
}