package thread

import (
	"learn/model"
	"learn/shortid"
	"time"
)

func GOLongUrl(url string) (int, string) {
	t_status, url := model.Read_url_redis(url)
	return t_status, url

}

func Savetoredis(url string) (int, string) {
	var re model.Redirect

	re = model.Redirect{}
	now := time.Now()
	now = now.In(time.Local)
	code := shortid.MustGenerate()

	re.Code = code
	re.URL = url
	re.CreatedAt = now
	status := model.Write_url_redis(re)

	return status, code
}
