package main

/*****************************
** Name: Chenfeng Nie
** Class: Crypto
** Assignment 1
** Date: Sep 9 2018
*****************************/
import (
	"bytes"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strings"

	"../package"
)

const MAX_KEY_LENGTH int = 20

//function to counter each letter"s freq
func countLetterFreq(ciphertext string) []int {
	N := 0
	//declare letter"s freq
	letterfreq := make([]int, 26)
	letterASCII := 65 //"A"
	for i := 0; i < 26; i++ {
		char := string(i + letterASCII)
		letterfreq[i] = strings.Count(ciphertext, char)
		N = N + letterfreq[i]
	}
	//fmt.Println(N)
	//fmt.Println(letterfreq)
	return letterfreq
}

//calcuate index of concedence
/**
If the Index of coincidence is low (close to
0.0385), i.e. similar to a random text, then the message has probably been crypted using a polyalphabetic cipher (a letter can be replaced by multiple other ones).

The more the coincidence count is low, the more alphabets have been used.

Example: Vigenere cipher with a key of length 4 to 8 letters have an IC of about
0.045 ± 0.05

Index of Coincidence in English = 1.73

**/
func calcIndexOfCoincidence(ciphertext string) float64 {
	IC := 0.0
	lengthOfText := float64(len(ciphertext) - 1.0)
	//there are 26 characters in total
	C := 26.0
	letterFreqArray := countLetterFreq(ciphertext)
	divider := (lengthOfText * (lengthOfText - 1)) / C
	var nSum float64
	for i := 0; i < 26; i++ {
		nSum = nSum + float64(letterFreqArray[i]*(letterFreqArray[i]-1))
	}

	//fmt.Println(letterFreqArray)
	//fmt.Println(divider)
	//if the length of ciphertext is 1 so the divider will be zero
	if divider == 0 {
		return 0.0
	}

	IC = nSum / divider
	return IC
}

//func to determine possible key length
func possibleKeyLength(ciphertext string) []int {
	var possibleKeyArray []float64
	var keyLenResult []int
	const englishIC float64 = 1.73
	for i := 1; i <= MAX_KEY_LENGTH; i++ {
		averageIC := 0.0
		IC := 0.0
		for j := 0; j < i; j++ {
			//var subString string
			var buffer bytes.Buffer
			for k := 0; k < len(ciphertext)-1; k++ {
				if k%i == j {
					buffer.WriteString(string(ciphertext[k]))
				}
			}
			IC = IC + calcIndexOfCoincidence(buffer.String())
		}
		averageIC = IC / float64(i)
		//fmt.Println("averageIC = " + strconv.FormatFloat(averageIC, "f", 4, 64))
		possibleKeyArray = append(possibleKeyArray, averageIC)
	}
	for index := range possibleKeyArray {
		if math.Abs(possibleKeyArray[index]-englishIC) < 0.20 {
			keyLenResult = append(keyLenResult, index+1)
		}
	}

	return keyLenResult
}

//split ciphertext with given length and transpose text
func transpose(ciphertext string, keylength int) []string {
	var substringSlice []string
	var substring bytes.Buffer
	for index := 0; index < keylength; index++ {
		for i := index; i < len(ciphertext)-1; i = i + keylength {
			substring.WriteString(string(ciphertext[i]))
		}
		//fmt.Println(substring.String())
		//fmt.Println()
		substringSlice = append(substringSlice, substring.String())
		substring.Reset()
	}
	return substringSlice
}

//guess the key
func guesskey(substringSlice []string) string {
	var key string
	characters := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	//reference: Letter frequency
	// https://en.wikipedia.org/wiki/Letter_frequency
	englistFreq := map[string]float64{
		"A": 8.167,
		"B": 1.492,
		"C": 2.782,
		"D": 4.253,
		"E": 12.702,
		"F": 2.228,
		"G": 2.015,
		"H": 6.094,
		"I": 6.966,
		"J": 0.153,
		"K": 0.772,
		"L": 4.025,
		"M": 2.406,
		"N": 6.749,
		"O": 7.507,
		"P": 1.929,
		"Q": 0.095,
		"R": 5.987,
		"S": 6.327,
		"T": 9.056,
		"U": 2.758,
		"V": 0.978,
		"W": 2.361,
		"X": 0.150,
		"Y": 1.974,
		"Z": 0.074,
	}

	for _, text := range substringSlice {
		//fmt.Println(i, text)
		var scoreSlice []float64

		for _, letter := range characters {
			//fmt.Println(index, string(letter))
			//score to determine the most possible letter as a key
			var letterFreqScore float64 = 0.0

			decryptMsg := vigenerefunc.Vigenere(string(letter), text, false)
			letterFreqinArray := countLetterFreq(decryptMsg)
			//fmt.Println("freq array =", letterFreqinArray)
			//fmt.Println("decryptMsg =", decryptMsg)
			for i := 0; i < 26; i++ {
				letterFreqScore += (float64(letterFreqinArray[i]) * englistFreq[string(i+65)])
				//fmt.Println(i, letterFreq)
			}
			scoreSlice = append(scoreSlice, letterFreqScore)
			//fmt.Println(string(letter), letterFreqScore)

		}
		//find the biggest freq socre in score slice
		biggest := scoreSlice[0]
		var letter string
		for index, v := range scoreSlice {
			if v >= biggest {
				biggest = v
				letter = string(index + 65)
			}
		}
		fmt.Println("The biggest score is ", biggest, "===>", letter)
		//fmt.Println()
		key += letter
	}
	return key
}

//main
func main() {
	var possibleKey string
	if len(os.Args) == 2 {
		fileName := os.Args[1]
		// check if the key length is between 1 to 32 characters
		ciphertext, err := ioutil.ReadFile(fileName)
		vigenerefunc.CheckFile(err)
		//fmt.Print(ciphertext)
		str := string(ciphertext) // convert content to a "string"
		//fmt.Println(str)
		//fmt.Println(calcIndexOfCoincidence(str))
		//fmt.Println(possibleKeyLength(str))

		for _, length := range possibleKeyLength(str) {
			substring := transpose(str, length)
			possibleKey = guesskey(substring)
			decryptMessage := vigenerefunc.Vigenere(possibleKey, str, false)
			fmt.Println("-----------------------------------")
			fmt.Println("Length: ", length)
			fmt.Println("Key: ", possibleKey)
			fmt.Println("message is :", decryptMessage)
			fmt.Println("-----------------------------------")
		}

		//fmt.Println(substring[1])

	} else {
		fmt.Println("check your input args!")
	}

}
