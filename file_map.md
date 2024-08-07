# 项目文件结构

## API

- [3_1_weichatPay.go](api/3_1_weichatPay.go) - 微信支付相关代码
- [3_2_aliPay.go](api/3_2_aliPay.go) - 支付宝支付相关代码
- [common.go](api/common.go) - 公共功能实现
- [const.go](api/const.go) - 常量定义
- [file_smsc_upload.go](api/file_smsc_upload.go) - 短信上传文件相关功能
- [file_smsc_upload_test.go](api/file_smsc_upload_test.go) - 短信上传文件功能测试
- [notify.go](api/notify.go) - 通知相关功能
- [notify_test.go](api/notify_test.go) - 通知功能测试
- [pregate_trade_findBankNameAndBankCode.go](api/pregate_trade_findBankNameAndBankCode.go) - 查询银行名称和银行代码功能
- [pregate_trade_findBankNameAndBankCode_test.go](api/pregate_trade_findBankNameAndBankCode_test.go) - 查询银行名称和银行代码功能测试
- [pregate_trade_findCmmtAreaInfoList.go](api/pregate_trade_findCmmtAreaInfoList.go) - 查询地区信息列表功能
- [pregate_trade_findCmmtAreaInfoList_test.go](api/pregate_trade_findCmmtAreaInfoList_test.go) - 查询地区信息列表功能测试
- [report_scan_union_appidAddOrUpdate.go](api/report_scan_union_appidAddOrUpdate.go) - 报告扫描联合 appid 添加或更新功能
- [report_scan_union_appidAddOrUpdate_test.go](api/report_scan_union_appidAddOrUpdate_test.go) - 报告扫描联合 appid 添加或更新功能测试
- [report_scan_union_reportAdd.go](api/report_scan_union_reportAdd.go) - 报告扫描联合报告添加功能
- [report_scan_union_upload.go](api/report_scan_union_upload.go) - 报告扫描联合上传功能
- [t1_smsc_addCustInfoApply.go](api/t1_smsc_addCustInfoApply.go) - 添加客户信息申请功能
- [t1_smsc_addCustInfoApply_test.go](api/t1_smsc_addCustInfoApply_test.go) - 添加客户信息申请功能测试
- [t1_smsc_auditCustInfoApply.go](api/t1_smsc_auditCustInfoApply.go) - 审核客户信息申请功能
- [t1_smsc_auditCustInfoApply_test.go](api/t1_smsc_auditCustInfoApply_test.go) - 审核客户信息申请功能测试
- [t1_smsc_sign.go](api/t1_smsc_sign.go) - 短信签名功能
- [t1_smsc_sign_test.go](api/t1_smsc_sign_test.go) - 短信签名功能测试

## Build

- [all.cov](build/all.cov) - 总覆盖率报告
- [c.cov](build/c.cov) - 覆盖率报告
- [coverage.cov](build/coverage.cov) - 覆盖率报告
- [coverage.html](build/coverage.html) - HTML 格式的覆盖率报告
- [coverage.json](build/coverage.json) - JSON 格式的覆盖率报告
- [coverage.xml](build/coverage.xml) - XML 格式的覆盖率报告

## Example

- [notify/server.go](example/notify/server.go) - 通知示例代码
- [pay/server.go](example/pay/server.go) - 支付示例代码

## Temp

- [cert_id.txt](temp/cert_id.txt) - 证书 ID 文件
- [private_key.pem](temp/private_key.pem) - 私钥文件
- [public_key.pem](temp/public_key.pem) - 公钥文件

## TestData

- [conf.json](testdata/conf.json) - 测试配置文件
- [test.jpeg](testdata/test.jpeg) - 测试图片文件

## Utils

- [utils.go](utils/utils.go) - 工具函数实现
- [utils_test.go](utils/utils_test.go) - 工具函数测试

## 其他文件

- [file_map.md](file_map.md) - 文件映射说明
- [go.mod](go.mod) - Go 模块配置文件
- [go.sum](go.sum) - Go 依赖项校验和文件
- [sync.sh](sync.sh) - 同步脚本
- [test.sh](test.sh) - 测试脚本
