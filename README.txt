Go Concurrency and Algorithms

Overview This repository contains a Go package implementing several
fundamental programming tasks, including mathematical computation,
string processing with Unicode support, data aggregation, and parallel
counting using goroutines and channels.

Project Structure - concurrency&algorithms.go : Implementation of all tasks - concurrency&algorithms_test.go
: Unit tests - list.txt : Input data file for concurrency test - go.mod
: Module definition

Implemented Tasks

1.  Sum of Squares Computes: 0^2 + 1^2 + 2^2 + … + |n|^2

Function: getSumSquares(n int) int

2.  Word Extraction Extracts all words from a string that end with a
    specified rune (Unicode-safe).

Function: getWords(text string, endLetter rune) []string

3.  Course Registration Aggregation Returns a map showing the number of
    unique students registered per course. Duplicate registration
    records are ignored.

Function: getCourseInfo(records []RegRecord) map[string]int

4.  Parallel Count Counts the occurrences of a key in a list of integers
    using multiple goroutines and channels for synchronization.

Function: count(list []int, key int, numThreads int) int

Requirements - Go 1.18 or newer

Verify installation: go version

Running Tests Navigate to the project directory and execute:

go test

For verbose output:

go test -v

Notes - list.txt must be located in the same directory when running
tests. - This project does not contain a main() function and is intended
to be executed using go test. - Demonstrates usage of goroutines,
channels, maps, and Unicode-safe string processing.
