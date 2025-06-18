package migrations

// Package migrations manages the DB schema migrations only.
// Data operations are handled by other services.

type Migration struct {
	ID   int
	Name string
	SQL  string
}

// migrations lists all migrations applied to the DB.
var migrations = []Migration{
	{
		ID:   1,
		Name: "create_coin_info_table",
		SQL: `
		CREATE TABLE IF NOT EXISTS coin_info (
			id SERIAL PRIMARY KEY,
			cmc_id INT NOT NULL,
			name VARCHAR(255) NOT NULL,
			symbol VARCHAR(255) NOT NULL,
			slug VARCHAR(255) NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		)
		`,
	},
}
