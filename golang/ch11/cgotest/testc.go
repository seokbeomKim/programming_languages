package main

// #include <stdio.h>
// #include <stdlib.h>
//
// static void myprint(char* s) {
//   printf("%s\n", s);
// }
//
// extern int print_in_c(int a);
import "C"

// 반드시 import "C"는 위의 코드와 반드시 반드시 붙여야 한다.
// 그렇지 않을 경우 에러 발생

import "unsafe"

func main() {
	cs := C.CString("Hello from stdio")
	C.myprint(cs)
	C.free(unsafe.Pointer(cs))
	C.print_in_c(1)
}
