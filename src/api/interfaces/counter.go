package interfaces

type Counter struct {
	ID  string `json:"-" bson:"_id"`
	Seq int    `json:"-" bson:"seq"`
}
