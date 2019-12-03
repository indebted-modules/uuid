package uuid_test

import (
	"testing"

	googleuuid "github.com/google/uuid"
	"github.com/indebted-modules/uuid"
	"github.com/stretchr/testify/suite"
)

type UUIDSuite struct {
	suite.Suite
}

func TestUUIDSuite(t *testing.T) {
	suite.Run(t, new(UUIDSuite))
}

func (s *UUIDSuite) TestNewID() {
	id := uuid.NewID()
	a, err := googleuuid.Parse(id)
	s.Nil(err)
	s.True(a.Variant() == googleuuid.RFC4122)
}

func (s *UUIDSuite) TestValidateID() {
	s.False(uuid.ValidateID("r4nd0m"))
	s.False(uuid.ValidateID("00000000-0000-0000-0000-000000000000"))
	s.True(uuid.ValidateID("f8e0db06-129d-4435-97fb-60c64f7ea4fb"))
}
