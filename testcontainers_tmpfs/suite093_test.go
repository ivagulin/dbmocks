package testcontainers_tmpfs

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type Suite093 struct {
	SuiteBase
}

func (s *Suite093) TestFunc1()  {}
func (s *Suite093) TestFunc2()  {}
func (s *Suite093) TestFunc3()  {}
func (s *Suite093) TestFunc4()  {}
func (s *Suite093) TestFunc5()  {}
func (s *Suite093) TestFunc6()  {}
func (s *Suite093) TestFunc7()  {}
func (s *Suite093) TestFunc8()  {}
func (s *Suite093) TestFunc9()  {}
func (s *Suite093) TestFunc10() {}

func TestSuite093(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(Suite093))
}
