package filereader_test

import (
	"clear-arch/internal/usecase/filereader"
	"context"
	"fmt"
	"testing"
)

func TestGetDataFromFile(t *testing.T) {
	a := filereader.New("../../../test/xml/example1.xml")

	d, err := a.GetDataFromFile(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println(string(d))
}
