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

func (t *Table) Row(index int) Row {
	if len(t.Rows) > index {
		return t.Rows[index]
	}
	return Row{}
}

func (h *Header) append(cell HeaderCell) {
	h.Cells = append(h.Cells, cell)
}

func (r *Row) append(cell RowCell) {
	r.Cells = append(r.Cells, cell)
}

func (r *Row) Cell(index int) RowCell {
	if len(r.Cells) > index {
		return r.Cells[index]
	}
	return RowCell{}
}

func (r *Row) CellByName(s string) RowCell {
	for _, cell := range r.Cells {
		if cell.Relation == s {
			return cell
		}
	}
	return RowCell{}
}
