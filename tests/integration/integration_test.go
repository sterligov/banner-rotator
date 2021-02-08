package integration

import (
	"fmt"
	"testing"

	"github.com/nats-io/nats.go"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-testfixtures/testfixtures/v3"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc"
)

type Suite struct {
	suite.Suite

	db       *sqlx.DB
	grpcConn *grpc.ClientConn
	natsConn *nats.Conn
}

func TestIntegration(t *testing.T) {
	suite.Run(t, new(Suite))
}

func (s *Suite) SetupSuite() {
	conn, err := grpc.Dial("localhost:8880", grpc.WithInsecure())
	s.Require().NoError(err)
	s.grpcConn = conn

	addr := "rotator_user:rotator_pass@tcp(localhost:3315)/rotator?parseTime=true"
	db, err := sqlx.Connect("mysql", addr)
	s.Require().NoError(err)

	fixtures, err := testfixtures.New(
		testfixtures.DangerousSkipTestDatabaseCheck(),
		testfixtures.Database(db.DB),
		testfixtures.Dialect("mysql"),
		testfixtures.Directory("fixtures"),
	)
	s.Require().NoError(err)

	err = fixtures.Load()
	s.Require().NoError(err)
	s.db = db

	s.natsConn, err = nats.Connect("nats://localhost:4222")
	s.Require().NoError(err)
}

func (s *Suite) TearDownSuite() {
	tables := []string{"statistic", "banner", "slot", "social_group", "banner_slot"}
	for _, t := range tables {
		_, err := s.db.Exec(fmt.Sprintf("DELETE FROM %s", t))
		s.Require().NoError(err)
	}

	s.Require().NoError(s.db.Close())
	s.Require().NoError(s.grpcConn.Close())
	s.natsConn.Close()
}
