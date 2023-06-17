# go-qoq



一款go语言开发的跨平台翻译软件

项目由 [wails](https://github.com/wailsapp/wails) 框架完成

# 进度

- [x] qoq翻译 macOS 系统核心功能完成
- [ ] macOS 打包为应用 待完成
- [ ] Window go-qoq 待开发 目前有qt6开发的 [window版本 qoq](https://github.com/duolabmeng6/qoq)





# [点击这里下载应用](https://github.com/duolabmeng6/go-qoq/releases)



安装教程



![image-20230618022842692](images/README/image-20230618022842692.png)

打开后会提示辅助功能访问 这是用于划词翻译的权限 如果不给权限会出现窗口不消失和快捷键无效的情况

![image-20230618022950349](images/README/image-20230618022950349.png)

![image-20230618023155525](images/README/image-20230618023155525.png)

如果使用截屏翻译功能 需要录制屏幕权限 否则截屏会失败

![image-20230618023219205](images/README/image-20230618023219205.png)

![image-20230618023338871](images/README/image-20230618023338871.png)



## macOS 系统翻译

如果想要调用苹果 🍎 系统翻译，系统需要在 **macOS 12.3.1** 以上。

此外还需要安装一个快捷指令（`Easydict-Translate-V1.2.0`），请用 **Safari 浏览器** 打开下方链接。

```
https://www.icloud.com/shortcuts/776f8a1d8e43471885e8a505eb9a9deb
```

安装后即可使用



# 开发使用

编译应用

```
wails build
```

调试应用

```
直接在goland 运行 main.go 即可
```



本项目参考

[tisfeng/Easydict: 一个简洁优雅的翻译词典 macOS App](https://github.com/tisfeng/Easydict)




![3000b7271008d1da9d4a5a7f7bb4fa96](images/README/3000b7271008d1da9d4a5a7f7bb4fa96.png)

![image-20230618022414195](images/README/image-20230618022414195.png)





![image-20230618022444684](images/README/image-20230618022444684.png)



![image-20230618022555084](images/README/image-20230618022555084.png)

## 声明

qoq 为 [GPL-3.0](https://github.com/tisfeng/Easydict/blob/main/LICENSE) 开源协议，仅供学习交流，任何人都可以免费获取该产品和源代码。如果你认为您的合法权益受到侵犯，请立即联系[作者](https://github.com/duolabmeng6)。你可以自由使用源代码，但必须附上相应的许可证和版权声明。
