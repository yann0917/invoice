# invoice

> WSDL（网络服务描述语言，Web Services Description Language）是一门基于 XML 的语言，用于描述 Web Services 以及如何对它们进行访问。
>
> SOAP（Simple Object Access Protocol ）简单对象访问协议是在分散或分布式的环境中交换信息的简单的协议，是一个基于XML的协议，它包括四个部分：SOAP封装(envelop)，封装定义了一个描述消息中的内容是什么，是谁发送的，谁应当接受并处理它以及如何处理它们的框架；SOAP编码规则（encoding rules），用于表示应用程序需要使用的数据类型的实例; SOAP RPC表示(RPC representation)，表示远程过程调用和应答的协定;SOAP绑定（binding），使用底层协议交换信息。

天津航信智税综合管理系统接口Go语言实现，实现的功能有：

* 发票数据上传
* 开具发票
* 打印发票

## 注意

网上关于 go 语言的 SOAP 协议很少，该仓库为对接时接口时做的记录，如果您对接了基于 SOAP 的遗留系统，希望该仓库对您有帮助。
该项目为 Demo，参数未从配置文件读取，部分开票数据固定写死的，如果您对接的是这个开票接口，请将开票参数修改成符合您项目要求的参数。

## References

* [Go与SOAP](https://tonybai.com/2019/01/08/go-and-soap/)
* [gowsdl](https://github.com/hooklift/gowsdl)

---
