package main

import "fmt"

// CheckNumber 小切手番号
type CheckNumber int

// CardNumber カード番号
type CardNumber string

// CardType カードの種類(OR型)
type CardType int

// CardTypeの定数
// iotaは連番を生成する (コンパイル時に決まるっぽい)
// ↑のためDBに保存するときには使わない方がいい
// https://speakerdeck.com/uji/5fen-dewan-quan-li-jie-surugofalseiota?slide=21
const (
	Visa       CardType = iota // 0
	MasterCard                 // 1
)

// CreditCardInfo カード情報(構造体)
type CreditCardInfo struct {
	CardType   CardType
	CardNumber CardNumber
}

// PaymentMethod 支払い方法
type PaymentMethod interface{}
type (
	Cash  struct{}
	Check struct{ Number CheckNumber }
	Card  struct{ Info CreditCardInfo }
)

// PaymentAmount 支払い総額
type PaymentAmount float64

// Currency 通貨
type Currency int

const (
	USD Currency = iota
	EUR
)

// Payment 支払い
type Payment struct {
	Amount   PaymentAmount
	Currency Currency
	Method   PaymentMethod
}

// UnpaidInvoice 未払い請求書
type UnpaidInvoice struct{}

// PaidInvoice 支払い済み請求書
type PaidInvoice struct{}

// PayInvoice 支払い
// 未払い請求書と支払いを受け取り、支払い済み請求書を返す
type PayInvoice func(UnpaidInvoice, Payment) PaidInvoice

// ConvertPaymentCurrency 通貨変換
// 支払いと通貨を受け取り、通貨を変換した支払いを返す
type ConvertPaymentCurrency func(Payment, Currency) Payment

func main() {
	// Example usage
	payment := Payment{
		Amount:   100.0,
		Currency: USD,
		Method: Card{
			Info: CreditCardInfo{
				CardType:   Visa,
				CardNumber: "1234-5678-9012-3456",
			},
		},
	}

	fmt.Println(payment)
}
