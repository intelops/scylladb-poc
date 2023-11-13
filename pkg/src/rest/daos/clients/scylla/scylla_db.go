package scylla

import (
	"fmt"
	"log"

	"github.com/MrAzharuddin/scylladb-gin/config"
	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/v2"
)

type Manager struct {
	cfg config.Config
}

func NewManager() (*Manager, error) {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
		return nil, err
	}
	return &Manager{cfg: cfg}, nil
}

func (m *Manager) Connect() (gocqlx.Session, error) {
	cluster := gocql.NewCluster(m.cfg.ScyllaHosts)
	cluster.Keyspace = m.cfg.ScyllaKeyspace
	cluster.Consistency = gocql.Quorum
	cluster.PoolConfig.HostSelectionPolicy = gocql.TokenAwareHostPolicy(gocql.RoundRobinHostPolicy())
	cluster.Compressor = &gocql.SnappyCompressor{}
	cluster.RetryPolicy = &gocql.ExponentialBackoffRetryPolicy{NumRetries: 3}
	return gocqlx.WrapSession(cluster.CreateSession())
}

func (m *Manager) GetKeyspace() string {
	return m.cfg.ScyllaKeyspace
}

func (m *Manager) CreateKeyspace(session *gocqlx.Session, keyspace string) error {
	tmt := fmt.Sprintf(`CREATE KEYSPACE IF NOT EXISTS %s  WITH replication = {'class': 'SimpleStrategy', 'replication_factor': 1}`, keyspace)
	return session.ExecStmt(tmt)
}

func (m *Manager) CreateTable(session *gocqlx.Session, query string, table string) error {
	tmt := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s.%s (%s)`, m.cfg.ScyllaKeyspace, table, query)
	return session.Query(tmt, []string{}).ExecRelease()
}
