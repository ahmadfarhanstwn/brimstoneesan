package brimstoneesan

import "database/sql"

type initPaths struct {
	rootPath    string
	folderNames []string
}

type cookieConfig struct {
	name     string
	lifetime string
	persist  string
	secure   string
	domain   string
}

type databaseConfig struct {
	databaseType string
	dsn          string
}

type Database struct {
	DatabaseType string
	Pool         *sql.DB
}
