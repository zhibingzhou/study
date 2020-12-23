package common

import (
	"fmt"
	"io/ioutil"
	"learn/model"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

var ratelimit = time.Tick(300 * time.Millisecond)

//<a href="/fenzhan.asp?sheng=黑龙江&shi=齐齐哈尔&quyu=龙沙区">龙沙区</a>

var citylistRe = `<a href="(/[^/]*/)" target="_blank">([^-]*)-&gt;</a>`

//<a href="/user/26057.html" target="_blank">威信sd1816</a>
var userlistRe = `<a href="(/user/[0-9]+.html)" target="_blank">([^<]+)</a>`

//<option value=/fenzhan.asp?Page=2&sheng=%CC%EC%BD%F2&shi=&quyu=>2</option>

var pagelistRe = `<option value=([^>]*)>([0-9]+)</option>`

//解析url
func Fetch(url string) ([]byte, error) {
	<-ratelimit //等待时间
	re, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if re.StatusCode == http.StatusOK {
		all, err := ioutil.ReadAll(re.Body)
		if err != nil {
			return nil, err
		}
		return all, nil
	}
	defer re.Body.Close()

	return nil, fmt.Errorf("wrong!!")
}

//城市列表解析
func ParseCityList(contents []byte) ParseResult {

	re := regexp.MustCompile(citylistRe)
	math := re.FindAllSubmatch(contents, -1)
	result := ParseResult{}
	for _, m := range math {
		result.Items = append(result.Items, "city:"+string(m[1]))
		result.Requests = append(result.Requests, Request{Url: GoUrl + string(m[1]), ParserFunc: func(c []byte) ParseResult {
			return PageUserList(c)
		}})
	}

	return result
}

//分页
func PageUserList(contens []byte) ParseResult {

	re := regexp.MustCompile(pagelistRe)
	math := re.FindAllSubmatch(contens, -1)
	result := ParseResult{}
	for _, m := range math {
		result.Items = append(result.Items, "分页:"+string(m[2]))
		result.Requests = append(result.Requests, Request{Url: GoUrl + string(m[1]), ParserFunc: func(c []byte) ParseResult {
			return ParseUserList(c)
		}})
	}

	return result

}

//用户列表解析
func ParseUserList(contents []byte) ParseResult {

	re := regexp.MustCompile(userlistRe)
	math := re.FindAllSubmatch(contents, -1)
	result := ParseResult{}
	for _, m := range math {
		Url := GoUrl + string(m[1])
		Name := string(m[2])
		result.Items = append(result.Items, "user:"+string(m[2]))
		result.Requests = append(result.Requests, Request{Url: GoUrl + string(m[1]), ParserFunc: func(c []byte) ParseResult {
			return ParseMerInfoList(c, Name, Url)
		}})
	}

	return result
}

//用户信息解析
func ParseMerInfoList(contents []byte, name, mer_url string) ParseResult {

	mer_Info := model.Information{}
	result := ParseResult{}

	mer_Info.Name = name
	mer_Info.GoUrl = mer_url
	mer_str_id := ""

	pare_id := `<span class="ja_c_uid">ID:([0-9]+)</span>`
	re := regexp.MustCompile(pare_id)
	math := re.FindAllSubmatch(contents, -1)
	for _, m := range math {
		mer_Info.Id, _ = strconv.Atoi(string(m[1]))
		mer_str_id = string(m[1])
	}

	dom, err := goquery.NewDocumentFromReader(strings.NewReader(string(contents)))
	if err != nil {
		fmt.Println("ParseMerInfoList  ", err.Error())
	}

	dom.Find(".new_pd_about_label").Each(func(i int, selection *goquery.Selection) {
		message := selection.Text()

		message = DeleteExtraSpace(message)

		mer_map := strings.Split(message, "\n")

		mer_Info.Age, _ = strconv.Atoi(strings.Replace(mer_map[1], "岁", "", -1))
		mer_Info.Height, _ = strconv.ParseFloat(strings.Replace(mer_map[2], "cm", "", -1), 64)
		mer_Info.Weight, _ = strconv.ParseFloat(strings.Replace(mer_map[3], "kg", "", -1), 64)
		mer_Info.Body = strings.Replace(mer_map[4], "身材：", "", -1)
		mer_Info.Address = strings.Replace(mer_map[5], "所在地：", "", -1)

	})

	dom.Find(".new_pd_about_label2").Each(func(i int, selection *goquery.Selection) {
		message := selection.Text()

		message = DeleteExtraSpace(message)

		mer_map := strings.Split(message, " ")

		mer_Info.Job = strings.Replace(strings.Replace(mer_map[0], "\n", "", -1), "职业：", "", -1)
		mer_Info.Marry = mer_map[1]
		mer_Info.School = mer_map[2]
		strmoney := strings.Replace(mer_map[3], "年收入：", "", -1)
		moneymap := strings.Split(strmoney, "万")
		moneymapend := strings.Split(moneymap[0], "-")

		if len(moneymapend) > 1 {
			mer_Info.MoneyMin, _ = strconv.Atoi(moneymapend[0])
			mer_Info.MoneyMax, _ = strconv.Atoi(moneymapend[1])
		} else {
			mer_Info.MoneyMin = 0
			mer_Info.MoneyMax, _ = strconv.Atoi(moneymapend[0])
		}

	})

	dom.Find(".new_nxdb").Each(func(i int, selection *goquery.Selection) {
		message := selection.Text()

		message = DeleteExtraSpace(message)

		mer_map := strings.Split(message, "\n")

		mer_Info.Note = mer_map[1]
	})

	imageurl := `<img src="(/photo/[^"]+)" id="cpic" class="cpic" [^>]*>`

	re = regexp.MustCompile(imageurl)
	math = re.FindAllSubmatch(contents, -1)
	for _, m := range math {
		mer_Info.ImgUrl = GoUrl + string(m[1])
	}

	gender := `<i class="i i_s1"></i>`

	mer_Info.Gender = "girl"

	if strings.Contains(string(contents), gender) {
		mer_Info.Gender = "man"
	}

	result.Items = append(result.Items, mer_Info)

	fmt.Println(mer_Info)

	//写入redis 和 数据库 去重
	model.MerInfoRedis(mer_str_id, mer_Info)
	// model.Job_listRedis(mer_Info.Job, model.Job_list{Job_title: mer_Info.Job})
	// model.Marry_listRedis(mer_Info.Marry, model.Marry_list{Marry_title: mer_Info.Marry})
	// model.School_listRedis(mer_Info.School, model.School_list{School_title: mer_Info.School})
	// model.Body_listRedis(mer_Info.Body, model.Body_list{Body_title: mer_Info.Body})

	return result
}
