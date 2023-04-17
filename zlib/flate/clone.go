package flate

import "io"

func (e *deflateFast) clone() *deflateFast {
	if e == nil {
		return nil
	}
	var prev []byte
	if e.prev != nil {
		prev = make([]byte, len(e.prev))
		copy(prev, e.prev)
	}
	return &deflateFast{
		table: e.table,
		prev:  prev,
		cur:   e.cur,
	}
}

func (e *huffmanEncoder) clone() *huffmanEncoder {
	if e == nil {
		return nil
	}
	var codes []hcode
	if e.codes != nil {
		codes = make([]hcode, len(e.codes))
		copy(codes, e.codes)
	}
	var freqcache []literalNode
	if e.freqcache != nil {
		freqcache = make([]literalNode, len(e.freqcache))
		copy(freqcache, e.freqcache)
	}
	var lns byLiteral
	if e.lns != nil {
		lns = make(byLiteral, len(e.lns))
		copy(lns, e.lns)
	}
	var lfs byFreq
	if e.lfs != nil {
		lfs = make(byFreq, len(e.lfs))
		copy(lfs, e.lfs)
	}
	return &huffmanEncoder{
		codes:     codes,
		freqcache: freqcache,
		bitCount:  e.bitCount,
		lns:       lns,
		lfs:       lfs,
	}
}

func (e *huffmanBitWriter) clone(w io.Writer) *huffmanBitWriter {
	if e == nil {
		return nil
	}
	var literalFreq []int32
	if e.literalFreq != nil {
		literalFreq = make([]int32, len(e.literalFreq))
		copy(literalFreq, e.literalFreq)
	}
	var offsetFreq []int32
	if e.offsetFreq != nil {
		offsetFreq = make([]int32, len(e.offsetFreq))
		copy(offsetFreq, e.offsetFreq)
	}
	var codegen []uint8
	if e.codegen != nil {
		codegen = make([]uint8, len(e.codegen))
		copy(codegen, e.codegen)
	}
	return &huffmanBitWriter{
		writer:          w,
		bits:            e.bits,
		nbits:           e.nbits,
		bytes:           e.bytes,
		codegenFreq:     e.codegenFreq,
		nbytes:          e.nbytes,
		literalFreq:     literalFreq,
		offsetFreq:      offsetFreq,
		codegen:         codegen,
		literalEncoding: e.literalEncoding.clone(),
		offsetEncoding:  e.offsetEncoding.clone(),
		codegenEncoding: e.codegenEncoding.clone(),
		err:             e.err,
	}
}

func (d *compressor) clone(w io.Writer) compressor {
	var window []byte
	if d.window != nil {
		window = make([]byte, len(d.window))
		copy(window, d.window)
	}
	var tokens []token
	if d.tokens != nil {
		tokens = make([]token, len(d.tokens))
		copy(tokens, d.tokens)
	}
	return compressor{
		compressionLevel: d.compressionLevel,
		w:                d.w.clone(w),
		bulkHasher:       d.bulkHasher,
		fill:             d.fill,
		step:             d.step,
		sync:             d.sync,
		bestSpeed:        d.bestSpeed.clone(),
		chainHead:        d.chainHead,
		hashHead:         d.hashHead,
		hashPrev:         d.hashPrev,
		hashOffset:       d.hashOffset,
		index:            d.index,
		window:           window,
		windowEnd:        d.windowEnd,
		blockStart:       d.blockStart,
		byteAvailable:    d.byteAvailable,
		tokens:           tokens,
		length:           d.length,
		offset:           d.offset,
		maxInsertIndex:   d.maxInsertIndex,
		err:              d.err,
		hashMatch:        d.hashMatch,
	}
}

func (w *Writer) Clone(wr io.Writer) *Writer {
	if w == nil {
		return nil
	}
	var dict []byte
	if w.dict != nil {
		dict = make([]byte, len(w.dict))
		copy(dict, w.dict)
	}
	return &Writer{
		d:    w.d.clone(wr),
		dict: dict,
	}
}
