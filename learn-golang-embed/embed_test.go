package learn_golang_embed

import (
	"embed"
	_ "embed"
	"fmt"
	"io/fs"
	"os"
	"testing"
)

//go:embed version.txt
var version string

func TestEmbed(t *testing.T) {
	fmt.Println(version)
}

//go:embed files/file_a.txt
//go:embed files/file_b.txt
//go:embed files/file_c.txt
//go:embed files/file_d.txt
var files embed.FS

func TestMultipleEmbed(t *testing.T) {
	a, _ := files.ReadFile("files/file_a.txt")
	fmt.Println(string(a))

	b, _ := files.ReadFile("files/file_b.txt")
	fmt.Println(string(b))

	c, _ := files.ReadFile("files/file_c.txt")
	fmt.Println(string(c))
}

//go:embed files/*.txt
var path embed.FS

func TestMultipleEmbedWildcard(t *testing.T) {
	dirEntries, _ := path.ReadDir("files")
	for _, entry := range dirEntries {
		if !entry.IsDir() {
			file, _ := path.ReadFile("files/" + entry.Name())
			fmt.Printf("%s %s\n", entry.Name(), string(file))
		}
	}
}

//go:embed quby_sikat.jpeg
var image []byte

func TestFileEmbed(t *testing.T) {
	err := os.WriteFile("quby_sikat_v2.jpeg", image, fs.ModePerm)
	if err != nil {
		panic(err)
	}
}
