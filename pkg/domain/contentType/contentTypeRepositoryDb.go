package contentType

import (
	"database/sql"

	"gorm.io/gorm"

	"github.com/tamiat/backend/pkg/errs"
)

type ContentTypeRepositoryDb struct {
	db    *gorm.DB
	sqlDB *sql.DB
}

func (r ContentTypeRepositoryDb) isTableExists(id string) (string, error) {
	var name string
	row := r.db.Table("contenttype").Select("name").Where("id= ?", id).Row()
	err := row.Scan(&name)
	if err != nil {
		return "", errs.ErrContentTypeNotFound
	}
	return name, nil
}

func (r ContentTypeRepositoryDb) isColExists(tableName string, colName string) error {
	var numOfCols int
	row :=r.sqlDB.QueryRow("SELECT COUNT(*) FROM information_schema.columns WHERE table_name=$1 and column_name=$2", tableName, colName)
	err := row.Scan(&numOfCols)
	if err != nil {
		return errs.ErrDb
	}
	if numOfCols == 0 {
		return errs.ErrColNotFound
	}
	return nil
}

func (r ContentTypeRepositoryDb) Create(n string, cols string) (string, error) {
	_, err := r.sqlDB.Exec("INSERT INTO contentType (name) VALUES ($1)", n)
	if err != nil {
		return "", errs.ErrDb
	}
	_, err = r.sqlDB.Exec("CREATE TABLE " + n + "(" + cols + ")")
	if err != nil {
		return "", errs.ErrDb
	}
	var id string
	err = r.sqlDB.QueryRow("SELECT currval(pg_get_serial_sequence('contentType','id'));").Scan(&id)
	if err != nil {
		return "", errs.ErrDb
	}
	return id, nil
}

func (r ContentTypeRepositoryDb) DeleteById(id string) error {
	name, err := r.isTableExists(id)
	if err != nil {
		return err
	}
	_, err = r.sqlDB.Exec(`DELETE FROM contentType WHERE id= $1` , id)
	if err != nil {
		return err
	}
	_, err = r.sqlDB.Exec("DROP TABLE " + name)
	if err != nil {
		return err
	}
	return nil
}

func (r ContentTypeRepositoryDb) UpdateColName(id string, oldName string, newName string) error {
	name, err := r.isTableExists(id)
	if err != nil {
		return err
	}
	err = r.isColExists(name, oldName)
	if err != nil {
		return err
	}
	query := `ALTER TABLE ` + name + ` RENAME COLUMN ` + oldName + ` TO ` + newName
	_, err = r.sqlDB.Exec(query)
	if err != nil {
		return errs.ErrDb
	}
	return nil
}

func (r ContentTypeRepositoryDb) AddCol(id string, col string) error {
	name, err := r.isTableExists(id)
	if err != nil {
		return err
	}
	query := `ALTER TABLE ` + name + ` ADD COLUMN ` + col
	_, err = r.sqlDB.Exec(query)
	if err != nil {
		return errs.ErrDb
	}
	return nil
}

func (r ContentTypeRepositoryDb) DeleteCol(id string, col string) error {
	name, err := r.isTableExists(id)
	if err != nil {
		return err
	}
	err = r.isColExists(name, col)
	if err != nil {
		return err
	}
	newErr := r.db.Raw("ALTER TABLE ? DROP COLUMN ?", name, col)
	if newErr != nil {
		return errs.ErrDb
	}
	return nil
}

func NewContentTypeRepositoryDb(gormDb *gorm.DB, sqlDB *sql.DB) ContentTypeRepositoryDb {
	return ContentTypeRepositoryDb{gormDb, sqlDB}
}
