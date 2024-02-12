package filesystem

const (
	MaxPhotoSize   = 15 << 20 // 15 MB
	MinPhotoSize   = 1 << 12  // 4 kB
	MaxRequestSize = 1 << 24  // 16 MB
	MaxTextLength  = 2000

	FileSystemPath      = "/tmp/filesystem/"
	AcceptedImageFormat = "image/png"
)
