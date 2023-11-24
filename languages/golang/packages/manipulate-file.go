package main

import (
	"fmt"
	"os"
    "bufio"
)

func operationFile(nameFile string){

	f, err := os.Create("files/" + nameFile)

	if err != nil {
		panic(err)
	}

	size, err := f.Write([]byte("write data in my file, because I Study"))

	if err != nil {
		panic(err)
	}

	fmt.Printf("%d bytes\n", size)

	f.Close()
		
}
func openfile(file string){
	data, err := os.ReadFile("files/" + file)

	if err != nil {
		panic(err)
	}

	//fmt.Println(data) //slice bytes
	fmt.Println(string(data))

    myFile, err := os.Open("files/" + file)

    if err != nil {
        panic(err)
    }
    
    reader := bufio.NewReader(myFile)

    buffer := make([]byte, 2)
    for {
        n, err := reader.Read(buffer)
        if err != nil {
            break
        }
        fmt.Println(string(buffer[:n]))
    }

}

func main(){
	var fileName string = "newfile.txt"
	operationFile(fileName)
	openfile(fileName)
    
}
