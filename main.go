package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
)

type keyPair struct {
	Public  string
	Private string
}

func main() {
	public := flag.String("pub", "", "eg: ~/.ssh/id_rsa.pub")
	private := flag.String("pvt", "", "eg: ~/.ssh/id_rsa")
	flag.Parse()

	if len(*public) == 0 && len(*private) == 0 {
		panic("must provide a certificate")
	}

	if len(*public) > 0 && len(*private) > 0 {
		pub, err := ioutil.ReadFile(*public)
		if err != nil {
			panic("public key unreadable")
		}
		pvt, err := ioutil.ReadFile(*private)
		if err != nil {
			panic("private key unreadable")
		}

		object, err := json.Marshal(keyPair{Public: string(pub), Private: string(pvt)})
		if err != nil {
			panic(err)
		}
		fmt.Println(string(object))
	} else {
		if len(*public) > 0 {
			pub, err := ioutil.ReadFile(*public)
			if err != nil {
				panic("public key unreadable")
			}
			object, err := json.Marshal(keyPair{Public: string(pub)})
			if err != nil {
				panic(err)
			}
			fmt.Println(string(object))
		} else {
			pvt, err := ioutil.ReadFile(*private)
			if err != nil {
				panic("private key unreadable")
			}

			object, err := json.Marshal(keyPair{Private: string(pvt)})
			if err != nil {
				panic(err)
			}
			fmt.Println(string(object))
		}
	}
}
