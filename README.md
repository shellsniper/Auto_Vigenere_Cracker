# Auto_Vigenere_Cracker

---
Usage:
---

# step 1:

## run: 
* ./vigenere-encrypt <encipherment key> <plaintext filename> > <ciphertext filename>
		- encrypt plainetext file with given key and output as a file in the vigenere-encrypt folder
* ./vigenere-decrypt <decipherment key> <ciphertext filename>
		- decrypt ciphertext file with given key in the vigenere-encrypt folder

# step 2: Finding key lengths and Full cryptanalysis
## run:
* ./vigenere-keylength <ciphertext filename>
		- run full cryptanalysis to figure out the possible key length, possbile key and decrypted ciphertext
