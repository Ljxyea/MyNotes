package composite

import "fmt"

//组合模式
//将对象组合成树状结构， 并像独立对象一样使用他们

// 组件接口
type File struct {
	name string
}

func (f *File) search(keyword string) {
	fmt.Printf("Searching file %s for keyword %s\n", f.name, keyword)
}

func (f *File) getName() string {
	return f.name
}

// 组合
type Folder struct {
	components []Component
	name       string
}

func (f *Folder) add(component Component) {
	f.components = append(f.components, component)
}

func (f *Folder) search(keyword string) {
	for _, component := range f.components {
		component.search(keyword)
	}
}

// 叶子
type Component interface {
	search(keyword string)
}
