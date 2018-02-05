package gotree

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func log(a ...interface{}) {
	fmt.Println(a...)
}
func error(e interface{}) {
	fmt.Printf("ERROR: %s", e)
}

// GoTree shows current folder structure
func GoTree() {
	cwd, err := os.Getwd()
	if err != nil {
		error(err)
		return
	}
	ctx := context{cwd, os.Args[1:]}
	draw(ctx.Root, 0, ctx.isDrawFiles())
}

func getPrefix(level int) string {
	prefix := ""
	for i := 1; i <= level; i++ {
		prefix += "   " // \t is big
	}
	prefix += "|-"
	return prefix
}

func draw(root string, level int, fileDraw bool) {
	infos, err := ioutil.ReadDir(root)
	if err != nil {
		error(err)
		return
	}
	prefix := getPrefix(level)
	for _, val := range infos {
		postfix := ""
		isDir := val.IsDir()
		if isDir {
			postfix = "/"
		}
		if isDir || fileDraw {
			log(prefix, val.Name()+postfix)
		}
		if val.IsDir() {
			draw(filepath.Join(root, val.Name()), level+1, fileDraw)
		}
	}
}
