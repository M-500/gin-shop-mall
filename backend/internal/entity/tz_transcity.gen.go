// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package entity

const TableNameTranscity = "tz_transcity"

// Transcity mapped from table <tz_transcity>
type Transcity struct {
	TranscityID int64 `gorm:"column:transcity_id;primaryKey;autoIncrement:true" json:"transcity_id"`
	TransfeeID  int64 `gorm:"column:transfee_id;comment:运费项id" json:"transfee_id"` // 运费项id
	CityID      int64 `gorm:"column:city_id;comment:城市id" json:"city_id"`          // 城市id
}

// TableName Transcity's table name
func (*Transcity) TableName() string {
	return TableNameTranscity
}
