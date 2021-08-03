package invoice

import (
	"fmt"
	"testing"
)

func TestSign(t *testing.T) {
	signStr := sign(InvoiceName, InvoicePwd, "A00004")

	t.Log(signStr)
}

func TestGetTax(t *testing.T) {
	tax := getTax(198.61, 0.06)
	t.Log(tax)
}
func TestInvoice(t *testing.T) {
	invoice := map[string]string{
		"kind":     "2", // 1-纸质，2-电子
		"type":     "2", // 0-专票,2-普票
		"order_id": "12345",
	}

	buyer := map[string]string{
		"name":  "测试有限公司",
		"type":  "01",
		"tax":   "91440101005CHUFC5B",
		"email": "386139859@qq.com",
	}
	goods := []map[string]float64{
		{"num": 2, "price": 100},
		{"num": 1, "price": 120}}
	data, _, err := Invoice.Apply(invoice, buyer, goods)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(data)
}

func TestInvoiceStatus(t *testing.T) {
	data, _, err := Invoice.Status("A00001")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(data)
}

func TestInvoicePrint(t *testing.T) {

	invoice := map[string]string{
		"type":     "2", // 0-专票,2-普票,41-卷票
		"order_id": "A00001",
		"code":     "012001800111",
		"number":   "91060377",
	}
	data, _, err := Invoice.Print(invoice)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(data)
}
