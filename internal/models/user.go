package models

type User struct {
	ID    string `json:"id" bson:"_id,omitempty"` // Change ID to string for JSON
	Name  string `json:"name" bson:"name"`
	Email string `json:"email" bson:"email"`
}
