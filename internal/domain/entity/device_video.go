package entity

type DeviceVideo struct {
	Model
	DeviceID    int64  `gorm:"not null" json:"device_id"`                     // 关联的设备ID
	EnableVideo bool   `gorm:"type:tinyint(1);default:0" json:"enable_video"` // 是否启用视频功能，0-禁用，1-启用
	VideoURL    string `gorm:"type:varchar(255)" json:"video_url"`            // 视频流地址或 ffmpeg 服务地址
	StreamType  string `gorm:"type:varchar(50)" json:"stream_type"`           // 视频流类型，如 RTSP、HLS 等
	Status      int64  `gorm:"type:tinyint(1);default:0" json:"status"`       // 状态，0-离线，1-在线
}
