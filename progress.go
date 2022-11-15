// Copyright (c) 2022 Andreas Auernhammer. All rights reserved.
// Use of this source code is governed by a license that can be
// found in the LICENSE file.

package mem

import (
	"errors"
	"io"
	"time"
)

// Progress represents the progress of an I/O operation,
// like reading data from a file or network connection.
type Progress struct {
	// N is the number of bytes since the last progress
	// update.
	N Size

	// Total is the number of bytes since the start
	// of the operation.
	Total Size

	// Err is any error that occurred during the operation.
	// Once the operation completes, Err is io.EOF.
	Err error
}

// Done reports whether the operation has been completed.
func (p *Progress) Done() bool { return errors.Is(p.Err, io.EOF) }

// NewProgressReader returns a new ProgressReader that wraps r and
// calls update periodically with the current progress while reading.
func NewProgressReader(r io.Reader, d time.Duration, update func(Progress)) *ProgressReader {
	return &ProgressReader{
		R:           r,
		Update:      update,
		UpdateEvery: d,
	}
}

// ProgressReader wraps an io.Reader and calls Update
// with the current status when reading makes progress.
type ProgressReader struct {
	R io.Reader // The underlying io.Reader

	// Update, if non-nil, is called with the current
	// progress whenever a read from R completes and
	// either the UpdateEvery period has ellapsed
	// since the last update or UpdateAfter bytes have
	// been read.
	//
	// As a special case, Update is called after every
	// read from R if UpdateEvery and UpdateAfter
	// are both <= 0.
	//
	// The passed progress contains the number of bytes
	// read since the last Update call, the number of
	// bytes read in total so far and any error that
	// has occurred while reading from R.
	// Once reading from R returns an non-nil error,
	// including io.EOF, Update is called immediately
	// one more time and then never again.
	//
	// Update is called by the goroutine reading from
	// R. A long-running or blocking Update function
	// defers or blocks reads, and therefore, impacts
	// read performance.
	// In such cases, sending the progress to another
	// concurrently running goroutine via a channel
	// may be viable solution.
	Update func(Progress)

	// UpdateEvery is the duration that has to ellapse
	// between two Update calls.
	//
	// If UpdateEvery <= 0, Update may be called after
	// every read.
	UpdateEvery time.Duration

	// UpdateAfter is the number of bytes that have to
	// be read from R before Update is called again.
	//
	// If UpdateAfter <= 0, Update may be called after
	// every read.
	UpdateAfter Size

	n, total   Size
	lastUpdate time.Time
	err        error
}

func (r *ProgressReader) Read(p []byte) (int, error) {
	if r.err != nil {
		return 0, r.err
	}

	n, err := r.R.Read(p)
	r.n += Size(n)
	r.total += Size(n)
	if err != nil {
		r.err = err
	}
	if r.Update != nil {
		switch {
		case (r.UpdateEvery <= 0 && r.UpdateAfter <= 0) || err != nil:
			r.Update(r.Progress())
			r.n = 0
		case r.UpdateAfter > 0 && r.n >= r.UpdateAfter:
			r.Update(r.Progress())
			r.n = 0
		case r.UpdateEvery > 0 && r.lastUpdate.IsZero():
			r.Update(r.Progress())
			r.n = 0
			r.lastUpdate = time.Now()
		case r.UpdateEvery > 0:
			now := time.Now()
			if diff := now.Sub(r.lastUpdate); diff >= r.UpdateEvery {
				r.Update(r.Progress())
				r.n = 0
				r.lastUpdate = now
			}
		}
	}
	return n, err
}

// Progress returns the current progress.
//
// It contains the number of bytes read since the
// last invocation of Update by Read, the total
// number of bytes read so far and any error that
// has occurred while reading from R.
func (r *ProgressReader) Progress() Progress {
	return Progress{
		N:     r.n,
		Total: r.total,
		Err:   r.err,
	}
}
