package vo

import "time"

type DeviceVideoAdminVo struct {
	ID          int       `json:"id"`
	DeviceName  string    `json:"device_name"`
	DeviceID    int       `json:"device_id"`    // 关联的设备ID
	EnableVideo bool      `json:"enable_video"` // 是否启用视频功能，0-禁用，1-启用
	VideoURL    string    `json:"video_url"`    // 视频流地址或 ffmpeg 服务地址
	StreamType  string    `json:"stream_type"`  // 视频流类型，如 RTSP、HLS 等
	Status      int64     `json:"status"`       // 状态，0-离线，1-在线
	CreatedAt   time.Time `json:"created_at"`
}
