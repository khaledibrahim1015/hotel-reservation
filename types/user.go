package types

type User struct {
	// ID        string `bson:"_id" json:"id,omitempty"`
	ID        string `bson:"_id" json:"id"`
	FirstName string `bson:"firstName" json:"firstName"`
	LastName  string `bson:"lastName"lastName"`
}
