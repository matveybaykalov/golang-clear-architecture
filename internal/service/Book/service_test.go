package book

import (
	"clear-arch/internal/entity"
	logger "clear-arch/internal/logger"
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

type test struct {
	name     string
	mock     func()
	expected interface{}
	err      error
}

// Чтобы сгенерировать тесты необходимо использовать следующую команду:
// mockgen -source=service.go -destination=service_mock.go -package=book
func TestGetBooks(t *testing.T) {
	mockCtl := gomock.NewController(t)
	defer mockCtl.Finish()

	repo := NewMockBookRepo(mockCtl)
	logger := logger.NewMockLogger(mockCtl)

	books := New(repo, logger)

	errorExample := errors.New("Some error")

	tests := []test{
		{
			name: "check methos call with nil output",
			mock: func() {
				repo.EXPECT().GetAllBooks(context.Background()).Return(nil, nil)
			},
			expected: []entity.Book(nil),
			err:      nil,
		},
		{
			name: "check method call with some outout",
			mock: func() {
				repo.EXPECT().GetAllBooks(context.Background()).Return([]entity.Book{
					{
						Name:         "Book name",
						PackageCount: 10,
						Author: entity.Author{
							Name:    "Author name",
							Surname: "Author surname",
						},
					},
				}, nil)
			},
			expected: []entity.Book{
				{
					Name:         "Book name",
					PackageCount: 10,
					Author: entity.Author{
						Name:    "Author name",
						Surname: "Author surname",
					},
				},
			},
			err: nil,
		},
		{
			name: "some error in output",
			mock: func() {
				repo.EXPECT().GetAllBooks(context.Background()).Return(nil, errorExample)
			},
			expected: []entity.Book(nil),
			err:      errorExample,
		},
	}

	for _, tc := range tests {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			tc.mock()

			actual, err := books.GetBooks(context.Background())

			require.ErrorIs(t, tc.err, err)
			require.Equal(t, tc.expected, actual)
			//if diff := cmp.Diff(tc.expect, actual); diff != "" {
			//	t.Fatal(diff, "Should be equal.")
			//}
		})
	}

}
