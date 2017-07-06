// CAUTION: Generated file - DO NOT EDIT.

// Copyright 2017 The Prometheus Authors
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package textparse

import (
	"fmt"
	"math"
	"strconv"
	"unicode/utf8"

	"github.com/prometheus/prometheus/pkg/value"
)

// Lex is called by the parser generated by "go tool yacc" to obtain each
// token. The method is opened before the matching rules block and closed at
// the end of the file.
func (l *lexer) Lex() int {
	const (
		lstateInit = iota
		lstateName
		lstateValue
		lstateTimestamp
		lstateLabels
		lstateLName
		lstateLValue
		lstateLValueIn
	)
	s := lstateInit

	if l.i >= len(l.b) {
		return eof
	}
	c := l.b[l.i]

	l.ts = nil
	l.mstart = l.nextMstart
	l.offsets = l.offsets[:0]

yystate0:

	switch yyt := s; yyt {
	default:
		panic(fmt.Errorf(`invalid start condition %d`, yyt))
	case 0: // start condition: INITIAL
		goto yystart1
	case 1: // start condition: lstateName
		goto yystart7
	case 2: // start condition: lstateValue
		goto yystart10
	case 3: // start condition: lstateTimestamp
		goto yystart16
	case 4: // start condition: lstateLabels
		goto yystart21
	case 5: // start condition: lstateLName
		goto yystart26
	case 6: // start condition: lstateLValue
		goto yystart30
	case 7: // start condition: lstateLValueIn
		goto yystart33
	}

	goto yystate0 // silence unused label error
	goto yystate1 // silence unused label error
yystate1:
	c = l.next()
yystart1:
	switch {
	default:
		goto yyabort
	case c == '#':
		goto yystate4
	case c == '\t' || c == '\n' || c == '\r' || c == ' ':
		goto yystate3
	case c == '\x00':
		goto yystate2
	case c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'z':
		goto yystate6
	}

yystate2:
	c = l.next()
	goto yyrule1

yystate3:
	c = l.next()
	switch {
	default:
		goto yyrule3
	case c == '\t' || c == '\n' || c == '\r' || c == ' ':
		goto yystate3
	}

yystate4:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c == '\n':
		goto yystate5
	case c >= '\x01' && c <= '\t' || c == '\v' || c == '\f' || c >= '\x0e' && c <= 'ÿ':
		goto yystate4
	}

yystate5:
	c = l.next()
	goto yyrule2

yystate6:
	c = l.next()
	switch {
	default:
		goto yyrule4
	case c >= '0' && c <= ':' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate6
	}

	goto yystate7 // silence unused label error
yystate7:
	c = l.next()
yystart7:
	switch {
	default:
		goto yyabort
	case c == '\t' || c == ' ':
		goto yystate8
	case c == '{':
		goto yystate9
	}

yystate8:
	c = l.next()
	switch {
	default:
		goto yyrule6
	case c == '\t' || c == ' ':
		goto yystate8
	case c == '{':
		goto yystate9
	}

yystate9:
	c = l.next()
	goto yyrule5

	goto yystate10 // silence unused label error
yystate10:
	c = l.next()
yystart10:
	switch {
	default:
		goto yyabort
	case c == 'N':
		goto yystate13
	case c == '\t' || c == ' ':
		goto yystate12
	case c >= '\x01' && c <= '\b' || c == '\v' || c == '\f' || c >= '\x0e' && c <= '\x1f' || c >= '!' && c <= 'M' || c >= 'O' && c <= 'ÿ':
		goto yystate11
	}

yystate11:
	c = l.next()
	switch {
	default:
		goto yyrule17
	case c >= '\x01' && c <= '\b' || c == '\v' || c == '\f' || c >= '\x0e' && c <= '\x1f' || c >= '!' && c <= 'ÿ':
		goto yystate11
	}

yystate12:
	c = l.next()
	switch {
	default:
		goto yyrule15
	case c == '\t' || c == ' ':
		goto yystate12
	}

yystate13:
	c = l.next()
	switch {
	default:
		goto yyrule17
	case c == 'a':
		goto yystate14
	case c >= '\x01' && c <= '\b' || c == '\v' || c == '\f' || c >= '\x0e' && c <= '\x1f' || c >= '!' && c <= '`' || c >= 'b' && c <= 'ÿ':
		goto yystate11
	}

yystate14:
	c = l.next()
	switch {
	default:
		goto yyrule17
	case c == 'N':
		goto yystate15
	case c >= '\x01' && c <= '\b' || c == '\v' || c == '\f' || c >= '\x0e' && c <= '\x1f' || c >= '!' && c <= 'M' || c >= 'O' && c <= 'ÿ':
		goto yystate11
	}

yystate15:
	c = l.next()
	switch {
	default:
		goto yyrule16
	case c >= '\x01' && c <= '\b' || c == '\v' || c == '\f' || c >= '\x0e' && c <= '\x1f' || c >= '!' && c <= 'ÿ':
		goto yystate11
	}

	goto yystate16 // silence unused label error
yystate16:
	c = l.next()
yystart16:
	switch {
	default:
		goto yyabort
	case c == '\n' || c == '\r':
		goto yystate19
	case c == '\t' || c == ' ':
		goto yystate18
	case c == '\x00':
		goto yystate17
	case c >= '0' && c <= '9':
		goto yystate20
	}

yystate17:
	c = l.next()
	goto yyrule21

yystate18:
	c = l.next()
	switch {
	default:
		goto yyrule18
	case c == '\t' || c == ' ':
		goto yystate18
	}

yystate19:
	c = l.next()
	switch {
	default:
		goto yyrule20
	case c == '\n' || c == '\r':
		goto yystate19
	}

yystate20:
	c = l.next()
	switch {
	default:
		goto yyrule19
	case c >= '0' && c <= '9':
		goto yystate20
	}

	goto yystate21 // silence unused label error
yystate21:
	c = l.next()
yystart21:
	switch {
	default:
		goto yyrule9
	case c == ',':
		goto yystate23
	case c == '\t' || c == ' ':
		goto yystate22
	case c == '}':
		goto yystate25
	}

yystate22:
	c = l.next()
	switch {
	default:
		goto yyrule7
	case c == '\t' || c == ' ':
		goto yystate22
	}

yystate23:
	c = l.next()
	switch {
	default:
		goto yyrule9
	case c == '\t' || c == ' ':
		goto yystate24
	case c == '}':
		goto yystate25
	}

yystate24:
	c = l.next()
	switch {
	default:
		goto yyrule9
	case c == '\t' || c == ' ':
		goto yystate24
	}

yystate25:
	c = l.next()
	goto yyrule8

	goto yystate26 // silence unused label error
yystate26:
	c = l.next()
yystart26:
	switch {
	default:
		goto yyabort
	case c == '=':
		goto yystate28
	case c == '\t' || c == ' ':
		goto yystate27
	case c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'z':
		goto yystate29
	}

yystate27:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c == '=':
		goto yystate28
	case c == '\t' || c == ' ':
		goto yystate27
	}

yystate28:
	c = l.next()
	goto yyrule11

yystate29:
	c = l.next()
	switch {
	default:
		goto yyrule10
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate29
	}

	goto yystate30 // silence unused label error
yystate30:
	c = l.next()
yystart30:
	switch {
	default:
		goto yyabort
	case c == '"':
		goto yystate32
	case c == '\t' || c == ' ':
		goto yystate31
	}

yystate31:
	c = l.next()
	switch {
	default:
		goto yyrule12
	case c == '\t' || c == ' ':
		goto yystate31
	}

yystate32:
	c = l.next()
	goto yyrule13

	goto yystate33 // silence unused label error
yystate33:
	c = l.next()
yystart33:
	switch {
	default:
		goto yystate34 // c >= '\x00' && c <= '!' || c >= '#' && c <= '[' || c >= ']' && c <= 'ÿ'
	case c == '"':
		goto yystate35
	case c == '\\':
		goto yystate36
	}

yystate34:
	c = l.next()
	switch {
	default:
		goto yystate34 // c >= '\x00' && c <= '!' || c >= '#' && c <= '[' || c >= ']' && c <= 'ÿ'
	case c == '"':
		goto yystate35
	case c == '\\':
		goto yystate36
	}

yystate35:
	c = l.next()
	goto yyrule14

yystate36:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c >= '\x01' && c <= '\t' || c >= '\v' && c <= 'ÿ':
		goto yystate34
	}

yyrule1: // \0
	{
		return eof
	}
yyrule2: // #[^\r\n]*\n
	{
		l.mstart = l.i
		goto yystate0
	}
yyrule3: // [\r\n \t]+
	{
		l.mstart = l.i
		goto yystate0
	}
yyrule4: // {S}({M}|{D})*
	{
		s = lstateName
		l.offsets = append(l.offsets, l.i)
		l.mend = l.i
		goto yystate0
	}
yyrule5: // ([ \t]*)\{
	{
		s = lstateLabels
		goto yystate0
	}
yyrule6: // [ \t]+
	{
		s = lstateValue
		l.vstart = l.i
		goto yystate0
	}
yyrule7: // [ \t]+

	goto yystate0
yyrule8: // ,?\}
	{
		s = lstateValue
		l.mend = l.i
		goto yystate0
	}
yyrule9: // (,?[ \t]*)
	{
		s = lstateLName
		l.offsets = append(l.offsets, l.i)
		goto yystate0
	}
yyrule10: // {S}({L}|{D})*
	{
		l.offsets = append(l.offsets, l.i)
		goto yystate0
	}
yyrule11: // [ \t]*=
	{
		s = lstateLValue
		goto yystate0
	}
yyrule12: // [ \t]+

	goto yystate0
yyrule13: // \"
	{
		s = lstateLValueIn
		l.offsets = append(l.offsets, l.i)
		goto yystate0
	}
yyrule14: // (\\.|[^\\"]|\0)*\"
	{
		s = lstateLabels
		if !utf8.Valid(l.b[l.offsets[len(l.offsets)-1] : l.i-1]) {
			l.err = fmt.Errorf("Invalid UTF-8 label value.")
			return -1
		}
		l.offsets = append(l.offsets, l.i-1)
		goto yystate0
	}
yyrule15: // [ \t]+
	{
		l.vstart = l.i
		goto yystate0
	}
yyrule16: // (NaN)
	{
		l.val = math.Float64frombits(value.NormalNaN)
		s = lstateTimestamp
		goto yystate0
	}
yyrule17: // [^\n \t\r]+
	{
		// We don't parse strictly correct floats as the conversion
		// repeats the effort anyway.
		l.val, l.err = strconv.ParseFloat(yoloString(l.b[l.vstart:l.i]), 64)
		if l.err != nil {
			return -1
		}
		s = lstateTimestamp
		goto yystate0
	}
yyrule18: // [ \t]+
	{
		l.tstart = l.i
		goto yystate0
	}
yyrule19: // {D}+
	{
		ts, err := strconv.ParseInt(yoloString(l.b[l.tstart:l.i]), 10, 64)
		if err != nil {
			l.err = err
			return -1
		}
		l.ts = &ts
		goto yystate0
	}
yyrule20: // [\r\n]+
	{
		l.nextMstart = l.i
		return 1
	}
yyrule21: // \0
	{
		return 1

	}
	panic("unreachable")

	goto yyabort // silence unused label error

yyabort: // no lexem recognized
	l.err = fmt.Errorf("no token found")
	return -1
}