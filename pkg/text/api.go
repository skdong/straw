package text
import(
	"io/ioutil"
	"golang.org/x/text/encoding/simplifiedchinese"
)

func GetText(filename string)string{
	return readText(filename)
}

func readText(filename string)string{
	context, err := ioutil.ReadFile(filename)
	if err != nil{
		return ""
	}
	context,_ = simplifiedchinese.GBK.NewDecoder().Bytes(context)
	return string(context)
}