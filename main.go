package main

import "fmt"

// Genericを使わない場合、型ごとに関数を定義する必要がある
// m の値を加算 (int64)
func SumInts(m map[string]int64) int64 {
    var s int64
    for _, v := range m {
        s += v
    }
    return s
}

// m の値を加算	(float64)
func SumFloats(m map[string]float64) float64 {
    var s float64
    for _, v := range m {
        s += v
    }
    return s
}

// m の値を加算 int64/float64の両方をサポート
func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
    var s V
    for _, v := range m {
        s += v
    }
    return s
}

func main() {
    // int値のマップを初期化
    ints := map[string]int64{
        "first":  34,
        "second": 12,
    }

    // float値のマップを初期化
    floats := map[string]float64{
        "first":  35.98,
        "second": 26.99,
    }

    fmt.Printf("Non-Generic Sums: %v and %v\n",
        SumInts(ints),
        SumFloats(floats))

	fmt.Printf("Generic Sums: %v and %v\n",
		SumIntsOrFloats[string, int64](ints),
		SumIntsOrFloats[string, float64](floats))
}