# 银盛支付 GO SDK 简介
    为了提高客户接入的便捷性，本系统提供 SDK 方式接入，使用本 SDK 将极大的简化开发者的工作，开发者将无需考虑通信、签名、验签等，只需要关注业务参数的输入。
    注意：本sdk支持银盛新版接口，接口文档地址：https://open.ysepay.com/openDocs/summary/B00288/B00156/A00169/I000285.html
    目前对接了二维码支付接口，正在补充其余接口，敬请期待。

## SDK 项目结构说明
- ysepaysdk -- SDK核心包, 包含通信, 加解签, 接口参数对象等
- example -- 接口调用/参数赋值演示demo
- testdata -- 测试数据

## SDK 接入说明 
以下两种方式任选其一
1. 直接在go.mod中引用(require github.com/code-serenade/ysepaysdk)
2. 直接下载源码文件, 将ysepaysdk(SDK核心包)源码放入项目中


## SDK 使用说明
    接口命名直接根据接口method来命名, 方便用户使用, 需要使用某接口时, 通过接口method进行搜索, 找到对应的struct, demo等

1. 配置初始化
- 初始化为全局配置(多商户模式下, 可初始化多份)
- 第一种通过配置文件路径初始化
- 第二种通过配置内容直接初始化
```
sdk, _ := ysepaysdk.NewConfigFromFile("temp.json")

```
```
sdk = NewConfigFromItem("partnerID", "ysPublicKey", "merPrivateKey")

```

2. 组装请求参数
- 为了接口使用更加方便, 我们将参数粗分为必填/非必填, 必填直接函数的传入参数内, 非必填可自行组装
```
    // 设置必填参数
	param := NewBizContentOnlineQrcodepay(outTradeNo, subject, sellerID, businessCode, bankType, totalAmount)
    // 设置非必填
	param.Currency = "CNY"
```

3. 发起API调用
```
resp, err := sdk.OnlineQrcodepayRequest(param)

```
