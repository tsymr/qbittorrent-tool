# qbittorrent-tool
qbittorrent 辅助工具

# 怎么使用

根据平台下载对应的 release 文件，复制 example.config.json 为 config.json，修改 qbittorrent 的配置信息以及自定义标签、分类映射。

使用`设置-下载-运行外部程序`或定时任务运行。

# 1. 根据tracker域名设置标签

根据`tracker`地址给种子设置标签，如 `https://tracker.baidu.com/announce.php` 则设置标签为 `baidu.com`。

如果需要自定义标签，需要在 config.json 配置好对应关系，已经存在的标签不会自动删除，可在 webui 右键删除，等待种子再次被处理。

# 2. 自动设置分类

根据种子保存路径设置分类，需要在 `config.json` 配置好对应关系，获取不到对应关系的种子不会被处理。

# 3. 自动控制做种时间、分享率、速度

扩展自带的`做种限制`功能，增加根据标签、分类、tracker、关键字精确限制停止做种、删种、例外继续保种等。

可配置多组规则，按顺序处理，后面的会覆盖前面的，一组规则里的所有规则都不生效时整租规则视为无效。

对于不启用的规则，可以将字段删除，以保持规则精炼和高效识别，详情参考配置文件和[Wiki](https://github.com/fengqi/qbittorrent-tool/wiki/%E8%87%AA%E5%8A%A8%E6%8E%A7%E5%88%B6%E5%81%9A%E7%A7%8D%E9%85%8D%E7%BD%AE%E8%AF%B4%E6%98%8E)。

# 4. 根据tracker工作状态设置标签

根据种子的tracker返回的状态设置标签，目前已有的状态有：关闭(不支持的协议如DHT)、正常、更新中、未开始工作(暂停的种子)、出错(原因五花八门)。

其中只有出错需要额外关心，然而出错的原因五花八门状态却都是4，所以需要提前配置错误信息对应的标签，已发现未配置的错误信息会在处理结果种给出，如果不想设置标签又不想在处理结果种显示可以配置为空。

已经存在的标签不会自动删除，可在 webui 右键删除，等待种子再次被处理。

# 5. 批量替换tracker

todo

Test