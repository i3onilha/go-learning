package prototype

type Inode interface {
	GetName() string
	GetType() string
	GetChildren() []Inode
	Clone() Inode
	Print(string)
}
