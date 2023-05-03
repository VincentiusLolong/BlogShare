package services

import (
	"fmt"
	"mailinglist/infrastructure/models"
)

func (s *services) CreateContent(contents models.Contents, userid string) (string, error) {
	data := fmt.Sprintf("insert into contents(account_id, content_type, title, contents) VALUES ('%v','%v','%v','%v')", contents.Account_id, contents.Content_type, contents.Title, contents.Contents)
	row, errdata := s.psql.PGExecQuery(data)
	return row, errdata
}
