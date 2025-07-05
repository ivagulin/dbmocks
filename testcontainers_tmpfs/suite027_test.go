package testcontainers_tmpfs

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type Suite027 struct {
	SuiteBase
}

func (s *Suite027) TestFunc1()  {}
func (s *Suite027) TestFunc2()  {}
func (s *Suite027) TestFunc3()  {}
func (s *Suite027) TestFunc4()  {}
func (s *Suite027) TestFunc5()  {}
func (s *Suite027) TestFunc6()  {}
func (s *Suite027) TestFunc7()  {}
func (s *Suite027) TestFunc8()  {}
func (s *Suite027) TestFunc9()  {}
func (s *Suite027) TestFunc10() {}

func TestSuite027(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(Suite027))
}
