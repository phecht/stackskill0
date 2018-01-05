package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Verb struct {
	Verb    string
	English string
	S1      string `json:"First Person Singular"`
	S2      string
	S3      string
	P1      string
	P2      string
	P3      string
}

func main() {
	var toBeCalled Verb
	var toSpeak Verb
	someBytes := []byte(`{"Verb":"llamarse", "English":"to be called","First Person Singular":"me llamo"}`)
	json.Unmarshal(someBytes, &toBeCalled)
	fmt.Println("-------")
	fmt.Println(toBeCalled.Verb)
	fmt.Println(toBeCalled.English)
	fmt.Println(toBeCalled.S1)
	fmt.Println("-------")
	v1 := Verb{"ir", "to go", "voy", "vas", "va", "vamos", "vais", "van"}
	jObj, _ := json.Marshal(v1)
	fmt.Println(v1)
	fmt.Printf("%T \n", jObj)
	fmt.Println(string(jObj))
	fmt.Println(jObj)
	fmt.Println("-------")
	json.NewEncoder(os.Stdout).Encode(v1)
	fmt.Println("-------")
	rdr := strings.NewReader(`{"Verb":"hablar","English":"to speak","S1":"hablo"}`)
	json.NewDecoder(rdr).Decode(&toSpeak)
	fmt.Println(toSpeak.Verb)
	fmt.Println(toSpeak.English)
	fmt.Println("-------")

}
