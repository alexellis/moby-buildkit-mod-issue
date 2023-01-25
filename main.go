package main

import (
	"bytes"
	"flag"
	"fmt"

	"os"

	"github.com/moby/buildkit/client"
	"github.com/moby/moby/pkg/archive"
)

func main() {

	var file string
	flag.StringVar(&file, "file", "", "Tar file to read")
	flag.Parse()

	if file == "" {
		panic("No file specified")
	}

	lchownEnabled := false

	opts := archive.TarOptions{
		NoLchown: !lchownEnabled,
	}

	fmt.Printf("Extracting %s\n", file)

	tmpWorkspace := os.TempDir()
	tarBytes, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}

	if err := archive.Untar(bytes.NewReader(tarBytes),
		tmpWorkspace, &opts); err != nil {
		panic(err)
	}

	var c *client.Client
	if c == nil {
		fmt.Println("Client is nil")
	}

	fmt.Printf("Check %s\n", tmpWorkspace)
}
