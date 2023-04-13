package prototype

import "fmt"

type Folder struct {
	children []Inode
	name     string
}

func (f *Folder) Print(indentation string) {
	fmt.Println(indentation + f.name)
	for _, child := range f.children {
		child.Print(indentation + indentation)
	}
}

func (f *Folder) Clone() Inode {
	cloneFolder := &Folder{name: f.name + "_clone"}
	var tempChildren []Inode
	for _, child := range f.children {
		clone := child.Clone()
		tempChildren = append(tempChildren, clone)
	}
	cloneFolder.children = tempChildren
	return cloneFolder
}

func (f *Folder) GetName() string {
	return f.name
}

func (f *Folder) GetType() string {
	return fmt.Sprintf("%T", f)
}

func (f *Folder) GetChildren() []Inode {
	return f.children
}
