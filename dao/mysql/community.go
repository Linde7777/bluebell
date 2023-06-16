package mysql

import (
	"bluebell/models"
	"database/sql"
	"go.uber.org/zap"
)

func GetCommunityList() (communityList []*models.CommunityBasicSelected, err error) {
	sqlStr := "select community_id,community_name from community"
	if err = db.Select(&communityList, sqlStr); err != nil {
		if err == sql.ErrNoRows {
			zap.L().Warn("table community is empty")
			err = nil
		}
	}

	return
}

// GetCommunityDetail will return rowIsEmpty:true
// and error:nil if it meets sql.ErrNoRows
func GetCommunityDetail(id int64) (
	communityDetail *models.CommunityDetailSelected,
	rowIsEmpty bool, err error) {

	communityDetail = new(models.CommunityDetailSelected)
	sqlStr := "select community_id,community_name,introduction," +
		"create_time from community where community_id=?"
	if err = db.Get(communityDetail, sqlStr, id); err != nil {
		if err == sql.ErrNoRows {
			err = nil
			return nil, true, err
		}
		return nil, false, err
	}

	return
}
