package bxylog

import "io"

type multiWriter struct {
	writers []io.Writer
}

func (mw *multiWriter) ChangeFile(w io.Writer) {
	mw.writers[1] = w
}

func (mw *multiWriter) Write(p []byte) (n int, err error) {
	for _, w := range mw.writers {
		n, err = w.Write(p)
		if err != nil {
			return n, err
		}
	}
	return len(p), nil
}
