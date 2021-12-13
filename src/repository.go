package main

import (
	"net/http"
)

func (repo UrlRepo) getLatestSequence() int {
	conn, err := repo.conn.db.Conn(repo.conn.ctx)
	if err != nil {
		repo.logger.Fatal("Error reading request body", http.StatusInternalServerError)
	}
	defer conn.Close()
	var UrlId int
	err = conn.QueryRowContext(repo.conn.ctx, "SELECT nextval('public.\"url_Id\"')").Scan(&UrlId)
	return UrlId
}
func (repo UrlRepo) insert(url UrlDetail) bool {
	repo.conn.db.NewInsert().
		Model(&url).
		Exec(repo.conn.ctx)
	return true
}

func (repo UrlRepo) get(shortUrlId int) (UrlDetail, error) {
	url := new(UrlDetail)
	if err := repo.conn.db.NewSelect().
		Model(url).
		Where("id = ?", shortUrlId).
		Scan(repo.conn.ctx); err != nil {
		repo.logger.Print(err)
		return UrlDetail{}, err
	}
	return *url , nil
}
