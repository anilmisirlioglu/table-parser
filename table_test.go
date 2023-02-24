package table

import (
	"reflect"
	"runtime"
	"testing"
)

const input = `
NAME                ROLES                      OS IMAGE             KERNEL-VERSION      CONTAINER RUNTIME
master-1            control-plane,master       Ubuntu 20.04.1 LTS   5.13.0-39-generic   containerd://1.4.1
monitor-1           monitor                    Ubuntu 20.04.2 LTS   5.13.0-39-generic   containerd://1.4.2
worker-1            worker                     Ubuntu 20.04.3 LTS   5.13.0-39-generic   containerd://1.4.3
`

func TestTable_Row(t *testing.T) {
	tx := ReadAll(input)

	tests := []struct {
		row  *Row
		want *Row
	}{
		{
			tx.Row(0),
			&Row{
				Text: "master-1            control-plane,master       Ubuntu 20.04.1 LTS   5.13.0-39-generic   containerd://1.4.1",
				Cells: []RowCell{
					{Relation: "NAME", Value: "master-1"},
					{Relation: "ROLES", Value: "control-plane,master"},
					{Relation: "OS IMAGE", Value: "Ubuntu 20.04.1 LTS"},
					{Relation: "KERNEL-VERSION", Value: "5.13.0-39-generic"},
					{Relation: "CONTAINER RUNTIME", Value: "containerd://1.4.1"},
				},
			}},
		{
			tx.Row(1),
			&Row{
				Text: "monitor-1           monitor                    Ubuntu 20.04.2 LTS   5.13.0-39-generic   containerd://1.4.2",
				Cells: []RowCell{
					{Relation: "NAME", Value: "monitor-1"},
					{Relation: "ROLES", Value: "monitor"},
					{Relation: "OS IMAGE", Value: "Ubuntu 20.04.2 LTS"},
					{Relation: "KERNEL-VERSION", Value: "5.13.0-39-generic"},
					{Relation: "CONTAINER RUNTIME", Value: "containerd://1.4.2"},
				},
			},
		},
		{
			tx.Row(2),
			&Row{
				Text: "worker-1            worker                     Ubuntu 20.04.3 LTS   5.13.0-39-generic   containerd://1.4.3",
				Cells: []RowCell{
					{Relation: "NAME", Value: "worker-1"},
					{Relation: "ROLES", Value: "worker"},
					{Relation: "OS IMAGE", Value: "Ubuntu 20.04.3 LTS"},
					{Relation: "KERNEL-VERSION", Value: "5.13.0-39-generic"},
					{Relation: "CONTAINER RUNTIME", Value: "containerd://1.4.3"},
				},
			},
		},
		{tx.Row(3), nil},
		{tx.Row(99), nil},
	}
	for _, tt := range tests {
		if !reflect.DeepEqual(tt.row, tt.want) {
			t.Errorf("Row failed\n got: %v\nwant: %v", tt.row.Text, tt.want)
		}
	}
}

func TestRow_Cell(t *testing.T) {
	tx := ReadAll(input)
	row := tx.Row(0)

	tests := []struct {
		cell *RowCell
		want *RowCell
	}{
		{row.Cell(0), &RowCell{
			Value:    "master-1",
			Relation: "NAME",
		}},
		{row.Cell(1), &RowCell{
			Value:    "control-plane,master",
			Relation: "ROLES",
		}},
		{row.Cell(2), &RowCell{
			Value:    "Ubuntu 20.04.1 LTS",
			Relation: "OS IMAGE",
		}},
		{row.Cell(3), &RowCell{
			Value:    "5.13.0-39-generic",
			Relation: "KERNEL-VERSION",
		}},
		{row.Cell(4), &RowCell{
			Value:    "containerd://1.4.1",
			Relation: "CONTAINER RUNTIME",
		}},
		{row.Cell(5), nil},
		{row.Cell(99), nil},
	}
	for _, tt := range tests {
		if !reflect.DeepEqual(tt.cell, tt.want) {
			t.Errorf("Cell failed\n got: %v\nwant: %v", tt.cell, tt.want)
		}
	}
}

func TestRow_CellByName(t *testing.T) {
	tx := ReadAll(input)
	row := tx.Row(0)

	tests := []struct {
		cell *RowCell
		want *RowCell
	}{
		{row.CellByName("NAME"), &RowCell{
			Value:    "master-1",
			Relation: "NAME",
		}},
		{row.CellByName("ROLES"), &RowCell{
			Value:    "control-plane,master",
			Relation: "ROLES",
		}},
		{row.CellByName("OS IMAGE"), &RowCell{
			Value:    "Ubuntu 20.04.1 LTS",
			Relation: "OS IMAGE",
		}},
		{row.CellByName("KERNEL-VERSION"), &RowCell{
			Value:    "5.13.0-39-generic",
			Relation: "KERNEL-VERSION",
		}},
		{row.CellByName("CONTAINER RUNTIME"), &RowCell{
			Value:    "containerd://1.4.1",
			Relation: "CONTAINER RUNTIME",
		}},
		{row.CellByName("NOTHING"), nil},
		{row.CellByName(runtime.GOOS), nil},
	}
	for _, tt := range tests {
		if !reflect.DeepEqual(tt.cell, tt.want) {
			t.Errorf("CellByName failed\n got: %v\nwant: %v", tt.cell, tt.want)
		}
	}
}
