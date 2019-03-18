package model

type PostTag struct {
	BaseModel
	PostId uint // post id
	TagId  uint // tag id
}
