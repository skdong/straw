package content

import(
	"os"
	"io/ioutil"
	"golang.org/x/text/encoding/simplifiedchinese"
)

func GetInfo(absultePath string)(Content, error){
	var c Content
	info, err := os.Stat(absultePath)
	if err != nil{
		return c, err
	}
	c.Name = info.Name()
	c.Size = info.Size()
	c.IsDir = info.IsDir()
	if c.IsDir{
		c.Contents = getDir(absultePath)
	}else{
		c.Info = readText(absultePath)
	}
	return c, nil
}

func readText(filename string)string{
	context, err := ioutil.ReadFile(filename)
	if err != nil{
		return ""
	}
	context,_ = simplifiedchinese.GBK.NewDecoder().Bytes(context)
	return string(context)
}