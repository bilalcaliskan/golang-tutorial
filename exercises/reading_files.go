package exercises

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/gobuffalo/packr"
	"io/ioutil"
	"log"
	"os"
)

func ReadFiles() {
	fmt.Printf("\nBeginning of reading an entire file into memory...\n")
	/*
	One of the most basic file operations is reading an entire file into memory. This is done with the help of the
	ReadFile function of the ioutil package.
	 */
	// Below method reads the file into memory and returns a byte slice which is stored in data
	data, err := ioutil.ReadFile("exercises/test.txt")
	if err != nil {
		fmt.Println("File reading error", err)
	}
	fmt.Println("Contents of file:", string(data))
	/*
	If you run the program from compiled binary at $GOROOT/bin program will complain that it cannot find test.txt.
	The reason is Go is a compiled language. What go install does is, it creates a binary from the source code. The
	binary is independent of the source code and it can be run from any location. Since test.txt is not found in the
	location from which the binary is run, the program complains that it cannot find the file specified.
	There are three ways to solve this problem:
	1- Using absolute file path
	2- Passing the file path as a command line flag
	3- Bundling the text file along with the binary
	 */
	// 1- Using absolute file path
	data2, err2 := ioutil.ReadFile("/Users/bilal.caliskan/go/src/golang-tutorial/exercises/test.txt")
	if err2 != nil {
		fmt.Println("File reading error", err2)
	} else {
		fmt.Println("Contents of file:", string(data2))
	}
	// 2- Passing the file path as a command line flag
	/*
	Another way to solve this problem is to pass the file path as a command line flag. Using the flag package, we can
	get the file path as input from the command line and then read its contents. The flag package has a String function.
	This function accepts 3 arguments. The first is the name of the flag, second is the default value and the third is
	a short description of the flag.
	 */
	// This function returns the address of the string variable that stores the value of the flag.
	fptr := flag.String("fpath", "test.txt", "file path to read from")
	// flag.Parse() should be called before any flag is accessed by the program.
	//flag.Parse()
	fmt.Println("value of fpath is", *fptr)
	data3, err3 := ioutil.ReadFile(*fptr)
	if err3 != nil {
		fmt.Println("File reading error", err3)
	} else {
		fmt.Println("Contents of file:", string(data3))
	}
	// 3- Bundling the text file along with the binary
	/*
	Wouldn't it be awesome if we are able to bundle the text file along with our binary? There are various packages
	that help us achieve this. We will be using packr. packr converts static files such as .txt to .go files which are
	then embedded directly into the binary. Packer is intelligent enough to fetch the static files from disk rather
	than from the binary during development. This prevents the need for re-compilation during development when only
	static files change.
	 */
	box := packr.NewBox("files/")
	data4 := box.String("test.txt")
	fmt.Println("Contents of file:", data4)
	/*
	Very interesting thing about packr is that if you change the content
	of file after binary compiled, you will see that your changes effected already built binary. You can see that the
	program prints the updated contents of test.txt without the need for any recompilation.
	 */
	// packr install -v filehandling
	/*
	If you run above command, bundles the static file along with the binary.
	*/

	fmt.Printf("\nBeginning of reading a file in small chunks...\n")
	/*
	In the last section, we learned how to load an entire file into memory. When the size of the file is extremely
	large it doesn't make sense to read the entire file into memory especially if you are running low on RAM. A more
	optimal way is to read the file in small chunks. This can be done with the help of the bufio package.
	 */
	fptr = flag.String("fpath2", "test.txt", "file path to read from")
	flag.Parse()

	f, err := os.Open(*fptr)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	r := bufio.NewReader(f)
	b := make([]byte, 3)
	for {
		n, err := r.Read(b)
		if err != nil {
			fmt.Println("Error reading file:", err)
			break
		}
		fmt.Println(string(b[0:n]))
	}

	fmt.Printf("\nBeginning of reading a file line by line...\n")
	f, err = os.Open(*fptr)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	s := bufio.NewScanner(f)
	for s.Scan() {
		fmt.Println(s.Text())
	}
	err = s.Err()
	if err != nil {
		log.Fatal(err)
	}
}
