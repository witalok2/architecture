package postgres

const (
	sqlCreateUser        = `INSERT INTO "user" (name) VALUES ($1) RETURNING id`
	sqlCreateUserAddress = `SELECT * FROM "user"`
)
