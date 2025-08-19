package pkg

import (
	"fmt"
	"github.com/smartwalle/alipay/v3"
)

type AliPay interface {
	Pay(name, ordersyn, price string) string
}

type AliPat struct {
	PrivateKey string
	AppId      string "9021000142691056"
	NotifyURL  string "http://10f1b3b.r18.cpolar.top"
	ReturnURL  string "http://localhost:9200/#/status/online"
}

func NewAlipay() *AliPat {
	return &AliPat{
		PrivateKey: "MIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQCDReeiLQuHuZoJCMNJKyJDn6fiYjiPzR+ObC9HcEl35zyMobzlc9/o35Fu67PPbnVA/qf1EIPQMQCRJHIb2FAHomy6xSaQmQgeou9nQYXJi4+lV/fd+hKx2r9hBHFV+r9wj7VtH6pe7aogPeSVH7c1gNzKCmOaYicaoto7n/wvUjoG2Xb1iIAN41EOv7aNxpsDQEu8v35BzlBG9QxkhN/nFhagofv3gbziFuxCxkGh02+GpDNBZEpQ51TbnTgM+XI98NP/C8C/swx0T4L5V78WnM4iQVVG8Un9C/y0TZI0G4RpybKSJTL1imEWEFdQ5pP5hLaLGe3kEu/8tUj372nvAgMBAAECggEAeV209SRuR5KaTgXy/v/JEvpV7iBfKNw9c6Cj6Ylv81Ivkdmq2fzSo+um+wUh3gLCl0+ZmyIkq+PSIV6vvPy1qQlLOzYxHHjPn/uGGVHjzl2gdf6ppGmixRp328uOuS41tmy3EOe4zUEwC9UNM6355ZZ5CramNcLPYnn0CDI2pswm5gb74bVr2zH4p8AAjn/ziub+yQP9oflwe38bxCZA8SIgKc3EnHme9YrSQzcE2hA+cEeraFNaVPmaoRpHlv9hrBv2XYHspPDkEoOrLHuz8T3KYFQCgX8mIUU0H5ABmotMzkUkInJDqGogSxDfq9LfWiInbQznpxyDCfja4cgrcQKBgQDdUrWPgkTT7MbkFejv1dVYlheCgKl1zkp5QUyB4i0L9JH7iLI9V0Bcmo2LU/E2JZ0IOHgqp4aDlOTnaMpW7v0F4D2wjMu7+/6pmw06CW1RHrLiNIliUe8Vbxjvb33mIBd/1qyZY7hWIcek8JM/DmFw3ZeVUjhI30kD7Q+qmDn2xwKBgQCX10fOAzTPfHYpgfKQJ3O3Vp/2n9cnMEkjroKOhq2KhumEuWn6UQSPE02jI6pO5lY/9aCOvOoDdJf++AxkXOfU8T+defa68eln/jmYZaSfGgyPZ1oU8gkfjvZHMfmzyEg2ghQX06x4ZRc0ELJMNluq47W5bcLAlktoFd3eCuyrmQKBgHjcqvOkswt8ORzjbiJTBr/qrwoAUwpXTkorQ6mGJw1pULMo8hUXabloWTWl8IqePcP4en7on0eJ7vsEOcMBncEwtN+N89BDi815M97muQDcn9lx0TiU/9gXpGlU3E7oZYzVYcoeL9MxfUHwcqizpLiq+hV8IeqBrPDs62pA2R6PAoGAZkdD/dgZDZ4ntePrcQCyKI/4JlcaxTH4Qkm9daXwZbxktdtzMIK8UGfPRxGyGX18IY8hvKQx+WvjKoMJTd1q5/wtPrU15k0nTL5pK2dkESupxDr46dzzGkfhSqm2KkzBn73VLQuPRHXLwG1yvHVtH9pMwX/WzIUwfAcmsPeo3hkCgYEAh2C6Au2iepPJWGZllGsX1cToX+wBpyQDZ0ZVDVWPgy0oNVjSYoIgbdPf9mSZupY1lsX+Exw6zRx0CT1gblx5kCT+GfjeuDBH4+qXVLZKDbcq61O6hTU7w6b0MYXFfbSNy0vnUF9XchNqb9EpJEcbVOsVTXj9tw3GEQPJrq3L0gA=",
		AppId:      "9021000142691056",
		NotifyURL:  "http://10f1b3b.r18.cpolar.top",
		ReturnURL:  "http://localhost:9200/#/status/online",
	}
}

func (a *AliPat) Pay(name, ordersyn, price string) string {
	var privateKey = a.PrivateKey // 必须，上一步中使用 RSA签名验签工具 生成的私钥
	var client, err = alipay.New(a.AppId, privateKey, false)
	var p = alipay.TradeWapPay{}
	p.NotifyURL = a.NotifyURL
	p.ReturnURL = a.ReturnURL
	p.Subject = name
	p.OutTradeNo = ordersyn
	p.TotalAmount = price
	p.ProductCode = "QUICK_WAP_WAY"

	url, err := client.TradeWapPay(p)
	if err != nil {
		fmt.Println(err)
	}

	// 这个 payURL 即是用于打开支付宝支付页面的 URL，可将输出的内容复制，到浏览器中访问该 URL 即可打开支付页面。
	var payURL = url.String()
	fmt.Println(payURL)
	return payURL
}
