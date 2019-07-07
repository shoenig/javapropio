// Copyright (c) 2013 - Richard Boyer. All rights reserved.
// Use of this source code is governed by a BSD-style license that can
// be found in the LICENSE file.
package javapropio

import (
	"bytes"
	"testing"
)

func Test_All(t *testing.T) {
	table := []struct {
		key    string
		val    string
		expect string
	}{
		{"rrr.ccc.percent", "19",
			"rrr.ccc.percent=19\n"},
		{"break.ab.d", "Iñtërnâtiônàlizætiøn",
			"break.ab.d=I\\u00F1t\\u00EBrn\\u00E2ti\\u00F4n\\u00E0liz\\u00E6ti\\u00F8n\n"},
	}

	for _, row := range table {
		var buf bytes.Buffer
		pw, err := NewWriter(&buf)
		if err != nil {
			t.Error(err)
			continue
		}
		err = pw.WriteProp(row.key, row.val)
		if err != nil {
			t.Error(err)
			continue
		}
		err = pw.Close()
		if err != nil {
			t.Error(err)
			continue
		}

		got := string(buf.Bytes())
		if got != row.expect {
			t.Errorf("Expected '%s' but got '%s'", row.expect, got)
			continue
		}
	}
}
