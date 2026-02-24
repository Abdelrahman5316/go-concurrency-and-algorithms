package concurrency

import (
	"math"
	"strings"
	"unicode/utf8"
)

// This function should output 0^2 + 1^2 + 2^2 + ... + (|n|)^2
func getSumSquares(n int) int {

	var sum float64
	var n_2 float64 = math.Abs(float64(n))
	for i := 0; i <= int(n_2); i++ {
		res := math.Pow(float64(i), 2)
		sum = sum + res
	}
	return int(sum)
}

// This function extracts all words ending with endLetter from the string text
func getWords(text string, endLetter rune) []string {

	fields := strings.Fields(text)
	txt := []string{}

	for _, field := range fields {

		for i, w := range field {

			if string(endLetter) == string(w) && i == utf8.RuneCountInString(field)-1 {

				txt = append(txt, field)

			}

		}

	}

	return txt
}

type RegRecord struct {
	studentId  int
	courseName string
}

// This method receives a list of student registration records and should return
// a map that shows the number of students registered per course.

func getCourseInfo(records []RegRecord) map[string]int {

	m := map[RegRecord]bool{}
	res := map[string]int{}
	for _, r := range records {
		m[r] = true
	}
	for rec, _ := range m {

		res[rec.courseName]++
	}
	return res
}

// Task 4
// This method is required to count the occurrences of an input key in a list of integers.

func count(list []int, key int, numThreads int) int {

	inputChan := make(chan int)
	outChan := make(chan int)
	for i := 0; i < numThreads; i++ {
		go countWorker(key, inputChan, outChan)
	}
	for _, data := range list {
		inputChan <- data
	}
	close(inputChan)
	count := 0
	for i := 0; i < numThreads; i++ {
		count = count + <-outChan
	}
	return count
}

// This worker method receives inputs via inputChan, and outputs the number of occurrences to outChan

func countWorker(key int, inputChan chan int, outChan chan int) {
	// To Do
	occurances := 0
	for num := range inputChan {
		if num == key {
			occurances++
		}
	}
	outChan <- occurances

}
