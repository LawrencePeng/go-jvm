package classpath

import (
	"path/filepath"
	"io/ioutil"
)

type DirEntry struct {
	absDir string
}

func newDirEntry(path string) *DirEntry {
	absPath, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &DirEntry{absDir:absPath}
}

func (self *DirEntry) readClass(className string) ([]byte, Entry, error) {
	fileName := filepath.Join(self.absDir, className)
	data, err := ioutil.ReadFile(fileName)
	return data, self, err
}

func (self *DirEntry) String() string {
	return self.absDir
}
