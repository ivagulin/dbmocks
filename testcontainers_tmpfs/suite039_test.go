package testcontainers_tmpfs

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type Suite039 struct {
	SuiteBase
}

func (s *Suite039) TestFunc1()  {}
func (s *Suite039) TestFunc2()  {}
func (s *Suite039) TestFunc3()  {}
func (s *Suite039) TestFunc4()  {}
func (s *Suite039) TestFunc5()  {}
func (s *Suite039) TestFunc6()  {}
func (s *Suite039) TestFunc7()  {}
func (s *Suite039) TestFunc8()  {}
func (s *Suite039) TestFunc9()  {}
func (s *Suite039) TestFunc10() {}

func TestSuite039(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(Suite039))
}
