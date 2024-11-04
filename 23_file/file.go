package main

import (
	"fmt"
	"os"
)

func main(){
// 	file,error:=os.Open("text_file.txt")

// if error!=nil{
// 	//log the error
// 	panic(error)
// }
// fileInfo,err:=file.Stat()

// if err!=nil {
// 	//log the error
// 	panic(err)
	
// }
// fmt.Println("file name:", fileInfo.Name())
// fmt.Println("file modified at: ",fileInfo.ModTime())

// defer file.Close()
//-----------------2nd method to read file using buffer---------------
// buf:=make([]byte,20)

// length,err:=file.Read(buf)

// if err!=nil {
// 	panic(err)	
// }

// for i := 0; i < len(buf); i++ {
// 	println("data: ", length, string(buf[i]))
// }

// println("data in the file:", length," ",buf)
//to read file in synce in one go, a better solution might be using ReadFile

data,err:=os.ReadFile("text_file.txt")
if err!=nil{
	panic(err)
}
fmt.Println(string(data))

//read folder
 

}

