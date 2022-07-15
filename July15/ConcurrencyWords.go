package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

var printNow chan bool
var wg sync.WaitGroup

func count(s []string) {

	if _, ok := <-printNow; ok {

		fmt.Println(len(s))
	}
	wg.Done()
}
func revese(s []string) {

	for i, x := 0, len(s)-1; i < x; i, x = i+1, j-1 {
		s[i], s[x] = s[x], s[i]
	}

	fmt.Println(s)
	printNow <- true
	time.Sleep(time.Millisecond * 45)
	wg.Done()

}

func main() {
	p := "Book One is geared toward teaching students how to write an instructional paragraph using the FNTF formula: First, Next, Then, Finally. The book first describes the How-to paragraph, gives examples, and has students dissect paragraphs using multiple choice and short answer questions. Once students are familiar with the purpose and style of a How-to paragraph, they make their own using provided prompts. Students then learn how to edit and format paragraphs with instruction given on margins, editing marks, and more. Students then edit sample paragraphs. The book then provides various lessons for using correct capitalization, punctuation, subjects, and verbs. Thought content of the paragraphs is then covered and students gain practice in reviewing if paragraphs make sense, correcting ones that do not, and making more of their own paragraphs. Once students have become familiar with reading, writing, and editing paragraphs, the last chapter of the book instructs students on how to write a whole How-to essay using the same step-by-step process they have learned for paragraphs."

	v := strings.Fields(p)
	var flag bool
	var temp []string
	for i := 0; i < len(v); i++ {
		if flag {
			wg.Add(2)
			printNow = make(chan bool)
			go revese(temp)
			go count(temp)
			wg.Wait()
			temp = nil

		}
		flag = false
		y := strings.Split(v[i], "")
		for x := 0; x < len(y); x++ {

			if x == len(y)-1 {

				if y[x] == "." {

					flag = true
				}
				temp = append(temp, strings.Join(strings.Fields(v[i]), ""))
			}

		}
	}

}
