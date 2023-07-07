package msuformat

import (
	"clear-arch/internal/entity"
	"clear-arch/internal/logger"
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/require"
)

type test struct {
	name string
	mock func()
	res  interface{}
	err  error
}

func TestGetBooks(t *testing.T) {
	mockCtl := gomock.NewController(t)
	defer mockCtl.Finish()

	repo := NewMockMsuFormatRepo(mockCtl)
	logger := logger.NewMockLogger(mockCtl)

	books := New(repo, logger)

	tests := []test{
		{
			name: "performe xml transformation",
			mock: func() {
				repo.EXPECT().GetBytes(context.Background()).Return([]byte("<library><authors><author><id>1</id><name>Alexander</name><surname>Pushkin</surname></author></authors><books><book><id>1</id><authorId>1</authorId><name>Captain's daughter</name><pageCount>352</pageCount></book></books></library>"), nil)
			},
			res: []entity.Book{
				{
					Name:         "Captain's daughter",
					PackageCount: 352,
					Author: entity.Author{
						Name:    "Alexander",
						Surname: "Pushkin",
					},
				},
			},
			err: nil,
		},
	}

	for _, tc := range tests {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			tc.mock()

			actual, err := books.GetBooks(context.Background())

			require.ErrorIs(t, tc.err, err)
			// require.Equal(t, tc.res, res)
			if diff := cmp.Diff(tc.res, actual); diff != "" {
				t.Fatal(diff, "Should be equal.")
			}
		})
	}

}
