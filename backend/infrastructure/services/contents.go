package services

import (
	"fmt"
	"mailinglist/infrastructure/models"
)

func (s *services) CreateContent(contents models.Contents, userid string) (string, error) {
	data := fmt.Sprintf("insert into contents(account_id, content_type, title, contents) VALUES ('%v','%v','%v','%v')", userid, contents.Content_type, contents.Title, contents.Content_data)
	row, errdata := s.psql.PGExecQuery(data)
	return row, errdata
}

func (s *services) GetUserContent(userid string) ([]models.Userconten, error) {
	data := fmt.Sprintf("SELECT u.username, c.content_id, c.content_create, c.content_type, c.title,c.contents FROM account u INNER JOIN contents c ON u.account_id = c.account_id WHERE u.account_id = '%v'LIMIT 5 OFFSET 0", userid)
	_, contentrow := s.psql.PGRowQuery(data, false)

	var allcontent []models.Userconten
	for contentrow.Next() {
		var content models.Userconten
		err := contentrow.Scan(&content.Username, &content.Content_id, &content.Content_create, &content.Content_type, &content.Title, &content.Content_data)
		if err == nil {
			allcontent = append(allcontent, content)
		}
	}

	if contentrow.Err() != nil {
		return nil, contentrow.Err()
	}

	return allcontent, nil
}
