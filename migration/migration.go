package migration

import (
	internalMigration "github.com/beyond3800/hawk/internal/console/migration"
)


type Migration = internalMigration.Migration

func Register(name string, m Migration) {
	
	internalMigration.Register(name, m)
}