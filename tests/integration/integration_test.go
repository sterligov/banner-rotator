package integration

import (
	"testing"

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
}

func TestIntegration(t *testing.T) {
	suite.Run(t, new(Suite))
}

func (s *Suite) SetupSuite() {
	//dbName := os.Getenv("DB_NAME")
	//dbUser := os.Getenv("DB_USER")
	//dbPass := os.Getenv("DB_PASSWORD")
	//dbHost := os.Getenv("DB_HOST")
	//dbPort := os.Getenv("DB_PORT")
	//grpcConn := os.Getenv("GRPC_SERVER")

	conn, err := grpc.Dial("localhost:8880", grpc.WithInsecure())
	s.Require().NoError(err)
	s.grpcConn = conn

	//addr := fmt.Sprintf(
	//	"%s:%s@tcp(%s:%s)/%s?parseTime=true",
	//	dbUser,
	//	dbPass,
	//	dbHost,
	//	dbPort,
	//	dbName,
	//)
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
}

func (s *Suite) TearDownSuite() {
	_, err := s.db.Exec("DELETE FROM statistic")
	s.Require().NoError(err)
}
