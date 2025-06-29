package models

type Category struct {
	ID    int    `gorm:"column:c_key;primaryKeY"`
	ObjId string `gorm:"column:c_id;unique"`
	Name  string `gorm:"column:c_name"`
}
