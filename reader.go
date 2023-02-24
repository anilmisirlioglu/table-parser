package table

import (
	"bufio"
	"io"
	"strings"
	"unicode"
)

type Reader struct {
	h Header
	s *bufio.Scanner
}

func NewReader(rd io.Reader) *Reader {
	s := bufio.NewScanner(rd)

	h := ""
	for h == "" && s.Scan() {
		h = strings.TrimSpace(s.Text())
	}

	if h == "" {
		return &Reader{s: s}
	}

	return &Reader{
		h: parseHeader(tabToSpace(h)),
		s: s,
	}
}

func (r *Reader) Next() bool {
	if r.s.Scan() {
		if strings.TrimSpace(r.s.Text()) == "" {
			return r.Next()
		}
		return true
	}
	return false
}

func (r *Reader) Row() Row {
	return parseRow(tabToSpace(r.s.Text()), r.h)
}

func (r *Reader) Header() Header {
	return r.h
}

func ReadAll(s string) *Table {
	items := strings.Split(tabToSpace(s), "\n")

	var rows []string
	for _, item := range items {
		if strings.TrimSpace(item) != "" {
			rows = append(rows, item)
		}
	}

	rl := len(rows)
	if rl <= 1 {
		t := &Table{Rows: make([]Row, 0)}
		if rl == 1 {
			t.Header = parseHeader(rows[0])
		}

		return t
	}

	table := &Table{
		Header: parseHeader(rows[0]),
		Rows:   make([]Row, 0, rl-1),
	}

	for _, row := range rows[1:] {
		table.Rows = append(table.Rows, parseRow(row, table.Header))
	}

	return table
}

func parseHeader(h string) Header {
	h = strings.TrimSpace(h)

	var (
		newc   bool
		space  = -1
		index  int
		header = Header{
			Text:  h,
			Cells: []HeaderCell{},
		}
	)

	runes := []rune(h)
	for i, ru := range runes {
		if i == len(runes)-1 {
			header.append(HeaderCell{
				Key:   string(runes[index:]),
				Index: index,
			})
			continue
		}

		if unicode.IsSpace(ru) {
			if space > 0 {
				if newc {
					header.append(HeaderCell{
						Key:   string(runes[index : i-1]),
						Index: index,
					})
					space = -1
					newc = false
				}
			} else {
				space++
			}
		} else {
			if !newc && space > 0 {
				space = -1
			}

			switch space {
			case -1:
				index = i
				space = 0
				newc = true
			default:
				space = 0
			}
		}
	}

	return header
}

func parseRow(s string, header Header) Row {
	cells := header.Cells
	if strings.TrimSpace(s) == "" {
		row := Row{}
		for _, cell := range cells {
			row.append(RowCell{Relation: cell.Key})
		}

		return row
	}

	row := Row{Text: strings.TrimSpace(s), Cells: make([]RowCell, 0, len(cells))}
	for i, h := range cells {
		curr := h.Index
		next := len(s)
		if i+1 != len(cells) {
			next = cells[i+1].Index
		}

		row.append(RowCell{
			Relation: h.Key,
			Value:    strings.TrimSpace(s[curr:next]),
		})
	}

	return row
}

func tabToSpace(s string) string {
	return strings.ReplaceAll(s, "\t", "    ")
}
