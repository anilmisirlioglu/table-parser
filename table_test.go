package table

import (
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
		row  Row
		want string
	}{
		{tx.Row(0), "master-1            control-plane,master       Ubuntu 20.04.1 LTS   5.13.0-39-generic   containerd://1.4.1"},
		{tx.Row(1), "monitor-1           monitor                    Ubuntu 20.04.2 LTS   5.13.0-39-generic   containerd://1.4.2"},
		{tx.Row(2), "worker-1            worker                     Ubuntu 20.04.3 LTS   5.13.0-39-generic   containerd://1.4.3"},
		{tx.Row(3), ""},
		{tx.Row(99), ""},
	}
	for _, tt := range tests {
		if tt.row.Text != tt.want {
			t.Errorf("Row: got = %v, want %v", tt.row.Text, tt.want)
		}
	}
}

func TestRow_Cell(t *testing.T) {
	tx := ReadAll(input)
	row := tx.Row(0)

	tests := []struct {
		cell RowCell
		want string
	}{
		{row.Cell(0), "master-1"},
		{row.Cell(1), "control-plane,master"},
		{row.Cell(2), "Ubuntu 20.04.1 LTS"},
		{row.Cell(3), "5.13.0-39-generic"},
		{row.Cell(4), "containerd://1.4.1"},
		{row.Cell(5), ""},
		{row.Cell(99), ""},
	}
	for _, tt := range tests {
		if tt.cell.Value != tt.want {
			t.Errorf("Cell: got = %v, want %v", tt.cell.Value, tt.want)
		}
	}
}

func TestRow_CellByName(t *testing.T) {
	tx := ReadAll(input)
	row := tx.Row(0)

	tests := []struct {
		cell RowCell
		want string
	}{
		{row.CellByName("NAME"), "master-1"},
		{row.CellByName("ROLES"), "control-plane,master"},
		{row.CellByName("OS IMAGE"), "Ubuntu 20.04.1 LTS"},
		{row.CellByName("KERNEL-VERSION"), "5.13.0-39-generic"},
		{row.CellByName("CONTAINER RUNTIME"), "containerd://1.4.1"},
		{row.CellByName("NOTHING"), ""},
		{row.CellByName(runtime.GOOS), ""},
	}
	for _, tt := range tests {
		if tt.cell.Value != tt.want {
			t.Errorf("CellByName: got = %v, want %v", tt.cell.Value, tt.want)
		}
	}
}
