package jsondb

import (
	"database/sql"
	"encoding/json"
	"log"
)

type JsonDB struct {
	db *sql.DB
}

func NewJsonDB(db *sql.DB) (*JsonDB, error) {
	if db == nil {
		return nil, "The database connection must be already initialized"
	}

	j := &JsonDB{db: db}

	return j, nil
}

func (j JsonDB) Json(statement string) string {
	rows, err := j.db.Query(statement)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		log.Fatal(err)
	}
	count := len(columns)
	tableData := make([]map[string]interface{}, 0)
	values := make([]interface{}, count)
	valuePtrs := make([]interface{}, count)

	for rows.Next() {
		for i := 0; i < count; i++ {
			valuePtrs[i] = &values[i]
		}
		rows.Scan(valuePtrs...)
		entry := make(map[string]interface{})
		for i, col := range columns {
			var v interface{}
			val := values[i]
			b, ok := val.([]byte)
			if ok {
				v = string(b)
			} else {
				v = val
			}
			entry[col] = v
		}
		tableData = append(tableData, entry)
	}
	jsonData, err := json.Marshal(tableData)
	if err != nil {
		log.Fatal(err)
	}
	return string(jsonData)
}
