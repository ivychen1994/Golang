package main

import (
	"fmt"
	"net/http"
	"net/url"
	"io/ioutil"
	"encoding/json"
	//"strconv"\
	"time"
)

func main() {
	ticker := time.NewTicker(time.Minute);
	for {
		/*if num > 5 {
			//大于5次关闭打点器
			ticker.Stop();
			break;
		}*/
		//否则从打点器中获取chan
		select {
		case <-ticker.C:
			request()
		}
	}

}

func request() {
	URL := "http://api.k780.com";

	//初始化参数
	param := url.Values{}

	//配置请求参数
	param.Set("app", "finance.rate_cnyquot")
	param.Set("curno", "USD") //AED阿联酋迪拉姆,AUD澳大利亚元,BRL巴西雷亚尔,CAD加拿大元,CHF瑞士法郎,DKK丹麦克朗,EUR欧元,GBP英镑,HKD港币,IDR印度尼西亚盾,INR印度卢比,JPY日元,KRW韩国元,KZT哈萨克斯坦坚戈,MOP澳门币,MYR马来西亚林吉特,NOK挪威克朗,NZD新西兰元,PHP菲律宾比索,RUB俄罗斯卢布,SEK瑞典克朗,SGD新加坡元,THB泰铢,TJS塔吉克斯坦索莫尼,TWD台币,USD美元,ZAR南非兰特,ASR沙特里亚尔
	param.Set("appkey", "10003")
	param.Set("sign", "b59bc3ef6191eb9f747dd4e83c99f2a4")
	param.Set("format", "json")

	//发送请求
	data, err := Get(URL, param)
	if err != nil {
		fmt.Errorf("请求失败,错误信息:\r\n%v", err)
	} else {
		var netReturn map[string]interface{}
		json.Unmarshal(data, &netReturn)
		fmt.Println("美元对应的人名币外汇")
		if netReturn["success"].(string) == "1" {
			if v, ok := netReturn["result"]; ok {
				result := v.(map[string]interface{})
				fmt.Printf("%-15s%-15s%-15s%-15s%-15s%-15s%-18s%-15s\n", "银行简称", "银行名称", "现汇卖出价", "现汇买入价", "现钞卖出价", "现钞买入价", "中间价", "数据更新时间")
				resultMap := result["USD"].(map[string]interface{})
				BOCMap := resultMap["BOC"].(map[string]interface{})
				print(BOCMap)
				CCBMap := resultMap["CCB"].(map[string]interface{})
				print(CCBMap)
				ICBCMap := resultMap["ICBC"].(map[string]interface{})
				print(ICBCMap)
				ABCMap := resultMap["ABC"].(map[string]interface{})
				print(ABCMap)
				CEBMap := resultMap["CEB"].(map[string]interface{})
				print(CEBMap)
			}
		}else{
			fmt.Println( netReturn["msg"].(string))
		}
	}
}

//get网络请求
func Get(apiURL string, params url.Values) (rs []byte, err error) {
	var Url *url.URL
	Url, err = url.Parse(apiURL)
	if err != nil {
		fmt.Printf("解析url错误:\r\n%v", err)
		return nil, err
	}
	//如果参数中有中文参数,这个方法会进行URLEncode
	Url.RequestURI()
	Url.RawQuery = params.Encode()
	resp, err := http.Get(Url.String())
	if err != nil {
		fmt.Println("err:", err)
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

//打印
func print(bankmap map[string]interface{}) {
	fmt.Printf("%-19s%-15s%-20s%-20s%-20s%-20s%-20s%-18s\n", bankmap["bankno"], bankmap["banknm"], bankmap["se_sell"], bankmap["se_buy"], bankmap["cn_sell"], bankmap["cn_buy"], bankmap["middle"], bankmap["upddate"])
}
