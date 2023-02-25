package sql

import (
	"context"
	"database/sql"
	"embed"
	"fmt"
	"testing"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/knadh/koanf"
	"github.com/mikeblum/teapotbot.dev/conf"
	"github.com/pressly/goose/v3"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

const (
	testConfFile = "../.env"
)

//go:embed internal/migrations/**.sql
var embedTestMigrations embed.FS

type SqlTestSuite struct {
	cfg *koanf.Koanf
	log *logrus.Entry
}

func setupSqlTestSuite(t *testing.T) (*SqlTestSuite, func(*testing.T, *SqlTestSuite) error, error) {
	var db *sql.DB
	var err error
	log := conf.NewLog(testConfFile)
	cfg, _ := conf.NewConf(conf.Provider(testConfFile))

	if db, err = sql.Open(driver, cfg.String(EnvDatabaseUrl)); err != nil {
		return nil, nil, err
	}
	if db == nil {
		return nil, nil, fmt.Errorf("db not configured")
	}

	defer db.Close()

	goose.SetBaseFS(embedMigrations)
	goose.SetTableName(migrationsTable)
	goose.SetLogger(log)
	goose.SetSequential(true)
	goose.SetVerbose(true)

	if err = goose.SetDialect(dialect); err != nil {
		return nil, nil, err
	}

	// step #1: reset any state issues
	goose.Reset(db, migrationsDir)

	// step #2: apply test harness migrations prefixed with 0000
	// step #3: rebase regular migrations on top of the test harness > 1000
	if err = Setup(cfg); err != nil {
		return nil, nil, err
	}

	suite := &SqlTestSuite{
		cfg: cfg,
		log: log,
	}
	return suite, teardownSqlSuite, nil
}

func teardownSqlSuite(t *testing.T, suite *SqlTestSuite) error {
	var err error
	var db *sql.DB
	if db, err = sql.Open(driver, suite.cfg.String(EnvDatabaseUrl)); err != nil {
		return err
	}
	if db == nil {
		return fmt.Errorf("db not configured")
	}

	defer db.Close()

	goose.SetTableName(migrationsTable)
	goose.SetLogger(suite.log)
	goose.SetSequential(true)
	goose.SetVerbose(true)

	if err := goose.SetDialect(dialect); err != nil {
		return err
	}

	// in LIFO order from SetupSqlTestSuite teardown the regular migrations first
	// lastly teardown the test harness
	if err := goose.Down(db, migrationsDir); err != nil {
		return err
	}
	return nil
}

func TestSql(t *testing.T) {
	suite, teardown, err := setupSqlTestSuite(t)
	assert.Nil(t, err)
	assert.NotNil(t, suite)
	assert.NotNil(t, teardown)
	defer func(t *testing.T) {
		err = nil
		// TODO: enable teardown via flag
		// err := teardown(t, suite)
		assert.Nil(t, err)
	}(t)
	t.Run("sql=created_updated_timestamps", suite.TimestampInsertUpdateTest)
}

func (s *SqlTestSuite) TimestampInsertUpdateTest(t *testing.T) {
	var conn *pgx.Conn
	var err error
	s.log.Info(s.cfg.String(EnvDatabaseUrl))
	if conn, err = pgx.Connect(context.Background(), s.cfg.String(EnvDatabaseUrl)); err != nil {
		assert.Nil(t, err)
	}
	defer conn.Close(context.Background())
	var out pgconn.CommandTag
	if out, err = conn.Exec(context.Background(),
		`INSERT INTO "public"."crawls" ("url", "status_code") 
		 VALUES ($1, $2);`,
		"example.com", 200); err != nil {
		assert.Nil(t, err)
	}
	assert.True(t, out.Insert())
	assert.Equal(t, int64(1), out.RowsAffected())
	// sleep one second to differentiate created vs updated
	time.Sleep(time.Second)
	if out, err = conn.Exec(context.Background(),
		`UPDATE "public"."crawls" 
			SET url = $2
		 WHERE url = $1;`,
		"example.com", "www.example.com"); err != nil {
		assert.Nil(t, err)
	}
	assert.True(t, out.Update())
	assert.Equal(t, int64(1), out.RowsAffected())
	var url string
	var created time.Time
	var updated time.Time
	if err = conn.QueryRow(context.Background(),
		`SELECT
				url,
				created,
				updated
			FROM
				crawls
			WHERE
				1 = 1
				AND url = $1;`,
		"www.example.com").Scan(&url, &created, &updated); err != nil {
		assert.Nil(t, err)
	}
	assert.NotEmpty(t, url)
	assert.NotEmpty(t, created)
	assert.NotEmpty(t, updated)
	assert.NotEqual(t, created, updated)
}
