# cobra-cli-proj
cobra-cli-proj is a CLI tool that empowers users to, make directories, files. It also allows users to add data to files, print all data from the file and also search through the file to print specific data.

## Examples:
- cobra-cli-proj.exe mkd Test  - Makes directory named Test. You can also assign full path while creating directory
- cobra-cli-proj.exe mkf test.txt  -  Makes test.txt file at current location. Can also provde full location.
- cobra-cli-proj.exe appendF "Hi, This is test file" test.txt  -  Appends user input text to the file.
- cobra-cli-proj.exe readF test.txt  -  Reads the content of file and prints it on the console.
- cobra-cli-proj.exe writestructfile .\Data.txt 9 Patel 52 9984126238 - Write user given input to the file and saves it in specific format using | to seperate data.
- cobra-cli-proj.exe datastruct .\Data.txt - Prints data from given file on console using structs to properly formatted.
- cobra-cli-proj.exe searchdata Tirth .\Data.txt - Searches the given file for "Tirth" and prints the row with all data.
