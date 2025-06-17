package pgtestpkg

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type Suite076 struct {
	SuiteBase
}

func (s *Suite076) TestFunc1()  {}
func (s *Suite076) TestFunc2()  {}
func (s *Suite076) TestFunc3()  {}
func (s *Suite076) TestFunc4()  {}
func (s *Suite076) TestFunc5()  {}
func (s *Suite076) TestFunc6()  {}
func (s *Suite076) TestFunc7()  {}
func (s *Suite076) TestFunc8()  {}
func (s *Suite076) TestFunc9()  {}
func (s *Suite076) TestFunc10() {}

func TestSuite076(t *testing.T) {
	suite.Run(t, new(Suite076))
}
