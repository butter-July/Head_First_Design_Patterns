package main

func main() {
	file1 := &file{name: "file1"}
	file2 := &file{name: "File2"}
	file3 := &file{name: "file3"}
	file3.search("rose")
	Folder1 := &Folder{name: "Folder1"}
	Folder1.add(file1)
	Folder1.add(file2)
	Folder1.search("sky")
}
