package content

type Content struct{
	Name string `json:"name"`
	IsDir bool `json:"isdir"`
	Info string `json:"info"`
	Size int64 `json:"size"`
	Contents []Content `json:"contents"`
}