package mysql

import (
	"bluebell/models"
	"github.com/jmoiron/sqlx"
	"strings"
)

func InsertPost(pc *models.PostBasic) error {
	// sorry for the following long text, if I spilt it,
	//the IDE will view it as error
	sqlStr := "insert into post(post_id, title, content, author_id, community_id) values (?,?,?,?,?)"
	_, err := db.Exec(sqlStr, pc.ID, pc.Title, pc.Content,
		pc.AuthorID, pc.CommunityID)
	return err
}

func GetPostBasicByID(id int64) (*models.PostBasic, error) {
	sqlStr := "select post_id,title,content,author_id," +
		"community_id,create_time from post where post_id=?"
	p := new(models.PostBasic)
	err := db.Get(p, sqlStr, id)
	return p, err
}

func GetPostBasicList(targetPageNumber, pageSize int64) (postBasicList []*models.PostBasic, err error) {
	sqlStr := "select post_id,title,content,author_id," +
		"community_id,create_time from post order by create_time desc limit ?,? "
	postBasicList = make([]*models.PostBasic, 0, 2)

	// in frontend, the initial page number is 1,
	// in database, the index of post start at 0,
	// that's why I write (targetPageNumber-1)*pageSize
	err = db.Select(&postBasicList, sqlStr, (targetPageNumber-1)*pageSize,
		pageSize)
	return
}

func GetPostBasicListByIDs(IDs []string) (postBasicList []*models.PostBasic, err error) {
	sqlStr := "select post_id,title,content,author_id," +
		"community_id,create_time from post where post_id in (?) " +
		"order by FIND_IN_SET(post_id,?)"
	query, args, err := sqlx.In(sqlStr, IDs, strings.Join(IDs, ","))
	if err != nil {
		return nil, err
	}

	query = db.Rebind(query)
	err = db.Select(&postBasicList, query, args...)
	return
}
