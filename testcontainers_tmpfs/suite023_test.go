package testcontainers_tmpfs

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type Suite023 struct {
	SuiteBase
}

func (s *Suite023) TestFunc1()  {}
func (s *Suite023) TestFunc2()  {}
func (s *Suite023) TestFunc3()  {}
func (s *Suite023) TestFunc4()  {}
func (s *Suite023) TestFunc5()  {}
func (s *Suite023) TestFunc6()  {}
func (s *Suite023) TestFunc7()  {}
func (s *Suite023) TestFunc8()  {}
func (s *Suite023) TestFunc9()  {}
func (s *Suite023) TestFunc10() {}

func TestSuite023(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(Suite023))
}
