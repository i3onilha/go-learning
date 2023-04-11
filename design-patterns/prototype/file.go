package prototype

import "fmt"

type File struct {
	name string
}

func (f *File) Print(identation string) {
	fmt.Println(identation + f.name)
}

func (f *File) Clone() Inode {
	return &File{name: f.name + "_clone"}
}

func (f *File) GetName() string {
	return f.name
}

func (f *File) GetType() string {
	return fmt.Sprintf("%T", f)
}

func (f *File) GetChildren() []Inode {
	return nil
}
