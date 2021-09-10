// 변수
// http://golang.site/go/article/4-Go-%EB%B3%80%EC%88%98%EC%99%80-%EC%83%81%EC%88%98
// https://jamielim.tistory.com/entry/1%EC%A3%BC%EC%B0%A8-Go%EC%96%B8%EC%96%B4-%EA%B8%B0%EB%B3%B8-%EB%AC%B8%EB%B2%95-%EC%9E%90%EB%A3%8C%ED%98%95-%EC%97%B0%EC%82%B0%EC%9E%90-%EB%93%B1

package main

func main() {
	// 일반적인 선언
	var a int
	a = 10
	println(a)

	var b int = 5
	println(b)

	var i, j, k int = 1, 2, 3
	println(i, j, k)

	var f float32
	f = 11.0
	println(f)

	// 자료형 생략
	var num = 9
	println(num)

	// 짧은 선언
	n := 100
	s := "string"
	println(n, s)

	// var()를 통해 변수 여러개 선언 및 초기화
	var (
		aa, bb int = 100, 200
		cc, dd     = 30, "dd!"
	)
	println(aa, bb, cc, dd)
}
