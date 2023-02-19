package clients

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/hiejulia/api-online-book-store/utils"
	"strings"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Constants
var (
	SQLFile    string
	SQLType    string
	SQLUser    string
	SQLPass    string
	SQLName    string
	SQLHost    string
	SQLPort    string
	configOnce sync.Once
	sqlOnce    sync.Once
)

// SQL ...
type SQL struct {
	cfg  Config
	conn *gorm.DB
}

// SQLConfig ...
func SQLConfig() Config {
	configOnce.Do(func() {
		SQLFile = utils.GetEnvStr("DB_FILE")
		SQLType = utils.GetEnvStr("DB_TYPE")
		SQLUser = utils.GetEnvStr("DB_USER")
		SQLPass = utils.GetEnvStr("DB_PASS")
		SQLName = utils.GetEnvStr("DB_NAME")
		SQLHost = utils.GetEnvStr("DB_HOST")
		SQLPort = utils.GetEnvStr("DB_PORT")
	})

	return Config{
		File:     SQLFile,
		Host:     SQLHost,
		Name:     SQLName,
		Password: SQLPass,
		Port:     SQLPort,
		Type:     SQLType,
		User:     SQLUser,
	}
}

// NewSQL ...
func NewSQL(cfg Config) (store *SQL) {
	sqlOnce.Do(func() {
		store = &SQL{cfg: cfg}
	})
	return
}

// AutoMigrate ...
func (store *SQL) AutoMigrate(v ...interface{}) (err error) {
	err = store.conn.AutoMigrate(v...)
	return
}

// BatchInsert ...
func (store *SQL) BatchInsert(vs []interface{}, size int) (err error) {
	i := 0
	j := 0
	for j < size {
		if (j+1)%10 == 0 || (j+1) == size {
			if err = store.conn.Create(vs[i : j+1]).Error; err != nil {
				return
			}

			i = j
		}
	}
	return
}

// Close ...
func (store *SQL) Close() (err error) {
	return
}

// Commit ...
func (store *SQL) Commit() (err error) {
	err = store.conn.Commit().Error
	return
}

// Config ...
func (store *SQL) Config() (cfg Config) {
	cfg = store.cfg
	return
}

// Create ...
func (store *SQL) Create(v interface{}) (err error) {
	err = store.conn.Create(v).Error
	return
}

// CreateInBatches ...
func (store *SQL) CreateInBatches(value interface{}, batchSize int) (err error) {
	err = store.conn.CreateInBatches(value, batchSize).Error
	return
}

// Count ...
func (store *SQL) Count(v interface{}, where map[string]interface{}, count *uint64) (err error) {
	tx := store.conn.Model(v)
	for k, v := range where {
		tx.Where(k, v)
	}

	var n int64
	if err = tx.Count(&n).Error; err != nil {
		return
	}

	*count = uint64(n)
	return
}

// CountSingle ...
func (store *SQL) CountSingleWhere(where interface{}, count *uint64) (err error) {
	var n int64
	if err = store.conn.Model(where).Where(where).Count(&n).Error; err != nil {
		return
	}

	*count = uint64(n)
	return
}

// DB ...
func (store *SQL) DB() (db *gorm.DB) {
	db = store.conn
	return
}

// Delete ...
func (store *SQL) Delete(v interface{}) (err error) {
	err = store.conn.Where(v).Delete(v).Error
	return
}

// DeleteIn ...
func (store *SQL) DeleteIn(v interface{}, field string, args []string) (err error) {
	query := fmt.Sprintf("%s IN ?", field)
	err = store.conn.Where(query, args).Delete(v).Error
	return
}

// Find ...
func (store *SQL) Find(v, vs interface{}) (err error) {
	err = store.conn.Where(v).Find(vs).Error
	return
}

// FindIn ...
func (store *SQL) FindIn(field string, args []string, v, vs interface{}) (err error) {
	err = store.conn.Where(v).Where(fmt.Sprintf("%s IN ?", field), args).Find(vs).Error
	return
}

// FindNotIn ...
func (store *SQL) FindNotIn(field string, args []string, v, vs interface{}) (err error) {
	err = store.conn.Where(v).Where(fmt.Sprintf("%s NOT IN ?", field), args).Find(vs).Error
	return
}

// FindWhere ...
func (store *SQL) FindWhere(v interface{}, where map[string]interface{}, vs interface{}) (err error) {
	tx := store.conn.Model(v)
	for k, v := range where {
		tx.Where(k, v)
	}
	err = tx.Find(vs).Error
	return
}

func (store *SQL) FindWhereOr(v interface{}, where map[string]interface{}, or map[string]interface{}, vs interface{}) (err error) {
	tx := store.conn.Model(v)
	for k, v := range where {
		tx.Where(k, v)
	}
	for k, v := range or {
		tx.Or(k, v)
	}
	err = tx.Find(vs).Error
	return
}

// FindWhereOrderBy ...
func (store *SQL) FindWhereOrderBy(v interface{}, where map[string]interface{}, vs interface{}, orderBy string) (err error) {
	tx := store.conn.Model(v)
	for k, v := range where {
		tx.Where(k, v)
	}
	tx.Order(orderBy)
	err = tx.Find(vs).Error
	return
}

// FindWhereLimitOrderBy ...
func (store *SQL) FindWhereLimitOrderBy(v interface{}, where map[string]interface{}, vs interface{}, limit int, orderBy string) (err error) {
	tx := store.conn.Model(v)
	for k, v := range where {
		tx.Where(k, v)
	}
	tx.Limit(limit)
	tx.Order(orderBy)
	err = tx.Find(vs).Error
	return
}

// First ...
func (store *SQL) First(v interface{}) (err error) {
	err = store.conn.Where(v).First(v).Error
	return
}

// FirstSelect ...
func (store *SQL) FirstSelect(v interface{}, cols ...string) (err error) {
	err = store.conn.Select(cols).Where(v).First(v).Error
	return
}

// Last ...
func (store *SQL) Last(v interface{}) (err error) {
	err = store.conn.Where(v).Last(v).Error
	return
}

func (store *SQL) Type() string {
	return store.cfg.Type
}

// Open ...
func (store *SQL) Open() (err error) {
	var dsn string
	if dsn, err = store.cfg.DSN(); err != nil {
		return
	}

	var dialector gorm.Dialector
	switch store.cfg.Type {
	case "mysql":
		dialector = mysql.Open(dsn)
	case "sqlite":
		dialector = sqlite.Open(dsn)
	default:
		err = fmt.Errorf("unknown sql clients dialect %s, options are: mysql, postgres, sqlite", store.cfg.Type)
		return
	}

	store.conn, err = gorm.Open(dialector, &gorm.Config{
		SkipDefaultTransaction:                   true,
		FullSaveAssociations:                     false,
		DisableForeignKeyConstraintWhenMigrating: store.cfg.Type == "sqlite",
		// PrepareStmt:            true,
	})
	if err != nil {
		return
	}

	var db *sql.DB
	if db, err = store.conn.DB(); err != nil {
		return
	}

	if store.cfg.Type == "mysql" || store.cfg.Type == "postgres" {
		fmt.Println("Configuring for MySQL or Postgres")
		db.SetMaxIdleConns(2)
		db.SetMaxOpenConns(500)
		// db.SetConnMaxIdleTime(1 * time.Minute)
		// db.SetConnMaxLifetime(1 * time.Second)
	} else if store.cfg.Type == "sqlite" {
		fmt.Println("Configuring for SQLite")
		db.SetMaxIdleConns(1)
		db.SetMaxOpenConns(1)
	}
	return
}

// Page ...
func (store *SQL) Page(vs, v interface{}, limit, offset int) (err error) {
	err = store.conn.Where(v).Limit(limit).Offset(offset).Find(vs).Error
	return
}

// Rollback ...
func (store *SQL) Rollback() (err error) {
	err = store.conn.Rollback().Error
	return
}

// Session ...
func (store *SQL) Session() *SQL {
	// return &SQL{
	// 	cfg:  store.cfg,
	// 	conn: store.conn,
	// }
	return store
}

// TX ...
func (store *SQL) TX() *SQL {
	return &SQL{
		cfg:  store.cfg,
		conn: store.conn.Begin(),
	}
}

// Update ...
func (store *SQL) Update(v interface{}) (err error) {
	err = store.conn.Updates(v).Error
	return
}

// UpdateAll ...
func (store *SQL) UpdateAll(col string, ids []string, v interface{}) (err error) {
	if len(ids) == 0 {
		return
	}

	tmp := make([]string, 0)
	for i, id := range ids {
		tmp = append(tmp, fmt.Sprintf("'%s'", id))

		if (i+1)%10 == 0 || (i+1) == len(ids) {
			where := fmt.Sprintf("%s IN (%s)", col, strings.Join(tmp, ","))
			if err = store.conn.Model(v).Where(where).Updates(v).Error; err != nil {
				return
			}

			tmp = tmp[:0]
		}
	}
	return
}

// UpdateOnly ...
func (store *SQL) UpdateOnly(v interface{}, updates map[string]interface{}, where map[string]interface{}) (err error) {
	tx := store.conn.Model(v)
	for k, v := range where {
		tx.Where(k, v)
	}
	err = tx.Updates(updates).Error
	return
}

// UpdateWhere ...
func (store *SQL) UpdateWhere(v interface{}, where map[string]interface{}) (err error) {
	tx := store.conn.Model(v)
	for k, v := range where {
		tx.Where(k, v)
	}
	err = tx.Updates(v).Error
	return
}

// Where ...
func (store *SQL) Where(v interface{}) *SQL {
	store.conn.Where(v)
	return store
}

// WithContext ...
func (store *SQL) WithContext(ctx context.Context) *SQL {
	return &SQL{
		cfg:  store.cfg,
		conn: store.conn.WithContext(ctx),
	}
}
