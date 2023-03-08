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
	var stats = mariadb_stats.GetUsersStats()
	cfg := collectorConfig{}
	if err := env.Parse(&cfg); err != nil {
		log.Println("ERROR: COL01 - %+v\n", err)
	}
	for _, stat := range stats {
		ch <- prometheus.MustNewConstMetric(collector.TotalConnections, prometheus.CounterValue, stat.TOTAL_CONNECTIONS, stat.USER, cfg.NODE)
		ch <- prometheus.MustNewConstMetric(collector.ConcurrentConnections, prometheus.CounterValue, stat.CONCURRENT_CONNECTIONS, stat.USER, cfg.NODE)
		ch <- prometheus.MustNewConstMetric(collector.ConnectedTime, prometheus.CounterValue, stat.CONNECTED_TIME, stat.USER, cfg.NODE)
		ch <- prometheus.MustNewConstMetric(collector.BusyTime, prometheus.CounterValue, stat.BUSY_TIME, stat.USER, cfg.NODE)
		ch <- prometheus.MustNewConstMetric(collector.CpuTime, prometheus.CounterValue, stat.CPU_TIME, stat.USER, cfg.NODE)
		ch <- prometheus.MustNewConstMetric(collector.BytesReceived, prometheus.CounterValue, stat.BYTES_RECEIVED, stat.USER, cfg.NODE)
		ch <- prometheus.MustNewConstMetric(collector.BytesSend, prometheus.CounterValue, stat.BYTES_SENT, stat.USER, cfg.NODE)
		ch <- prometheus.MustNewConstMetric(collector.BinlogBytesWritten, prometheus.CounterValue, stat.BINLOG_BYTES_WRITTEN, stat.USER, cfg.NODE)
		ch <- prometheus.MustNewConstMetric(collector.RowsRead, prometheus.CounterValue, stat.ROWS_READ, stat.USER, cfg.NODE)
		ch <- prometheus.MustNewConstMetric(collector.RowsSent, prometheus.CounterValue, stat.ROWS_SENT, stat.USER, cfg.NODE)
		ch <- prometheus.MustNewConstMetric(collector.RowsDeleted, prometheus.CounterValue, stat.ROWS_DELETED, stat.USER, cfg.NODE)
		ch <- prometheus.MustNewConstMetric(collector.RowsInserted, prometheus.CounterValue, stat.ROWS_INSERTED, stat.USER, cfg.NODE)
		ch <- prometheus.MustNewConstMetric(collector.RowsUpdated, prometheus.CounterValue, stat.ROWS_UPDATED, stat.USER, cfg.NODE)
		ch <- prometheus.MustNewConstMetric(collector.SelectCommands, prometheus.CounterValue, stat.SELECT_COMMANDS, stat.USER, cfg.NODE)
		ch <- prometheus.MustNewConstMetric(collector.UpdateCommands, prometheus.CounterValue, stat.UPDATE_COMMANDS, stat.USER, cfg.NODE)
		ch <- prometheus.MustNewConstMetric(collector.OtherCommands, prometheus.CounterValue, stat.OTHER_COMMANDS, stat.USER, cfg.NODE)
		ch <- prometheus.MustNewConstMetric(collector.CommitTransactions, prometheus.CounterValue, stat.COMMIT_TRANSACTIONS, stat.USER, cfg.NODE)
		ch <- prometheus.MustNewConstMetric(collector.RollbackTransactions, prometheus.CounterValue, stat.ROLLBACK_TRANSACTIONS, stat.USER, cfg.NODE)
		ch <- prometheus.MustNewConstMetric(collector.DeniedConnections, prometheus.CounterValue, stat.DENIED_CONNECTIONS, stat.USER, cfg.NODE)
		ch <- prometheus.MustNewConstMetric(collector.Lost_Connections, prometheus.CounterValue, stat.LOST_CONNECTIONS, stat.USER, cfg.NODE)
		ch <- prometheus.MustNewConstMetric(collector.AccessDenied, prometheus.CounterValue, stat.ACCESS_DENIED, stat.USER, cfg.NODE)
		ch <- prometheus.MustNewConstMetric(collector.EmptyQueries, prometheus.CounterValue, stat.EMPTY_QUERIES, stat.USER, cfg.NODE)
		ch <- prometheus.MustNewConstMetric(collector.TotalSslConnections, prometheus.CounterValue, stat.TOTAL_SSL_CONNECTIONS, stat.USER, cfg.NODE)
		ch <- prometheus.MustNewConstMetric(collector.MaxStatementExceeded, prometheus.CounterValue, stat.MAX_STATEMENT_TIME_EXCEEDED, stat.USER, cfg.NODE)
	}
}
