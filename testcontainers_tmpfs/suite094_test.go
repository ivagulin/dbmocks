package testcontainers_tmpfs

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type Suite094 struct {
	SuiteBase
}

func (s *Suite094) TestFunc1()  {}
func (s *Suite094) TestFunc2()  {}
func (s *Suite094) TestFunc3()  {}
func (s *Suite094) TestFunc4()  {}
func (s *Suite094) TestFunc5()  {}
func (s *Suite094) TestFunc6()  {}
func (s *Suite094) TestFunc7()  {}
func (s *Suite094) TestFunc8()  {}
func (s *Suite094) TestFunc9()  {}
func (s *Suite094) TestFunc10() {}

func TestSuite094(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(Suite094))
}
