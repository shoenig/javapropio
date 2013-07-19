// Copyright (c) 2013 - Richard Boyer. All rights reserved.
// Use of this source code is governed by a BSD-style license that can
// be found in the LICENSE file.
package javapropio

import (
	"bufio"
	"code.google.com/p/go-charset/charset"
	_ "code.google.com/p/go-charset/data"
	"io"
)

const unhex = "0123456789ABCDEF"

type Writer struct {
	w   io.WriteCloser
	buf *bufio.Writer
}

func NewWriter(w io.Writer) (*Writer, error) {
	cw, err := charset.NewWriter("iso-8859-1", w)
	if err != nil {
		return nil, err
	}
	return &Writer{cw, bufio.NewWriter(cw)}, nil
}

func (pw *Writer) Close() error {
	return pw.w.Close()
}

func (pw *Writer) WriteProp(k, v string) error {
	escapePropKey(pw.buf, k)
	pw.buf.WriteRune('=')
	escapePropKey(pw.buf, v)
	pw.buf.WriteRune('\n')

	return pw.buf.Flush()
}

func escapePropKey(buf *bufio.Writer, s string)   { escapeProp(buf, s, true) }
func escapePropValue(buf *bufio.Writer, s string) { escapeProp(buf, s, false) }

func escapeProp(buf *bufio.Writer, s string, isKey bool) {
	for x, r := range s {
		if r > 61 && r < 127 {
			if r == '\\' {
				buf.WriteRune('\\')
				buf.WriteRune('\\')
				continue
			}
			buf.WriteRune(r)
			continue
		}

		switch {
		case r == ' ':
			if x == 0 || isKey {
				buf.WriteRune('\\')
			}
			buf.WriteRune(' ')
		case r == '\t':
			buf.WriteRune('\\')
			buf.WriteRune('t')
		case r == '\n':
			buf.WriteRune('\\')
			buf.WriteRune('n')
		case r == '\r':
			buf.WriteRune('\\')
			buf.WriteRune('r')
		case r == '\f':
			buf.WriteRune('\\')
			buf.WriteRune('f')
		case r == '=', r == ':', r == '#', r == '!':
			buf.WriteRune('\\')
			buf.WriteRune(r)
		case int32(r) < 0x20, int32(r) > 0x7e:
			buf.WriteRune('\\')
			buf.WriteRune('u')

			buf.WriteRune(rune(unhex[int32(r)>>12&0xF]))
			buf.WriteRune(rune(unhex[int32(r)>>8&0xF]))
			buf.WriteRune(rune(unhex[int32(r)>>4&0xF]))
			buf.WriteRune(rune(unhex[int32(r)&0xF]))
		default:
			buf.WriteRune(r)
		}
	}
}
