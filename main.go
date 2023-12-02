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

// １. m の値を加算	(float64)
func SumFloats(m map[string]float64) float64 {
	var s float64
	for _, v := range m {
		s += v
	}
	return s
}

// 2. m の値を加算 int64/float64の両方をサポート
func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

// 型制約を定義
type Number interface {
	int64 | float64
}

// 3. 定義した型制約を用いる m の値を加算 int64/float64の両方をサポート
func SumNumbers[K comparable, V Number](m map[K]V) V {
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

	// 基本的には、型引数を指定する必要がある
	fmt.Printf("Generic Sums: %v and %v\n",
		SumIntsOrFloats[string, int64](ints),
		SumIntsOrFloats[string, float64](floats))

	// Go コンパイラーが使用する型を推測できる場合は、呼び出しコードで型引数を省略できる
	fmt.Printf("Generic Sums, type parameters inferred: %v and %v\n",
		SumIntsOrFloats(ints),
		SumIntsOrFloats(floats))

	// 型制約を定義した場合の呼び出し
	fmt.Printf("Generic Sums with Constraint: %v and %v\n",
		SumNumbers(ints),
		SumNumbers(floats))

}
