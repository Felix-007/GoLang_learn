package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"time"
	"xorm.io/core"
	"xorm.io/xorm"
)

func main() {

	fmt.Println("跟单系统试试")
	// 注册一个默认的路由器
	router := gin.Default()
	// 最基本的用法
	router.POST("/test", func1)
	// 绑定端口是8888
	router.Run(":8888")

}

// func1: 处理最基本的POST
func func1(c *gin.Context) {
	// 回复一个200 OK, 在client的http-post的resp的body中获取数据
	//c.String(http.StatusOK, "test OK")
	/**
	[{
		"CarNum": "京A88888",
		"RouteNum": "123"
	}, {
		"CarNum": "京A88888",
		"RouteNum": "123"
	}]
	*/
	raw, err := c.GetRawData() //raw 写json字符串
	if err != nil {
		fmt.Printf("失败:请求json参数获取失败!error=%v\n", err)
	}
	var paras []Para
	json.Unmarshal(raw, &paras)
	if len(paras) == 0 {
		c.String(http.StatusAccepted, "json请求参数为:=", paras)
	}

	//查询数据库
	db := NewDb()
	var result []PositionInfo
	for _, v := range paras {
		var temp PositionInfo
		sql_list := ` SELECT t4.n_jd jd,t4.n_wd wd,t1.c_yldm yldm ,t1.c_ldlsh ldlsh,t1.c_sfjdm sfjdm,t6.c_zdmc sfjmc,t2.c_zdjdm zdjdm,t7.c_zdmc zdjmc,t1.c_pcdh pcdh,t4.d_wzsj wzsj
			 	FROM	yygk_yw.t_rw_cldqrw t1 
				LEFT JOIN yygk_core.t_jc_ylxx t2 ON t1.c_yldm = t2.c_yldm
				LEFT JOIN yygk_core.t_sb_sbxx t3 ON t1.c_cldm = t3.c_cldm
				LEFT JOIN yygk_core.t_sb_sbzxzt t4 ON t3.c_sbxlh = t4.c_sbxlh
				LEFT JOIN yygk_core.t_jc_clxx t5 ON t1.c_cldm = t5.c_cldm
				LEFT JOIN yygk_core.t_jc_zdxx t6 on t1.c_sfjdm=t6.c_zddm 
				LEFT JOIN yygk_core.t_jc_zdxx t7 on t1.c_zdjdm=t7.c_zddm 
				WHERE
				t5.c_cph = ?
				AND t1.c_yldm = ? `

		_, err := db.SQL(sql_list, v.CarNum, v.RouteNum).Get(&temp)
		if err != nil {
			fmt.Printf("DB查询失败:error:=%v\n", err)
			c.String(http.StatusServiceUnavailable, "Db查询连接建立失败error:=", err)
		}
		//时间格式化
		formatTime("20060102150405", &temp)

		result = append(result, temp)
	}

	s, err := json.Marshal(result)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("返回的json字符串", string(s))
	c.String(http.StatusOK, string(s))
}

func NewDb() (engine *xorm.Engine) {
	Url := "username:pssword@(host:port)/db?parseTime=true"

	ShowSQL := true
	Debug := true
	MaxIdle := 10
	MaxOpen := 10
	engine, err := xorm.NewEngine("mysql", Url)
	if err != nil {
		fmt.Printf("数据库 Url(%s) error: %v\n", Url, err)
		panic(err)
	}
	engine.ShowSQL(ShowSQL)
	if Debug {
		engine.Logger().SetLevel(core.LOG_DEBUG)
	}
	engine.SetMaxIdleConns(MaxIdle)
	engine.SetMaxOpenConns(MaxOpen)
	return
}

type Para struct {
	CarNum   string
	RouteNum string
}
type PositionInfo struct {
	Longitude           float64 `xorm:"jd" ` //经度
	Latitude            float64 `xorm:"wd"`  //纬度
	Route_code          string  `xorm:"yldm"`
	Waybill_num         string  `xorm:"ldlsh"` //路单流水号
	Depart_station_code string  `xorm:"sfjdm"` //出发站点代码
	Depart_station_name string  `xorm:"sfjmc"` //出发站点名称
	Arrive_station_code string  `xorm:"zdjdm"` //到达站点代码
	Arrive_station_name string  `xorm:"zdjmc"` // 到达站点名称
	Dispatch_num        string  `xorm:"pcdh"`  // 派车单号
	Gather_time         string  `xorm:"wzsj"`  // 数据采集时间 格式：寄递平台YYYYMMDDHH24MISS
}

func formatTime(outFormat string, p *PositionInfo) {
	timeTemplate := "2006-01-02T15:04:05Z"
	t1, _ := time.ParseInLocation(timeTemplate, p.Gather_time, time.Local)
	if !t1.IsZero() {
		p.Gather_time = t1.Format(outFormat)
	}
}
