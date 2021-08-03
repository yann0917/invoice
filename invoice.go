package invoice

import (
	"bytes"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"encoding/xml"
	"log"
	"math"
	"strings"

	"github.com/yann0917/invoice/soap"
)

var (
	// URL 开发票请求地址
	URL = "http://121.36.219.110:9081/Ninvoicejs/service/invoice?wsdl"
	// InvoiceName 请求用户
	InvoiceName = "1201010000022260"
	// InvoicePwd 请求密码
	InvoicePwd = "123456"
)

// Client request invoice
type Client struct {
	Client   *soap.Client
	Name     string
	Password string
}

// NewClient new invoice Client
func NewClient(url, name, pwd string) *Client {
	return &Client{
		Client:   soap.NewClient(url),
		Name:     name,
		Password: pwd,
	}
}

// Invoice soap协议解析客户端
var Invoice *Client

func init() {
	Invoice = NewClient(URL, InvoiceName, InvoicePwd)
}

// Body 上传发票数据 body
type Body struct {
	XMLName xml.Name `xml:"http://webservice.cn.com/ invoice"`
	Arg0    string   `xml:"arg0"`
}

// StatusBody 开发票数据 body
type StatusBody struct {
	XMLName xml.Name `xml:"http://webservice.cn.com/ invoiceStatus"`
	Arg0    string   `xml:"arg0"`
}

// Response 上传发票数据resp
type Response struct {
	XMLName xml.Name `xml:"http://webservice.cn.com/ invoiceResponse"`
	Return  string   `xml:"return"`
}

// StatusResponse 开发票 resp
type StatusResponse struct {
	XMLName xml.Name `xml:"http://webservice.cn.com/ invoiceStatusResponse"`
	Return  string   `xml:"return"`
}

// PrintResponse 打印发票 resp
type PrintResponse struct {
	XMLName xml.Name `xml:"http://webservice.cn.com/ invoicePrintResponse"`
	Return  string   `xml:"return"`
}

// PrintBody 打印发票 body
type PrintBody struct {
	XMLName xml.Name `xml:"http://webservice.cn.com/ invoicePrint"`
	Arg0    string   `xml:"arg0"`
}

// Interface 节点解析
type Interface struct {
	XMLName xml.Name `xml:"interface"`
	Return  struct {
		ReturnCode    string `xml:"returnCode"`
		ReturnMessage string `xml:"returnMessage"`
	} `xml:"return"`
}

// StatusInterface 节点解析 开发票 resp
type StatusInterface struct {
	XMLName xml.Name `xml:"interface"`
	Return  struct {
		ReturnCode    string                `xml:"returnCode"`
		ReturnMessage string                `xml:"returnMessage"`
		Fpxx          invoiceStatusRespFpxx `xml:"fpxx"`
	} `xml:"return"`
}

// invoiceStatusRespFpxx 开发票返回开票信息
type invoiceStatusRespFpxx struct {
	Djbh string `xml:"djbh"` // 单据编号
	Fpdm string `xml:"fpdm"` // 发票代码
	Fphm string `xml:"fphm"` // 发票号码
	Kprq string `xml:"kprq"` // 开票日期
	URL  string `xml:"url"`  // 下载地址
}

// Sfxx Sfxx
type Sfxx struct {
	Wsname string `xml:"wsname"`
	WsPwd  string `xml:"wspwd"`
}

// Send 开票数据
type Send struct {
	XMLName xml.Name `xml:"interface"`
	Sfxx    Sfxx     `xml:"sfxx"`
	Fpxx    Fpxx     `xml:"fpxx"`
	Fpmxs   []Fpmx   `xml:"fpmxs>fpmx"`
}

// StatusSend 请求发票报文
type StatusSend struct {
	XMLName xml.Name `xml:"interface"`
	Sfxx    Sfxx     `xml:"sfxx"`
	Fpxx    Fpxx     `xml:"fpxx"`
}

// PrintSend 请求打印发票报文
type PrintSend struct {
	XMLName xml.Name  `xml:"interface"`
	Sfxx    Sfxx      `xml:"sfxx"`
	Fpxx    printFpxx `xml:"fpxx"`
}

type printFpxx struct {
	Djbh string `xml:"djbh"` // 单据编号
	Fpdm string `xml:"fpdm"` // 发票代码
	Fphm string `xml:"fphm"` // 发票号码
	Dylx string `xml:"dylx"` // 打印类型 0-无预览 1-有预览
	Fplx string `xml:"fplx"` // 发票类型 0-专票 2-普票 41-卷票
	Dyzl string `xml:"dyzl"` // 打印种类 0发票 1清单
}

// Fpxx 发票信息
type Fpxx struct {
	Djbh   string `xml:"djbh"`   // 单据编号
	Fpzl   string `xml:"fpzl"`   // 发票种类 1-纸质，2-电子
	Fplx   string `xml:"fplx"`   // 发票类型 0-专票，2-普票
	Fpzf   string `xml:"fpzf"`   // 发票正负 1-正票，2-红票
	Xfmc   string `xml:"xfmc"`   // 销方名称
	Xfsh   string `xml:"xfsh"`   // 销方税号
	Xfkpjh string `xml:"xfkpjh"` // 销方开票机号 **
	Xfdz   string `xml:"xfdz"`   // 销方地址
	Xfdh   string `xml:"xfdh"`   // 销方电话
	Xfyhzh string `xml:"xfyhzh"` // 销方引号账号
	Gflx   string `xml:"gflx"`   // 购方企业类型 01企业、02机关事业单位、03个人、04其它
	Gfmc   string `xml:"gfmc"`   // 购方名称
	Gfsh   string `xml:"gfsh"`   // 购方税号
	Gfdz   string `xml:"gfdz"`   // 购方地址
	Gfdh   string `xml:"gfdh"`   // 购方电话
	Gfyhzh string `xml:"gfyhzh"` // 购方银行账号
	Gfsj   string `xml:"gfsj"`   // 购方手机
	Gfyx   string `xml:"gfyx"`   // 购方邮箱
	Tsfs   string `xml:"tsfs"`   // 推送方式 0-不推送，1-手机，2-邮箱，3-手机+邮箱
	Yysbz  string `xml:"yysbz"`  // 营业税标识
	Kce    string `xml:"kce"`    // 扣除额
	Byzd1  string `xml:"byzd1"`  // 备用字段1
	Byzd2  string `xml:"byzd2"`  // 备用字段2
	Bz     string `xml:"bz"`     // 备注
	Kpr    string `xml:"kpr"`    // 开票人
	Skr    string `xml:"skr"`    // 收款人
	Fhr    string `xml:"fhr"`    // 复核人
	Hcfpdm string `xml:"hcfpdm"` // 正数发票代码
	Hcfphm string `xml:"hcfphm"` // 正数发票号码
	Hztzdh string `xml:"hztzdh"` // 红字通知单号
	Qdbz   string `xml:"qdbz"`   // 清单标识 0-非清单，1-清单
}

// Fpmx 发票商品明细
type Fpmx struct {
	Spmc    string  `xml:"spmc"`
	Spbm    string  `xml:"spbm"` // 必填 19位，位数不够后面补0
	Ggxh    string  `xml:"ggxh"`
	Dw      string  `xml:"dw"`
	Hsbz    string  `xml:"hsbz"`    // 必填，1含税 0 不含税
	Sl      float64 `xml:"sl"`      // 数量 开红票时为负数，折扣时为空。
	Dj      float64 `xml:"dj"`      // 单价 精确到小数点号后四位，折扣时为空。
	Je      float64 `xml:"je"`      // 金额 精确到小数点号后2位，开红票与折扣时为负数
	Slv     string  `xml:"slv"`     // 税率 0.06
	Se      float64 `xml:"se"`      // 税额 小数点后 2 位，以元为单位精确到分开，红票时为负数，折扣行时负数。
	Yhzcbs  string  `xml:"yhzcbs"`  // 优惠政策标识 0不使用，1 使用<SFYHZC>
	Lslbs   string  `xml:"lslbs"`   // 零税率标识,空：非零税率，0：出口零税，1：免税，2：不征税，3普通零税率
	Zzstsgl string  `xml:"zzstsgl"` // 增值税特殊管理 当yhzcbs是1时必填<YHZCNR>
	Zkhbz   string  `xml:"zkhbz"`   // 折扣行标识必填，0商品行，3被折扣行，4折扣行（当为折扣行时，上一行必须是被折扣行）
	Byzd1   string  `xml:"byzd1"`   // 备用字段1
	Byzd2   string  `xml:"byzd2"`   // 备用字段2
}

// sign 签名算法
//
// md5(用户 + djbh单据编号+ 平台密码)
func sign(name, pwd, djbh string) string {

	var buff bytes.Buffer

	buff.WriteString(name)
	buff.WriteString(djbh)
	buff.WriteString(pwd)

	s := md5.New()
	s.Write(buff.Bytes())

	return hex.EncodeToString(s.Sum(nil))
}

// getTax 税额
// 税额=(含税金额/（1+税率)*税率
// 含税金额即价税合计
func getTax(money, taxRate float64) float64 {
	tax := (money / (1 + taxRate)) * taxRate
	n10 := math.Pow10(2)
	taxx := math.Trunc((tax+0.5/n10)*n10) / n10

	return taxx
}

// base64Body 报文加密
func base64Body(params interface{}) (data string, err error) {
	// 结构体转 xml
	req, err := xml.Marshal(params)
	if err != nil {
		return
	}
	// log.Info("请求发票数据:" + string(req))
	// base64 加密
	data = base64.StdEncoding.EncodeToString(req)
	return
}

// Apply 申请发票
// invoice["kind"] = "Fpzl", invoice["type"] = "Fplx"
// buyer["name"] = "Gfmc" buyer["type"] = "Gflx" buyer["tax"] = "Gfsh" buyer["email"] = "email"
// [goods["num"] = "Sl",goods["price"] = "Dj"]
func (c *Client) Apply(invoice, buyer map[string]string, goods []map[string]float64) (resp Interface, raw string, err error) {
	var fp Send
	fp.Sfxx.Wsname = InvoiceName
	fp.Sfxx.WsPwd = sign(InvoiceName, InvoicePwd, invoice["order_id"])
	fp.Fpxx.Djbh = invoice["order_id"] // 单据编号
	fp.Fpxx.Fpzl = invoice["kind"]
	fp.Fpxx.Fplx = invoice["type"]
	if mark, ok := invoice["mark"]; ok {
		fp.Fpxx.Fpzf = mark
	} else {
		fp.Fpxx.Fpzf = "1"
	}

	if machine, ok := invoice["machine"]; ok {
		fp.Fpxx.Xfkpjh = machine
	}

	if code, ok := invoice["invoice_code"]; ok {
		fp.Fpxx.Hcfpdm = code
	}

	if code, ok := invoice["invoice_no"]; ok {
		fp.Fpxx.Hcfphm = code
	}

	fp.Fpxx.Xfmc = "天津汉威信恒文化传播有限公司"
	fp.Fpxx.Xfsh = "91120222300725265H"

	fp.Fpxx.Xfdz = "天津市武清区黄花店镇政府南路189号"
	fp.Fpxx.Xfdh = "022-22152133"
	fp.Fpxx.Xfyhzh = "招商银行股份有限公司天津武清支行 122904911810801"
	fp.Fpxx.Kpr = "陈燕艳"
	fp.Fpxx.Skr = "富娅娜"
	fp.Fpxx.Fhr = "朱明"

	// 购方
	name := strings.ReplaceAll(strings.ReplaceAll(buyer["name"], " ", ""), "\r\n", "")
	tax := strings.ReplaceAll(strings.ReplaceAll(strings.ToUpper(buyer["tax"]), " ", ""), "\r\n", "")
	fp.Fpxx.Gfmc = name
	fp.Fpxx.Gflx = buyer["type"]
	fp.Fpxx.Gfsh = tax
	fp.Fpxx.Gfyx = buyer["email"]
	fp.Fpxx.Tsfs = "2"

	// 商品清单

	var spbm = ""
	// 税收分类编码
	spbm = "3040304020000000000"
	for _, item := range goods {
		amount := item["num"] * item["price"]
		fp.Fpmxs = append(fp.Fpmxs, Fpmx{
			Spmc:   "票款",
			Spbm:   spbm,
			Hsbz:   "1",
			Slv:    "0.06",
			Yhzcbs: "0",
			Sl:     item["num"],
			Dj:     item["price"],
			Je:     amount,
			Se:     getTax(amount, 0.06),
			Zkhbz:  "0",
		})
	}

	// 构造 body
	var body Body
	body.Arg0, _ = base64Body(fp)
	reqBody, _ := xml.Marshal(body)
	respStr := &Response{}
	err = Invoice.Client.Call("", reqBody, respStr)
	if err != nil {
		// log.Info(err)
		return
	}

	// 返回值解析
	raw = respStr.Return
	// log.Info(invoice["order_id"] + ":上传发票resp:" + raw)
	xml.Unmarshal([]byte(raw), &resp)
	return
}

// Status 开票
func (c *Client) Status(orderID string) (resp StatusInterface, raw string, err error) {
	var fp StatusSend
	fp.Sfxx.Wsname = InvoiceName
	fp.Sfxx.WsPwd = sign(InvoiceName, InvoicePwd, orderID)
	fp.Fpxx.Djbh = orderID

	// 构造 body
	var body StatusBody
	body.Arg0, _ = base64Body(fp)
	reqBody, _ := xml.Marshal(body)

	respStr := &StatusResponse{}
	err = Invoice.Client.Call("", reqBody, respStr)
	if err != nil {
		log.Fatal(err)
		// log.Info(err)
		return
	}

	// 返回值解析
	raw = respStr.Return
	// log.Info(orderID + ":开发票resp:" + raw)
	xml.Unmarshal([]byte(raw), &resp)
	return
}

// Print 打印发票
func (c *Client) Print(param map[string]string) (resp Interface, raw string, err error) {
	var fp PrintSend
	fp.Sfxx.Wsname = InvoiceName
	fp.Sfxx.WsPwd = sign(InvoiceName, InvoicePwd, param["order_id"])
	fp.Fpxx.Djbh = param["order_id"]
	fp.Fpxx.Fpdm = param["code"]
	fp.Fpxx.Fphm = param["number"]
	fp.Fpxx.Fplx = param["type"]
	fp.Fpxx.Dylx = "0"
	fp.Fpxx.Dyzl = "0"

	// 构造 body
	var body PrintBody
	body.Arg0, _ = base64Body(fp)
	reqBody, _ := xml.Marshal(body)

	respStr := &PrintResponse{}
	err = Invoice.Client.Call("", reqBody, respStr)
	if err != nil {
		// log.Info(err)
		return
	}

	// 返回值解析
	// log.Info(respStr)
	raw = respStr.Return
	xml.Unmarshal([]byte(raw), &resp)
	return
}
