package bmsturepo

import (
	"context"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"runtime"
)

func (br *BmstuRepo) GetBytes(context.Context) ([]byte, error) {
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)

	br.log.Debug(fmt.Sprintf("project root : %s", basepath))

	path := filepath.Join(basepath, "../../../test/yml/example1.yml")

	d, err := ioutil.ReadFile(path)

	br.log.Info("Open file test/yml/example1.yml")

	if err != nil {
		return nil, err
	}

	return d, nil
}
