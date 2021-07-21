package config

type Upload struct {
	UploadMaxSize int
	SaveSameFile  bool
	UploadPath    string
	Url           string
	AllowExts     []string
}
