package main

import (
	"crypto/rand"
	"fmt"
	"io"
)

func main() {
	slice1()
	slice2()
	slice3()
	slice4()
}

func slice4() {
	var nonce [24]byte
	fmt.Println(nonce)
	io.ReadFull(rand.Reader, nonce[:])
	fmt.Println(nonce)

}

func slice3() {
	verbs := make([][]string, 0)
	// ser
	word := make([]string, 8)
	word[0] = "ser"
	word[1] = "to be (perm)"
	word[2] = "soy"
	word[3] = "eres"
	word[4] = "es"
	word[5] = "somos"
	word[6] = "sois"
	word[7] = "son"
	verbs = append(verbs, word)
	word1 := make([]string, 8)
	word1[0] = "ir"
	word1[1] = "to go"
	word1[2] = "voy"
	word1[3] = "vas"
	word1[4] = "va"
	word1[5] = "vamos"
	word1[6] = "vais"
	word1[7] = "van"
	verbs = append(verbs, word1)
	word2 := make([]string, 8)
	word2[0] = "tomar"
	word2[1] = "to take, drink"
	word2[2] = "tomo"
	word2[3] = "tomas"
	word2[4] = "toma"
	word2[5] = "tomamos"
	word2[6] = "tomais"
	word2[7] = "toman"
	verbs = append(verbs, word2)

	fmt.Println(verbs)
	fmt.Println(word)
	fmt.Println(word1)
	fmt.Println(verbs[0])
	fmt.Println(verbs[1])
	word = verbs[1]
	fmt.Println(verbs)
	fmt.Println(word)
	fmt.Println(word1)

}
func slice2() {
	mySlice := make([]string, 0)
	mySlice = append(mySlice, "lunes")
	mySlice = append(mySlice, "martes")
	mySlice = append(mySlice, "miercules")
	mySlice = append(mySlice, "jueves")
	mySlice = append(mySlice, "veirnes")
	mySlice = append(mySlice, "sabado")
	mySlice = append(mySlice, "domingo")

	finDeSemana := append(mySlice[5:])
	diasDeLaSemana := append(mySlice[:5])

	fmt.Println("----------------")
	fmt.Println(mySlice)
	fmt.Println("len:", len(mySlice), "cap:", cap(mySlice))
	fmt.Println("----------------")
	fmt.Println(finDeSemana)
	fmt.Println("len:", len(finDeSemana), "cap:", cap(finDeSemana))
	fmt.Println("----------------")
	fmt.Println(diasDeLaSemana)
	fmt.Println("len:", len(diasDeLaSemana), "cap:", cap(diasDeLaSemana))

}

func slice1() {
	for sliceCap := 0; sliceCap < 15; sliceCap = sliceCap + 3 {
		mySlice := make([]int, 0, sliceCap)

		fmt.Println("----------------")
		fmt.Println(mySlice)
		fmt.Println("len:", len(mySlice), "cap:", cap(mySlice))

		for ind := 0; ind < 15; ind++ {
			mySlice = append(mySlice, ind)
			fmt.Println("len:", len(mySlice), "cap:", cap(mySlice), "value:", ind)

		}
		fmt.Println("----------------")
		fmt.Println(mySlice)
		fmt.Println("----------------")

	}

}
