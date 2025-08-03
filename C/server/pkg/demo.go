package pkg

import (
	"encoding/json"
	"fmt"
	"github.com/Baidu-AIP/golang-sdk/aip/censor"
)

type TextCensorResponse struct {
	Conclusion     string      `json:"conclusion"`
	LogID          int64       `json:"log_id"`
	PhoneRisk      interface{} `json:"phoneRisk"`
	IsHitMd5       bool        `json:"isHitMd5"`
	ConclusionType uint        `json:"conclusionType"`
}

func Textcen(content string) bool {
	fmt.Printf("开始检测内容: '%s'\n", content)
	
	// 检查内容是否为空
	if content == "" {
		fmt.Println("内容为空，返回true（通过检测）")
		return true
	}
	
	client := censor.NewClient("0z8MYNewsrHDiHdHFflVsSpZ", "WYNTIJZSq9uhJnN6rMRXbEbh34AdVl3G")
	res := client.TextCensor(content)
	fmt.Printf("百度API原始响应: %s\n", res)

	var resp TextCensorResponse
	if err := json.Unmarshal([]byte(res), &resp); err != nil {
		fmt.Printf("解析JSON错误: %v\n", err)
		return false // 解析失败，认为不通过
	}

	fmt.Printf("解析结果: %+v\n", resp)
	
	// 百度内容审核API返回值说明：
	// conclusionType: 1-合规，2-不合规，3-疑似，4-审核失败
	// 只有conclusionType为1时才认为内容合规
	isClean := resp.ConclusionType == 1
	fmt.Printf("内容是否合规: %v (conclusionType: %d)\n", isClean, resp.ConclusionType)
	
	return isClean
}
