// https://jamielim.tistory.com/entry/1%EC%A3%BC%EC%B0%A8-Go%EC%96%B8%EC%96%B4-%EA%B8%B0%EB%B3%B8-%EB%AC%B8%EB%B2%95-%EC%9E%90%EB%A3%8C%ED%98%95-%EC%97%B0%EC%82%B0%EC%9E%90-%EB%93%B1

package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var s1 = "한글이ㄷ"
	fmt.Println(len(s1))                    // len()은 한글은 한 글자당 길이를 3으로 인식한다
	fmt.Println(utf8.RuneCountInString(s1)) // RuneCountInString()을 사용하여 utf로 저장한 문자열의 실제 길이를 구할 수 있다

	var a = "A, "
	var b = "HahahA"
	fmt.Println(a + b)         // 문자열 합치기
	fmt.Println(b == "HahahA") // 문자열 내용 일치 여부. 답은 true/false이다.
	fmt.Println(a == b)        // 문자열 내용 일치 여부. 답은 true/false이다.

	// 각 문자 출력하기 (한글은 이렇게 못함)
	// printf()와 %c 키워드 사용시 문자 출력
	var eng = "Hello"
	fmt.Printf("%c\n", eng[0]) // H
	fmt.Printf("%c\n", eng[1]) // E
	fmt.Printf("%c\n", eng[2]) // L
	fmt.Printf("%c\n", eng[3]) // L
	fmt.Printf("%c\n", eng[4]) // O

	// Println()을 사용하면 아스키코드가 출력된다
	fmt.Println(eng[0]) // 72
	fmt.Println(eng[1]) // 101
	fmt.Println(eng[2]) // 108
	fmt.Println(eng[3]) // 108
	fmt.Println(eng[4]) // 111

	fmt.Printf("\a")
}
