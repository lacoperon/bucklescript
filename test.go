package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"sync"
)

func checkError(err error) {
	if err != nil {
		log.Fatalf("Error: %s", err.Error())
	}
}

func testTheme(theme string) {
	fmt.Println("Removing", theme)
	os.RemoveAll(theme)
	cmd := exec.Command("bsb", "-theme", theme, "-init", theme)
	output, err := cmd.CombinedOutput()

	fmt.Println(string(output))
	checkError(err)

	fmt.Println("Started to build ")
	cmd2 := exec.Command("npm", "run", "build")
	cmd2.Dir = theme
	output2, err := cmd2.CombinedOutput()
	fmt.Println(string(output2))
	checkError(err)
	os.RemoveAll(theme)
}

func runMoCha(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Running Mocha tests")
	cmd := exec.Command("mocha", "./jscomp/test/**/*test.js")
	output, err := cmd.CombinedOutput()
	fmt.Println(string(output))
	checkError(err)
	fmt.Println("Running mocha finished")

}

func installGlobal(wg *sync.WaitGroup) {
	defer wg.Done()

	cmd := exec.Command("npm", "install", "-g", ".")

	output, err := cmd.CombinedOutput()
	fmt.Printf(string(output))
	checkError(err)
}

var cmd = exec.Command

func main() {
	vendorOCamlPath, _ := filepath.Abs(filepath.Join(".", "vendor", "ocaml", "bin"))
	os.Setenv("PATH",
		vendorOCamlPath+string(os.PathListSeparator)+os.Getenv("PATH"))
	// Avoid rebuilding OCaml again
	output, _ := cmd("which", "ocaml").CombinedOutput()
	fmt.Println(string(output))

	var wg sync.WaitGroup
	wg.Add(2)
	go runMoCha(&wg)
	go installGlobal(&wg)
	wg.Wait()

	for _, theme := range []string{"basic", "basic-reason", "generator", "minimal"} {
		fmt.Println("Test theme", theme)
		wg.Add(1)
		go (func(theme string) {
			defer wg.Done()
			testTheme(theme)
		})(theme)
	}
	wg.Wait()
}
