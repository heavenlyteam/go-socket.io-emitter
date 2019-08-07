package socketioemitter

import (
	"github.com/ugorji/go/codec"
	"io"
)

func EncodeMessage(m interface{}) (b []byte, err error) {
	var (
		w  io.Writer
		mh codec.MsgpackHandle
	)

	enc := codec.NewEncoder(w, &mh)
	enc = codec.NewEncoderBytes(&b, &mh)
	err = enc.Encode(m)
	return
}
