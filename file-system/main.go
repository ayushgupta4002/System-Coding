package main

import (
	"fmt"
	"strings"
)

type Node interface {
	getName() string
}

type Directory struct {
	name     string
	parent   *Directory
	children map[string]Node
}

func (d *Directory) removeDirectory() error {
	_, exists := d.parent.children[d.name]
	if !exists {
		return fmt.Errorf("no such file exists to be deleted")
	}
	delete(d.parent.children, d.name)
	return nil
}

func (d *Directory) addChildren(child Node) {

	name := child.getName()

	d.children[name] = child

}

func (d *Directory) makeDirectory(name string) (*Directory, error) {
	newDir := &Directory{
		name:     name,
		parent:   d,
		children: make(map[string]Node),
	}

	d.addChildren(newDir)
	return newDir, nil
}

func (d *Directory) getName() string {
	return d.name
}

type File struct {
	name    string
	parent  *Directory
	content string
}

func (f *File) removeFile() error {

	_, exists := f.parent.children[f.name]
	if !exists {
		return fmt.Errorf("no such file exists to be deleted")
	}

	delete(f.parent.children, f.name)
	return nil
}

func (f *File) getContent() string {
	return f.content
}

func (f *File) addContent(content string) error {
	f.content = content
	return nil
}

func (f *File) getName() string {
	return f.name
}

func mkdir(path string, parent *Directory) error {
	parts := strings.Split(path, "/")
	curr := parent
	for _, v := range parts {
		if v == "" {
			continue
		}
		child, exists := curr.children[v]

		if !exists {
			curr.makeDirectory(v)
		}
		dir, ok := child.(*Directory)
		if !ok {
			return fmt.Errorf("%s is a file", v)
		}
		curr = dir

	}
	return nil
}

func touch(path string, parent *Directory) error {
	parts := strings.Split(path, "/")
	fileName := parts[len(parts)-1]
	curr := parent

	for _, dirName := range parts[1 : len(parts)-1] {

		child, exists := curr.children[dirName]
		if !exists {
			return fmt.Errorf("directory %s not found", dirName)
		}

		dir, ok := child.(*Directory)
		if !ok {
			return fmt.Errorf("%s is a file", dirName)
		}

		curr = dir
	}

	if _, exists := curr.children[fileName]; exists {
		return fmt.Errorf("file already exists")
	}

	curr.addChildren(&File{
		name:    fileName,
		parent:  curr,
		content: "",
	})

	return nil
}

func ls(path string, curr *Directory) error {
	parts := strings.Split(path, "/")

	for _, v := range parts {
		if v == "" {
			continue
		}
		child, exists := curr.children[v]

		if !exists {
			return fmt.Errorf("no directory with name %s found", v)
		}
		dir, ok := child.(*Directory)
		if !ok {
			return fmt.Errorf("%s is a file", v)
		}
		curr = dir

	}

	for name, node := range curr.children {

		_, ok := node.(*Directory)

		if ok {
			fmt.Printf("\n [DIR] : %s", name)
		} else {
			fmt.Printf("\n [FILE] : %s", name)
		}

	}

	return nil
}

func rm(path string, curr *Directory) error {

	parts := strings.Split(path, "/")

	for _, v := range parts[0 : len(parts)-1] {
		if v == "" {
			continue
		}
		child, exists := curr.children[v]

		if !exists {
			return fmt.Errorf("no directory with name %s found", v)
		}
		dir, ok := child.(*Directory)
		if !ok {
			return fmt.Errorf("%s is a file", v)
		}
		curr = dir

	}

	node, exists := curr.children[parts[len(parts)-1]]
	if !exists {
		return fmt.Errorf("no file or directory found")
	}

	switch n := node.(type) {
	case *Directory:
		n.removeDirectory()
	case *File:
		n.removeFile()
	default:
		return fmt.Errorf("unknown node type")
	}

	return nil

}

func cat(path string, curr *Directory) (string, error) {

	parts := strings.Split(path, "/")

	for _, v := range parts[0 : len(parts)-1] {
		if v == "" {
			continue
		}
		child, exists := curr.children[v]

		if !exists {
			return "", fmt.Errorf("no directory with name %s found", v)
		}
		dir, ok := child.(*Directory)
		if !ok {
			return "", fmt.Errorf("%s is a file", v)
		}
		curr = dir

	}

	node, exists := curr.children[parts[len(parts)-1]]
	if !exists {
		return "", fmt.Errorf("no file or directory found")
	}

	switch n := node.(type) {
	case *Directory:
		return "", fmt.Errorf("not a file")
	case *File:
		v := n.getContent()
		fmt.Printf("\n content in file %s is:", n.name)
		fmt.Printf("\n %s", v)
		return v, nil
	default:
		return "", fmt.Errorf("unknown file type")
	}

}

func vim(path string, curr *Directory, content string) error {
	parts := strings.Split(path, "/")

	for _, v := range parts[0 : len(parts)-1] {
		if v == "" {
			continue
		}
		child, exists := curr.children[v]

		if !exists {
			return fmt.Errorf("no directory with name %s found", v)
		}
		dir, ok := child.(*Directory)
		if !ok {
			return fmt.Errorf("%s is a file", v)
		}
		curr = dir

	}

	node, exists := curr.children[parts[len(parts)-1]]
	if !exists {
		return fmt.Errorf("no file or directory found")
	}

	switch n := node.(type) {
	case *Directory:
		return fmt.Errorf("no file with this name found")
	case *File:
		n.addContent(content)

	default:
		return fmt.Errorf("unknown node type")
	}

	return nil

}
func main() {

	baseDir := Directory{
		name:     "/",
		parent:   nil,
		children: make(map[string]Node, 0),
	}
	mkdir("/ayush", &baseDir)
	mkdir("/ayush/work", &baseDir)
	mkdir("/ayush/work/workFile.txt", &baseDir)
	touch("/ayush/file1.txt", &baseDir)
	touch("/ayush/file2.txt", &baseDir)

	ls("/ayush", &baseDir)

	rm("/ayush/file2.txt", &baseDir)
	ls("/ayush", &baseDir)

	vim("/ayush/file1.txt", &baseDir, "hello this is a test message")
	cat("/ayush/file1.txt", &baseDir)
}
