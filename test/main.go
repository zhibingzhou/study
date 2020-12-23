package main

import (
	// "encoding/json"

	"fmt"
	"image"
	"image/png"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"sort"

	// "strconv"
	"testing"

	//	"strconv"
	"strings"
	"time"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/robertkrimen/otto"

	// "github.com/jinzhu/gorm"
	"gitlab.stagingvip.net/publicGroup/public/common"
)

type abcd interface {
}

var workerPool *WorkerPool

func main() {

	shpay()
	// lognurl := "https://www.alipay.com/?appId=09999988&actionType=toCard&sourceId=bill&cardNo=621771*****24648&bankAccount=%e9%bb%84%e6%b5%a9&money=300&amount=300&bankMark=CITIC&bankName=%e4%b8%ad%e4%bf%a1%e9%93%b6%e8%a1%8c&cardIndex=&cardNoHidden=true&cardChannel=HISTORY_CARD&orderSource=from"

	// shorurl := GetShorUrl(lognurl, "imkey")
	// if shorurl != "" {
	// 	fmt.Println(shorurl)
	// }

	// Runjs()
	// abc := 1
	// mon := 300 - float64(abc)/100.00

	// mm := strconv.FormatFloat(mon, 'f', 2, 64)

	// fmt.Println(mm)

	// // for i := 1; i < 32; i++ {
	// // 	str := fmt.Sprintf(`SELECT SUM(amount) AS total FROM pay_list WHERE
	// // 	status = 3 AND pay_code = "bfpay" AND mer_code IN ("NB2020","XP001","hm2020","DW2020","1568xsddc") AND create_time > '2020-07-%d 00:00:00' AND create_time <  '2020-07-%d 23:59:59';`, i, i)
	// // 	fmt.Println(str)
	// // }
	// //shpay()
	r := gin.Default()
	//r.Use(Cors())
	r.GET("/tranfer", tranpay)
	r.GET("/query_pay", qpay)
	r.GET("/pay", pay)
	r.GET("/show", show)
	r.POST("/test", posttest)
	r.LoadHTMLGlob("view/**/*")
	r.LoadHTMLFiles("./view/success.tpl", "./view/jump.tpl")

	// msg_b := `amount=20000&merchantId=200280&orderNo=16029325171ldtrtm&repCode=0001&repMsg=讯息成功&resultUrl=http://47.96.72.217:8081/api/payPage.html?id=110742&sign=1d37a517eb4a6cb67e0cf58fb5d27b1e&tradeDate=20201017&tradeTime=190157`
	// var json_res map[string]string
	// json_res = UrlToMap(string(msg_b))
	// if fmt.Sprintf("%v", json_res["repCode"]) != "0001" {

	// }

	// img_url := fmt.Sprintf("%v", json_res["resultUrl"])

	// fmt.Println(img_url)
	// // a:= 7390457813701373

	// // agent_arr := strings.Split("123", "_")
	// // fmt.Println(len(agent_arr))

	// //db()
	// //  content := `您存款账户1743于8月18日16:39跨行转账转入人民币500.34元，活期余额人民币807.34元。【平安银行】`
	// //  a,b := ttt(content)

	// //  fmt.Println(a,b)
	// //YBPayFor()
	// //KsPayFor()
	// //LRXPayFor()

	r.Run(":8080")

}

func Runjs() {

	url := "http://www.baidu.com"
	jsfile := "./ttt.js"
	bytes, err := ioutil.ReadFile(jsfile)
	vm := otto.New()
	_, err = vm.Run(string(bytes))
	enc, err := vm.Call("base64encode", nil, url)
	fmt.Println(enc, err)
}

func Testfunc(t *testing.T) {
	fmt.Println("123")
}

func writePng(filename string, img image.Image) {
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	err = png.Encode(file, img)
	// err = jpeg.Encode(file, img, &jpeg.Options{100})      //图像质量值为100，是最好的图像显示
	if err != nil {
		log.Fatal(err)
	}
	file.Close()
	log.Println(file.Name())
}

func show(c *gin.Context) {

	img_url := "http://47.96.72.217:8081/api/payPage.html?id=110063"

	c.HTML(200, "jump.tpl", gin.H{
		"pay_url": img_url,
		"title":   "后台接口调试",
	})
}

var nk *StudentPool

func ttt(content string) (string, string) {

	card_num := ""
	amount := ""
	//通过切割，获取尾号信息
	c_arr := strings.Split(content, "尾号")
	if len(c_arr) > 1 {
		card_num = common.Substr(c_arr[1], 0, 4)
	}
	if !strings.Contains(content, "向您尾号") && !strings.Contains(content, "转账转入") {
		return amount, card_num
	}
	c_arr1 := strings.Split(content, "转账")
	if len(c_arr1) < 2 {
		c_arr1 = strings.Split(content, "转入")
		if len(c_arr1) < 2 {
			return amount, card_num
		}
	}
	//通过元去切割
	c_arr2 := strings.Split(c_arr1[1], "元")
	reg := regexp.MustCompile(`[\d]|[.]`) // 查找连续的数字
	amount_arr := reg.FindAllString(c_arr2[0], -1)
	if len(amount_arr) < 1 {
		return amount, card_num
	}
	for _, a_val := range amount_arr {
		amount = amount + a_val
	}
	return amount, card_num

}

func t2(out chan int, inin chan chan int) {
	for {
		inin <- out
		time.Sleep(time.Second * 5)
	}
}

func posttest(c *gin.Context) {
	c.JSON(200, &PGPAY{Notify_url: "123", Pay_url: "456"})
}
func pay(c *gin.Context) {
	//xwspay()
	//bfPay()
	//xPay()
	// cfpay()
	//MdPay()
	//klqdpay()
	//testchannel()
	// nk.WorkGO <- GotoChange{Score: 10, Gander: "girl"}
	//fmt.Println(<-nk.Result)
	base64 := "alipays://platformapi/startapp?appId=20000067&url=alipays%3A%2F%2Fplatformapi%2Fstartapp%3FappId%3D20000123%26actionType%3Dscan%26biz_data%3D%7B%22s%22%3A%22money%22%2C%22u%22%3A%222088402759328420%22%2C%22a%22%3A%226.66%22%2C%22m%22%3A%22PID%E8%BD%AC%E8%B4%A6%22%7D"
	log.Println("Original data:", base64)
	code, err := qr.Encode(base64, qr.L, qr.Unicode)
	// code, err := code39.Encode(base64)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Encoded data: ", code.Content())

	if base64 != code.Content() {
		log.Fatal("data differs")
	}

	code, err = barcode.Scale(code, 300, 300)
	if err != nil {
		log.Fatal(err)
	}

	c.Writer.Header().Set("Content-Type", "image/png")
	png.Encode(c.Writer, code)

}

func tranpay(c *gin.Context) {
	PayFor()
}

func qpay(c *gin.Context) {
	PayQuery()
}

func chanel() {
	workerPool = NewWorkerPool(1)
	workerPool.Run()
}

func MerRateInfo(mer_code, pay_code, class_code, bank_code string) MerRate {
	var m_info MerRate
	gdb.DB.Where("mer_code=? and pay_code=? and class_code=? and bank_code=?", mer_code, pay_code, class_code, bank_code).First(&m_info)
	return m_info
}

func db() {

	conn_str := fmt.Sprintf("%s:%s@%s(%s:%s)/%s?charset=utf8", "root", "Eric_191021", "tcp", "127.0.0.1", "3306", "pay_system")
	db, err := gorm.Open("mysql", conn_str)
	if err != nil {
		fmt.Println("conn_str->", conn_str)
		panic(err)
	}

	//db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&Address{})
	//db.SingularTable(true)
	gdb = Gorm{DB: db}
	mer_list := []map[string]string{}
	fmer_list := []map[string]string{}
	lieke_field := "private_key like ? and code like ?"
	like_first := "U" + "%"
	like_second := "hm20" + "%"
	table_name := "mer_list"
	p_where := map[string]interface{}{}
	p_where["code"] = "abcd"
	fields := []string{"code", "title", "domain", "title", "qq", "skype", "telegram", "phone", "email", "private_key", "amount", "total_in", "total_out", "is_agent"}
	//fmer_list, _ = PageList(table_name, "", 10, 0, fields, p_where)
	delete(p_where, "code")
	mer_list, _ = RLikePageList(table_name, lieke_field, like_first, like_second, 10, 0, fields, p_where)
	mer_list = append(mer_list, fmer_list...)
	fmt.Println(mer_list)
}

func Query(sql_str string) error {
	res := gdb.DB.Exec(sql_str)
	return res.Error
}

/**
*  根据系统订单查询
 */
func OrderById(order_number string) PayList {
	var p_list PayList
	gdb.DB.Model(&PayList{}).Where("id = ?", order_number).First(&p_list)
	return p_list
}

func RateList(mer_code, class_code, bank_code string) []MerRate {
	var m_info []MerRate
	gdb.DB.Where("mer_code = ?", mer_code).Find(&m_info)
	return m_info
}

func DateListTotal(table_name, date_field, s_date, e_date, like_sql, field string, p_where map[string]interface{}) (int, float64) {
	var c_total CountTotal

	gdb.DB.Table(table_name).Select(field).Where(p_where).Where(date_field+">=? and "+date_field+"<=? "+like_sql, s_date, e_date).Scan(&c_total)

	return c_total.Num, c_total.Total
}

func Trans(sql_arr []string) error {
	tx := gdb.DB.Begin()
	// 注意，一旦你在一个事务中，使用tx作为数据库句柄
	var err error
	for _, sql := range sql_arr {
		//更新订单状态
		if err = tx.Exec(sql).Error; err != nil {
			tx.Rollback()
			break
		}
	}

	if err != nil {
		return err
	}

	tx.Commit()
	return nil
}

func RLikePageList(table_name, like_field, like_first, like_second string, page_size, offset int, fields []string, p_where map[string]interface{}) ([]map[string]string, error) {
	records := []map[string]string{}

	u_rows, err := gdb.DB.Table(table_name).Select(fields).Where(like_field, like_first, like_second).Where(p_where).Limit(page_size).Offset(offset).Rows()
	if err != nil {
		return records, err
	}
	//创建有效切片
	values := make([]interface{}, len(fields))
	//行扫描，必须复制到这样切片的内存地址中去
	scanArgs := make([]interface{}, len(fields))

	for j := range values {
		scanArgs[j] = &values[j]
	}

	for u_rows.Next() {
		err = u_rows.Scan(scanArgs...)
		if err != nil {
			break
		}
		record := map[string]string{}
		for i, col := range values {
			col_s, ok := col.([]byte)
			if ok {
				record[fields[i]] = string(col_s)
			} else {
				record[fields[i]] = fmt.Sprintf("%v", col)
			}
		}
		records = append(records[0:], record)
	}
	return records, err
}

func LikeListTotal(table_name, like_field, like_where, field string, p_where map[string]interface{}) (int, float64) {
	var c_total CountTotal
	gdb.DB.Table(table_name).Select(field).Where(like_field+" like ?", like_where).Where(p_where).Scan(&c_total)

	return c_total.Num, c_total.Total
}

//苹果支付
type PGPAY struct {
	Notify_url string
	Pay_url    string
	Mer_code   string
	Key        string
}

type Re struct {
	Amount string
	Number string
}

var http_header map[string]string

func init() {

	http_header = make(map[string]string)
	http_header["Content-type"] = "application/x-www-form-urlencoded; charset=UTF-8"
	http_header["User-Agent"] = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.79 Safari/537.36"
}

/**
*  日期分页
 */
func SecondPageDateList(table_name, date_field, s_date, e_date, like_sql, field string, typevalue []string, page_size, offset int, fields []string,
	p_where map[string]interface{}) (
	[]map[string]string,
	error) {
	records := []map[string]string{}

	u_rows, err := gdb.DB.Table(table_name).Select(fields).Where(p_where).Where(
		date_field+">=? and "+date_field+"<=? "+like_sql, s_date, e_date).Where(field, typevalue).
		Limit(page_size).Order(date_field + " desc").
		Offset(offset).Rows()
	if err != nil {
		return records, err
	}
	//创建有效切片
	values := make([]interface{}, len(fields))
	//行扫描，必须复制到这样切片的内存地址中去
	scanArgs := make([]interface{}, len(fields))

	for j := range values {
		scanArgs[j] = &values[j]
	}

	for u_rows.Next() {
		err = u_rows.Scan(scanArgs...)
		if err != nil {
			break
		}
		record := map[string]string{}
		for i, col := range values {
			col_s, ok := col.([]byte)
			if ok {
				record[fields[i]] = string(col_s)
			} else {
				record[fields[i]] = fmt.Sprintf("%v", col)
			}
		}
		records = append(records[0:], record)
	}
	return records, err
}

//苹果支付
func pgPay(c *gin.Context) {

	tdpay := PGPAY{
		Notify_url: "https://www.baidu.com/",
		Pay_url:    "http://api.haofuqian.com/api/pay/index",
		Mer_code:   "2352452743",
		Key:        "2765f0c283617e9ce52754bfb8455a85d35c747c",
	}

	p := PayData{
		Amount:       "500.00",
		Order_number: "45628723449518823112365",
		Pay_bank:     "10109",
		Ip:           "127.0.0.1",
	}

	param_form := map[string]string{
		"partner":   tdpay.Mer_code,
		"amount":    p.Amount,
		"tradeNo":   p.Order_number,
		"notifyUrl": tdpay.Notify_url,
		"service":   p.Pay_bank,
	}

	//拼接
	result_url := MapCreatLinkSort(param_form, "&", true, false)
	result_url += fmt.Sprintf("&key=%s", tdpay.Key)
	sign := common.HexMd5(result_url)
	param_form["sign"] = sign
	fmt.Println(result_url)
	sbForm := "<html><head></head><body onload=\"document.forms[0].submit()\"><form name=\"order\" method=\"post\" action=\"PayBuildDomain\">"
	sbForm = strings.Replace(sbForm, "PayBuildDomain", tdpay.Pay_url, -1)
	for item := range param_form {
		sbForm += "<input name=\"" + item + "\" type=\"hidden\" value=\"" + param_form[item] + "\" />"
	}
	sbForm += "</form></body></html>"

	fmt.Println(sbForm)

}

//排序拼接
// mid 拼接符号
// url 是否是url模式
// conectnil 空值是否参与拼接 为true 参与拼接
func MapCreatLinkSort(m map[string]string, mid string, url bool, conectnil bool) string {
	var result = ""

	var keys []string
	for key := range m {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	for _, k := range keys {
		if conectnil {
			if url {
				result += k + "=" + fmt.Sprintf("%s", m[k]) + mid
			} else {
				result += k + fmt.Sprintf("%s", m[k]) + mid
			}
		} else if m[k] != "" {
			if url {
				result += k + "=" + fmt.Sprintf("%s", m[k]) + mid
			} else {
				result += k + fmt.Sprintf("%s", m[k]) + mid
			}
		}

	}

	if mid != "" {
		result = strings.TrimRight(result, mid)
	}

	return result
}

func CommonFieldsRow(table_name string, fields []string, c_w map[string]interface{}) (map[string]string, error) {
	record := map[string]string{}
	u_rows := gdb.DB.Table(table_name).Select(fields).Where(c_w).Row()

	//创建有效切片
	values := make([]interface{}, len(fields))
	//行扫描，必须复制到这样切片的内存地址中去
	scanArgs := make([]interface{}, len(fields))
	for j := range values {
		scanArgs[j] = &values[j]
	}

	err := u_rows.Scan(scanArgs...)

	for i, col := range values {
		col_s, ok := col.([]byte)
		if ok {
			record[fields[i]] = string(col_s)
		} else {
			record[fields[i]] = fmt.Sprintf("%v", col)
		}
	}

	return record, err
}

func testchannel() {

	sc := &UpdateCash{Order_number: "123", Pay_order: "123", Note: "", Order_type: 1}
	workerPool.JobQueue <- sc
	fmt.Println("im testchannel()")
	pool_res := <-workerPool.PoolRes
	fmt.Println("im testchannel() finished", pool_res)
}

func PageList(table_name, order_by string, page_size, offset int, fields []string, p_where map[string]interface{}) ([]map[string]string, error) {
	records := []map[string]string{}
	if order_by == "" {
		order_by = fields[0] + " desc"
	}
	u_rows, err := gdb.DB.Table(table_name).Select(fields).Where(p_where).Limit(page_size).Order(order_by).Offset(offset).Rows()
	if err != nil {
		return records, err
	}
	//创建有效切片
	values := make([]interface{}, len(fields))
	//行扫描，必须复制到这样切片的内存地址中去
	scanArgs := make([]interface{}, len(fields))

	for j := range values {
		scanArgs[j] = &values[j]
	}

	for u_rows.Next() {
		err = u_rows.Scan(scanArgs...)
		if err != nil {
			break
		}
		record := map[string]string{}
		for i, col := range values {
			col_s, ok := col.([]byte)
			if ok {
				record[fields[i]] = string(col_s)
			} else {
				record[fields[i]] = fmt.Sprintf("%v", col)
			}
		}
		records = append(records[0:], record)
	}
	return records, err
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method               //请求方法
		origin := c.Request.Header.Get("Origin") //请求头部
		var headerKeys []string                  // 声明请求头keys
		for k, _ := range c.Request.Header {
			headerKeys = append(headerKeys, k)
		}
		headerStr := strings.Join(headerKeys, ", ")
		if headerStr != "" {
			headerStr = fmt.Sprintf("access-control-allow-origin, access-control-allow-headers, %s", headerStr)
		} else {
			headerStr = "access-control-allow-origin, access-control-allow-headers"
		}
		if origin != "" {
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Origin", "*")                                       // 这是允许访问所有域
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE") //服务器支持的所有跨域请求的方法,为了避免浏览次请求的多次'预检'请求
			//  header的类型
			c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session,X_Requested_With,Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Pragma")
			//              允许跨域设置                                                                                                      可以返回其他子段
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers,Cache-Control,Content-Language,Content-Type,Expires,Last-Modified,Pragma,FooBar") // 跨域关键设置 让浏览器可以解析
			c.Header("Access-Control-Max-Age", "172800")                                                                                                                                                           // 缓存请求信息 单位为秒
			c.Header("Access-Control-Allow-Credentials", "false")                                                                                                                                                  //  跨域请求是否需要带cookie信息 默认设置为true
			c.Set("content-type", "application/json")                                                                                                                                                              // 设置返回格式是json
		}

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "Options Request!")
		}
		// 处理请求
		c.Next() //  处理请求
	}
}
