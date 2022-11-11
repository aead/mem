// Copyright (c) 2022 Andreas Auernhammer. All rights reserved.
// Use of this source code is governed by a license that can be
// found in the LICENSE file.

package mem

import "io"

// LimitReader returns a io.LimitedReader that reads from r
// but stops with io.EOF after n bytes.
func LimitReader(r io.Reader, n Size) *io.LimitedReader {
	return &io.LimitedReader{
		R: r,
		N: int64(n),
	}
}
