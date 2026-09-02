package database

import "fmt"


func executeDropIndex(t *Table) error {

	if len(t.err) > 0 {
		for _, err := range t.err {
			if err != nil {
				return err
			}
		}
	}

	for _, index := range t.dropIndexes {

		query := fmt.Sprintf(
			"ALTER TABLE %s DROP INDEX %s;",
			t.name,
			index,
		)

		if _, err := t.db.Exec(query); err != nil {
			return err
		}
	}

	return nil
}

func (t *Table) DropIndex(name ...string) error {

	for _,indexName := range name{
		indx := index{
			Name:    indexName,
			Type:    "DROPINDEX",
		}
		t.err = append(t.err, t.addToIndex(indx))
	}
	return executeDropIndex(t)
}
