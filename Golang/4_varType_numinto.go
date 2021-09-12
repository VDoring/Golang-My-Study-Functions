// https://jamielim.tistory.com/entry/1%EC%A3%BC%EC%B0%A8-Go%EC%96%B8%EC%96%B4-%EA%B8%B0%EB%B3%B8-%EB%AC%B8%EB%B2%95-%EC%9E%90%EB%A3%8C%ED%98%95-%EC%97%B0%EC%82%B0%EC%9E%90-%EB%93%B1

package main

import (
	"fmt"
	"math"
	"unsafe"
)

func main() {
	// 정수, 실수
	var i1 int = 3
	var f1 float32 = 0.2
	var f2 float64 = 9.9
	fmt.Println(i1, f1, f2)

	// 복소수
	var n1 complex64 = 1 + 2i         // (실수+허수 구조)
	var n2 complex64 = 1.123 + 2.345i // 고정 소수점
	var n3 complex64 = 1.e+3i         // 부동 소수점
	var n4 complex64 = .123i          // 실수부 생략
	fmt.Println(n1, n2, n3, n4)

	fmt.Println("n1 real:", real(n1)) // real() = 복소수에서 실수부 가져옴
	fmt.Println("n1 real:", imag(n1)) // imag() = 복소수에서 허수부 가져옴
	fmt.Println("n2 real:", real(n2))
	fmt.Println("n2 real:", imag(n2))
	fmt.Println("n3 real:", real(n3))
	fmt.Println("n3 real:", imag(n3))
	fmt.Println("n4 real:", real(n4))
	fmt.Println("n4 real:", imag(n4))

	var n5 complex64 = complex(1, 2) // 복소수에서 + 대신 complex()를 사용할 수 있다.
	fmt.Println("n5", n5)

	// 바이트
	var b1 byte = 10   // 10진수
	var b2 byte = 0x10 // 16진수
	var b3 byte = 'a'  // 문자
	fmt.Println(b1, b2, b3)

	// 룬
	// 유니코드 문자 코드 저장시 사용
	var r1 rune = '한'
	var r2 rune = '\ud55c'
	var r3 rune = '\U0000d55c'
	fmt.Println(r1, r2, r3)

	// 자료형이 지원하는 최대값 확인
	var chk1 uint8 = math.MaxUint8
	var chk2 uint16 = math.MaxUint16
	var chk3 uint32 = math.MaxUint32
	var chk4 uint64 = math.MaxUint64
	fmt.Println(chk1, chk2, chk3, chk4)

	// 변수 크기 구하기
	fmt.Println(unsafe.Sizeof(n1))
	fmt.Println(unsafe.Sizeof(n2))
	fmt.Println(unsafe.Sizeof(n3))
	fmt.Println(unsafe.Sizeof(n4))

}
