(()=>{var ne=Object.defineProperty;var w=(e,t)=>{for(var n in t)ne(e,n,{get:t[n],enumerable:!0})};var C={};w(C,{SetText:()=>re,Text:()=>le});var ie=window.location.origin+"/wails/runtime";function oe(e,t,n){let i=new URL(ie);i.searchParams.append("method",e),n&&i.searchParams.append("args",JSON.stringify(n));let o={headers:{}};return t&&(o.headers["x-wails-window-name"]=t),new Promise((l,f)=>{fetch(i,o).then(a=>{if(a.ok)return a.headers.get("Content-Type")&&a.headers.get("Content-Type").indexOf("application/json")!==-1?a.json():a.text();f(Error(a.statusText))}).then(a=>l(a)).catch(a=>f(a))})}function r(e,t){return function(n,i=null){return oe(e+"."+n,t,i)}}var R=r("clipboard");function re(e){R("SetText",{text:e})}function le(){return R("Text")}var M={};w(M,{Hide:()=>ae,Quit:()=>ue,Show:()=>se});var b=r("application");function ae(){b("Hide")}function se(){b("Show")}function ue(){b("Quit")}var S={};w(S,{Log:()=>fe});var ce=r("log");function fe(e){return ce("Log",e)}var L={};w(L,{GetAll:()=>de,GetCurrent:()=>pe,GetPrimary:()=>me});var E=r("screens");function de(){return E("GetAll")}function me(){return E("GetPrimary")}function pe(){return E("GetCurrent")}var we="useandom-26T198340PX75pxJACKVERYMINDBUSHWOLF_GQZbfghjklqvwyzrict";var g=(e=21)=>{let t="",n=e;for(;n--;)t+=we[Math.random()*64|0];return t};var ge=r("call"),u=new Map;function xe(){let e;do e=g();while(u.has(e));return e}function A(e,t,n){let i=u.get(e);i&&(n?i.resolve(JSON.parse(t)):i.resolve(t),u.delete(e))}function T(e,t){let n=u.get(e);n&&(n.reject(t),u.delete(e))}function P(e,t){return new Promise((n,i)=>{let o=xe();t=t||{},t["call-id"]=o,u.set(o,{resolve:n,reject:i}),ge(e,t).catch(l=>{i(l),u.delete(o)})})}function D(e){return P("Call",e)}function y(e,t,...n){return P("Call",{packageName:"wails-plugins",structName:e,methodName:t,args:n})}function N(e){let t=r("window",e);return{Center:()=>void t("Center"),SetTitle:n=>void t("SetTitle",{title:n}),Fullscreen:()=>void t("Fullscreen"),UnFullscreen:()=>void t("UnFullscreen"),SetSize:(n,i)=>t("SetSize",{width:n,height:i}),Size:()=>t("Size"),SetMaxSize:(n,i)=>void t("SetMaxSize",{width:n,height:i}),SetMinSize:(n,i)=>void t("SetMinSize",{width:n,height:i}),SetAlwaysOnTop:n=>void t("SetAlwaysOnTop",{alwaysOnTop:n}),SetPosition:(n,i)=>t("SetPosition",{x:n,y:i}),Position:()=>t("Position"),Screen:()=>t("Screen"),Hide:()=>void t("Hide"),Maximise:()=>void t("Maximise"),Show:()=>void t("Show"),Close:()=>void t("Close"),ToggleMaximise:()=>void t("ToggleMaximise"),UnMaximise:()=>void t("UnMaximise"),Minimise:()=>void t("Minimise"),UnMinimise:()=>void t("UnMinimise"),Restore:()=>void t("Restore"),SetBackgroundColour:(n,i,o,l)=>void t("SetBackgroundColour",{r:n,g:i,b:o,a:l})}}var he=r("events"),k=class{constructor(t,n,i){this.eventName=t,this.maxCallbacks=i||-1,this.Callback=o=>(n(o),this.maxCallbacks===-1?!1:(this.maxCallbacks-=1,this.maxCallbacks===0))}},x=class{constructor(t,n=null){this.name=t,this.data=n}},s=new Map;function h(e,t,n){let i=s.get(e)||[],o=new k(e,t,n);return i.push(o),s.set(e,i),()=>ve(o)}function F(e,t){return h(e,t,-1)}function U(e,t){return h(e,t,1)}function ve(e){let t=e.eventName,n=s.get(t).filter(i=>i!==e);n.length===0?s.delete(t):s.set(t,n)}function z(e){console.log("dispatching event: ",{event:e});let t=s.get(e.name);if(t){let n=[];t.forEach(i=>{i.Callback(e)&&n.push(i)}),n.length>0&&(t=t.filter(i=>!n.includes(i)),t.length===0?s.delete(e.name):s.set(e.name,t))}}function G(e,...t){[e,...t].forEach(i=>{s.delete(i)})}function H(){s.clear()}function v(e){he("Emit",e)}var Ce=r("dialog"),c=new Map;function be(){let e;do e=g();while(c.has(e));return e}function I(e,t,n){let i=c.get(e);i&&(n?i.resolve(JSON.parse(t)):i.resolve(t),c.delete(e))}function B(e,t){let n=c.get(e);n&&(n.reject(t),c.delete(e))}function d(e,t){return new Promise((n,i)=>{let o=be();t=t||{},t["dialog-id"]=o,c.set(o,{resolve:n,reject:i}),Ce(e,t).catch(l=>{i(l),c.delete(o)})})}function Q(e){return d("Info",e)}function Y(e){return d("Warning",e)}function J(e){return d("Error",e)}function m(e){return d("Question",e)}function j(e){return d("OpenFile",e)}function q(e){return d("SaveFile",e)}var Me=r("contextmenu");function Se(e,t,n,i){return Me("OpenContextMenu",{id:e,x:t,y:n,data:i})}function V(e){e?window.addEventListener("contextmenu",X):window.removeEventListener("contextmenu",X)}function X(e){_(e.target,e)}function _(e,t){let n=e.getAttribute("data-contextmenu");if(n)t.preventDefault(),Se(n,t.clientX,t.clientY,e.getAttribute("data-contextmenu-data"));else{let i=e.parentElement;i&&_(i,t)}}function K(e,t=null){let n=new x(e,t);v(n)}function Ee(){document.querySelectorAll("[data-wml-event]").forEach(function(t){let n=t.getAttribute("data-wml-event"),i=t.getAttribute("data-wml-confirm"),o=t.getAttribute("data-wml-trigger")||"click",l=function(){if(i){m({Title:"Confirm",Message:i,Buttons:[{Label:"Yes"},{Label:"No",IsDefault:!0}]}).then(function(f){f!=="No"&&K(n)});return}K(n)};t.removeEventListener(o,l),t.addEventListener(o,l)})}function Z(e){wails.Window[e]===void 0&&console.log("Window method "+e+" not found"),wails.Window[e]()}function Le(){document.querySelectorAll("[data-wml-window]").forEach(function(t){let n=t.getAttribute("data-wml-window"),i=t.getAttribute("data-wml-confirm"),o=t.getAttribute("data-wml-trigger")||"click",l=function(){if(i){m({Title:"Confirm",Message:i,Buttons:[{Label:"Yes"},{Label:"No",IsDefault:!0}]}).then(function(f){f!=="No"&&Z(n)});return}Z(n)};t.removeEventListener(o,l),t.addEventListener(o,l)})}function O(){Ee(),Le()}var $=function(e){chrome.webview.postMessage(e)};var p=!1;function ke(e){if(window.wails.Capabilities.HasNativeDrag===!0)return!1;let t=window.getComputedStyle(e.target).getPropertyValue("app-region");return t&&(t=t.trim()),t!=="drag"||e.buttons!==1?!1:e.detail===1}function ee(){window.addEventListener("mousedown",Oe),window.addEventListener("mousemove",Re),window.addEventListener("mouseup",We)}function Oe(e){if(ke(e)){if(e.offsetX>e.target.clientWidth||e.offsetY>e.target.clientHeight)return;p=!0}else p=!1}function We(e){(e.buttons!==void 0?e.buttons:e.which)>0&&W()}function W(){document.body.style.cursor="default",p=!1}function Re(e){p&&(p=!1,(e.buttons!==void 0?e.buttons:e.which)>0&&$("drag"))}window.wails={...te(null),Capabilities:{}};fetch("/wails/capabilities").then(e=>{e.json().then(t=>{window.wails.Capabilities=t})});window._wails={dialogCallback:I,dialogErrorCallback:B,dispatchWailsEvent:z,callCallback:A,callErrorCallback:T,endDrag:W};function te(e){return{Clipboard:{...C},Application:{...M,GetWindowByName(t){return te(t)}},Log:S,Screens:L,Call:D,Plugin:y,WML:{Reload:O},Dialog:{Info:Q,Warning:Y,Error:J,Question:m,OpenFile:j,SaveFile:q},Events:{Emit:v,On:F,Once:U,OnMultiple:h,Off:G,OffAll:H},Window:N(e)}}console.log("Wails v3.0.0 Debug Mode Enabled");V(!0);ee();document.addEventListener("DOMContentLoaded",function(e){O()});})();