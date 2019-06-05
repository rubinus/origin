package handlers

import (
	"context"
	//"fmt"
	"math/rand"
	"time"

	"git.zhugefang.com/gocore/zgo"
	"github.com/kataras/iris"
)

var answers = []string{
	"It is certain",
	"It is decidedly so",
	"Without a doubt",
	"Yes definitely",
	"You may rely on it",
	"As I see it yes",
	"Most likely",
	"Outlook good",
	"Yes",
	"Signs point to yes",
	"Reply hazy try again",
	"Ask again later",
	"Better not tell you now",
	"Cannot predict now",
	"Concentrate and ask again",
	"Don't count on it",
	"My reply is no",
	"My sources say no",
	"Outlook not so good",
	"Very doubtful",
}

var cityareas = []string{
	"朝阳", "海淀", "东城", "西城", "丰台",
}

func createData() string {
	rand.Seed(time.Now().Unix())
	var cityarea_id, answer_id int
	cityarea_id = rand.Intn(len(cityareas))
	answer_id = rand.Intn(len(answers))
	var tmp = make(map[string]interface{}, 4)
	tmp["cityarea_name"] = cityareas[cityarea_id]
	tmp["cityarea_id"] = cityarea_id
	tmp["answer_id"] = answer_id
	tmp["answer"] = answers[answer_id]

	js, err := zgo.Utils.MarshalMap(tmp)
	if err != nil {
		panic(err)
	}
	return js
}

func AddDataV1(ctx iris.Context) {
	city := ctx.Params().GetStringDefault("city", "")
	if city == "" {
		zgo.Http.JsonpErr(ctx, "city is null")
		return
	}

	var err error

	index := "answer_bj" //+ city
	//fmt.Println("index", index, "type: type")

	data := createData()
	//fmt.Println(data)
	result, err := zgo.Es.AddOneData(context.TODO(), index, "type", data, "")

	if err != nil {
		zgo.Http.JsonpErr(ctx, "add es error: "+err.Error())
		return
	}
	zgo.Http.JsonpOK(ctx, result)
}

func EsQueryV1(ctx iris.Context) {
	city := ctx.Params().GetStringDefault("city", "")
	if city == "" {
		zgo.Http.JsonpErr(ctx, "city is null")
		return
	}

	var err error

	index := "answer_bj" // + city
	//fmt.Println("index", index, "type: type")

	dsl := zgo.Es.NewDsl()
	dsl.Must(dsl.TermField("cityarea_id", 3))
	dslstr := dsl.QueryDsl()

	//fmt.Println("dsl: ", dslstr)
	result, err := zgo.Es.SearchDsl(context.TODO(), index, "type", dslstr, nil)

	if err != nil {
		zgo.Http.JsonpErr(ctx, "get es error: "+err.Error())
		return
	}
	zgo.Http.JsonpOK(ctx, result)
}

func CityareaAggsV1(ctx iris.Context) {
	city := ctx.Params().GetStringDefault("city", "")
	if city == "" {
		zgo.Http.JsonpErr(ctx, "city is null")
		return
	}

	var err error

	index := "answer_bj" // + city
	//fmt.Println("index", index, "type: type")

	dsl := zgo.Es.NewDsl()
	dsl.SetAggs(dsl.SimpleAggs("cityarea_id", 10))
	dslstr := dsl.QueryDsl()

	//fmt.Println("dsl: ", dslstr)
	result, err := zgo.Es.SearchDsl(context.TODO(), index, "type", dslstr, nil)

	if err != nil {
		zgo.Http.JsonpErr(ctx, "get es error: "+err.Error())
		return
	}

	zgo.Http.JsonpOK(ctx, result)
	return
}
