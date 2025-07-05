package testcontainers_tmpfs

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type Suite072 struct {
	SuiteBase
}

func (s *Suite072) TestFunc1()  {}
func (s *Suite072) TestFunc2()  {}
func (s *Suite072) TestFunc3()  {}
func (s *Suite072) TestFunc4()  {}
func (s *Suite072) TestFunc5()  {}
func (s *Suite072) TestFunc6()  {}
func (s *Suite072) TestFunc7()  {}
func (s *Suite072) TestFunc8()  {}
func (s *Suite072) TestFunc9()  {}
func (s *Suite072) TestFunc10() {}

func TestSuite072(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(Suite072))
}
