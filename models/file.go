package models

import (
	"os"
	"time"

	"github.com/jedib0t/go-pretty/v6/table"
)

type File struct {
	Name      string    `gorm:"primaryKey" json:"name"`
	Size      int64     `json:"size"`
	MD5       string    `gorm:"primaryKey" json:"hash"`
	CreatedAt time.Time `json:"-"`
	Devices   []*Device `gorm:"many2many:file_devices;" json:"devices"`
}

type NodeList struct {
	Nodes []File
}

func (nl *NodeList) Append(n File) {
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
