package mysql

import "bluebell/models"

func InsertPost(pc *models.Post) error {
	sqlStr := "insert into post(post_id, title, content, " +
		"author_id, community_id) values (?,?,?,?,?)"
	_, err := db.Exec(sqlStr, pc.ID, pc.Title, pc.Content,
		pc.AuthorID, pc.CommunityID)
	return err
}

func GetPostDetailByID(id int64) (*models.Post, error) {
	sqlStr := "select post_id,title,content,author_id," +
		"community_id,create_time from post where post_id=?"
	p := new(models.Post)
	err := db.Get(p, sqlStr, id)
	return p, err
}
