# Pandora-Box

一个简易的 Mihomo 桌面客户端

[下载APP](https://github.com/snakem982/Pandora-Box/releases/latest)

## 功能特点

- 支持 本地 HTTP/HTTPS/SOCKS 代理
- 支持 Vmess, Vless, Shadowsocks, Trojan, Tuic, Hysteria, Hysteria2, Wireguard, [Mieru](./Mieru.md) 协议
- 支持 分享链接, 订阅链接, Base64格式，Yaml格式 的数据输入解析 
- 内置将节点和订阅转换为 Mihomo 配置
- 支持 节点爬取，以及爬取后按国别和节点类型进行筛选
- 自动添加极简规则分组以及防DNS泄露配置
- 支持统一所有订阅的规则和分组
- 支持Tun模式

## 支持的系统平台

- Windows 10/11 AMD64/ARM64
- MacOS 10.13+ AMD64
- MacOS 11.0+ ARM64
- Linux AMD64/ARM64 [需要你自己从源代码构建](https://github.com/snakem982/Pandora-Box?tab=readme-ov-file#build)

## 使用手册

- [基本使用](Manual-CN.md)
- [统一所有订阅的规则和分组](UnifiedRuleGrouping.md)
- 自定义配置同 [Mihomo](https://wiki.metacubex.one/config/)

## 友情提示

- 若软件运行异常，尝试以管理员身份运行。
- 提示需要网络连接，点击允许。
- 如有疑问可留言。

## 界面预览

| Tab | 黑色主题                           | 白色主题                       |
|-----|--------------------------------|----------------------------|
| 通用  | ![General](img%2Fdark1.png)    | ![General](img%2F1.png)    |
| 节点  | ![Proxies](img%2Fdark2.png)    | ![Proxies](img%2F2.png)    |
| 配置  | ![Profiles](img%2Fdark3.png)   | ![Profiles](img%2F3.png)   |
| 连接  | ![Connection](img%2Fdark4.png) | ![Connection](img%2F4.png) |
