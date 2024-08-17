# 银盛支付 GO SDK 简介
    为了提高客户接入的便捷性，本系统提供 SDK 方式接入，使用本 SDK 将极大的简化开发者的工作，开发者将无需考虑通信、签名、验签等，只需要关注业务参数的输入。
    注意：本sdk目前仅支持银盛老版接口，接口文档地址：https://gateway-doc.ysepay.com/gatewayDocs/summary/N0000230/N0000231/I0000187.html
    未来会支持新接口，敬请期待。

## SDK 项目结构说明
- api -- SDK核心包, 包含通信, 加解签, 接口参数对象等
- example -- 接口调用/参数赋值演示demo
- testdata -- 测试数据

## SDK 接入说明 
以下两种方式任选其一
1. 直接在go.mod中引用(require ysepaysdk github.com/code-serenade/ysepaysdk/api)
2. 直接下载源码文件, 将ysepaysdk(SDK核心包)源码放入项目中


## SDK 使用说明
    接口命名直接根据接口method来命名, 方便用户使用, 需要使用某接口时, 通过接口method进行搜索, 找到对应的struct, demo等

1. 配置初始化
- 初始化为全局配置(多商户模式下, 可初始化多份)
- 第一种通过配置文件路径初始化
- 第二种通过配置内容直接初始化
```
sdk, _ := ysepaysdk.NewConfigFromFile("../../testdata/conf.json")

```
```
sdk = NewConfig(xmap.M{
		"cert_id":     "cert_id",
		"private_key": "private_key",
		"public_key":  "public_key",
	})

```

2. 组装请求参数
- 为了接口使用更加方便, 我们将参数粗分为必填/非必填, 必填直接函数的传入参数内, 非必填可自行组装
```
    // 设置必填参数
	param := NewPregateTradeFindCmmtAreaInfoListParam("北京")
    // 设置非必填
	param.PageNumber = "1"
	param.PageSize = "1"
```

3. 发起API调用
```
resp, err := sdk.PregateTradeFindCmmtAreaInfoListRequest(param)

```
