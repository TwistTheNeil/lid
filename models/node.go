package models

import (
	"lid/services/file_ops"
	"lid/services/logger"
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
)

type Node struct {
	Name string
	Size int64
	MD5  string
}

func (n *Node) Hash() {
	log := logger.CreateLogger("node.Hash")
	log.Trace("node.Hash called")

	f, err := os.Open(n.Name)
	if err != nil {
		log.Error("something went wrong opening file", err, "file", n.Name)
	}
	defer f.Close()

	n.MD5 = file_ops.MD5Hash(f)
}

type NodeList struct {
	Nodes []Node
}

func (nl *NodeList) Append(n Node) {
	nl.Nodes = append(nl.Nodes, n)
}

func (nl *NodeList) Pretty() {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Hash", "File Name", "Size (bytes)"})

	for _, n := range nl.Nodes {
		t.AppendRow([]interface{}{n.MD5, n.Name, n.Size})
	}
	t.Render()
}
