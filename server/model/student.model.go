package model

type Student struct {
	Name           string `bson:"name,omitempty" json:"name,omitempty"`
	Gender         string `bson:"gender,omitempty" json:"gender,omitempty"`
	IdentityNumber string `bson:"_id,omitempty" json:"id,omitempty"`
	Class          string `bson:"class,omitempty" json:"class,omitempty"`
	Pin            string `bson:"pin,omitempty" json:"pin,omitempty"`
}
