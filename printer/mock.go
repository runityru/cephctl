package printer

import (
	"github.com/stretchr/testify/mock"
)

type Mock struct {
	mock.Mock
}

func NewMock() *Mock {
	return &Mock{}
}

func (m *Mock) Green(format string, a ...any) {
	m.Called(format, a)
}

func (m *Mock) HiRed(format string, a ...any) {
	m.Called(format, a)
}

func (m *Mock) Println(a ...any) {
	m.Called(a)
}

func (m *Mock) Red(format string, a ...any) {
	m.Called(format, a)
}

func (m *Mock) Yellow(format string, a ...any) {
	m.Called(format, a)
}
