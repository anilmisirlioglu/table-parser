package table

type Table struct {
	Header Header
	Rows   []Row
}

type Header struct {
	Text  string
	Cells []HeaderCell
}

type HeaderCell struct {
	Key   string
	Index int
}

type Row struct {
	Text  string
	Cells []RowCell
}

type RowCell struct {
	Value    string
	Relation string
}

func (h *Header) append(cell HeaderCell) {
	h.Cells = append(h.Cells, cell)
}

func (h *Row) append(cell RowCell) {
	h.Cells = append(h.Cells, cell)
}
