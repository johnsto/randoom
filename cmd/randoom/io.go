package main

import (
	"bytes"
	"io"

	"github.com/pkg/errors"
)

type PrefixWriter struct {
	io.Writer
	Separator []byte
	Prefix    []byte
}

func NewPrefixWriter(w io.Writer, separator, prefix []byte) *PrefixWriter {
	return &PrefixWriter{
		Writer:    w,
		Prefix:    prefix,
		Separator: separator,
	}
}

func (w *PrefixWriter) Write(p []byte) (n int, err error) {
	lines := bytes.SplitAfter(p, []byte{'\n'})
	for _, line := range lines {
		if len(line) == 0 {
			break
		}
		n1, err1 := w.Writer.Write(w.Prefix)
		n = n + n1
		if err1 != nil {
			return n, errors.Wrapf(err1, "couldn't write prefix")
		}
		n2, err2 := w.Writer.Write(line)
		n = n + n2
		if err2 != nil {
			return n, errors.Wrapf(err2, "couldn't write line")
		}
	}
	return n, nil
}
