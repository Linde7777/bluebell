package models

// CommunityBasicSelected store the community
// basic info, where come from SQL query
type CommunityBasicSelected struct {
	ID   int    `json:"id" db:"community_id"`
	Name string `json:"name" db:"community_name"`
}

// CommunityDetailSelected store the community
// detail info, where come from SQL query
type CommunityDetailSelected struct {
	ID           int    `json:"id" db:"community_id"`
	Name         string `json:"name" db:"community_name"`
	Introduction string `json:"introduction" db:"introduction"`
}
