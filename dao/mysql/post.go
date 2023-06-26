package mysql

import (
	"bluebell/models"
	"github.com/jmoiron/sqlx"
	"strings"
)

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

func GetPostDetailList(targetPageNumber, pageSize int64) (postList []*models.Post, err error) {
	sqlStr := "select post_id,title,content,author_id," +
		"community_id,create_time from post order by create_time desc limit ?,? "
	postList = make([]*models.Post, 0, 2)

	// in frontend, the initial page number is 1,
	// in database, the index of post start at 0,
	// that's why I write (targetPageNumber-1)*pageSize
	err = db.Select(&postList, sqlStr, (targetPageNumber-1)*pageSize,
		pageSize)
	return
}

func GetPostDetailListByIDs(IDs []string) (postList []*models.Post, err error) {
	sqlStr := "select post_id,title,content,author_id," +
		"community_id,create_time from post where post_id in (?) " +
		"order by FIND_IN_SET(post_id,?)"
	query, args, err := sqlx.In(sqlStr, IDs, strings.Join(IDs, ","))
	if err != nil {
		return nil, err
	}

	query = db.Rebind(query)
	err = db.Select(&postList, query, args...)
	return
}
