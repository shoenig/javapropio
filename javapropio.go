// Copyright (c) 2013 - Richard Boyer. All rights reserved.
// Use of this source code is governed by a BSD-style license that can
// be found in the LICENSE file.
package javapropio // import "gophers.dev/pkgs/javapropio"

import (
	"bufio"
	"io"

	"github.com/rogpeppe/go-charset/charset"
	_ "github.com/rogpeppe/go-charset/data"
)

const unHex = "0123456789ABCDEF"

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
	write(pw.buf, '=')

	escapePropValue(pw.buf, v)
	write(pw.buf, '\n')

	return pw.buf.Flush()
}

func escapePropKey(buf *bufio.Writer, s string)   { escapeProp(buf, s, true) }
func escapePropValue(buf *bufio.Writer, s string) { escapeProp(buf, s, false) }

func escapeProp(w *bufio.Writer, s string, isKey bool) {
	for x, r := range s {
		if r > 61 && r < 127 {
			if r == '\\' {
				write(w, '\\')
				write(w, '\\')
				continue
			}
			write(w, r)
			continue
		}

		switch {
		case r == ' ':
			if x == 0 || isKey {
				write(w, '\\')
			}
			write(w, ' ')
		case r == '\t':
			write(w, '\\')
			write(w, 't')
		case r == '\n':
			write(w, '\\')
			write(w, 'n')
		case r == '\r':
			write(w, '\\')
			write(w, 'r')
		case r == '\f':
			write(w, '\\')
			write(w, 'f')
		case r == '=', r == ':', r == '#', r == '!':
			write(w, '\\')
			write(w, r)
		case int32(r) < 0x20, int32(r) > 0x7e:
			write(w, '\\')
			write(w, 'u')

			write(w, rune(unHex[int32(r)>>12&0xF]))
			write(w, rune(unHex[int32(r)>>8&0xF]))
			write(w, rune(unHex[int32(r)>>4&0xF]))
			write(w, rune(unHex[int32(r)&0xF]))
		default:
			write(w, r)
		}
	}
}

func write(w *bufio.Writer, r rune) {
	_, _ = w.WriteRune(r)
}
