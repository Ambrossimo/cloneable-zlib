package zlib

import (
	"encoding"
	"hash"
	"hash/adler32"
	"io"
)

func (z *Writer) Clone(w io.Writer) (*Writer, error) {
	if z == nil {
		return nil, nil
	}
	var digest hash.Hash32
	if z.digest != nil {
		bm := z.digest.(encoding.BinaryMarshaler)
		b, err := bm.MarshalBinary()
		if err != nil {
			return nil, err
		}
		digest = adler32.New()
		bu := digest.(encoding.BinaryUnmarshaler)
		err = bu.UnmarshalBinary(b)
		if err != nil {
			return nil, err
		}
	}
	var dict []byte
	if z.dict != nil {
		dict = make([]byte, len(z.dict))
		copy(dict, z.dict)
	}
	return &Writer{
		w:           w,
		level:       z.level,
		dict:        dict,
		compressor:  z.compressor.Clone(w),
		digest:      digest,
		err:         z.err,
		scratch:     z.scratch,
		wroteHeader: z.wroteHeader,
	}, nil
}
