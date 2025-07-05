package testcontainers_tmpfs

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type Suite025 struct {
	SuiteBase
}

func (s *Suite025) TestFunc1()  {}
func (s *Suite025) TestFunc2()  {}
func (s *Suite025) TestFunc3()  {}
func (s *Suite025) TestFunc4()  {}
func (s *Suite025) TestFunc5()  {}
func (s *Suite025) TestFunc6()  {}
func (s *Suite025) TestFunc7()  {}
func (s *Suite025) TestFunc8()  {}
func (s *Suite025) TestFunc9()  {}
func (s *Suite025) TestFunc10() {}

func TestSuite025(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(Suite025))
}
