package usecase_test

import (
	"clear-arch/internal/entity"
	"clear-arch/internal/usecase"
	"context"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

type test struct {
	name string
	mock func()
	res  interface{}
	err  error
}

func TestGetXMLBooks(t *testing.T) {
	mockCtl := gomock.NewController(t)
	defer mockCtl.Finish()

	repo := NewMockBooksParserRepo(mockCtl)
	books := usecase.NewXMLParser(repo)

	tests := []test{
		{
			name: "performe xml transformation",
			mock: func() {
				repo.EXPECT().GetDataFromFile(context.Background()).Return([]byte("<library><authors><author><id>1</id><name>Alexander</name><surname>Pushkin</surname></author></authors><books><book><id>1</id><authorId>1</authorId><name>Captain's daughter</name><pageCount>352</pageCount></book></books></library>"), nil)
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

			res, err := books.GetBooks(context.Background())

			require.Equal(t, res, tc.res)
			require.ErrorIs(t, err, tc.err)
		})
	}

	//	if diff := cmp.Diff(expected, actual); diff != "" {
	//	t.Fatal(diff, "Should be equal.")
	//}
}

func TestGetYMLBooks(t *testing.T) {
	mockCtl := gomock.NewController(t)
	defer mockCtl.Finish()

	repo := NewMockBooksParserRepo(mockCtl)
	books := usecase.NewYMLParser(repo)

	tests := []test{
		{
			name: "performe yml transformation",
			mock: func() {
				repo.EXPECT().GetDataFromFile(context.Background()).Return([]byte("books:\n  - id: 1\n    name: Woe from wit\n    authorName: Alexander\n    authorSurname: Griboyedov\n    pageCount: 256"), nil)
			},
			res: []entity.Book{
				{
					Name:         "Woe from wit",
					PackageCount: 256,
					Author: entity.Author{
						Name:    "Alexander",
						Surname: "Griboyedov",
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

			res, err := books.GetBooks(context.Background())

			require.ErrorIs(t, err, tc.err)
			require.Equal(t, res, tc.res)
		})
	}
}
