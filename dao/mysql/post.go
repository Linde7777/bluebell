package mysql

import "bluebell/models"

func InsertPost(pc *models.PostCreated) error {
	sqlStr := "insert into post(post_id, title, content, " +
		"author_id, community_id) values (?,?,?,?,?)"
	_, err := db.Exec(sqlStr, pc.ID, pc.Title, pc.Content,
		pc.AuthorID, pc.CommunityID)
	return err
}
