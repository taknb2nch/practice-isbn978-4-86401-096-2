package main

import (
	"fmt"
	"runtime"
	"math/rand"
	"time"
)

type example struct {
	Str string
	Value int
}

func finalizer(e *example) {
	fmt.Printf("Finalizing %s: %v\n", e.Str, e.Value)
}

func NewExample() *example {
	e := new(example)
	e.Value = rand.Int()

	fmt.Printf("%d\n", e.Value)

	runtime.SetFinalizer(e, finalizer)
	return e
}

func main() {
	e := NewExample()
	//NewExample()
	//e = NewExample()
	e.Str = "a structure"
	//e = nil

	time.Sleep(5 * time.Microsecond)

	e = NewExample()
	//NewExample()

	fmt.Println(e.Str)

	//e = NewExample()

	runtime.GC()

	// 本文より
	// 2つ目の呼び出しを単に値をnilに設定するように変更したら
	// ファイナライザは呼ばれないかもしれません。
	// 変数がnilに設定されると、
	// コンパイラは意味のない値の保存だと気づき、
	// その代入を単純に取り除き、
	// eは古い値を指したままとなります。

}

