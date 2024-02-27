//go:build tinygo

package main

//go:export test_int32
func test_int32(i, j int32) int32 {
	return i + j
}

//go:export test_uint32
func test_uint32(i, j uint32) uint32 {
	return i + j
}

//go:export test_int64
func test_int64(i, j int64) int64 {
	return i + j
}

//go:export test_uint64
func test_uint64(i, j uint64) uint64 {
	return i + j
}

//go:export test_float32
func test_float32(i, j float32) float32 {
	return i + j
}

//go:export test_float64
func test_float64(i, j float64) float64 {
	return i + j
}

func main() {}
