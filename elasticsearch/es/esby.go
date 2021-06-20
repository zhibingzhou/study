package es

import (
	"context"
	"elastric/model"
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"strconv"

	"github.com/olivere/elastic/v7"
)

// 索引mapping定义，这里仿微博消息构定义
const mapping = `
{
  "mappings": {
    "properties": {
      "user": {
        "type": "keyword"
      },
      "message": {
        "type": "text"
      },
      "image": {
        "type": "keyword"
      },
      "created": {
        "type": "date"
      },
      "tags": {
        "type": "keyword"
      },
      "location": {
        "type": "geo_point"
      },
      "suggest_field": {
        "type": "completion"
      }
    }
  }
}`

type Job interface {
	Insert()
	Select()
	Update()
	Delete()
}

//根据id
type ElaById struct {
	Id        string
	Index     string
	Msg       model.Weibo
	Iditems   []string
	Search    string
	SearchMap map[string]string
}

func (e ElaById) Insert() {
	// 执行ES请求需要提供一个上下文对象
	ctx := context.Background()

	// 首先检测下model.Weibo索引是否存在
	exists, err := model.ElaClient.IndexExists(e.Index).Do(ctx)
	if err != nil {
		// Handle error
		panic(err)
	}
	if !exists {
		// model.Weibo索引不存在，则创建一个
		_, err := model.ElaClient.CreateIndex(e.Index).Do(ctx)
		if err != nil {
			// Handle error
			panic(err)
		}
	}

	// 使用client创建一个新的文档
	put1, err := model.ElaClient.Index().
		Index(e.Index).  // 设置索引名称
		Id(e.Id).        // 设置文档id
		BodyJson(e.Msg). // 指定前面声明的微博内容
		Do(ctx)          // 执行请求，需要传入一个上下文对象
	if err != nil {
		// Handle error
		panic(err)
	}

	fmt.Printf("文档Id %s, 索引名 %s\n", put1.Id, put1.Index)
}

func (e ElaById) Select() {

	ctx := context.Background()

	switch e.Search {
	case "id":
		// 根据单个id查询文档
		get1, err := model.ElaClient.Get().
			Index(e.Index). // 指定索引名
			Id(e.Id).       // 设置文档id
			Do(ctx)         // 执行请求
		if err != nil {
			// Handle error
			panic(err)
		}
		if get1.Found {
			fmt.Printf("文档id=%s 版本号=%d 索引名=%s\n", get1.Id, get1.Version, get1.Index)
		}

		// 手动将文档内容转换成go struct对象
		msg2 := model.Weibo{}
		// 提取文档内容，原始类型是json数据
		data, _ := get1.Source.MarshalJSON()
		// 将json转成struct结果
		json.Unmarshal(data, &msg2)
		// 打印结果
		fmt.Println(string(data))
	case "id_group":
		if len(e.Iditems) > 0 { //多id 查询
			var result []*elastic.MultiGetItem
			for key, value := range e.Iditems {
				result = append(result, elastic.NewMultiGetItem().Index(value).Id(strconv.Itoa(key)))
			}
			res, err := model.ElaClient.MultiGet().Add(result...).Do(ctx)
			if err != nil {
				panic(err)
			}

			// 遍历文档
			for _, doc := range res.Docs {
				// 转换成struct对象
				var content model.Weibo
				tmp, _ := doc.Source.MarshalJSON()
				err := json.Unmarshal(tmp, &content)
				if err != nil {
					panic(err)
				}

				fmt.Println("多id获取结果 ：", content)
			}
			return
		}
	case "select", "select_in", "select_like", "select_between", "select_and", "select_or", "select_groupby", "select_count": //查询单个条件可以用

		if e.Search == "select_in" { // 相當于 sql  where  user in(olivere,olivere1)

			termQuery := elastic.NewTermsQuery("user", "olivere", "olivere1")
			//termQuery := elastic.NewTermsQuery("location", "here", "olivere1")
			searchResult, err := model.ElaClient.Search().
				Index(e.Index).        // 设置索引名
				Query(termQuery).      // 设置查询条件
				Sort("created", true). // 设置排序字段，根据Created字段升序排序，第二个参数false表示逆序
				From(0).               // 设置分页参数 - 起始偏移量，从第0行记录开始
				Size(10).              // 设置分页参数 - 每页大小
				Pretty(true).          // 查询结果返回可读性较好的JSON格式
				Do(ctx)                // 执行请求

			if err != nil {
				// Handle error
				panic(err)
			}

			fmt.Printf("查询消耗时间 %d ms, 结果总数: %d\n", searchResult.TookInMillis, searchResult.TotalHits())
			ShowSelect(searchResult)

		}

		if e.Search == "select_like" { // 相当于 where  message like %酱油%

			// 创建match查询条件
			matchQuery := elastic.NewMatchQuery("message", "酱油")

			searchResult, err := model.ElaClient.Search().
				Index(e.Index).        // 设置索引名
				Query(matchQuery).     // 设置查询条件
				Sort("created", true). // 设置排序字段，根据Created字段升序排序，第二个参数false表示逆序
				From(0).               // 设置分页参数 - 起始偏移量，从第0行记录开始
				Size(10).              // 设置分页参数 - 每页大小
				Do(ctx)

			if err != nil {
				// Handle error
				panic(err)
			}
			fmt.Printf("查询消耗时间 %d ms, 结果总数: %d\n", searchResult.TookInMillis, searchResult.TotalHits())
			ShowSelect(searchResult)

		}

		if e.Search == "select_between" {
			// 例1 等价表达式： Created > "2020-07-20" and Created < "2020-07-29"
			rangeQuery := elastic.NewRangeQuery("created").
				Gt("2021-06-01").
				Lt("2021-07-29")

			// // 例2 等价表达式： id >= 1 and id < 10
			// rangeQuery :=  model.ElaClient.NewRangeQuery("id").
			// 	Gte(1).
			// 	Lte(10)

			searchResult, err := model.ElaClient.Search().
				Index(e.Index).        // 设置索引名
				Query(rangeQuery).     // 设置查询条件
				Sort("created", true). // 设置排序字段，根据Created字段升序排序，第二个参数false表示逆序
				From(0).               // 设置分页参数 - 起始偏移量，从第0行记录开始
				Size(10).              // 设置分页参数 - 每页大小
				Pretty(true).          // 查询结果返回可读性较好的JSON格式
				Do(ctx)                // 执行请求

			if err != nil {
				// Handle error
				panic(err)
			}
			fmt.Printf("查询消耗时间 %d ms, 结果总数: %d\n", searchResult.TookInMillis, searchResult.TotalHits())
			ShowSelect(searchResult)

		}

		if e.Search == "select_and" { //联合查询
			// 创建bool查询  为组合查询
			boolQuery := elastic.NewBoolQuery().Must() //Must

			// 创建term查询
			termQuery := elastic.NewTermQuery("retweets", "0")
			matchQuery := elastic.NewMatchQuery("message", "打酱油") //like  相当于 where retweets = 0 and message like "%打酱%"
			// 创建term查询
			termQueryN := elastic.NewTermQuery("user", "olivere1")
			// 设置bool查询的must条件, 组合了两个子查询
			// 表示搜索匹配Author=tizi且Title匹配"golang es教程"的文档
			boolQuery.Must(termQuery, matchQuery) //为 and 条件
			boolQuery.MustNot(termQueryN)         //意思是 user != olivere1

			searchResult, err := model.ElaClient.Search().
				Index(e.Index).        // 设置索引名
				Query(boolQuery).      // 设置查询条件
				Sort("created", true). // 设置排序字段，根据Created字段升序排序，第二个参数false表示逆序
				From(0).               // 设置分页参数 - 起始偏移量，从第0行记录开始
				Size(10).              // 设置分页参数 - 每页大小
				Do(ctx)                // 执行请求

			if err != nil {
				// Handle error
				panic(err)
			}
			fmt.Printf("查询消耗时间 %d ms, 结果总数: %d\n", searchResult.TookInMillis, searchResult.TotalHits())
			ShowSelect(searchResult)

		}

		if e.Search == "select_or" {
			// 创建bool查询  为组合查询
			boolQuery := elastic.NewBoolQuery().Must() //Must

			// 创建term查询
			termQuery := elastic.NewTermQuery("retweets", "12")
			matchQuery := elastic.NewMatchQuery("message", "打酱油") //like  相当于 where retweets = 0 or message like "%打酱%"

			// 设置bool查询的must条件, 组合了两个子查询
			// 表示搜索匹配Author=tizi且Title匹配"golang es教程"的文档
			boolQuery.Should(termQuery, matchQuery) //为 or 条件

			searchResult, err := model.ElaClient.Search().
				Index(e.Index).        // 设置索引名
				Query(boolQuery).      // 设置查询条件
				Sort("created", true). // 设置排序字段，根据Created字段升序排序，第二个参数false表示逆序
				From(0).               // 设置分页参数 - 起始偏移量，从第0行记录开始
				Size(10).              // 设置分页参数 - 每页大小
				Do(ctx)                // 执行请求

			if err != nil {
				// Handle error
				panic(err)
			}
			fmt.Printf("查询消耗时间 %d ms, 结果总数: %d\n", searchResult.TookInMillis, searchResult.TotalHits())
			ShowSelect(searchResult)

		}

		if e.Search == "select_groupby" {

			// 创建Avg指标聚合 可综合
			avgs := elastic.NewAvgAggregation().
				Field("retweets") // 设置统计字段

			// 创建Terms桶聚合
			aggs := elastic.NewTermsAggregation().
				Field("user") // 根据shop_id字段值，对数据进行分组

			aggs.SubAggregation("zonghe", avgs) //分组后求平均

			// 创建Histogram桶聚合
			aHis := elastic.NewHistogramAggregation().
				Field("retweets"). // 根据retweets字段值，对数据进行分组
				Interval(10)       //  分桶的间隔为10，意思就是price字段值按10间隔分组 每10 一组 用在特殊的地方

			atime := elastic.NewDateHistogramAggregation().
				Field("created"). // 根据date字段值，对数据进行分组
				//  分组间隔：month代表每月、支持minute（每分钟）、hour（每小时）、day（每天）、week（每周）、year（每年)
				CalendarInterval("month").
				// 设置返回结果中桶key的时间格式
				Format("yyyy-MM-dd")

				// 创Range桶聚合
			arange := elastic.NewRangeAggregation().
				Field("retweets").   // 根据price字段分桶
				AddUnboundedFrom(5). // 范围配置, 0 - 100
				AddRange(5.0, 10.0). // 范围配置, 100 - 200
				AddUnboundedTo(20.0) // 范围配置，> 200的值

			searchResult, err := model.ElaClient.Search().
				Index(e.Index).                         // 设置索引名
				Query(elastic.NewMatchAllQuery()).      // 设置查询条件
				Aggregation("group_by_user", aggs).     //   groupby
				Aggregation("group_by_retweets", aHis). //   groupby
				Aggregation("group_by_time", atime).    //   groupby
				Aggregation("group_by_range", arange).  //   groupby
				Size(0).                                //  设置分页参数 - 每页大小,设置为0代表不返回搜索结果，仅返回聚合分析结果
				Do(ctx)                                 // 执行请求

			if err != nil {
				// Handle error
				panic(err)
			}
			fmt.Printf("查询消耗时间 %d ms, 结果总数: %d\n", searchResult.TookInMillis, searchResult.TotalHits())

			// 使用Terms函数和前面定义的聚合条件名称，查询结果
			agg, found := searchResult.Aggregations.Terms("group_by_user")
			if !found {
				log.Fatal("没有找到聚合数据")
			}

			// 遍历桶数据
			for _, bucket := range agg.Buckets {

				// // 使用Terms函数和前面定义的聚合条件名称，查询结果
				// aggs, found := bucket.Aggregations.Terms("zonghe")
				aggs, found := bucket.AvgBucket("zonghe")

				if found {
					fmt.Printf("bucket %q 文档总数 %d 平均值 in %v\n", bucket.Key, bucket.DocCount, *aggs.Value)
				}

				// // 每一个桶都有一个key值，其实就是分组的值，可以理解为SQL的group by值
				// bucketValue := bucket.Key
				// // 打印结果， 默认桶聚合查询，都是统计文档总数
				// fmt.Printf("bucket = %q 文档总数 = %d\n", bucketValue, bucket.DocCount)
			}

			// 使用Terms函数和前面定义的聚合条件名称，查询结果
			agg1, found1 := searchResult.Aggregations.Histogram("group_by_retweets")
			if !found1 {
				log.Fatal("没有找到聚合数据")
			}

			// 遍历桶数据
			for _, bucket := range agg1.Buckets {
				// 每一个桶都有一个key值，其实就是分组的值，可以理解为SQL的group by值
				bucketValue := bucket.Key
				// 打印结果， 默认桶聚合查询，都是统计文档总数
				fmt.Printf("bucket = %q 文档总数 = %d\n", bucketValue, bucket.DocCount)
			}

			// 使用Terms函数和前面定义的聚合条件名称，查询结果
			agg2, found2 := searchResult.Aggregations.DateHistogram("group_by_time")
			if !found2 {
				log.Fatal("没有找到聚合数据")
			}

			// 遍历桶数据
			for _, bucket := range agg2.Buckets {
				// 每一个桶都有一个key值，其实就是分组的值，可以理解为SQL的group by值
				bucketValue := bucket.Key
				// 打印结果， 默认桶聚合查询，都是统计文档总数
				fmt.Printf("bucket = %q 文档总数 = %d\n", bucketValue, bucket.DocCount)
			}

			// 使用Terms函数和前面定义的聚合条件名称，查询结果
			agg3, found3 := searchResult.Aggregations.DateHistogram("group_by_range")
			if !found3 {
				log.Fatal("没有找到聚合数据")
			}

			// 遍历桶数据
			for _, bucket := range agg3.Buckets {
				// 每一个桶都有一个key值，其实就是分组的值，可以理解为SQL的group by值
				bucketValue := bucket.Key
				// 打印结果， 默认桶聚合查询，都是统计文档总数
				fmt.Printf("bucket = %q 文档总数 = %d\n", bucketValue, bucket.DocCount)
			}

		}

		if e.Search == "select" { //查詢单个条件  where user = olivere  精确查询

			// 创建term查询条件，用于精确查询
			termQuery := elastic.NewTermQuery("user", "olivere0")
			searchResult, err := model.ElaClient.Search().
				Index(e.Index).        // 设置索引名
				Query(termQuery).      // 设置查询条件
				Sort("created", true). // 设置排序字段，根据Created字段升序排序，第二个参数false表示逆序
				From(0).               // 设置分页参数 - 起始偏移量，从第0行记录开始
				Size(10).              // 设置分页参数 - 每页大小
				Pretty(true).          // 查询结果返回可读性较好的JSON格式
				Do(ctx)                // 执行请求

			if err != nil {
				// Handle error
				panic(err)
			}

			//  searchResult.TookInMillis  == 上面拿 count 的方法
			fmt.Printf("查询消耗时间 %d ms, 结果总数: %d\n", searchResult.TookInMillis, searchResult.TotalHits())
			ShowSelect(searchResult)
		}

		if e.Search == "select_count" {

			// // // 创建Value Count指标聚合
			// aggs := elastic.NewValueCountAggregation().
			// 	Field("user") // 设置统字段

			// 创建Value Count指标聚合  **这种方法会去重 , 只有user 值不一样的时候可以记录
			aggs := elastic.NewCardinalityAggregation().
				Field("user") // 设置统计字段

			// 创建Avg指标聚合
			avgs := elastic.NewAvgAggregation().
				Field("retweets") // 设置统计字段

			amax := elastic.NewMaxAggregation().
				Field("retweets") // 设置统计字段

			// 创建Min指标聚合
			amin := elastic.NewMinAggregation().
				Field("retweets") // 设置统计字段

			// 创建Sum指标聚合
			asum := elastic.NewSumAggregation().
				Field("retweets") // 设置统计字段

			searchResult, err := model.ElaClient.Search().
				Index(e.Index).                    // 设置索引名
				Query(elastic.NewMatchAllQuery()). // 设置查询条件
				Aggregation("total", aggs).        //求 count  自己命名
				Aggregation("avg", avgs).          //求平均值
				Aggregation("total_price", asum).  //求合
				Aggregation("max", amax).          //最大值
				Aggregation("min", amin).          //最大值
				Size(0).                           // 设置分页参数 - 每页大小,设置为0代表不返回搜索结果，仅返回聚合分析结果
				Do(ctx)                            // 执行请求

			if err != nil {
				// Handle error
				panic(err)
			}

			// // // 使用ValueCount函数和前面定义的聚合条件名称，查询结果
			// agg, found := searchResult.Aggregations.ValueCount("total")
			// if found {
			// 	// 打印结果，注意：这里使用的是取值运算符
			// 	fmt.Println("id 总数", *agg.Value)
			// }

			// // 使用ValueCount函数和前面定义的聚合条件名称，查询结果
			agg, found := searchResult.Aggregations.Cardinality("total")
			if found {
				// 打印结果，注意：这里使用的是取值运算符
				fmt.Println("id 总数", *agg.Value)
			}

			// // 平均值
			agg, found = searchResult.Aggregations.Avg("avg")
			if found {
				// 打印结果，注意：这里使用的是取值运算符
				fmt.Println("retweets 平均值", *agg.Value)
			}

			// 平均值
			agg, found = searchResult.Aggregations.Sum("total_price")
			if found {
				// 打印结果，注意：这里使用的是取值运算符
				fmt.Println("retweets 求和", *agg.Value)
			}

			// 最大值
			agg, found = searchResult.Aggregations.Max("max")
			if found {
				// 打印结果，注意：这里使用的是取值运算符
				fmt.Println("retweets 最大值", *agg.Value)
			}

			// 最小值
			agg, found = searchResult.Aggregations.Max("min")
			if found {
				// 打印结果，注意：这里使用的是取值运算符
				fmt.Println("retweets 最小值", *agg.Value)
			}

			//  searchResult.TotalHits  == 上面拿 count 第一种的方法
			fmt.Printf("查询消耗时间 %d ms, 结果总数: %d\n", searchResult.TookInMillis, searchResult.TotalHits())
			ShowSelect(searchResult)
		}

	}

}

func ShowSelect(searchResult *elastic.SearchResult) {
	if searchResult.TotalHits() > 0 {
		// 查询结果不为空，则遍历结果
		var b1 model.Weibo
		sum := 0
		// 通过Each方法，将es结果的json结构转换成struct对象
		for _, item := range searchResult.Each(reflect.TypeOf(b1)) {
			// 转换成Article对象
			if t, ok := item.(model.Weibo); ok {
				fmt.Println(t.User)
				sum = sum + t.Retweets
			}
		}
		fmt.Println(sum)
	}
}

func (e ElaById) Update() {

	ctx := context.Background()

	//多条件查询和修改
	// _, err := model.ElaClient.UpdateByQuery(e.Index).
	// 	// 设置查询条件，这里设置Author=tizi
	// 	Query(elastic.NewMatchQuery("id", e.Id)).
	// 	// 通过脚本更新内容，将Title字段改为1111111
	// 	Script(elastic.NewScriptInline("ctx._source.retweets='124'")).
	// 	// 如果文档版本冲突继续执行
	// 	ProceedOnVersionConflict().
	// 	Do(ctx)

	_, err := model.ElaClient.Update().
		Index(e.Index).                                                          // 设置索引名
		Id(e.Id).                                                                // 文档id
		Doc(map[string]interface{}{"retweets": 12, "message": "1231233213123"}). // 更新retweets=0，支持传入键值结构
		Do(ctx)                                                                  // 执行ES查询
	if err != nil {
		// Handle error
		panic(err)
	}

}

//可删除数据或者索引
func (e ElaById) Delete() {

	ctx := context.Background()

	//result, err := model.ElaClient.DeleteByQuery(e.Index).Query(elastic.NewTermQuery("location", "here")).Do(ctx) // 设置索引名
	// 设置查询条件为: Author = tizi

	// 文档冲突也继续删除

	// 根据id删除一条数据
	_, err := model.ElaClient.Delete().
		Index(e.Index).
		Id(e.Id).
		Do(ctx)
	if err != nil {
		// Handle error
		panic(err)
	}

	//删除索引
	//model.ElaClient.DeleteIndex("blog").Do(ctx)

	//fmt.Println(result,err)
}
