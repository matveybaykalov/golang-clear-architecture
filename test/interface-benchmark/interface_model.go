package interfacebenchmark

type A struct {
	data int
}

func (a *A) inc() {
	a.data += 1
}

type B struct {
	data int
}

func (b *B) inc() {
	b.data += 1
}

type C interface {
	inc()
}

func StaticTest(arr1 []A, arr2 []B) {
	length := len(arr1)

	for i := 0; i < length; i++ {
		arr1[i].inc()
	}

	for i := 0; i < length; i++ {
		arr2[i].inc()
	}
}

func DynamicTest(arr []C) {
	length := len(arr)

	for i := 0; i < length; i++ {
		arr[i].inc()
	}
}
