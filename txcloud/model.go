package txcloud

type GetImageInfoResp struct {
	Format        string `json:"format"`
	Width         string `json:"width"`
	Height        string `json:"height"`
	Size          string `json:"size"`
	Md5           string `json:"md5"`
	FrameCount    string `json:"frame_count"`
	BitDepth      string `json:"bit_depth"`
	VerticalDpi   string `json:"vertical_dpi"`
	HorizontalDpi string `json:"horizontal_dpi"`
}
