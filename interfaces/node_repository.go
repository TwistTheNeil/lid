package interfaces

type NodeRepository interface {
	Create(node_name string, node_hash string, node_size int64)
	FindByName(node_name string)
	FindByMD5(node_md5_hash string)
	FindAll()
	DeleteByName(node_name string)
	DeleteByMD5(node_md5_hash string)
}
