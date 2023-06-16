//Global JS function for greeting
function PyAction(name, data, fn) {
    jsondata = {
        "action": name,
        "data": data
    }
    console.log("PyAction 发送数据", jsondata)
    wails.Events.Emit({
        name: "PyAction",
        data: JSON.stringify(jsondata)
    })
    wails.Events.Once(name, function (data) {
        console.log("PyAction 收到数据", data)

        fn(data.data)
    })

}

function js_translate(翻译平台, 内容, 源语言, 目标语言, fn) {
    jsondata = {
        "翻译平台": 翻译平台,
        "内容": 内容,
        "源语言": 源语言,
        "目标语言": 目标语言,
        "callfc": "js_translate" + 翻译平台,
    }
    console.log("js_translate" + 翻译平台 + " 发出数据", jsondata)
    wails.Events.Emit({
        name: "js_translate",
        data: JSON.stringify(jsondata)
    })
    wails.Events.Once("js_translate" + 翻译平台, function (data) {
        console.log("js_translate" + 翻译平台 + " 收到数据", data)
        jdata = JSON.parse(data.data)
        fn(jdata)
    })
}


tailwind.config = {
    theme: {
        extend: {
            colors: {
                clifford: '#da373d',
            }
        }
    },
}

window.runtime.EventsOn("setUrl", function (data) {
    console.log("setUrl", data)
    window.location.href = data

})