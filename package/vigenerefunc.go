package vigenerefunc

/*****************************
** Name: Chenfeng Nie
** Class: Crypto
** Assignment 1
** Date: Sep 9 2018
*****************************/
import (
	"fmt"
	"os"
	"strings"
)

// check if the file exists
func CheckFile(e error) {
	if e != nil {
		panic(e)
	}
}

// check file size
func CheckFileSize(fi os.FileInfo) bool {
	// convert byte to KB
	fileSizeInKB := float64(fi.Size()) / 1024.0
	// text file sizes yp to a maximum of 100 KB
	if fileSizeInKB <= 100.0 && fileSizeInKB >= 0.0 {
		//fmt.Printf("(Valid File)The file is %f KB", fileSizeInKB)
		//fmt.Println()
		return true
	} else {
		fmt.Printf("(Invalid File)Size of the file is %f KB, too large to process", fileSizeInKB)
		return false
	}

}

// check whether byte is character or not
func CheckLetter(plainChar byte) bool {
	if plainChar >= 65 && plainChar <= 90 {
		return true
	}
	return false

}

//Encrypt each character
func GetEncryptedChar(plainChar byte, keyChar byte) string {

	encryptedChar := ((plainChar - 65) + (keyChar - 65)) % 26

	return string(encryptedChar + 65)
}

// Decrypt each character
func GetDecryptedChar(cipherChar int, keyChar int) string {

	decryptedChar := (cipherChar - keyChar) % 26

	for decryptedChar < 0 {
		decryptedChar += 26
	}

	return string(decryptedChar + 65)
}

// encrypt or decrypt interface
func Vigenere(key string, inputtext string, encrypt bool) string {

	resultText := ""
	keyIndex := 0
	count := 0

	// sanitize the input
	keyUpper := strings.ToUpper(key)
	inputUpper := strings.ToUpper(inputtext)

	for i := 0; i < len(inputUpper); i++ {

		if CheckLetter(inputUpper[i]) {

			keyIndex = count % len(keyUpper)

			if encrypt == true {
				// encrypt
				resultText += GetEncryptedChar(inputUpper[i], keyUpper[keyIndex])
			} else {
				// decrypt
				resultText += GetDecryptedChar(int(inputUpper[i]), int(keyUpper[keyIndex]))
			}
			count++

		} else {

			//resultText += string(inputUpper[i])

		}

	}

	return resultText

}
