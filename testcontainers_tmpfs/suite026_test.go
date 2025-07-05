package testcontainers_tmpfs

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type Suite026 struct {
	SuiteBase
}

func (s *Suite026) TestFunc1()  {}
func (s *Suite026) TestFunc2()  {}
func (s *Suite026) TestFunc3()  {}
func (s *Suite026) TestFunc4()  {}
func (s *Suite026) TestFunc5()  {}
func (s *Suite026) TestFunc6()  {}
func (s *Suite026) TestFunc7()  {}
func (s *Suite026) TestFunc8()  {}
func (s *Suite026) TestFunc9()  {}
func (s *Suite026) TestFunc10() {}

func TestSuite026(t *testing.T) {
	t.Parallel()
	suite.Run(t, new(Suite026))
}
