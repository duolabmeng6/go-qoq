package mymodel

import (
	"fmt"
	"github.com/duolabmeng6/goefun/ecore"
	"github.com/duolabmeng6/goefun/etool"
	"github.com/duolabmeng6/goefun/etranslation"
	"os"
	"testing"
)

func TestMacOcr(t *testing.T) {

	图片地址 := SystemScreenshot()
	t.Log(图片地址)
	if 图片地址 == "" {
		return
	}
	txt := MacOcr(图片地址)
	t.Log(txt)
	os.Remove(图片地址)
}

func Test取剪切板内容(t *testing.T) {
	println("MacGetCmdC()", MacGetCmdC())
}
func Test翻译(t *testing.T) {
	println("Test翻译()", MacGetFanyi(`apple 
-=你好
hello
`, "en_US", "zh_CN"))
}

func Test取配置文件(t *testing.T) {
	println("E加载配置文件()", E加载配置文件("/Users/chensuilong/Desktop/goproject/v3fanyi/"))
	//ecore.E调试输出(GConfig.E翻译服务列表)

	//ecore.E调试输出(GConfig.E默认翻译服务列表)
	//eco
	// json 格式化

	ecore.E调试输出(etool.Json美化(GConfig))

	//ecore.E调试输出(GConfig.E默认翻译服务列表)

}

func TestDetectContentLanguage(t *testing.T) {
	println(etranslation.E内容语言类型检测("hello world"))
	println(etranslation.E内容语言类型检测("你好世界"))
	println(etranslation.E内容语言类型检测("こんにちは世界"))
	println(etranslation.E内容语言类型检测("你好"))
	println(etranslation.E内容语言类型检测("こんにちは"))
}

func TestE取全部名称(t *testing.T) {
	datalist := etranslation.New语言列表().E取全部名称对照数组()
	ecore.E调试输出(datalist)
	ecore.E调试输出(ecore.E到文本(datalist))

}

func Test热键检查(t *testing.T) {
	s := []string{"command", "keypad0", "control"}
	allIncluded := ContainsAll(s, "Keypad0", "Command", "Control")
	fmt.Println(allIncluded)
}
