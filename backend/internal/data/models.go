package data

import (
	"database/sql"
	"errors"
)

var (
	ErrRecordNotFound = errors.New("record not found")
	ErrEditConflict   = errors.New("edit conflict")
)

type Models struct {
	Users    UserModelInterface
	Tokens   TokenModelInterface
	Stations StationModelInterface
	TempData TempDataModelInterface
}

func NewModels(db *sql.DB) *Models {
	return &Models{
		Users:    UserModel{DB: db},
		Tokens:   TokenModel{DB: db},
		Stations: StationModel{DB: db},
		TempData: TempDataModel{DB: db},
	}
}
