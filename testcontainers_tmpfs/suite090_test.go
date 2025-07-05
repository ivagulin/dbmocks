package testcontainers_tmpfs

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type Suite090 struct {
	SuiteBase
}

func (s *Suite090) TestFunc1()  {}
func (s *Suite090) TestFunc2()  {}
func (s *Suite090) TestFunc3()  {}
func (s *Suite090) TestFunc4()  {}
func (s *Suite090) TestFunc5()  {}
func (s *Suite090) TestFunc6()  {}
func (s *Suite090) TestFunc7()  {}
func (s *Suite090) TestFunc8()  {}
func (s *Suite090) TestFunc9()  {}
func (s *Suite090) TestFunc10() {}

func TestSuite090(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(Suite090))
}
