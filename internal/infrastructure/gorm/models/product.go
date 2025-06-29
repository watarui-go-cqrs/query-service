package models

// 商品データの格納
type Product struct {
	ID           int    `gorm:"column:p_key;primaryKey"`
	ObjId        string `gorm:"column:p_id;unique"`
	Name         string `gorm:"column:p_name"`
	Price        uint32 `gorm:"column:p_price"`
	CategoryId   string `gorm:"column:c_id"`
	CategoryName string `gorm:"column:c_name"`
}
