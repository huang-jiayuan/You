package pkg

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	gourl "net/url"
	"strings"
	"time"
)

type RealNames struct {
	ErrorCode int    `json:"error_code"`
	Reason    string `json:"reason"`
	Result    struct {
		Realname    string `json:"realname"`
		Idcard      string `json:"idcard"`
		Isok        bool   `json:"isok"`
		IdCardInfor struct {
			Province string `json:"province"`
			City     string `json:"city"`
			District string `json:"district"`
			Area     string `json:"area"`
			Sex      string `json:"sex"`
			Birthday string `json:"birthday"`
		} `json:"IdCardInfor"`
	} `json:"result"`
	Sn string `json:"sn"`
}

var Res RealNames

func calcAuthorization(source string, secretId string, secretKey string) (auth string, datetime string, err error) {
	timeLocation, _ := time.LoadLocation("Etc/GMT")
	datetime = time.Now().In(timeLocation).Format("Mon, 02 Jan 2006 15:04:05 GMT")
	signStr := fmt.Sprintf("x-date: %s\nx-source: %s", datetime, source)

	// hmac-sha1
	mac := hmac.New(sha1.New, []byte(secretKey))
	mac.Write([]byte(signStr))
	sign := base64.StdEncoding.EncodeToString(mac.Sum(nil))

	auth = fmt.Sprintf("hmac id=\"%s\", algorithm=\"hmac-sha1\", headers=\"x-date x-source\", signature=\"%s\"",
		secretId, sign)

	return auth, datetime, nil
}

func urlencode(params map[string]string) string {
	var p = gourl.Values{}
	for k, v := range params {
		p.Add(k, v)
	}
	return p.Encode()
}

func RealName(idCard, realName string) (bool, int, error) {
	// 云市场分配的密钥Id
	secretId := "AKIDctwu6w9j1dxrq69mVk5h5V6g79zhvr0qge9o"
	// 云市场分配的密钥Key
	secretKey := "5yu5wrbm1rjmZkbamkv045NQ3mde6xabtAhS2vVc"
	source := "market"

	// 签名
	auth, datetime, _ := calcAuthorization(source, secretId, secretKey)

	// 请求方法
	method := "POST"
	// 请求头
	headers := map[string]string{"X-Source": source, "X-Date": datetime, "Authorization": auth}

	// 查询参数
	queryParams := make(map[string]string)

	// body参数
	bodyParams := make(map[string]string)
	bodyParams["cardNo"] = idCard
	bodyParams["realName"] = realName
	// url参数拼接
	url := "https://service-18c38npd-1300755093.ap-beijing.apigateway.myqcloud.com/release/idcard/VerifyIdcardv2"
	if len(queryParams) > 0 {
		url = fmt.Sprintf("%s?%s", url, urlencode(queryParams))
	}

	bodyMethods := map[string]bool{"POST": true, "PUT": true, "PATCH": true}
	var body io.Reader = nil
	if bodyMethods[method] {
		body = strings.NewReader(urlencode(bodyParams))
		headers["Content-Type"] = "application/x-www-form-urlencoded"
	}

	client := &http.Client{
		Timeout: 5 * time.Second,
	}
	request, err := http.NewRequest(method, url, body)
	if err != nil {
		return false, 0, fmt.Errorf("创建HTTP请求失败: %v", err)
	}
	for k, v := range headers {
		request.Header.Set(k, v)
	}
	response, err := client.Do(request)
	if err != nil {
		return false, 0, fmt.Errorf("HTTP请求失败: %v", err)
	}
	defer response.Body.Close()

	bodyBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return false, 0, fmt.Errorf("读取响应失败: %v", err)
	}
	fmt.Println(string(bodyBytes))

	err = json.Unmarshal(bodyBytes, &Res)
	if err != nil {
		return false, 0, fmt.Errorf("解析JSON响应失败: %v", err)
	}
	if Res.ErrorCode == 0 && Res.Result.Isok {
		// 计算年龄 - 假设从身份证号提取，或从接口返回中获取
		age, err := calculateAgeFromIDCard(idCard)
		if err != nil {
			return true, 0, err
		}
		return true, age, nil
	}
	return false, 0, fmt.Errorf("实名认证失败: %s", Res.Reason)
}

// 从身份证号计算年龄
func calculateAgeFromIDCard(idCard string) (int, error) {
	if len(idCard) != 18 {
		return 0, fmt.Errorf("身份证号码格式不正确")
	}

	// 从身份证号中提取出生日期 (第7-14位)
	birthStr := idCard[6:14]
	birthDate, err := time.Parse("20060102", birthStr)
	if err != nil {
		return 0, err
	}

	// 计算年龄
	now := time.Now()
	age := now.Year() - birthDate.Year()

	// 检查生日是否已过
	if now.Month() < birthDate.Month() ||
		(now.Month() == birthDate.Month() && now.Day() < birthDate.Day()) {
		age--
	}

	return age, nil
}
