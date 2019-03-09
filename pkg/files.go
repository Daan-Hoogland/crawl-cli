package pkg

import (
	"fmt"
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"

	"github.com/spf13/afero"
)

//File struct to handle simple file information easily.
//Maybe move this to a seperate file/package?
type File struct {
	name string
	path string
	size int64
	mode os.FileMode
}

//AppFs filesystem for application
var AppFs = afero.NewOsFs()

//WalkDirectory walks over a directory and returns a File struct array for files
//and directories found.
func WalkDirectory(dir string) ([]File, []File, error) {
	log.WithField("Directory", dir).Debugln("inside walkdirectory")
	afs := &afero.Afero{Fs: AppFs}
	var files []File
	var directories []File
	err := afs.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			directories = append(directories,
				File{
					name: filepath.Base(info.Name()),
					path: path,
					size: info.Size(),
					mode: info.Mode(),
				})
		} else {
			files = append(files,
				File{
					name: filepath.Base(info.Name()),
					path: path,
					size: info.Size(),
					mode: info.Mode(),
				})
		}

		return nil
	})

	return files, directories, err
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
