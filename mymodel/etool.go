package mymodel

import (
	"fmt"
	"github.com/duolabmeng6/goefun/ecore"
	"github.com/duolabmeng6/goefun/etranslation"
	"github.com/spf13/viper"
	"io/ioutil"
	"os"
	"os/exec"
)

func SystemScreenshot() string {
	file, err := ioutil.TempFile("", "screenshot")
	if err != nil {
		return ""
	}
	filePath := file.Name()
	file.Close()

	cmd := exec.Command("screencapture", "-i", "-s", "-x", filePath)
	cmd.Run()

	fileInfo, err := os.Stat(filePath)
	if err != nil {
		return ""
	}

	if fileInfo.Size() == 0 {
		os.Remove(filePath)
		return ""
	}

	return filePath
}

func MacOcr(图片路径 string) string {
	// 调用 ./llocr 图片路径 获取结果
	cmd := exec.Command("./llocr", 图片路径)
	out, err := cmd.Output()
	if err != nil {
		return ""
	}
	return string(out)
}

func MacGetCmdC() string {
	// 构造AppleScript代码
	script := `
tell application "System Events"
    delay 0.2
    keystroke "c" using {command down}
    delay 0.2 -- 等待复制操作完成，可以根据实际情况调整延迟时间
end tell

set clipboardContent to the clipboard as text
if clipboardContent is equal to missing value then
    set clipboardContent to ""
end if

return clipboardContent
`
	//println(script)
	cmd := exec.Command("osascript", "-e", script)
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("执行AppleScript时出错:", err)
	}
	resultValue := string(output)

	return ecore.E删首尾空(string(resultValue))
}
func MacGetFanyi(s, from, to string) string {
	if from == "auto" {
		from = ""
	}
	if to == "auto" {
		to = ""
	}
	if from == "" {
		from = "&from=en_US"
	} else {
		from = "&from=" + from
	}
	if to == "" {
		to = "zh_CN"
	}
	// 构造AppleScript代码
	script := `
tell application "Shortcuts"
	set myInput to "text=` + s + from + `&to=` + to + `"
	set resultValue to run shortcut named "Easydict-Translate-V1.2.0" with input myInput
end tell
return resultValue
`
	//println(script)
	cmd := exec.Command("osascript", "-e", script)
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("执行AppleScript时出错:", err)
	}
	resultValue := string(output)
	return resultValue
}

var GConfig struct {
	E划词翻译快捷键           string          `json:"划词翻译快捷键" mapstructure:"划词翻译快捷键"`
	E截图翻译快捷键           string          `json:"截图翻译快捷键" mapstructure:"截图翻译快捷键"`
	E显示输入框             bool            `json:"显示输入框" mapstructure:"显示输入框"`
	E显示语言切换栏           bool            `json:"显示语言切换栏" mapstructure:"显示语言切换栏"`
	E将翻译原文的换行符替换为空格    bool            `json:"将翻译原文的换行符替换为空格" mapstructure:"将翻译原文的换行符替换为空格"`
	E将翻译原文的空格去掉        bool            `json:"将翻译原文的空格去掉" mapstructure:"将翻译原文的空格去掉"`
	E将翻译原文的代码注释符号去掉    bool            `json:"将翻译原文的代码注释符号去掉" mapstructure:"将翻译原文的代码注释符号去掉"`
	E自动复制截图翻译OCR结果     bool            `json:"自动复制截图翻译OCR结果" mapstructure:"自动复制截图翻译OCR结果"`
	E自动复制首个翻译结果        bool            `json:"自动复制首个翻译结果" mapstructure:"自动复制首个翻译结果"`
	E点击翻译结果中的蓝色单词时将其复制 bool            `json:"点击翻译结果中的蓝色单词时将其复制" mapstructure:"点击翻译结果中的蓝色单词时将其复制"`
	E翻译服务列表            []GConfig翻译服务列表 `json:"翻译服务列表" mapstructure:"翻译服务列表"`
	E默认翻译服务列表          []GConfig翻译服务列表 `json:"默认翻译服务列表" mapstructure:"默认翻译服务列表"`
}

type GConfig翻译服务列表 struct {
	Name       string   `json:"name"`
	App_id     string   `json:"app_id"`
	Secret_key string   `json:"secret_key"`
	Enable     bool     `json:"enable"`
	Order      int      `json:"order"`
	Logo       string   `json:"logo"`
	Params     []string `json:"params"`
}

var G翻译接口 *etranslation.E翻译

func E加载配置文件(dir string) bool {
	defaultFile := dir + "/default_config.json"
	viper.SetConfigType("json")
	viper.SetConfigFile(defaultFile)
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("读取配置文件失败", err)
		return false
	}
	err = viper.Unmarshal(&GConfig)
	if err != nil {
		fmt.Println("读取配置文件失败2", err)
		return false
	}

	viper.SetConfigFile(dir + "/user_config.json")
	err = viper.ReadInConfig()
	if err != nil {
		fmt.Println("读取配置文件失败", err)
		//return false
	}

	err = viper.Unmarshal(&GConfig)
	if err != nil {
		fmt.Println("读取配置文件失败2", err)
		return false
	}

	G翻译接口 = etranslation.New翻译()
	G翻译接口.E注册服务("系统翻译", etranslation.New系统翻译())
	G翻译接口.E注册服务("必应免费翻译", etranslation.New必应免费翻译())
	G翻译接口.E注册服务("搜狗免费翻译", etranslation.New搜狗免费翻译())
	G翻译接口.E注册服务("爱词霸免费翻译", etranslation.New爱词霸免费翻译())
	G翻译接口.E注册服务("DeepL免费翻译", etranslation.NewDeepL免费翻译())
	G翻译接口.E注册服务("阿里云免费翻译", etranslation.New阿里云免费翻译())
	G翻译接口.E注册服务("百度翻译", etranslation.New百度翻译(ecore.Env("百度翻译_App_id", ""), ecore.Env("百度翻译_Secret_key", "")))
	G翻译接口.E注册服务("彩云小译", etranslation.New彩云小译(ecore.Env("彩云小译_App_id", "")))
	G翻译接口.E注册服务("有道翻译", etranslation.New有道翻译(ecore.Env("有道翻译_App_id", ""), ecore.Env("有道翻译_Secret_key", "")))
	G翻译接口.E注册服务("火山翻译", etranslation.New火山翻译(ecore.Env("火山翻译_App_id", ""), ecore.Env("火山翻译_Secret_key", "")))
	G翻译接口.E注册服务("腾讯翻译", etranslation.New腾讯翻译(ecore.Env("腾讯翻译_App_id", ""), ecore.Env("腾讯翻译_Secret_key", "")))
	G翻译接口.E注册服务("阿里云翻译", etranslation.New阿里云翻译(ecore.Env("阿里云翻译_App_id", ""), ecore.Env("阿里云翻译_Secret_key", "")))

	// 遍历 G翻译接口.E列出翻译模块() 把他加入 GConfig.E翻译服务列表
	for _, v := range G翻译接口.E列出翻译模块和初始化参数() {
		// 检查是否已经存在 GConfig.E默认翻译服务列表 的 Name
		var isExist bool
		for i, v2 := range GConfig.E默认翻译服务列表 {
			if v2.Name == v.Name {
				isExist = true
				GConfig.E默认翻译服务列表[i].Params = v.Params

				break
			}
		}
		if isExist {
			continue
		}

		GConfig.E默认翻译服务列表 = append(GConfig.E默认翻译服务列表, GConfig翻译服务列表{
			Name:   v.Name,
			Order:  0,
			Logo:   "",
			Params: v.Params,
			Enable: ecore.E判断文本(v.Name, "免费"),
		})
	}

	// 遍历 E默认翻译服务列表 把用户的配置覆盖上去
	for i, v := range GConfig.E默认翻译服务列表 {
		for _, v2 := range GConfig.E翻译服务列表 {
			if v.Name == v2.Name {
				GConfig.E默认翻译服务列表[i].App_id = v2.App_id
				GConfig.E默认翻译服务列表[i].Secret_key = v2.Secret_key
				GConfig.E默认翻译服务列表[i].Enable = v2.Enable
				GConfig.E默认翻译服务列表[i].Order = v2.Order
				if GConfig.E默认翻译服务列表[i].Params == nil {
					GConfig.E默认翻译服务列表[i].Params = make([]string, 0)
				}
				ecore.E写环境变量(v.Name+"_App_id", v2.App_id)
				ecore.E写环境变量(v.Name+"_Secret_key", v2.Secret_key)
			}
		}
	}

	G翻译接口.E注册服务("百度翻译", etranslation.New百度翻译(ecore.Env("百度翻译_App_id", ""), ecore.Env("百度翻译_Secret_key", "")))
	G翻译接口.E注册服务("彩云小译", etranslation.New彩云小译(ecore.Env("彩云小译_App_id", "")))
	G翻译接口.E注册服务("有道翻译", etranslation.New有道翻译(ecore.Env("有道翻译_App_id", ""), ecore.Env("有道翻译_Secret_key", "")))
	G翻译接口.E注册服务("火山翻译", etranslation.New火山翻译(ecore.Env("火山翻译_App_id", ""), ecore.Env("火山翻译_Secret_key", "")))
	G翻译接口.E注册服务("腾讯翻译", etranslation.New腾讯翻译(ecore.Env("腾讯翻译_App_id", ""), ecore.Env("腾讯翻译_Secret_key", "")))
	G翻译接口.E注册服务("阿里云翻译", etranslation.New阿里云翻译(ecore.Env("阿里云翻译_App_id", ""), ecore.Env("阿里云翻译_Secret_key", "")))

	//c := viper.AllSettings()
	//ecore.E调试输出(c)
	//ecore.E调试输出(GConfig)
	//ecore.E调试输出(GConfig.E划词翻译快捷键)
	//ecore.E调试输出(GConfig.E划词翻译快捷键)

	return true

}

func Say(txt string) {
	cmd := exec.Command("say", txt)
	err := cmd.Run()
	if err != nil {
		fmt.Println("执行AppleScript时出错:", err)
	}

}
