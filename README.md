Clash规则预处理工具
# 功能
* 无缝对接Clash订阅配置，支持Clash全局代理
* 支持PC、Mac、移动设备
* 支持token认证
* 只支持socks5代理
# 架构
* 办公软件 <-> 本地clash客户端 <-> ClashParser服务 <-> Clash订阅服务
* ClashParser会根据传递参数，自动从Clash订阅服务获取配置，然后在规则中添加一系列办公软件的配置，最后返回给本地clash客户端，解决手动修改Clash订阅配置没办法持久化的问题。
# 使用方法
* 编译clashparser部署到vps上
* 使用 https://${IP}:${PORT}/v0.1/parser/${token}/${base64_clash_subscribe_link}/${OS} 作为Clash订阅链接
# 参数说明
* IP: 服务器IP,字符串
* PORT: 服务器端口，int
* token: 认证token，写在配置文件里，字符串
* base64_clash_subscribe_link: Base64编码后的Clash原始订阅链接，字符串
* OS: 操作系统，枚举型字符串，win/mac/ios