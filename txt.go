package pendu

import (
	"time"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

func Words(word string) []string {

	list := []string{}
	CompletA := []string{}
	CompletB := ""

	for _, letter := range word {
		list = append(list, string(letter))
	}

	for b := 0; b < len(list); b++ {
		CompletB = CompletB + string(list[b])
	}

	for _, a := range CompletB {
		CompletA = append(CompletA, string(a))
	}

	return CompletA

}


func SelectWord(s string) []string {
	content, err := os.ReadFile(s)
	if err != nil {
		fmt.Println("file error")
		return []string(nil)
	}
	list := byteToString(content)
	mot := list[rand.Intn(len(list)-1)]
	run := []rune(mot)
	str := []string{}
	for i := 0; i < len(mot); i++ {
		str = append(str, string(run[i]))
	}
	return str
}

func byteToString(b []byte) []string {
	word := ""
	result := []string{}
	for _, i := range b {
		if string(i) == "\n" {
			result = append(result, word)
			word = ""
		} else {
			word += string(i)
		}
	}
	return result
}

func RandomWord(file []byte) string {
	words := strings.Split(string(file), "\n")
	rand.Seed(time.Now().UnixNano())
	word := rand.Intn(len(words))
	return words[word]
}


func RandomLetters(word string) []string {
	listWord := []string{}
	listReply := make([]string, len(word)/2-1)
	for _, strWord := range word {
		listWord = append(listWord, string(strWord))
	}
	if len(listWord) < 4 {
		listReply = append(listReply, "")
		return listReply
	} else {
		for i := 0; i < len(listWord)/2-1; i++ {
			rand.Seed(time.Now().UnixNano())
			letter := rand.Intn(len(listWord))
			letterAppend := listWord[letter]
			test := 1
			for j := 0; j < len(listReply); j++ {
				if letterAppend == listReply[j] {
					test = 0
				}
			}
			if test == 1 {
				listReply = append(listReply, letterAppend)
			} else {
				i -= 1
			}
		}
	}
	return listReply
}
