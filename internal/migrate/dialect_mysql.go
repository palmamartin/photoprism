// Code generated by go generate; DO NOT EDIT.
package migrate

var DialectMySQL = Migrations{
	{
		ID:         "20211121-094727",
		Dialect:    "mysql",
		Statements: []string{"DROP INDEX uix_places_place_label ON `places`;"},
	},
	{
		ID:         "20211124-120008",
		Dialect:    "mysql",
		Statements: []string{"DROP INDEX idx_places_place_label ON `places`;", "DROP INDEX uix_places_label ON `places`;"},
	},
	{
		ID:         "20220103-115400",
		Dialect:    "mysql",
		Statements: []string{"ALTER TABLE files MODIFY file_projection VARBINARY(40) NULL;", "ALTER TABLE files MODIFY file_color_profile VARBINARY(40) NULL;"},
	},
	{
		ID:         "20220118-172400",
		Dialect:    "mysql",
		Statements: []string{"ALTER TABLE albums MODIFY album_filter VARBINARY(767) DEFAULT '';", "CREATE INDEX IF NOT EXISTS idx_albums_album_filter ON albums (album_filter);"},
	},
}
