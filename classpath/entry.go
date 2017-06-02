package classpath

import (
    "os"
    "strings"
)

const pathListSep = string(os.PathListSeparator)

type Entry interface {
    readClass(className string) ([]byte, Entry, error)
    String() string
}


func newEntry(path string) Entry {
    if strings.Contains(path, pathListSep) {
        return newCompositeEntry(path)
    }
    if strings.HasSuffix(path, "*") {
        return newWildcardEntry(path)
    }
    if strings.HasSuffix(path, ".jar") ||
            strings.HasSuffix(path, ".zip") {
        return newZipEntry(path)
    }
    return newDirEntry(path)
}