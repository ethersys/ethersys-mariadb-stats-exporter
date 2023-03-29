package collector

import (
	"log"
	"mariadb_stats_exporter/mariadb_stats"
	"github.com/caarlos0/env/v7"
	"github.com/prometheus/client_golang/prometheus"
)

type collectorConfig struct {
	NODE				string		`env:"NODE" envDefault:"localhost"`
}

type RessourcesCollector struct {
	MemoryUsed				*prometheus.Desc
	TotalConnections		*prometheus.Desc
	ConcurrentConnections	*prometheus.Desc
	ConnectedTime			*prometheus.Desc
	BusyTime				*prometheus.Desc
	CpuTime					*prometheus.Desc
	BytesReceived			*prometheus.Desc
	BytesSend				*prometheus.Desc
	BinlogBytesWritten		*prometheus.Desc
	RowsRead				*prometheus.Desc
	RowsSent				*prometheus.Desc
	RowsDeleted				*prometheus.Desc
	RowsInserted			*prometheus.Desc
	RowsUpdated				*prometheus.Desc
	SelectCommands			*prometheus.Desc
	UpdateCommands			*prometheus.Desc
	OtherCommands			*prometheus.Desc
	CommitTransactions		*prometheus.Desc
	RollbackTransactions	*prometheus.Desc
	DeniedConnections		*prometheus.Desc
	Lost_Connections		*prometheus.Desc
	AccessDenied			*prometheus.Desc
	EmptyQueries			*prometheus.Desc
	TotalSslConnections		*prometheus.Desc
	MaxStatementExceeded	*prometheus.Desc
}

func NewRessourcesCollector() *RessourcesCollector {
	return &RessourcesCollector{
		MemoryUsed: prometheus.NewDesc("mariadb_stats_memory_used",
			"Total Memory used by MariaDB.",
			[]string{"node"}, nil,
		),
		TotalConnections: prometheus.NewDesc("mariadb_stats_total_connections",
			"The number of connections created for this user.",
			[]string{"user","node"}, nil,
		),
		ConcurrentConnections: prometheus.NewDesc("mariadb_stats_concurrent_connections",
			"The number of concurrent connections for this user.",
			[]string{"user","node"}, nil,
		),
		ConnectedTime: prometheus.NewDesc("mariadb_stats_connected_time",
			"The cumulative number of seconds elapsed while there were connections from this user.",
			[]string{"user","node"}, nil,
		),
		BusyTime: prometheus.NewDesc("mariadb_stats_busy_time",
			"The cumulative number of seconds there was activity on connections from this user.",
			[]string{"user","node"}, nil,
		),
		CpuTime: prometheus.NewDesc("mariadb_stats_cpu_time",
			"The cumulative CPU time elapsed while servicing this user's connections.",
			[]string{"user","node"}, nil,
		),
		BytesReceived: prometheus.NewDesc("mariadb_stats_bytes_received",
			"The number of bytes received from this user's connections.",
			[]string{"user","node"}, nil,
		),
		BytesSend: prometheus.NewDesc("mariadb_stats_bytes_send",
			"The number of bytes sent to this user's connections.",
			[]string{"user","node"}, nil,
		),
		BinlogBytesWritten: prometheus.NewDesc("mariadb_stats_binlog_bytes_written",
			"The number of bytes written to the binary log from this user's connections.",
			[]string{"user","node"}, nil,
		),
		RowsRead: prometheus.NewDesc("mariadb_stats_rows_read",
			"The number of rows read by this user's connections.",
			[]string{"user","node"}, nil,
		),
		RowsSent: prometheus.NewDesc("mariadb_stats_rows_sent",
			"The number of rows sent by this user's connections.",
			[]string{"user","node"}, nil,
		),
		RowsDeleted: prometheus.NewDesc("mariadb_stats_rows_deleted",
			"The number of rows deleted by this user's connections.",
			[]string{"user","node"}, nil,
		),
		RowsInserted: prometheus.NewDesc("mariadb_stats_rows_inserted",
			"The number of rows inserted by this user's connections.",
			[]string{"user","node"}, nil,
		),
		RowsUpdated: prometheus.NewDesc("mariadb_stats_rows_updated",
			"The number of rows updated by this user's connections.",
			[]string{"user","node"}, nil,
		),
		SelectCommands: prometheus.NewDesc("mariadb_stats_select_commands",
			"The number of SELECT commands executed from this user's connections.",
			[]string{"user","node"}, nil,
		),
		UpdateCommands: prometheus.NewDesc("mariadb_stats_update_commands",
			"The number of UPDATE commands executed from this user's connections.",
			[]string{"user","node"}, nil,
		),
		OtherCommands: prometheus.NewDesc("mariadb_stats_other_commands",
			"The number of other commands executed from this user's connections.",
			[]string{"user","node"}, nil,
		),
		CommitTransactions: prometheus.NewDesc("mariadb_stats_commit_transactions",
			"The number of COMMIT commands issued by this user's connections.",
			[]string{"user","node"}, nil,
		),
		RollbackTransactions: prometheus.NewDesc("mariadb_stats_rollback_transactions",
			"The number of ROLLBACK commands issued by this user's connections.",
			[]string{"user","node"}, nil,
		),
		DeniedConnections: prometheus.NewDesc("mariadb_stats_denied_connections",
			"The number of connections denied to this user.",
			[]string{"user","node"}, nil,
		),
		Lost_Connections: prometheus.NewDesc("mariadb_stats_lost_connections",
			"The number of this user's connections that were terminated uncleanly.",
			[]string{"user","node"}, nil,
		),
		AccessDenied: prometheus.NewDesc("mariadb_stats_access_denied",
			"The number of times this user's connections issued commands that were denied.",
			[]string{"user","node"}, nil,
		),
		EmptyQueries: prometheus.NewDesc("mariadb_stats_empty_queries",
			"The number of times this user's connections sent empty queries to the server.",
			[]string{"user","node"}, nil,
		),
		TotalSslConnections: prometheus.NewDesc("mariadb_stats_total_ssl_connections",
			"The number of TLS connections created for this user. (>= MariaDB 10.1.1)",
			[]string{"user","node"}, nil,
		),
		MaxStatementExceeded: prometheus.NewDesc("mariadb_stats_max_statement_exceeded",
			"The number of times a statement was aborted, because it was executed longer than its MAX_STATEMENT_TIME threshold. (>= MariaDB 10.1.1)",
			[]string{"user","node"}, nil,
		),
	}
}

func (collector *RessourcesCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- collector.MemoryUsed
	ch <- collector.TotalConnections
	ch <- collector.ConcurrentConnections
	ch <- collector.ConnectedTime
	ch <- collector.BusyTime
	ch <- collector.CpuTime
	ch <- collector.BytesReceived
	ch <- collector.BytesSend
	ch <- collector.BinlogBytesWritten
	ch <- collector.RowsRead
	ch <- collector.RowsSent
	ch <- collector.RowsDeleted
	ch <- collector.RowsInserted
	ch <- collector.RowsUpdated
	ch <- collector.SelectCommands
	ch <- collector.UpdateCommands
	ch <- collector.OtherCommands
	ch <- collector.CommitTransactions
	ch <- collector.RollbackTransactions
	ch <- collector.DeniedConnections
	ch <- collector.Lost_Connections
	ch <- collector.AccessDenied
	ch <- collector.EmptyQueries
	ch <- collector.TotalSslConnections
	ch <- collector.MaxStatementExceeded
}

func (collector *RessourcesCollector) Collect(ch chan<- prometheus.Metric) {
	var users_stats = mariadb_stats.GetUsersStats()
	var memory_used = mariadb_stats.GetMemoryStats()
	cfg := collectorConfig{}
	if err := env.Parse(&cfg); err != nil {
		log.Println("ERROR: COL01 - %+v\n", err)
	}
	ch <- prometheus.MustNewConstMetric(collector.MemoryUsed, prometheus.CounterValue, memory_used, cfg.NODE)
	for _, user_stats := range users_stats {
		ch <- prometheus.MustNewConstMetric(collector.TotalConnections, prometheus.CounterValue, user_stats.TOTAL_CONNECTIONS, user_stats.USER, cfg.NODE)
		ch <- prometheus.MustNewConstMetric(collector.ConcurrentConnections, prometheus.CounterValue, user_stats.CONCURRENT_CONNECTIONS, user_stats.USER, cfg.NODE)
		ch <- prometheus.MustNewConstMetric(collector.ConnectedTime, prometheus.CounterValue, user_stats.CONNECTED_TIME, user_stats.USER, cfg.NODE)
		ch <- prometheus.MustNewConstMetric(collector.BusyTime, prometheus.CounterValue, user_stats.BUSY_TIME, user_stats.USER, cfg.NODE)
		ch <- prometheus.MustNewConstMetric(collector.CpuTime, prometheus.CounterValue, user_stats.CPU_TIME, user_stats.USER, cfg.NODE)
		ch <- prometheus.MustNewConstMetric(collector.BytesReceived, prometheus.CounterValue, user_stats.BYTES_RECEIVED, user_stats.USER, cfg.NODE)
		ch <- prometheus.MustNewConstMetric(collector.BytesSend, prometheus.CounterValue, user_stats.BYTES_SENT, user_stats.USER, cfg.NODE)
		ch <- prometheus.MustNewConstMetric(collector.BinlogBytesWritten, prometheus.CounterValue, user_stats.BINLOG_BYTES_WRITTEN, user_stats.USER, cfg.NODE)
		ch <- prometheus.MustNewConstMetric(collector.RowsRead, prometheus.CounterValue, user_stats.ROWS_READ, user_stats.USER, cfg.NODE)
		ch <- prometheus.MustNewConstMetric(collector.RowsSent, prometheus.CounterValue, user_stats.ROWS_SENT, user_stats.USER, cfg.NODE)
		ch <- prometheus.MustNewConstMetric(collector.RowsDeleted, prometheus.CounterValue, user_stats.ROWS_DELETED, user_stats.USER, cfg.NODE)
		ch <- prometheus.MustNewConstMetric(collector.RowsInserted, prometheus.CounterValue, user_stats.ROWS_INSERTED, user_stats.USER, cfg.NODE)
		ch <- prometheus.MustNewConstMetric(collector.RowsUpdated, prometheus.CounterValue, user_stats.ROWS_UPDATED, user_stats.USER, cfg.NODE)
		ch <- prometheus.MustNewConstMetric(collector.SelectCommands, prometheus.CounterValue, user_stats.SELECT_COMMANDS, user_stats.USER, cfg.NODE)
		ch <- prometheus.MustNewConstMetric(collector.UpdateCommands, prometheus.CounterValue, user_stats.UPDATE_COMMANDS, user_stats.USER, cfg.NODE)
		ch <- prometheus.MustNewConstMetric(collector.OtherCommands, prometheus.CounterValue, user_stats.OTHER_COMMANDS, user_stats.USER, cfg.NODE)
		ch <- prometheus.MustNewConstMetric(collector.CommitTransactions, prometheus.CounterValue, user_stats.COMMIT_TRANSACTIONS, user_stats.USER, cfg.NODE)
		ch <- prometheus.MustNewConstMetric(collector.RollbackTransactions, prometheus.CounterValue, user_stats.ROLLBACK_TRANSACTIONS, user_stats.USER, cfg.NODE)
		ch <- prometheus.MustNewConstMetric(collector.DeniedConnections, prometheus.CounterValue, user_stats.DENIED_CONNECTIONS, user_stats.USER, cfg.NODE)
		ch <- prometheus.MustNewConstMetric(collector.Lost_Connections, prometheus.CounterValue, user_stats.LOST_CONNECTIONS, user_stats.USER, cfg.NODE)
		ch <- prometheus.MustNewConstMetric(collector.AccessDenied, prometheus.CounterValue, user_stats.ACCESS_DENIED, user_stats.USER, cfg.NODE)
		ch <- prometheus.MustNewConstMetric(collector.EmptyQueries, prometheus.CounterValue, user_stats.EMPTY_QUERIES, user_stats.USER, cfg.NODE)
		ch <- prometheus.MustNewConstMetric(collector.TotalSslConnections, prometheus.CounterValue, user_stats.TOTAL_SSL_CONNECTIONS, user_stats.USER, cfg.NODE)
		ch <- prometheus.MustNewConstMetric(collector.MaxStatementExceeded, prometheus.CounterValue, user_stats.MAX_STATEMENT_TIME_EXCEEDED, user_stats.USER, cfg.NODE)
	}

}
