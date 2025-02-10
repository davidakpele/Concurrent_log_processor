package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
	"sync"
)

// ProcessLogFile reads the log file in chunks and counts keyword occurrences concurrently.
func ProcessLogFile(filePath string, keywords []string) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	keywordMap := make(map[string]int)
	for _, keyword := range keywords {
		keywordMap[strings.ToLower(keyword)] = 0
	}

	var wg sync.WaitGroup
	mutex := &sync.Mutex{}
	ch := make(chan string, 100)

	go func() {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			ch <- scanner.Text()
		}
		close(ch)
	}()

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for line := range ch {
				for keyword := range keywordMap {
					matched, _ := regexp.MatchString("\\b"+keyword+"\\b", strings.ToLower(line))
					if matched {
						mutex.Lock()
						keywordMap[keyword]++
						mutex.Unlock()
					}
				}
			}
		}()
	}

	wg.Wait()

	// Sort and display results
	type kv struct {
		Key   string
		Value int
	}
	var sorted []kv
	for k, v := range keywordMap {
		sorted = append(sorted, kv{k, v})
	}
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].Value > sorted[j].Value
	})

	for _, item := range sorted {
		fmt.Printf("%s: %d\n", strings.ToUpper(item.Key), item.Value)
	}
}

// Prime and Palindrome Challenge
func isPrime(num int) bool {
	if num < 2 {
		return false
	}
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func isPalindrome(num int) bool {
	s := fmt.Sprintf("%d", num)
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}

func FindPrimePalindromes(N int) int {
	var wg sync.WaitGroup
	ch := make(chan int, 100)
	resultCh := make(chan int, N)
	count := 0

go func() {
		num := 2
		for count < N {
			ch <- num
			num++
		}
		close(ch)
	}()

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for num := range ch {
				if isPrime(num) && isPalindrome(num) {
					resultCh <- num
					count++
				}
			}
		}()
	}

	wg.Wait()
	close(resultCh)

	sum := 0
	for num := range resultCh {
		sum += num
	}
	return sum
}

func main() {
	// Example usage of log processor
	keywords := []string{"INFO", "ERROR", "DEBUG"}
	ProcessLogFile("log.txt", keywords)

	// Example usage of prime palindrome finder
	N := 10
	fmt.Println("Sum of first", N, "prime palindromes:", FindPrimePalindromes(N))
}
