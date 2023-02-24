package table

import (
	"reflect"
	"strings"
	"testing"
)

var tests = []struct {
	name string
	text string
	want *Table
}{
	{
		name: "test full table",
		text: `
NAME                ROLES                      OS IMAGE             KERNEL-VERSION      CONTAINER RUNTIME
master-1            control-plane,master       Ubuntu 20.04.1 LTS   5.13.0-39-generic   containerd://1.4.1
monitor-1           monitor                    Ubuntu 20.04.2 LTS   5.13.0-39-generic   containerd://1.4.2
worker-1            worker                     Ubuntu 20.04.3 LTS   5.13.0-39-generic   containerd://1.4.3
`,
		want: &Table{
			Header: Header{
				Text: "NAME                ROLES                      OS IMAGE             KERNEL-VERSION      CONTAINER RUNTIME",
				Cells: []HeaderCell{
					{Index: 0, Key: "NAME"},
					{Index: 20, Key: "ROLES"},
					{Index: 47, Key: "OS IMAGE"},
					{Index: 68, Key: "KERNEL-VERSION"},
					{Index: 88, Key: "CONTAINER RUNTIME"},
				},
			},
			Rows: []Row{
				{
					Text: "master-1            control-plane,master       Ubuntu 20.04.1 LTS   5.13.0-39-generic   containerd://1.4.1",
					Cells: []RowCell{
						{Relation: "NAME", Value: "master-1"},
						{Relation: "ROLES", Value: "control-plane,master"},
						{Relation: "OS IMAGE", Value: "Ubuntu 20.04.1 LTS"},
						{Relation: "KERNEL-VERSION", Value: "5.13.0-39-generic"},
						{Relation: "CONTAINER RUNTIME", Value: "containerd://1.4.1"},
					},
				},
				{
					Text: "monitor-1           monitor                    Ubuntu 20.04.2 LTS   5.13.0-39-generic   containerd://1.4.2",
					Cells: []RowCell{
						{Relation: "NAME", Value: "monitor-1"},
						{Relation: "ROLES", Value: "monitor"},
						{Relation: "OS IMAGE", Value: "Ubuntu 20.04.2 LTS"},
						{Relation: "KERNEL-VERSION", Value: "5.13.0-39-generic"},
						{Relation: "CONTAINER RUNTIME", Value: "containerd://1.4.2"},
					},
				},
				{
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
		},
	},
	{
		name: "test slipped table",
		text: `
NAME                ROLES                      OS IMAGE             KERNEL-VERSION      CONTAINER RUNTIME
master-1            control-plane,master       Ubuntu 20.04.1 LTS   5.13.0-39-generic   

      monitor-1     monitor                    Ubuntu 20.04.2 LTS   5.13.0-39-generic       containerd://1.4.2

                          worker                                    5.13.0-39-generic   containerd://1.4.3
`,
		want: &Table{
			Header: Header{
				Text: "NAME                ROLES                      OS IMAGE             KERNEL-VERSION      CONTAINER RUNTIME",
				Cells: []HeaderCell{
					{Index: 0, Key: "NAME"},
					{Index: 20, Key: "ROLES"},
					{Index: 47, Key: "OS IMAGE"},
					{Index: 68, Key: "KERNEL-VERSION"},
					{Index: 88, Key: "CONTAINER RUNTIME"},
				},
			},
			Rows: []Row{
				{
					Text: "master-1            control-plane,master       Ubuntu 20.04.1 LTS   5.13.0-39-generic",
					Cells: []RowCell{
						{Relation: "NAME", Value: "master-1"},
						{Relation: "ROLES", Value: "control-plane,master"},
						{Relation: "OS IMAGE", Value: "Ubuntu 20.04.1 LTS"},
						{Relation: "KERNEL-VERSION", Value: "5.13.0-39-generic"},
						{Relation: "CONTAINER RUNTIME"},
					},
				},
				{
					Text: "monitor-1     monitor                    Ubuntu 20.04.2 LTS   5.13.0-39-generic       containerd://1.4.2",
					Cells: []RowCell{
						{Relation: "NAME", Value: "monitor-1"},
						{Relation: "ROLES", Value: "monitor"},
						{Relation: "OS IMAGE", Value: "Ubuntu 20.04.2 LTS"},
						{Relation: "KERNEL-VERSION", Value: "5.13.0-39-generic"},
						{Relation: "CONTAINER RUNTIME", Value: "containerd://1.4.2"},
					},
				},
				{
					Text: "worker                                    5.13.0-39-generic   containerd://1.4.3",
					Cells: []RowCell{
						{Relation: "NAME"},
						{Relation: "ROLES", Value: "worker"},
						{Relation: "OS IMAGE"},
						{Relation: "KERNEL-VERSION", Value: "5.13.0-39-generic"},
						{Relation: "CONTAINER RUNTIME", Value: "containerd://1.4.3"},
					},
				},
			},
		},
	},
	{
		name: "test empty table",
		text: `NAME                ROLES                      OS IMAGE             KERNEL-VERSION      CONTAINER RUNTIME`,
		want: &Table{
			Header: Header{
				Text: "NAME                ROLES                      OS IMAGE             KERNEL-VERSION      CONTAINER RUNTIME",
				Cells: []HeaderCell{
					{Index: 0, Key: "NAME"},
					{Index: 20, Key: "ROLES"},
					{Index: 47, Key: "OS IMAGE"},
					{Index: 68, Key: "KERNEL-VERSION"},
					{Index: 88, Key: "CONTAINER RUNTIME"},
				},
			},
			Rows: make([]Row, 0),
		},
	},
	{
		name: "test empty input",
		text: "",
		want: &Table{
			Header: Header{},
			Rows:   make([]Row, 0),
		},
	},
}

func TestNewReader(t *testing.T) {
	r := NewReader(strings.NewReader("A\tB\tC"))
	if r == nil {
		t.Error("reader must be not nil")
	}

	r = NewReader(strings.NewReader("     "))
	if r == nil {
		t.Errorf("reader must not be nil")
	}
}

func TestReader(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewReader(strings.NewReader(tt.text))

			rows := make([]Row, 0)
			for r.Next() {
				rows = append(rows, r.Row())
			}

			got := &Table{Header: r.Header(), Rows: rows}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewReader() failed\n got: %v\nwant: %v", got, tt.want)
			}
		})
	}
}

func TestReadAll(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReadAll(tt.text); !reflect.DeepEqual(*got, *tt.want) {
				t.Errorf("ReadAll() failed\n got: %v\nwant: %v", got, tt.want)
			}
		})
	}
}
