GO CONCURRENCY AND ALGORITHMS
=============================

OVERVIEW
--------
This repository contains a Go package implementing several core programming
tasks. The project demonstrates mathematical computation, Unicode-safe string
processing, map-based aggregation, and parallel counting using goroutines
and channels.

The code is structured as a testable Go module and is intended to be executed
using the Go testing framework.


PROJECT STRUCTURE
-----------------
asg1.go          Implementation of all tasks
asg1_test.go     Unit tests
list.txt         Input data file used for concurrency testing
go.mod           Module definition


IMPLEMENTED TASKS
-----------------

1) Sum of Squares
   Computes:
   0^2 + 1^2 + 2^2 + ... + |n|^2

   Function:
   getSumSquares(n int) int


2) Word Extraction
   Extracts all words from a string that end with a specified rune.
   The implementation is Unicode-safe.

   Function:
   getWords(text string, endLetter rune) []string


3) Course Registration Aggregation
   Returns a map containing the number of unique students registered
   per course. Duplicate registration records are ignored.

   Function:
   getCourseInfo(records []RegRecord) map[string]int


4) Parallel Count
   Counts the number of occurrences of a key in a list of integers.
   The computation is performed in parallel using multiple goroutines
   and channels for synchronization.

   Function:
   count(list []int, key int, numThreads int) int


REQUIREMENTS
------------
Go 1.18 or newer.

Verify installation:
    go version


RUNNING TESTS
-------------
From the project directory:

    go test

Verbose mode:

    go test -v

Run a specific test:

    go test -run TestGetWords -v


NOTES
-----
- The file "list.txt" must be present in the same directory when running tests.
- This project does not contain a main() function.
- The code demonstrates use of goroutines, channels, maps, and
  Unicode-aware string handling.
