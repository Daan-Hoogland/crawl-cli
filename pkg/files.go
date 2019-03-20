package pkg

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/spf13/afero"
)

//File struct to handle simple file information easily.
//Maybe move this to a seperate file/package?
type File struct {
	Name string
	Path string
	Size int64
	Mode os.FileMode
}

//AppFs filesystem for application
var AppFs = afero.NewOsFs()

var jobs = make(chan string)

//WalkDirectory walks over a directory and returns a File struct array for files
//and directories found.
func WalkDirectory(dir string) {
	log.WithField("Directory", dir).Debugln("inside walkdirectory")

}

func scanDirectory(path string) ([]os.FileInfo, error) {
	fs, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	// files and directories found in directory
	fileInfo, err := fs.Readdir(-1)

	//todo: analyse the files


}



//diskUsage returns disk usage of a given path.
func diskUsage(currentPath string, info os.FileInfo) int64 {
	if info == nil {
		info, _ = os.Lstat(currentPath)
	}

	size := info.Size()

	if !info.IsDir() {
		return size
	}

	dir, err := os.Open(currentPath)

	if err != nil {
		fmt.Println(err)
		return size
	}
	defer dir.Close()

	files, err := dir.Readdir(-1)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, file := range files {
		if file.Name() == "." || file.Name() == ".." {
			continue
		}
		size += diskUsage(currentPath+"/"+file.Name(), file)
	}

	return size
}
