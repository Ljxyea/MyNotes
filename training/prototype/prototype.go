package prototype

import "fmt"

// 原型模式
type Inode interface {
	print(string)
	clone() Inode
}

// 具体原型
type File struct {
	name string
}

func (f *File) print(s string) {
	fmt.Println(s, f.name)
}

func (f *File) clone() Inode {
	return &File{
		name: f.name + "_clone",
	}
}
