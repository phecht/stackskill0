package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
)

func hashBucket(word string, buckets int) int {
	var sum int
	for _, v := range word {
		sum += int(v)
	}
	return sum % buckets
}

func makeHash1() {
	hashSize := 12
	book := "http://www.gutenberg.org/files/2701/2701.txt"
	buckets := make([]int, hashSize)

	res, err := http.Get(book)

	if err != nil {
		log.Fatalln(err)
	}
	defer res.Body.Close()
	bs := bufio.NewScanner(res.Body)
	bs.Split(bufio.ScanWords)
	for bs.Scan() {
		n := hashBucket(bs.Text(), hashSize)
		buckets[n]++
	}
	fmt.Println(buckets)

}

func getWordsS() map[string]string {
	words := make(map[string]string)
	res, err := http.Get("https://raw.githubusercontent.com/dwyl/english-words/master/words.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer res.Body.Close()

	bs := bufio.NewScanner(res.Body)
	bs.Split(bufio.ScanWords)
	for bs.Scan() {
		words[bs.Text()] = ""
	}

	if err = bs.Err(); err != nil {
		fmt.Println("Read error in getWordsS", err)
	}
	return words
}

/*
func getWords() string {
	res, err := http.Get("https://raw.githubusercontent.com/dwyl/english-words/master/words.txt")
	if err != nil {
		log.Fatalln(err)
	}
	bs, _ := ioutil.ReadAll(res.Body)

	//bs := bufio.NewScanner(res.Body)
	//fmt.Println(bs)
	//bs.Split(bufio.ScanWords)
	str := string(bs)
	return str
}
*/
func main() {

	//	scanned := getWords()
	//	fmt.Println(scanned)
	makeHash1()
	newWords := getWordsS()
	stopAt := 0
	for key := range newWords {
		fmt.Println(key)
		if stopAt > 200 {
			break
		}
		stopAt++
	}
}
