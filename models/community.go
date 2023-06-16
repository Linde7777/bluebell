package models

// CommunitySelected store the community info,
// where come from SQL query
type CommunitySelected struct {
	ID   int    `json:"id" db:"community_id"`
	Name string `json:"name" db:"community_name"`
}
