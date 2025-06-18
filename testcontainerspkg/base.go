package pgtestpkg

import (
	"github.com/ivagulin/dbmocks/testinfra"
	"github.com/stretchr/testify/suite"
)

type SuiteBase struct {
	suite.Suite
	gDB testinfra.DB
}

func (s *SuiteBase) SetupSuite() {
	s.gDB = testinfra.MustStartDB()
}

func (s *SuiteBase) TearDownSuite() {
	s.gDB.Close()
}

func (s *SuiteBase) TearDownTest() {
	s.gDB.MustCleanup()
}
