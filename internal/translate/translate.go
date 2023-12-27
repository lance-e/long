package translate

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type responseResult struct {
	From         string   `json:"from"`
	To           string   `json:"to"`
	TransResult  []result `json:"trans_result"`
	ErrorCode    int      `json:"error_code"`
	ErrorMessage string   `json:"error_msg"`
}
type result struct {
	Src string `json:"src"`
	Dst string `json:"dst"`
}
type errorResult struct {
	ErrorCode    int    `json:"error_code"`
	ErrorMessage string `json:"error_msg"`
}

func Translate(query string, to string) {
	basicUrl := "https://fanyi-api.baidu.com/api/trans/vip/translate"
	from := "auto" //可指定从哪种语言翻译
	//先拼接sign
	appid := ""  //替换为自己的appid
	secret := "" //替换为自己的密钥
	salt := "random"
	s := strings.Join([]string{appid, query, salt, secret}, "")
	//md5加密
	sign := md5.Sum([]byte(s))
	finalSign := hex.EncodeToString(sign[:])
	//再进行url的拼接
	//http://api.fanyi.baidu.com/api/trans/vip/translate?q=apple&from=en&to=zh&appid=2015063000000001&salt=1435660288&sign=f89f9594663708c1605f3d736d01d2d4
	finalUrl := strings.Join([]string{basicUrl, "?", "q=", query, "&from=", from, "&to=", to, "&appid=", appid, "&salt=", salt, "&sign=", finalSign}, "")

	//开始进行网络请求
	req, err := http.NewRequest("GET", finalUrl, nil)
	if err != nil {
		fmt.Println("could not produce a http request")
		return
	}
	response, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("could not get the response ")
		return
	}
	defer response.Body.Close()
	info, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("could not get the body of the response")
		return
	}
	var result responseResult
	fmt.Println(string(info))
	err = json.Unmarshal(info, &result)
	if err != nil {
		fmt.Println("the format is wrong")
		return
	}
	if result.ErrorCode != 0 {
		fmt.Println("warning! there are some error happened")
		fmt.Printf("the error code is %d,the error message is %s \n", result.ErrorCode, result.ErrorMessage)
		return
	}
	fmt.Println("the type of the word at beginning: ", result.From)
	fmt.Println("the type of the word after translate: ", result.To)
	fmt.Println("the source word: ", result.TransResult[0].Src)
	fmt.Println("the target word: ", result.TransResult[0].Dst)

}
