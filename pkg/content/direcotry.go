package content

import (
	"io/ioutil"
)

func getDir(absolutePath string)[]Content{
	var all []Content
	files, err := ioutil.ReadDir(absolutePath)
	if err != nil{
		return all
	}
	for _, file := range files{
		c := Content{Name: file.Name(),
		IsDir: file.IsDir(),
		Size: file.Size(),
	}
		all = append(all, c)
	}
	return all
}