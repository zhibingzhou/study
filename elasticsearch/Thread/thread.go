package thread

import (
	"elastric/es"
)

func Run() {
	var job es.Job
	// job = es.ElaById{Id: "0", Index: "xiha", Msg: model.Weibo{User: "olivere3", Message: "打酱油的一天我是xiha", Image: "123450", Tags: []string{"abc", "efg"}, Location: "here", Created: time.Now()}}
	// job.Insert()
	// job = es.ElaById{Id: "1", Index: "weibo", Msg: model.Weibo{User: "olivere2", Message: "打酱油的一天我是Weibo", Image: "12345e6", Tags: []string{"abc", "efg"}, Location: "here", Created: time.Now()}}
	// job.Insert()
	// job = es.ElaById{Id: "2", Index: "weibo", Msg: model.Weibo{User: "olivere1", Message: "打酱油的一天", Image: "123q456", Tags: []string{"abc", "efg"}, Location: "here", Created: time.Now()}}
	// job.Insert()
	// job = es.ElaById{Id: "3", Index: "weibo", Msg: model.Weibo{User: "olivere4", Message: "打酱油的一天", Image: "1234w56", Tags: []string{"abc", "efg"}, Location: "here", Created: time.Now()}}
	// job.Insert()
	// for i := 4; i <= 50; i++ {
	// 	if i%2 == 0 {
	// 		job = es.ElaById{Id: strconv.Itoa(i), Index: "weibo", Msg: model.Weibo{User: fmt.Sprintf("olivere%d", 0), Message: fmt.Sprintf("打酱油的%d天", i), Image: "1234w56", Tags: []string{"abc", "efg"}, Location: "here", Retweets: rand.Intn(50), Created: time.Now()}}
	// 	} else {
	// 		job = es.ElaById{Id: strconv.Itoa(i), Index: "weibo", Msg: model.Weibo{User: fmt.Sprintf("olivere%d", 1), Message: fmt.Sprintf("打酱油的%d天", i), Image: "1234w56", Tags: []string{"abc", "efg"}, Location: "here", Retweets: rand.Intn(50), Created: time.Now()}}
	// 	}

	// 	job.Insert()
	// }
	// job = es.ElaById{Iditems: []string{"xiha", "weibo"}, Search: "id_group"}
	// job.Select()
	// job = es.ElaById{Id: "3", Index: "weibo", Msg: model.Weibo{User: "olivere7", Message: "打酱油的三天", Image: "1234w56", Tags: []string{"abc", "efg"}, Location: "here", Created: time.Now()}}
	// job.Update()
	job = es.ElaById{Id: "3", Index: "weibo", Search: "select_groupby"}
	job.Select()

}
