package repsitories

type NodeRepository interface {
	Create(string, string, int64)
	FindByName(string)
	FindByMD5(string)
	FindAll()
	DeleteByName(string)
	DeleteByMD5(string)
}
