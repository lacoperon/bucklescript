

package main

import (
	"log"
	"os"
	"os/exec"
	"fmt"
)

func checkError(err error){
	if err != nil {
		log.Fatalf("Error: %s", err.Error())
	}
}

func testTheme(theme string){
	fmt.Println("Removing",theme)
	os.RemoveAll(theme)
	cmd := exec.Command("bsb", "-theme",theme,"-init",theme)
	output, err:= cmd.CombinedOutput()
	
	fmt.Println(string(output))
	checkError(err)
	
	fmt.Println("Started to build ")
	cmd2 := exec.Command("npm", "run", "build")		
	cmd2.Dir = theme
	output2, err:= cmd2.CombinedOutput()
	fmt.Println(string(output2))
	checkError(err)		
	os.RemoveAll(theme)
}


func runMoCha(){
	fmt.Println("Running Mocha tests")
	cmd:=exec.Command("mocha", "./jscomp/test/**/*test.js")
	output,err:=cmd.CombinedOutput()
	fmt.Println(string(output))
	checkError(err)
	fmt.Println("Running mocha finished")
	
}
func main(){
	runMoCha()
	exec.Command("npm", "link")
}