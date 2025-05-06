package err

var FileType = map[string]string{
	"TEXT":      "文本(限制pdf类型)(限制文件大小10M)",
	"IMAGE":     "图片(限制jpg、jpeg、png类型)(限制文件大小5M)(图片像素不低于400x400，不高于5000x5000)",
	"AUDIO":     "音频(限制mp3、wav类型)(限制文件大小50M)",
	"VIDEO":     "视频(限制mp4、avi、wmv类型)(限制文件大小200M)",
	"PDF_IMAGE": "系列图(限制pdf类型)(限制文件大小10M)(限制子图数量2-30)",
}
