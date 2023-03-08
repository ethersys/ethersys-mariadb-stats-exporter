package mariadb_stats

import (
	"log"
	"github.com/caarlos0/env/v7"
	_ "github.com/go-sql-driver/mysql"
    "github.com/jmoiron/sqlx"
)


type sqlConfig struct {
	SQL_HOST			string		`env:"SQL_HOST" envDefault:"localhost"`
	SQL_PORT			string     	`env:"SQL_PORT" envDefault:"3306"`
	SQL_USER			string		`env:"SQL_USER" envDefault:"root"`
	SQL_PASSWORD		string		`env:"SQL_PASSWORD" envDefault:""`
}

type statsRow struct {
    USER                        string      `db:"USER"`
    TOTAL_CONNECTIONS           float64     `db:"TOTAL_CONNECTIONS"`
    CONCURRENT_CONNECTIONS      float64     `db:"CONCURRENT_CONNECTIONS"`
    CONNECTED_TIME              float64     `db:"CONNECTED_TIME"`
    BUSY_TIME                   float64     `db:"BUSY_TIME"`
    CPU_TIME                    float64     `db:"CPU_TIME"`
    BYTES_RECEIVED              float64     `db:"BYTES_RECEIVED"`
    BYTES_SENT                  float64     `db:"BYTES_SENT"`
    BINLOG_BYTES_WRITTEN        float64     `db:"BINLOG_BYTES_WRITTEN"`
    ROWS_READ                   float64     `db:"ROWS_READ"`
    ROWS_SENT                   float64     `db:"ROWS_SENT"`
    ROWS_DELETED                float64     `db:"ROWS_DELETED"`
    ROWS_INSERTED               float64     `db:"ROWS_INSERTED"`
    ROWS_UPDATED                float64     `db:"ROWS_UPDATED"`
    SELECT_COMMANDS             float64     `db:"SELECT_COMMANDS"`
    UPDATE_COMMANDS             float64     `db:"UPDATE_COMMANDS"`
    OTHER_COMMANDS              float64     `db:"OTHER_COMMANDS"`
    COMMIT_TRANSACTIONS         float64     `db:"COMMIT_TRANSACTIONS"`
    ROLLBACK_TRANSACTIONS       float64     `db:"ROLLBACK_TRANSACTIONS"`
    DENIED_CONNECTIONS          float64     `db:"DENIED_CONNECTIONS"`
    LOST_CONNECTIONS            float64     `db:"LOST_CONNECTIONS"`
    ACCESS_DENIED               float64     `db:"ACCESS_DENIED"`
    EMPTY_QUERIES               float64     `db:"EMPTY_QUERIES"`
    TOTAL_SSL_CONNECTIONS       float64     `db:"TOTAL_SSL_CONNECTIONS"`
    MAX_STATEMENT_TIME_EXCEEDED float64     `db:"MAX_STATEMENT_TIME_EXCEEDED"`
}

func GetUsersStats() []statsRow {
	stats := []statsRow{}
	cfg := sqlConfig{}
	if err := env.Parse(&cfg); err != nil {
		log.Println("ERROR: GUS01 - %+v\n", err)
	}
    db, err := sqlx.Open("mysql", cfg.SQL_USER + ":" + cfg.SQL_PASSWORD + "@tcp(" + cfg.SQL_HOST + ":" + cfg.SQL_PORT + ")/information_schema")
    if err != nil {
        log.Println("ERROR: GUS02 - ", err.Error())
    }
    defer db.Close()
    err = db.Select(&stats, "SELECT * FROM USER_STATISTICS;")
    if err !=nil {
        log.Println("ERROR: GUS03 - ", err.Error())
    }
    return stats
}