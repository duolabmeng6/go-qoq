package main

import (
	"embed"
	_ "embed"
	"encoding/json"
	"fmt"
	"github.com/duolabmeng6/goefun/ecore"
	"github.com/duolabmeng6/goefun/etool"
	"github.com/duolabmeng6/goefun/etranslation"
	"github.com/go-vgo/robotgo"
	"github.com/ncruces/zenity"
	hook "github.com/robotn/gohook"
	"github.com/wailsapp/wails/v3/pkg/application"
	"github.com/wailsapp/wails/v3/pkg/events"
	"log"
	"os"
	"runtime"
	"strconv"
	"strings"
	"v3fanyi/mymodel"
)

//go:embed assets
var assets embed.FS
var AppData struct {
	置顶      bool
	可视      bool
	第一次复制   bool
	Hook    chan hook.Event
	划词翻译快捷键 []string
	截图翻译快捷键 []string
	调整位置    bool
	窗口_设置   *application.WebviewWindow
}

func 窗口_设置_显示(app *application.App) {
	if AppData.窗口_设置 == nil {
		AppData.窗口_设置 = app.NewWebviewWindowWithOptions(application.WebviewWindowOptions{
			Title: "qoq翻译 - 设置",
			Mac: application.MacWindow{
				Backdrop: application.MacBackdropTranslucent,
				TitleBar: application.MacTitleBar{
					AppearsTransparent:   true,
					Hide:                 false,
					HideTitle:            true,
					FullSizeContent:      true,
					UseToolbar:           false,
					HideToolbarSeparator: true,
				},
				InvisibleTitleBarHeight: 50,
				DisableShadow:           true,
			},
			Width:  500,
			Height: 600,
			URL:    "index_set2.html",
			Hidden: true,
		})
		ecore.E延时(300) //等待页面渲染完成再显示

		AppData.窗口_设置.On(events.Mac.WindowShouldClose, func(ctx *application.WindowEventContext) {
			fmt.Println("我想在这里禁止窗口销毁 我只隐藏他")
			AppData.窗口_设置 = nil
		})
	}
	AppData.窗口_设置.Show().Focus()
}
func 窗口_设置_隐藏(app *application.App) {
	if AppData.窗口_设置 == nil {
		return
	}
	AppData.窗口_设置.Hide()
}

func main() {
	fmt.Println("运行目录", ecore.E取运行目录())
	if !ecore.E文件是否存在(ecore.E取运行目录() + "/default_config.json") {
		字节集, _ := assets.ReadFile("assets/config/default_config.json")
		ecore.E写到文件(ecore.E取运行目录()+"/default_config.json", 字节集)
	}
	if !ecore.E文件是否存在(ecore.E取运行目录() + "/llocr") {
		字节集, _ := assets.ReadFile("assets/config/llocr")
		ecore.E写到文件(ecore.E取运行目录()+"/llocr", 字节集)
	}

	if mymodel.E加载配置文件("./") == false {
		panic("加载配置文件失败")
	}

	ecore.E调试输出(mymodel.GConfig)
	app := application.New(application.Options{
		Name:        "qoq翻译",
		Description: "",
		Mac: application.MacOptions{
			ActivationPolicy: application.ActivationPolicyAccessory,
			ApplicationShouldTerminateAfterLastWindowClosed: false,
		},
		Assets: application.AssetOptions{
			FS: assets,
		},
	})

	window := app.NewWebviewWindowWithOptions(application.WebviewWindowOptions{
		Title: "qoq翻译",
		Mac: application.MacWindow{
			Backdrop: application.MacBackdropTranslucent,
			TitleBar: application.MacTitleBar{
				AppearsTransparent:   true,
				Hide:                 true,
				HideTitle:            true,
				FullSizeContent:      true,
				UseToolbar:           false,
				HideToolbarSeparator: true,
			},
			InvisibleTitleBarHeight: 50,
			DisableShadow:           true,
		},
		Width:  500,
		Height: 400,
		URL:    "index.html",
	})
	AppData.调整位置 = true
	fmt.Println("window", window)
	//AppData.可视 = true
	window.SetAlwaysOnTop(true)
	//AppData.置顶 = true
	window.Hide()

	// 定义窗口创建完毕事件
	//app.On(events.Mac.ApplicationDidFinishLaunching, func() {
	//	println("ApplicationDidFinishLaunching")
	//	window.ExecJS("app.init()")
	//})

	//window_set.SetAlwaysOnTop(true)
	//window_set.Hide()

	app.Events.On("js_translate", func(e *application.WailsEvent) {
		jsondata := ecore.E到文本(e.Data)

		翻译平台 := etool.Json解析文本(jsondata, "翻译平台")
		内容 := etool.Json解析文本(jsondata, "内容")
		源语言 := etool.Json解析文本(jsondata, "源语言")   // auto
		目标语言 := etool.Json解析文本(jsondata, "目标语言") //auto
		callfc := etool.Json解析文本(jsondata, "callfc")

		目标语言x := "zh_CN"
		if 源语言 == "auto" {
			源语言 = etranslation.E内容语言类型检测(内容)
			if 源语言 == "zh_CN" {
				目标语言x = "en_US"
			}
			//识别语言 := mymodel.E简写转中文(源语言)
		}
		if 目标语言 == "auto" {
			目标语言 = 目标语言x
		}

		var 翻译结果 string

		if mymodel.G翻译接口.E模块是否存在(翻译平台) {
			翻译结果, _ = mymodel.G翻译接口.E取翻译模块(翻译平台).E翻译(内容, 源语言, 目标语言)
		} else {
			翻译结果 = "翻译平台不存在"
		}

		fmt.Println("js_translate==", "翻译平台:", 翻译平台, "内容:", 内容, "源语言:", 源语言, "目标语言:", 目标语言, "翻译结果", 翻译结果)

		if 翻译结果 == "" {
			翻译结果 = "翻译失败"
		}
		if mymodel.GConfig.E自动复制首个翻译结果 {
			if AppData.第一次复制 {
				AppData.第一次复制 = false
				app.Clipboard().SetText(翻译结果)
			}
		} else {
			if mymodel.GConfig.E自动复制截图翻译OCR结果 {
				app.Clipboard().SetText(内容)
			}
		}

		var rtdata struct {
			E名称 string `json:"名称"`
			E结果 string `json:"结果"`
		}
		rtdata.E名称 = 翻译平台
		rtdata.E结果 = 翻译结果
		jsontxt := ecore.E到文本(rtdata)
		println("js_translate", jsontxt)

		app.Events.Emit(&application.WailsEvent{
			Name: callfc,
			Data: jsontxt,
		})
	})

	app.Events.On("PyAction", func(e *application.WailsEvent) {
		log.Printf("[Go] WailsEvent received: %+v\n", e.Data)
		jsondata := ecore.E到文本(e.Data)
		action := etool.Json解析文本(jsondata, "action")
		data := etool.Json解析文本(jsondata, "data")
		fmt.Println("action", action, "data", data)
		if action == "操作_弹出菜单" {
			窗口_设置_显示(app)

			window.Hide()
		}
		if action == "操作_钉着窗口" {
			AppData.置顶 = data == "true"
			//window.SetAlwaysOnTop(AppData.置顶)
		}
		if action == "操作_截图翻译" {
			window.Hide()
			图片地址 := mymodel.SystemScreenshot()
			if 图片地址 != "" {
				txt := mymodel.MacOcr(图片地址)
				os.Remove(图片地址)
				fmt.Println(txt)
				window.Show().Focus()
				app.Events.Emit(&application.WailsEvent{
					Name: "setContent",
					Data: txt,
				})
			}
		}

		if action == "操作_播放声音" {
			//window.Show().SetSize(500, 1000)
			mymodel.Say(data)

		}
		if action == "操作_复制" {
			//复制内容到剪切板
			app.Clipboard().SetText(data)
		}
		if action == "操作_内容高度" {
			//设置窗口高度
			//文本转数值
			高度, _ := strconv.Atoi(data)
			调整窗口位置_限定高度(window, 高度)
			window.Reload()

			//窗口可视(window, true)

		}

		if action == "操作_获取平台列表" {
			// 遍历mymodel.GConfig.E翻译服务列表
			names := make([]string, 0)
			for _, v := range mymodel.GConfig.E默认翻译服务列表 {
				if v.Enable {
					names = append(names, v.Name)
				}
			}
			//根据 names 生产 isOpen 为 true
			isOpen := make([]bool, len(names))
			for i := 0; i < len(names); i++ {
				isOpen[i] = true
			}

			app.Events.Emit(&application.WailsEvent{
				Name: action,
				Data: etool.E到Json(map[string]interface{}{
					"平台列表":    names,
					"isOpen":  isOpen,
					"显示输入框":   mymodel.GConfig.E显示输入框,
					"显示语言切换栏": mymodel.GConfig.E显示语言切换栏,
					"语言列表":    etranslation.New语言列表().E取全部名称对照数组(),
				}),
			})
		}
		if action == "操作_获取全部配置" {
			app.Events.Emit(&application.WailsEvent{
				Name: action,
				Data: etool.E到Json(mymodel.GConfig),
			})
		}
		if action == "操作_保存配置" {
			运行目录 := ecore.E取运行目录()
			println("运行目录", 运行目录)
			ecore.E写到文件(运行目录+"/user_config.json", []byte(data))
			窗口_设置_隐藏(app)

			//重新加载配置
			mymodel.E加载配置文件("./")
			window.ExecJS("app.init()")

			绑定热键(app, window)
		}
		if action == "操作_设置取消" {
			窗口_设置_隐藏(app)
		}
		if action == "操作_检查更新" {
			检查更新()
		}
		if action == "操作_测试API" {
			fmt.Print("操作_测试API")
			var rtdata mymodel.GConfig翻译服务列表
			json.Unmarshal([]byte(data), &rtdata)
			fmt.Print("测试API", rtdata.Secret_key, rtdata.App_id)

			x翻译接口 := etranslation.New翻译()
			if rtdata.Name == "百度翻译" {
				x翻译接口.E注册服务("百度翻译", etranslation.New百度翻译(rtdata.App_id, rtdata.Secret_key))
			}
			if rtdata.Name == "彩云小译" {
				x翻译接口.E注册服务("彩云小译", etranslation.New彩云小译(rtdata.App_id))
			}
			if rtdata.Name == "有道翻译" {
				x翻译接口.E注册服务("有道翻译", etranslation.New有道翻译(rtdata.App_id, rtdata.Secret_key))
			}
			if rtdata.Name == "火山翻译" {
				x翻译接口.E注册服务("火山翻译", etranslation.New火山翻译(rtdata.App_id, rtdata.Secret_key))
			}
			if rtdata.Name == "腾讯翻译" {
				x翻译接口.E注册服务("腾讯翻译", etranslation.New腾讯翻译(rtdata.App_id, rtdata.Secret_key))
			}
			if rtdata.Name == "阿里云翻译" {
				x翻译接口.E注册服务("阿里云翻译", etranslation.New阿里云翻译(rtdata.App_id, rtdata.Secret_key))
			}
			if rtdata.Name == "DeepL翻译" {
				x翻译接口.E注册服务("DeepL翻译", etranslation.NewDeepL翻译(rtdata.App_id))
			}
			翻译内容, _ := x翻译接口.E取翻译模块(rtdata.Name).E翻译("我是一个小可爱", "zh_CN", "en_US")
			app.Events.Emit(&application.WailsEvent{
				Name: action,
				Data: etool.E到Json(map[string]interface{}{
					"翻译内容": 翻译内容,
				}),
			})
		}

		if action == "操作_获取原文本语言" {
			AppData.第一次复制 = true
			AppData.调整位置 = true
			翻译内容 := etool.Json解析文本(data, "翻译内容")
			源语言 := etool.Json解析文本(data, "源语言")
			翻译内容 = ecore.E删首尾空(翻译内容)
			if 翻译内容 == "" {
				return
			}

			println("翻译内容", 翻译内容, "源语言", 源语言)

			目标语言 := "zh_CN"

			识别语言 := etranslation.E内容语言类型检测(翻译内容)
			if 识别语言 == "zh_CN" {
				目标语言 = "en_US"
			}
			if mymodel.GConfig.E将翻译原文的换行符替换为空格 {
				翻译内容 = strings.Replace(翻译内容, "\n", " ", -1)
			}
			if mymodel.GConfig.E将翻译原文的空格去掉 {
				翻译内容 = strings.Replace(翻译内容, "-", "", -1)
			}
			if mymodel.GConfig.E将翻译原文的代码注释符号去掉 {
				翻译内容 = strings.Replace(翻译内容, "//", "", -1)
				翻译内容 = strings.Replace(翻译内容, "#", "", -1)
				翻译内容 = strings.Replace(翻译内容, "/*", "", -1)
				翻译内容 = strings.Replace(翻译内容, "*/", "", -1)
			}

			识别语言 = etranslation.E简写转中文(识别语言)
			app.Events.Emit(&application.WailsEvent{
				Name: action,
				Data: etool.E到Json(map[string]interface{}{
					"翻译内容": 翻译内容,
					"源语言":  源语言,
					"识别语言": 识别语言,
					"目标语言": 目标语言,
				}),
			})
		}

	})

	systemTray := app.NewSystemTray()
	if runtime.GOOS == "darwin" {
		//systemTray.SetTemplateIcon(icons.SystrayMacTemplate)
		b, _ := assets.ReadFile("assets/images/app.ico")
		systemTray.SetTemplateIcon(b)
	}

	myMenu := app.NewMenu()
	myMenu.Add("qoq " + mymodel.Version).OnClick(func(ctx *application.Context) {
		println("Hello World!")
	})
	myMenu.Add("输入翻译").OnClick(func(ctx *application.Context) {
		println("输入翻译")
		window.Show().Focus()
	})
	myMenu.Add("截图翻译 " + mymodel.GConfig.E截图翻译快捷键).OnClick(func(ctx *application.Context) {
		println("截图翻译")
		//window.Hide()
		window.ExecJS("app.操作_截图翻译()")
	})
	myMenu.Add("划词翻译 " + mymodel.GConfig.E划词翻译快捷键).OnClick(func(ctx *application.Context) {
		println("划词翻译")
		划词翻译(app, window)

	})
	myMenu.Add("静默截图OCR").OnClick(func(ctx *application.Context) {
		println("静默截图OCR")
		图片地址 := mymodel.SystemScreenshot()
		if 图片地址 != "" {
			txt := mymodel.MacOcr(图片地址)
			os.Remove(图片地址)
			fmt.Println(txt)
			app.Clipboard().SetText(txt)
		}
	})
	myMenu.AddSeparator()

	myMenu.Add("偏好设置").OnClick(func(ctx *application.Context) {
		println("偏好设置")
		窗口_设置_显示(app)
	})
	myMenu.Add("检查更新").OnClick(func(ctx *application.Context) {
		println("检查更新")

		//mymodel.OpenURL("https://github.com/duolabmeng6/go-qoq")
		检查更新()

	})
	myMenu.AddSeparator()

	myMenu.Add("退出").OnClick(func(ctx *application.Context) {
		println("退出")
		app.Quit()

	})

	systemTray.SetMenu(myMenu)

	绑定热键(app, window)

	err := app.Run()
	if err != nil {
		log.Fatal(err.Error())
	}
}
func 划词翻译(app *application.App, window *application.WebviewWindow) {
	text := mymodel.MacGetCmdC()
	fmt.Println("剪贴板中的文本数据:", text)
	text = ecore.E删首尾空(text)
	窗口可视(window, true)
	app.Events.Emit(&application.WailsEvent{
		Name: "setContent",
		Data: text,
	})
}

func 窗口可视(window *application.WebviewWindow, b bool) {
	AppData.可视 = b
	if AppData.可视 {
		// 检查 鼠标处x 加上窗口宽度 是否超出屏幕
		鼠标处x, 鼠标处y := robotgo.GetMousePos()
		fmt.Println("鼠标坐标：", 鼠标处x, 鼠标处y)
		screen, _ := window.GetScreen()
		屏幕宽度 := screen.Size.Width
		屏幕高度 := screen.Size.Height
		窗口宽度 := window.Width()
		窗口高度 := window.Height()
		if 鼠标处x+窗口宽度 > 屏幕宽度 {
			鼠标处x = 屏幕宽度 - 窗口宽度
		}
		if 鼠标处y+窗口高度 > 屏幕高度 {
			SetWindowPosition(window, 鼠标处x, 屏幕高度-窗口高度)
		} else {
			SetWindowPosition(window, 鼠标处x, 鼠标处y)
		}
	}

	if b {
		window.Show().Focus()
	} else {
		window.Hide()
	}
}

func 绑定热键(app *application.App, window *application.WebviewWindow) {
	mymodel.E初始化mac键代码()
	n划词翻译快捷键 := mymodel.GConfig.E划词翻译快捷键
	n截图翻译快捷键 := mymodel.GConfig.E截图翻译快捷键
	if runtime.GOOS == "darwin" {
		n划词翻译快捷键 = strings.Replace(n划词翻译快捷键, "Cmd", "command", -1)
		n截图翻译快捷键 = strings.Replace(n截图翻译快捷键, "Cmd", "command", -1)
		AppData.划词翻译快捷键 = ecore.E分割文本(n划词翻译快捷键, "+")
		AppData.截图翻译快捷键 = ecore.E分割文本(n截图翻译快捷键, "+")
	} else {
		AppData.划词翻译快捷键 = ecore.E分割文本(n划词翻译快捷键, "+")
		AppData.截图翻译快捷键 = ecore.E分割文本(n截图翻译快捷键, "+")
	}
	ecore.E调试输出("================绑定热键 划词 快捷键", AppData.划词翻译快捷键)
	ecore.E调试输出("================绑定热键 截图 快捷键", AppData.截图翻译快捷键)
	if AppData.Hook != nil {
		return
	}
	go func() {
		var G按下的热键列表 = []string{}
		robotgo.EventHook(hook.KeyHold, []string{}, func(e hook.Event) {
			G按下的热键列表 = append(G按下的热键列表, mymodel.Mac键代码取键名称(e.Rawcode))
		})
		robotgo.EventHook(hook.KeyUp, []string{}, func(e hook.Event) {
			//ecore.E调试输出("G按下的热键列表", G按下的热键列表)
			// Esc 关闭窗口
			if e.Keychar == rune(65535) {
				window.Hide()
			}
			if mymodel.ContainsAll(G按下的热键列表, AppData.划词翻译快捷键...) {
				println("划词翻译")
				go 划词翻译(app, window)
			}
			if mymodel.ContainsAll(G按下的热键列表, AppData.截图翻译快捷键...) {
				println("截图翻译")
				go window.ExecJS("app.操作_截图翻译()")
			}
			//清空  G按下的热键列表
			G按下的热键列表 = []string{}
		})

		robotgo.EventHook(hook.MouseDown, []string{}, func(e hook.Event) {
			if AppData.置顶 {
				return
			}
			if !AppData.可视 {
				return
			}

			fmt.Println("hook: ", e)

			// 获取窗口的位置信息 判断鼠标是否还在窗口内
			x, y := GetWindowPosition(window)
			w, h := window.Size()
			w += 20
			h += 20
			fmt.Println("窗口位置：", x, y, w, h, e.X, e.Y)
			if int(e.X) < x || int(e.X) > x+w || int(e.Y) < y || int(e.Y) > y+h {
				window.Hide()
				AppData.可视 = false
				fmt.Println("鼠标在窗口外")
			} else {
				//鼠标在窗口内
				fmt.Println("鼠标在窗口内")

				//runtime.WindowHide(a.ctx)
			}
		})
		AppData.Hook = robotgo.EventStart()
		<-robotgo.EventProcess(AppData.Hook)
	}()
}

func GetWindowPosition(app *application.WebviewWindow) (int, int) {
	//screen, err := app.GetScreen()
	//if err != nil {
	//	return 0, 0
	//}
	//wwidth := app.Width()
	//wheight := app.Height()
	//swidth := screen.Size.Width
	//sheight := screen.Size.Height
	x, y := app.Position()
	return x, y
	//newX := x
	//newY := sheight - wheight - y - 1
	//return newX, newY
}
func SetWindowPosition(app *application.WebviewWindow, x int, y int) {
	app.SetPosition(x, y)
	//获取屏幕大小
	//screen, err := app.GetScreen()
	//if err != nil {
	//	return
	//}
	////wwidth := app.Width()
	//wheight := app.Height()
	////swidth := screen.Size.Width
	//sheight := screen.Size.Height
	//newX := x
	//newY := sheight - wheight - y - 1
	//app.SetPosition(newX, newY)
}

func 调整窗口位置_限定高度(window *application.WebviewWindow, 内容高度 int) {
	if AppData.调整位置 == false {
		return
	}
	AppData.调整位置 = false
	window.SetSize(window.Width(), 内容高度)
	window.Reload()
	//
	//var n左边, n顶边, n宽度, n高度 int
	//
	//保留上边距 := 24
	//保留下边距 := 36
	//screen, err := window.GetScreen()
	//if err != nil {
	//	return
	//}
	//屏幕宽度 := screen.Size.Width
	//屏幕高度 := screen.Size.Height
	//窗口宽度 := window.Width()
	//窗口高度 := window.Height()
	//窗口左边, 窗口顶边 := GetWindowPosition(window)
	//println("窗口左边", 窗口左边, "窗口顶边", 窗口顶边, "窗口宽度", 窗口宽度, "窗口高度", 窗口高度, "屏幕宽度", 屏幕宽度, "屏幕高度", 屏幕高度)
	//
	//if 窗口左边 < 0 {
	//	窗口左边 = 0
	//}
	//if 窗口顶边 < 0 {
	//	窗口顶边 = 0
	//}
	//
	//屏幕高度 = 屏幕高度 - 保留上边距 - 保留下边距
	//if 内容高度 > 屏幕高度 {
	//	内容高度 = 屏幕高度
	//}
	//n高度 = 内容高度 + 12
	//n高度 = 窗口宽度
	//n左边 = 窗口左边
	//n顶边 = 窗口顶边
	//
	//if 窗口左边+内容高度 > 屏幕高度 {
	//	n左边 = 窗口左边
	//	n顶边 = 屏幕高度 - 内容高度
	//}
	//
	//if n宽度 < 0 {
	//	n宽度 = 300
	//}
	//if n高度 < 0 {
	//	n高度 = 400
	//}
	//ecore.E调试输出("n左边", n左边, "n顶边", n顶边, "n宽度", n宽度, "n高度", n高度)
	//
	//SetWindowPosition(window, n左边, n顶边)
	//window.SetSize(n宽度, n高度)
	//window.Reload()
}
func 检查更新() {

	下载文件夹路径 := mymodel.E取用户下载文件夹路径()
	info := mymodel.E获取Github仓库Releases版本和更新内容()
	println(info.MacDownloadURL)
	println(下载文件夹路径)
	if info.Version == mymodel.Version {
		err := zenity.Info("当前已经是最新版本")
		if err != nil {
			return
		}
		return
	}

	err := zenity.Question("软件有新版本可用，是否更新？\n当前版本:"+
		mymodel.Version+
		"\n最新版本:"+info.Version,
		zenity.Title("更新提示"),
		zenity.Icon(zenity.QuestionIcon),
		zenity.OKLabel("更新"),
		zenity.CancelLabel("取消"))
	ecore.E调试输出(err)
	println("更新", err)
	if err != nil {
		return
	}
	progress, _ := zenity.Progress(
		zenity.Title("软件更新"),
		zenity.MaxValue(100), // 设置最大进度值为100
	)
	//for i := 1; i <= 100; i++ {
	//	// 更新进度对话框的进度
	//	progress.Value(i)
	//	time.Sleep(100 * time.Millisecond) // 模拟任务执行时间
	//}
	progress.Text("正在下载...")

	err = mymodel.E下载带进度回调(info.MacDownloadURL, 下载文件夹路径+"/qoq_MacOS.zip", func(进度 float64) {
		fmt.Println("正在下载...", 进度)
		progress.Text("正在下载..." + fmt.Sprintf("%.2f", 进度) + "%")
		progress.Value(int(进度))
	})
	if err != nil {
		fmt.Println("下载出错:", err)
		zenity.Info("下载错误,检查你的网络")
		progress.Close()
		return
	}
	progress.Text("下载完成 即将完成更新")
	if progress.Close() != nil {
		fmt.Println("点击了取消")
		return
	}
	fmt.Println("下载完成了")
	flag, s := mymodel.E更新自己MacOS应用(下载文件夹路径+"/qoq_MacOS.zip", "qoq.app")
	println(flag, s)
}
