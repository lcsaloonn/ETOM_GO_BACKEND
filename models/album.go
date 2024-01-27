package models


type Adress struct {
	State string `json:"state" bson:"user_state"`
}


type Album struct {
	Name string `json:"name" bson:"user_name"`
	Age int `json:"age" bson:"user_age"`
	Adress Adress  `json:"adress" bson:"user_adress"`
}

