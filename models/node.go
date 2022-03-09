package models

import (
	"os"
	"time"

	"github.com/jedib0t/go-pretty/v6/table"
)

type Node struct {
	Name      string    `gorm:"primaryKey" json:"name"`
	Size      int64     `json:"size"`
	MD5       string    `gorm:"primaryKey" json:"hash"`
	CreatedAt time.Time `json:"-"`
	Devices   []*Device `gorm:"many2many:node_devices;" json:"-"`
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
