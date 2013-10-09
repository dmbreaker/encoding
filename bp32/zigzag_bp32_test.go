/*
 * Copyright (c) 2013 Zhen, LLC. http://zhen.io. All rights reserved.
 * Use of this source code is governed by the Apache 2.0 license.
 *
 */

package bp32

import (
	"testing"
	"log"
	"github.com/reducedb/encoding"
)

func init() {
	log.Printf("zigzag_bp32_test/init: generating %d int32s\n", size)
	data = encoding.GenerateClustered(size, size*2)
	log.Printf("zigzag_bp32_test/init: generated %d integers for test", size)
}

func TestZigZagBP32(t *testing.T) {
	sizes := []int{128, 128*10, 128*100, 128*1000, 128*10000}
	encoding.TestCodec(NewZigZagBP32(), data, sizes, t)
}

func BenchmarkZigZagBP32Compress(b *testing.B) {
	encoding.BenchmarkCompress(NewZigZagBP32(), data, b)
}

func BenchmarkZigZagBP32Uncompress(b *testing.B) {
	encoding.BenchmarkUncompress(NewZigZagBP32(), data, b)
}