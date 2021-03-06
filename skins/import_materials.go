package skins

import (
	"database/sql"
	"io"
	"time"

	"github.com/lflux/eve-sdeloader/statements"
	"github.com/lflux/eve-sdeloader/utils"
)

type SkinMaterial struct {
	DisplayNameID int64 `yaml:"displayNameID"`
	MaterialSetID int64 `yaml:"materialSetID"`
	ID            int64 `yaml:"skinMaterialID"`
}

func ImportMaterials(db *sql.DB, r io.Reader) error {
	defer utils.TimeTrack(time.Now(), "skin materials")

	entries := make(map[string]*SkinMaterial)

	err := utils.LoadFromReader(r, entries)
	if err != nil {
		return err
	}

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	stmt, err := statements.InsertMaterialStmt(tx)
	if err != nil {
		return err
	}

	for materialID, material := range entries {
		_, err = stmt.Exec(materialID, material.DisplayNameID, material.MaterialSetID)
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}
