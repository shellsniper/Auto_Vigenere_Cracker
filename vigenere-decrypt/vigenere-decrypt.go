package main

/*****************************
** Name: Chenfeng Nie
** Class: Crypto
** Assignment 1
** Date: Sep 9 2018
*****************************/

import (
	"fmt"
	"io/ioutil"
	"os"

	"../package"
)

// main func
func main() {
	//check if there are exact three args are input by user
	if len(os.Args) == 3 {
		key := os.Args[1]
		// check if the key length is between 1 to 32 characters
		if len(key) <= 32 && len(key) > 0 {
			fileName := os.Args[2]
			ciphertext, err := ioutil.ReadFile(fileName)
			// read file's size and catch error
			fi, e := os.Stat(fileName)
			if e != nil {
				fmt.Println("Could not obtain stat, handle error")
			} else {
				if vigenerefunc.CheckFileSize(fi) == true {
					vigenerefunc.CheckFile(err)
					//fmt.Println(string(plaintext))
					//ciphertext := vigenere(key, string(plaintext), true)
					newplaintext := vigenerefunc.Vigenere(key, string(ciphertext), false)
					//fmt.Println("Plaintext: " + newplaintext)
					//print recoverd text in terminal
					fmt.Println(newplaintext)
				}
			}

		} else {
			fmt.Println("check your encipherment KEY! (length: 1-32)")
		}
	} else {
		fmt.Println("check your input args")
	}

}
