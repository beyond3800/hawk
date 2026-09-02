package database

import (
	"fmt"
	"strings"
)

type index struct {
	Name    string
	Columns []string
	Type    string
}
type check struct {
	Name      string
	Condition string
	column    string
}

func (t *Table) Index(name string, columns ...string) *Table {
	indx := index{
		Name:    name,
		Columns: columns,
		Type:    "INDEX",
	}
	t.err = append(t.err, t.addToIndex(indx))
	return t
}

func (t *Table) Unique(name string, columns ...string) *Table {
	indx := index{
		Name:    name,
		Columns: columns,
		Type:    "UNIQUE",
	}
	t.err = append(t.err, t.addToIndex(indx))
	return t
}

func (t *Table) FullText(name string, columns ...string) *Table {
	indx := index{
		Name:    name,
		Columns: columns,
		Type:    "FULLTEXT",
	}
	t.err = append(t.err, t.addToIndex(indx))
	return t
}


func (t *Table) addToIndex(indx index) error {

	if t.indexNameExists(indx.Name) {
        return fmt.Errorf(
            "index %q already exists on table %q",
            indx.Name,
            t.name,
        )
    }
	switch indx.Type {
	case "UNIQUE":
		t.uniques = append(t.uniques, indx)

	case "FULLTEXT":
		t.fullTexts = append(t.fullTexts, indx)

	case "INDEX":
		t.indexes = append(t.indexes, indx)

	case "DROPINDEX":
		t.dropIndexes = append(t.dropIndexes, indx.Name)

	default:
        return fmt.Errorf(
            "unsupported index type %q",
            indx.Type,
        )
	}
	return nil
}

func (t *Table) Check(name, condition string) *Table {

	t.checks = append(
		t.checks,
		check{
			Condition: condition,
			column: "",
		},
	)

	return t
}

func (t *Table) createIndexesQuery() ([]string, error){
	var createdIndex []string
	// at table level
	if len(t.indexes) > 0 {
		for _, index := range t.indexes {
			column := strings.Join(index.Columns, ",")
			createdIndex = append(createdIndex, fmt.Sprintf(" INDEX %s (%s)", index.Name, column))
			if err := t.validateEmptyIndex(index.Name, index.Columns...); err != nil {
				return nil,err
			}
		}
	}
	if len(t.uniques) > 0 {
		for _, index := range t.uniques {
			column := strings.Join(index.Columns, ",")
			createdIndex = append(createdIndex, fmt.Sprintf(" UNIQUE INDEX %s (%s)", index.Name, column))
			if err := t.validateEmptyIndex(index.Name, index.Columns...); err != nil {
				return nil,err
			}
		}
	}
	if len(t.fullTexts) > 0 {
		for _, index := range t.fullTexts {
			column := strings.Join(index.Columns, ",")
			createdIndex = append(createdIndex, fmt.Sprintf(" FULLTEXT INDEX %s (%s)", index.Name, column))
			if err := t.validateEmptyIndex(index.Name, index.Columns...); err != nil {
				return nil,err
			}
		}
	}
	if len(t.checks) > 0 {
		for _, check := range t.checks{
			if check.column != "" {
				query := fmt.Sprintf("%s INT CONSTRAINT %s CHECK (%s)", check.column, check.Name, check.Condition)
				createdIndex = append(createdIndex, query)
			}
			query := fmt.Sprintf(" CONSTRAINT %s CHECK (%s)", check.Name, check.Condition)
			createdIndex = append(createdIndex, query)
		}
	}
	if len(t.foreignKeys) > 0 {
		for _, foreignKey := range t.foreignKeys{
			column := strings.Join(foreignKey.Columns, ", ")
			
			var name string
			if foreignKey.Name == ""{
				name = fmt.Sprintf("%s_%s_foreign", t.name, foreignKey.Columns[0])
			}else{
				name = foreignKey.Name
			}
			// if t.indexNameExists(name){
			// 	return nil, fmt.Errorf("index %q already exists on table %q", name,t.name)
			// }

			query := fmt.Sprintf("CONSTRAINT %s FOREIGN KEY (%s) REFERENCES %s(%s)",
			name, column, foreignKey.Table, foreignKey.References)

			if foreignKey.OnDelete != ""{
				query += fmt.Sprintf(" ON DELETE %s", foreignKey.OnDelete)
			}
			if foreignKey.OnUpdate != ""{
				query += fmt.Sprintf(" ON UPDATE %s", foreignKey.OnUpdate)
			}
			
			createdIndex = append(createdIndex, query)
		}
		
	}
	return createdIndex,nil
}


