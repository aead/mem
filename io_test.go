// Copyright (c) 2022 Andreas Auernhammer. All rights reserved.
// Use of this source code is governed by a license that can be
// found in the LICENSE file.

package mem

import (
	"fmt"
	"io"
	"testing"
)

func TestProgress_Done(t *testing.T) {
	for i, test := range progressDoneTests {
		if done := test.Progress.Done(); done != test.Done {
			t.Fatalf("Test %d: got %v - want %v", i, done, test.Done)
		}
	}
}

var progressDoneTests = []struct {
	Progress Progress
	Done     bool
}{
	{Progress: Progress{Err: nil}, Done: false},
	{Progress: Progress{Err: io.ErrUnexpectedEOF}, Done: false},
	{Progress: Progress{Err: io.EOF}, Done: true},
	{Progress: Progress{Err: fmt.Errorf("wrapped %w", io.EOF)}, Done: true},
}
