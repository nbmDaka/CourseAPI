package data

import (
	"context"
	"database/sql"
	"errors"
	_ "github.com/lib/pq"
	"time"
)

type ModuleInfo struct {
	ID             int64     `json:"id"`
	CreatedAt      time.Time `json:"-"`
	UpdatedAt      time.Time `json:"-"`
	ModuleName     string    `json:"moduleName"`
	ModuleDuration int32     `json:"moduleDuration"`
	ExamType       string    `json:"examType"`
	Version        int32     `json:"version"`
}
type ModuleInfoModel struct {
	DB *sql.DB
}

//func ValidateMovie(v *validator.Validator, moduleInfo *ModuleInfo) {
//	v.Check(moduleInfo.ModuleName != "", "moduleName", "must be provided")
//	v.Check(len(moduleInfo.ModuleName) <= 500, "moduleName", "must not be more than 500 bytes long")
//	v.Check(moduleInfo.ExamType != nil, "examType", "must be provided")
//	v.Check(validator.Unique(moduleInfo.Version), "genres", "must not contain duplicate values")
//}

func (m ModuleInfoModel) Insert(moduleInfo *ModuleInfo) error {

	query := `
	INSERT INTO n_beketovDB (moduleName, moduleDuration, examType)
	VALUES ($1, $2, $3)
	RETURNING id, created_at, version`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	args := []any{moduleInfo.ModuleName, moduleInfo.ModuleDuration, moduleInfo.ExamType}

	return m.DB.QueryRowContext(ctx, query, args...).Scan(&moduleInfo.ID, &moduleInfo.CreatedAt, &moduleInfo.Version)

}

func (m ModuleInfoModel) Get(id int64) (*ModuleInfo, error) {
	if id < 1 {
		return nil, ErrRecordNotFound
	}

	query := `
	SELECT  id, created_at, title, year, runtime, genres, version
	FROM n_beketovDB
	WHERE id = $1`

	var moduleInfo ModuleInfo

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, id).Scan(
		&moduleInfo.ID,
		&moduleInfo.CreatedAt,
		&moduleInfo.ModuleName,
		&moduleInfo.ModuleDuration,
		&moduleInfo.ExamType,
		&moduleInfo.Version,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &moduleInfo, nil
}

func (m ModuleInfoModel) Update(moduleInfo *ModuleInfo) error {
	// Add the 'AND version = $6' clause to the SQL query.
	query := `
	UPDATE n_beketovDB
	SET moduleName = $1, moduleDuration = $2, examType = $3, version = version + 1
	WHERE id = $5 AND version = $6
	RETURNING version`
	args := []any{
		moduleInfo.ModuleName,
		moduleInfo.ModuleDuration,
		moduleInfo.ExamType,
		moduleInfo.ID,
		moduleInfo.Version,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, args...).Scan(&moduleInfo.Version)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return ErrEditConflict
		default:
			return err
		}
	}
	return nil
}

func (m ModuleInfoModel) Delete(id int64) error {
	// Return an ErrRecordNotFound error if the movie ID is less than 1.
	if id < 1 {
		return ErrRecordNotFound
	}
	// Construct the SQL query to delete the record.
	query := `
	DELETE FROM n_beketovDB
	WHERE id = $1`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	result, err := m.DB.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrRecordNotFound
	}
	return nil
}
