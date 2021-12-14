package main

import (
	"database/sql"
	"net/http"
)

func (repo UrlRepoImpl) getLatestSequence() int {
	conn, err := repo.conn.db.Conn(repo.conn.ctx)
	if err != nil {
		repo.logger.Fatal("Error in connection with db", http.StatusInternalServerError)
	}
	defer conn.Close()
	var UrlId int
	err = conn.QueryRowContext(repo.conn.ctx, "SELECT nextval('url_Id')").Scan(&UrlId)
	return UrlId
}
func (repo UrlRepoImpl) insert(url UrlDetail) bool {
	repo.conn.db.NewInsert().
		Model(&url).
		Exec(repo.conn.ctx)
	return true
}

func (repo UrlRepoImpl) get(shortUrlId int) (UrlDetail, error) {
	url := new(UrlDetail)
	if err := repo.conn.db.NewSelect().
		Model(url).
		Where("id = ?", shortUrlId).
		Scan(repo.conn.ctx); err != nil {
		repo.logger.Println(err)
		return UrlDetail{}, err
	}
	return *url, nil
}

func (repo *UrlRepoImpl) init() (sql.Result, error) {
	conn, err := repo.conn.db.Conn(repo.conn.ctx)
	if err != nil {
		repo.logger.Fatal("Error in connection with db", http.StatusInternalServerError)
	}
	defer conn.Close()
	conn.ExecContext(repo.conn.ctx, "CREATE SEQUENCE IF NOT EXISTS url_Id")
	return repo.conn.db.NewCreateTable().Model((*UrlDetail)(nil)).IfNotExists().Exec(repo.conn.ctx)
}
