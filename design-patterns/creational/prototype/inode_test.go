package prototype

import (
	"testing"
)

func TestInodePrototype(t *testing.T) {
	file1 := &File{name: "File1"}
	file2 := &File{name: "File2"}
	file3 := &File{name: "File3"}
	folder1 := &Folder{
		children: []Inode{file1},
		name:     "Folder1",
	}
	folder2 := &Folder{
		children: []Inode{folder1, file2, file3},
		name:     "Folder2",
	}
	cloneFolder := folder2.Clone()
	if !isClone(folder2, cloneFolder) {
		t.Error("Test failed: folder2 is equal to cloneFolder")
	}
	if isClone(folder2, folder1) {
		t.Error("Test failed: folder1 is not equal to folder2")
	}
	if isClone(folder1, folder2) {
		t.Error("Test failed: file1 is not equal to file2")
	}
}

func isClone(origin, clone Inode) bool {
	if origin.GetType() != clone.GetType() {
		return false
	}
	if origin.GetName()+"_clone" != clone.GetName() {
		return false
	}
	if origin.GetType() == "*prototype.File" {
		return true
	}
	if len(origin.GetChildren()) != len(clone.GetChildren()) {
		return false
	}
	for i := 0; i < len(origin.GetChildren()); i++ {
		if !isClone(origin.GetChildren()[i], clone.GetChildren()[i]) {
			return false
		}
	}
	return true
}
