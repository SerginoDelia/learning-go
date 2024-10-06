/ go run hello.go --> to run the code
// go build hello.go --> to compile the code
// go build -o hello_world hello.go --> use the -o flag to give the binary a different name or location

package main

import "fmt"

func main() {
	fmt.Println("Hello, world!")
}

// Getting 3rd Party Go Tools
// The "go install" command takes one argument which is the location of the source code repo you want to install
// followed by an @ and the version  of the tool you want (use @latest if you want the latest version)
// It then downloads, compiles, and intall the tool in your $GOPATH/bin directory

// Install hey, a tool that load tests HTTP servers. You can point it at the website of your choosing or an
// application that you've written
// go install github.com/rakyll/hey@latest

// run hey
// hey https://www.golang.org

// update an installed tool to the latest version
// rerun the install command again with the newer version specified or with @latest

// go fmt automatically reformats your code to match the standard format
// "goimports" cleans up your import statements, it puts them in alphabetical order and removes unused imports
// go install golang.org/x/ tools/cmd/goimports@latest
// run goimports "goimports -l -w .
// -l flag tells goimports to print the files with incorrect formatting to the console
// -w tells goimports to modify the files in-place
// . specifies the directory to be scanned

// always run go fmt and goimports before compiling your code

// Linting and Vetting

// Intall golint
// go install golang.org/x/lint/golint@latest
// and run it with "golint ./..." that runs golint over the entire project

// run "go vet ./..."
// The code is syntactically valid, but there are mistakes that are not what you meant to do
// including passing the wrong number of parameters to formatting method or assigning values to variables
// that are never used "go vet" detect these kind of errors

// run mulitple tools together with "golangci-lin run". it combines go lint and go vet and many other tools
// https://golangci-lint.run/usage/configuration/
