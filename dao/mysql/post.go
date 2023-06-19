package mysql

import "bluebell/models"

func InsertPost(pc *models.Post) error {
	// sorry for the following long text, if I spilt it,
	//the IDE will view it as error
	sqlStr := "insert into post(post_id, title, content, author_id, community_id) values (?,?,?,?,?)"
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

func GetPostDetailList() (postList []*models.Post, err error) {
	sqlStr := "select post_id,title,content,author_id," +
		"community_id,create_time from post limit 2"
	postList = make([]*models.Post, 0, 2)
	err = db.Select(&postList, sqlStr)
	return
}
