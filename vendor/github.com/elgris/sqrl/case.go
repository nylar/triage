package sqrl

import (
	"bytes"
	"errors"
)

// sqlizerBuffer is a helper that allows to write many Sqlizers one by one
// without constant checks for errors that may come from Sqlizer
type sqlizerBuffer struct {
	bytes.Buffer
	args []interface{}
	err  error
}

// WriteSql converts Sqlizer to SQL strings and writes it to buffer
func (b *sqlizerBuffer) WriteSql(item Sqlizer) {
	if b.err != nil {
		return
	}

	var str string
	var args []interface{}
	str, args, b.err = item.ToSql()

	if b.err != nil {
		return
	}

	b.WriteString(str)
	b.WriteByte(' ')
	b.args = append(b.args, args...)
}

func (b *sqlizerBuffer) ToSql() (string, []interface{}, error) {
	return b.String(), b.args, b.err
}

// whenPart is a helper structure to describe SQLs "WHEN ... THEN ..." expression
type whenPart struct {
	when Sqlizer
	then Sqlizer
}

func newWhenPart(when interface{}, then interface{}) whenPart {
	return whenPart{newPart(when), newPart(then)}
}

// CaseBuilder builds SQL CASE construct which could be used as parts of queries.
type CaseBuilder struct {
	whatPart  Sqlizer
	whenParts []whenPart
	elsePart  Sqlizer
}

// ToSql implements Sqlizer
func (b *CaseBuilder) ToSql() (sqlStr string, args []interface{}, err error) {
	if len(b.whenParts) == 0 {
		err = errors.New("case expression must contain at lease one WHEN clause")

		return
	}

	sql := sqlizerBuffer{}

	sql.WriteString("CASE ")
	if b.whatPart != nil {
		sql.WriteSql(b.whatPart)
	}

	for _, p := range b.whenParts {
		sql.WriteString("WHEN ")
		sql.WriteSql(p.when)
		sql.WriteString("THEN ")
		sql.WriteSql(p.then)
	}

	if b.elsePart != nil {
		sql.WriteString("ELSE ")
		sql.WriteSql(b.elsePart)
	}

	sql.WriteString("END")

	return sql.ToSql()
}

// what sets optional value for CASE construct "CASE [value] ..."
func (b *CaseBuilder) what(expr interface{}) *CaseBuilder {
	b.whatPart = newPart(expr)
	return b
}

// When adds "WHEN ... THEN ..." part to CASE construct
func (b *CaseBuilder) When(when interface{}, then interface{}) *CaseBuilder {
	// TODO: performance hint: replace slice of WhenPart with just slice of parts
	// where even indices of the slice belong to "when"s and odd indices belong to "then"s
	b.whenParts = append(b.whenParts, newWhenPart(when, then))
	return b
}

// Else sets optional "ELSE ..." part for CASE construct
func (b *CaseBuilder) Else(expr interface{}) *CaseBuilder {
	b.elsePart = newPart(expr)
	return b

}
