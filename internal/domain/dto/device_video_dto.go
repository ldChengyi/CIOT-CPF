package dto

type DeviceVideoDTO struct {
	ID          int64  `json:"id" xml:"id" form:"id" query:"id"`
	DeviceID    int64  `json:"device_id" xml:"device_id" form:"device_id" query:"device_id"`             // 关联的设备ID
	EnableVideo bool   `json:"enable_video" xml:"enable_video" form:"enable_video" query:"enable_video"` // 是否启用视频功能，0-禁用，1-启用
	VideoURL    string `json:"video_url" xml:"video_url" form:"video_url" query:"video_url"`             // 视频流地址或 ffmpeg 服务地址
	StreamType  string `json:"stream_type" xml:"stream_type" form:"stream_type" query:"stream_type"`     // 视频流类型，如 RTSP、HLS 等
	Status      int64  `json:"status" xml:"status" form:"status" query:"status"`                         // 状态，0-离线，1-在线
}
