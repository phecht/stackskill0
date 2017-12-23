package main

import "fmt"

type palabra struct {
	verbo string
	verb  string
	s1    string
	s2    string
	s3    string
	p1    string
	p2    string
	p3    string
}

func main() {
	p := palabra{"ir", "to go", "voy", "vas", "va", "vamos", "vais", "van"}
	actualVerb := ""
	actualTense := ""
	fmt.Println("if")
	fmt.Scanln()
	tense := "s1"
	if tense == "p1" {
		actualTense = "first person plural"
		actualVerb = p.p1
	} else if tense == "s1" {
		actualTense = "first person singular"
		actualVerb = p.s1
	}
	printV(p.verbo, p.verb, actualVerb, actualTense)
}

func printV(sv string, ev string, av string, at string) {
	fmt.Println(sv, " means ", ev, ". In the ", at, " and translates to", av)

}
