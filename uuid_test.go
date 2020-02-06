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
	id := googleuuid.MustParse(uuid.NewID())
	s.True(id.Variant() == googleuuid.RFC4122)
}

func (s *UUIDSuite) TestNewIDFor() {
	s.Equal(uuid.NewIDFor("ns1", "A"), uuid.NewIDFor("ns1", "A"))
	s.Equal(uuid.NewIDFor("ns1", "B"), uuid.NewIDFor("ns1", "B"))
	s.Equal(uuid.NewIDFor("ns1", "C"), uuid.NewIDFor("ns1", "C"))

	s.NotEqual(uuid.NewIDFor("ns1", "A"), uuid.NewIDFor("ns2", "A"))
	s.NotEqual(uuid.NewIDFor("ns1", "B"), uuid.NewIDFor("ns2", "B"))
	s.NotEqual(uuid.NewIDFor("ns1", "C"), uuid.NewIDFor("ns2", "C"))

	s.NotEqual(uuid.NewIDFor("ns1", "A"), uuid.NewIDFor("ns1", "B"))
	s.NotEqual(uuid.NewIDFor("ns1", "B"), uuid.NewIDFor("ns1", "C"))
	s.NotEqual(uuid.NewIDFor("ns1", "C"), uuid.NewIDFor("ns1", "A"))

	id := googleuuid.MustParse(uuid.NewIDFor("ns1", "A"))
	s.True(id.Variant() == googleuuid.RFC4122)
}

func (s *UUIDSuite) TestValidateID() {
	s.False(uuid.ValidateID("r4nd0m"))
	s.False(uuid.ValidateID("00000000-0000-0000-0000-000000000000"))
	s.True(uuid.ValidateID("f8e0db06-129d-4435-97fb-60c64f7ea4fb"))
}
