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

type FileList struct {
	Files []File
}

func (fl *FileList) Append(n File) {
	fl.Files = append(fl.Files, n)
}

func (fl *FileList) Pretty() {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Hash", "File Name", "Size (bytes)"})

	for _, f := range fl.Files {
		t.AppendRow([]interface{}{f.MD5, f.Name, f.Size})
	}
	t.Render()
}
