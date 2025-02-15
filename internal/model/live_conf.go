package model

const TableNameLiveConf = "live_confs"

// LiveConf mapped from table <live_confs>
type LiveConf struct {
	ID            int64  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	RoomDisplayID string `gorm:"column:room_display_id;not null" json:"room_display_id"`
	URL           string `gorm:"column:url;not null" json:"url"`
	Name          string `gorm:"column:name" json:"name"`
	ModifiedOn    int64  `gorm:"column:modified_on" json:"modified_on"`
	CreatedOn     int64  `gorm:"column:created_on" json:"created_on"`
	ModifiedBy    string `gorm:"column:modified_by" json:"modified_by"`
	CratedBy      string `gorm:"column:crated_by" json:"crated_by"`
	Enable        bool   `gorm:"column:enable;not null;" json:"enable"`
	Cron          string `gorm:"column:cron" json:"cron"`
}

// TableName LiveConf's table name
func (*LiveConf) TableName() string {
	return TableNameLiveConf
}

// select enable=1
func SelectEnableLiveConf() ([]*LiveConf, error) {
	var liveConfs []*LiveConf
	if err := DB.Where("enable = ?", 1).Find(&liveConfs).Error; err != nil {
		return nil, err
	}
	return liveConfs, nil
}

func GetAllLiveConf() ([]*LiveConf, error) {
	var liveConfs []*LiveConf
	if err := DB.Find(&liveConfs).Error; err != nil {
		return nil, err
	}
	return liveConfs, nil
}

// get by id
func GetLiveConfByID(id int64) (*LiveConf, error) {
	var liveConf *LiveConf
	if err := DB.Where("id = ?", id).First(&liveConf).Error; err != nil {
		return nil, err
	}
	return liveConf, nil
}
