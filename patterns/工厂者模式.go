package main

import "io"

/*


*/
type Store interface {
	Open(string) (io.ReadWriteCloser, error)
}

type StorageType int

const (
	DiskStorage StorageType = 1 << iota
	TempStorage
	MemoryStorage
)

func newMemoryStorage() Store {
	var f Store
	// 新增逻辑
	return f
}

func nowDiskStorage() Store {
	var f Store
	// 新增逻辑
	return f
}

func newTempStroage() Store {
	var f Store
	// 新增逻辑
	return f
}


func NewStore(t StorageType) Store {
	switch t{
	case MemoryStorage:
		return newMemoryStorage()
	case DiskStorage:
		return nowDiskStorage()
	default:
		return newTempStroage()
	}

}


func main(){
	s := NewStore(MemoryStorage)
	f, _ := s.Open("file")
	f.Write([]byte("data"))
	defer f.Close()
}