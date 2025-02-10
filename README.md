# Concurrent_log_processor
Golang Concurrent Log Processor & Prime Palindrome Finder

## Description
**This project implements two concurrent solutions in Go:**
1. **Concurrent Log Processor:** Reads a log file in chunks, counts keyword occurrences concurrently, and outputs the sorted keyword frequency.
2. **Prime Palindrome Finder:** Finds the first N prime palindromic numbers concurrently and returns their sum.

## Features
- Uses **Goroutines** for concurrency.
- Utilizes **channels** for communication between workers.
- Implements **sync.Mutex** and **sync.WaitGroup** to avoid race conditions.
- Processes large log files efficiently.
- Finds prime palindromes in an optimized manner.


## Installation
Ensure you have Go installed. If not, download and install it from Go's official website.
### Clone the Repository
```
git clone <repository_url>
cd golang-concurrent-processor
```
### Initialize the Go Module
```
go mod init golang-concurrent-processor
go mod tidy
```

### Usage
#### Run the Log Processor

```
go run main.go
```

Make sure you have a _log.txt_ file in the same directory containing logs. Modify _keywords_ in _main.go_ as needed.

### Run Prime Palindrome Finder
```
go run main.go
```

Modify the *N* value inside *main.go* to find different amounts of prime palindromes.

### File Structure
C:/
├── main.go          # Main application logic
├── go.mod          # Go module file
├── log.txt         # Sample log file for testing
└── README.md       # Project documentation

### Example Output
#### Log Processor
```
INFO: 3
ERROR: 2
DEBUG: 1
```
### Prime Palindrome Finder
For **N = 10**, the output:
```
Sum of first 10 prime palindromes: 1132
```



