package html

import (
	"fmt"
)

// Window (body) event handlers
func ExampleOnAfterPrint() {
 el := Body(OnAfterPrint("handler()"))
	fmt.Println(el.String())
	// Output: <body onafterprint="handler()"></body>
}

func ExampleOnBeforePrint() {
 el := Body(OnBeforePrint("handler()"))
	fmt.Println(el.String())
	// Output: <body onbeforeprint="handler()"></body>
}

func ExampleOnBeforeUnload() {
	el := Body(Attributes(OnBeforeUnload("handler()")))
	fmt.Println(el.String())
	// Output: <body onbeforeunload="handler()"></body>
}

func ExampleOnHashChange() {
	el := Body(Attributes(OnHashChange("handler()")))
	fmt.Println(el.String())
	// Output: <body onhashchange="handler()"></body>
}

func ExampleOnLanguageChange() {
	el := Body(Attributes(OnLanguageChange("handler()")))
	fmt.Println(el.String())
	// Output: <body onlanguagechange="handler()"></body>
}

func ExampleOnMessage() {
	el := Body(Attributes(OnMessage("handler()")))
	fmt.Println(el.String())
	// Output: <body onmessage="handler()"></body>
}

func ExampleOnMessageError() {
	el := Body(Attributes(OnMessageError("handler()")))
	fmt.Println(el.String())
	// Output: <body onmessageerror="handler()"></body>
}

func ExampleOnOffline() {
	el := Body(Attributes(OnOffline("handler()")))
	fmt.Println(el.String())
	// Output: <body onoffline="handler()"></body>
}

func ExampleOnOnline() {
	el := Body(Attributes(OnOnline("handler()")))
	fmt.Println(el.String())
	// Output: <body ononline="handler()"></body>
}

func ExampleOnPageSwap() {
	el := Body(Attributes(OnPageSwap("handler()")))
	fmt.Println(el.String())
	// Output: <body onpageswap="handler()"></body>
}

func ExampleOnPageHide() {
	el := Body(Attributes(OnPageHide("handler()")))
	fmt.Println(el.String())
	// Output: <body onpagehide="handler()"></body>
}

func ExampleOnPageReveal() {
	el := Body(Attributes(OnPageReveal("handler()")))
	fmt.Println(el.String())
	// Output: <body onpagereveal="handler()"></body>
}

func ExampleOnPageShow() {
	el := Body(Attributes(OnPageShow("handler()")))
	fmt.Println(el.String())
	// Output: <body onpageshow="handler()"></body>
}

func ExampleOnPopState() {
	el := Body(Attributes(OnPopState("handler()")))
	fmt.Println(el.String())
	// Output: <body onpopstate="handler()"></body>
}

func ExampleOnRejectionHandled() {
	el := Body(Attributes(OnRejectionHandled("handler()")))
	fmt.Println(el.String())
	// Output: <body onrejectionhandled="handler()"></body>
}

func ExampleOnStorage() {
	el := Body(Attributes(OnStorage("handler()")))
	fmt.Println(el.String())
	// Output: <body onstorage="handler()"></body>
}

func ExampleOnUnhandledRejection() {
	el := Body(Attributes(OnUnhandledRejection("handler()")))
	fmt.Println(el.String())
	// Output: <body onunhandledrejection="handler()"></body>
}

func ExampleOnUnload() {
	el := Body(Attributes(OnUnload("handler()")))
	fmt.Println(el.String())
	// Output: <body onunload="handler()"></body>
}

// Media event handlers (use <video>)
func ExampleOnCanPlay() {
	el := Video(Attributes(OnCanPlay("handler()")))
	fmt.Println(el.String())
	// Output: <video oncanplay="handler()"></video>
}

func ExampleOnCanplayThrough() {
	el := Video(Attributes(OnCanplayThrough("handler()")))
	fmt.Println(el.String())
	// Output: <video oncanplaythrough="handler()"></video>
}

func ExampleOnCueChange() {
	el := Track(Attributes(OnCueChange("handler()")))
	fmt.Println(el.String())
	// Output: <track oncuechange="handler()">
}

func ExampleOnDurationChange() {
	el := Video(Attributes(OnDurationChange("handler()")))
	fmt.Println(el.String())
	// Output: <video ondurationchange="handler()"></video>
}

func ExampleOnEmptied() {
	el := Video(Attributes(OnEmptied("handler()")))
	fmt.Println(el.String())
	// Output: <video onemptied="handler()"></video>
}

func ExampleOnEnded() {
	el := Video(Attributes(OnEnded("handler()")))
	fmt.Println(el.String())
	// Output: <video onended="handler()"></video>
}

func ExampleOnError() {
	el := Video(Attributes(OnError("handler()")))
	fmt.Println(el.String())
	// Output: <video onerror="handler()"></video>
}

func ExampleOnLoadedData() {
	el := Video(Attributes(OnLoadedData("handler()")))
	fmt.Println(el.String())
	// Output: <video onloadeddata="handler()"></video>
}

func ExampleOnLoadedMetaData() {
	el := Video(Attributes(OnLoadedMetaData("handler()")))
	fmt.Println(el.String())
	// Output: <video onloadedmetadata="handler()"></video>
}

func ExampleOnLoadStart() {
	el := Video(Attributes(OnLoadStart("handler()")))
	fmt.Println(el.String())
	// Output: <video onloadstart="handler()"></video>
}

func ExampleOnPause() {
	el := Video(Attributes(OnPause("handler()")))
	fmt.Println(el.String())
	// Output: <video onpause="handler()"></video>
}

func ExampleOnPlay() {
	el := Video(Attributes(OnPlay("handler()")))
	fmt.Println(el.String())
	// Output: <video onplay="handler()"></video>
}

func ExampleOnPlaying() {
	el := Video(Attributes(OnPlaying("handler()")))
	fmt.Println(el.String())
	// Output: <video onplaying="handler()"></video>
}

func ExampleOnProgress() {
	el := Video(Attributes(OnProgress("handler()")))
	fmt.Println(el.String())
	// Output: <video onprogress="handler()"></video>
}

func ExampleOnRatechange() {
	el := Video(Attributes(OnRatechange("handler()")))
	fmt.Println(el.String())
	// Output: <video onratechange="handler()"></video>
}

func ExampleOnSeeked() {
	el := Video(Attributes(OnSeeked("handler()")))
	fmt.Println(el.String())
	// Output: <video onseeked="handler()"></video>
}

func ExampleOnSeeking() {
	el := Video(Attributes(OnSeeking("handler()")))
	fmt.Println(el.String())
	// Output: <video onseeking="handler()"></video>
}

func ExampleOnStalled() {
	el := Video(Attributes(OnStalled("handler()")))
	fmt.Println(el.String())
	// Output: <video onstalled="handler()"></video>
}

func ExampleOnSuspend() {
	el := Video(Attributes(OnSuspend("handler()")))
	fmt.Println(el.String())
	// Output: <video onsuspend="handler()"></video>
}

func ExampleOnTimeUpdate() {
	el := Video(Attributes(OnTimeUpdate("handler()")))
	fmt.Println(el.String())
	// Output: <video ontimeupdate="handler()"></video>
}

func ExampleOnVolumeChange() {
	el := Video(Attributes(OnVolumeChange("handler()")))
	fmt.Println(el.String())
	// Output: <video onvolumechange="handler()"></video>
}

func ExampleOnWaiting() {
	el := Video(Attributes(OnWaiting("handler()")))
	fmt.Println(el.String())
	// Output: <video onwaiting="handler()"></video>
}

// Form-related handlers
func ExampleOnFormData() {
	el := Form(Attributes(OnFormData("handler()")))
	fmt.Println(el.String())
	// Output: <form onformdata="handler()"></form>
}

func ExampleOnReset() {
	el := Form(Attributes(OnReset("handler()")))
	fmt.Println(el.String())
	// Output: <form onreset="handler()"></form>
}

func ExampleOnSubmit() {
	el := Form(Attributes(OnSubmit("handler()")))
	fmt.Println(el.String())
	// Output: <form onsubmit="handler()"></form>
}

// Slot handler
func ExampleOnSlotChange() {
	el := Slot(Attributes(OnSlotChange("handler()")))
	fmt.Println(el.String())
	// Output: <slot onslotchange="handler()"></slot>
}

// Canvas/WebGL context handlers
func ExampleOnContextLost() {
	el := Canvas(Attributes(OnContextLost("handler()")))
	fmt.Println(el.String())
	// Output: <canvas oncontextlost="handler()"></canvas>
}

func ExampleOnContextRestored() {
	el := Canvas(Attributes(OnContextRestored("handler()")))
	fmt.Println(el.String())
	// Output: <canvas oncontextrestored="handler()"></canvas>
}

// Generic element handlers (use <div>)
func ExampleOnAuxClick() {
	el := Div(Attributes(OnAuxClick("handler()")))
	fmt.Println(el.String())
	// Output: <div onauxclick="handler()"></div>
}

func ExampleOnBeforeInput() {
	el := Div(Attributes(OnBeforeInput("handler()")))
	fmt.Println(el.String())
	// Output: <div onbeforeinput="handler()"></div>
}

func ExampleOnBeforeMatch() {
	el := Div(Attributes(OnBeforeMatch("handler()")))
	fmt.Println(el.String())
	// Output: <div onbeforematch="handler()"></div>
}

func ExampleOnBeforeToggle() {
	el := Div(Attributes(OnBeforeToggle("handler()")))
	fmt.Println(el.String())
	// Output: <div onbeforetoggle="handler()"></div>
}

func ExampleOnBlur() {
	el := Div(Attributes(OnBlur("handler()")))
	fmt.Println(el.String())
	// Output: <div onblur="handler()"></div>
}

func ExampleOnCancel() {
	el := Div(Attributes(OnCancel("handler()")))
	fmt.Println(el.String())
	// Output: <div oncancel="handler()"></div>
}

func ExampleOnChange() {
	el := Div(Attributes(OnChange("handler()")))
	fmt.Println(el.String())
	// Output: <div onchange="handler()"></div>
}

func ExampleOnClick() {
	el := Button(Attributes(OnClick("alert('hi')")), Text("Click"))
	fmt.Println(el.String())
	// Output: <button onclick="alert('hi')">Click</button>
}

func ExampleOnClose() {
	el := Div(Attributes(OnClose("handler()")))
	fmt.Println(el.String())
	// Output: <div onclose="handler()"></div>
}

func ExampleOnCommand() {
	el := Div(Attributes(OnCommand("handler()")))
	fmt.Println(el.String())
	// Output: <div oncommand="handler()"></div>
}

func ExampleOnContextMenu() {
	el := Div(Attributes(OnContextMenu("handler()")))
	fmt.Println(el.String())
	// Output: <div oncontextmenu="handler()"></div>
}

func ExampleOnCopy() {
	el := Div(Attributes(OnCopy("handler()")))
	fmt.Println(el.String())
	// Output: <div oncopy="handler()"></div>
}

func ExampleOnCut() {
	el := Div(Attributes(OnCut("handler()")))
	fmt.Println(el.String())
	// Output: <div oncut="handler()"></div>
}

func ExampleOnDblClick() {
	el := Div(Attributes(OnDblClick("handler()")))
	fmt.Println(el.String())
	// Output: <div ondblclick="handler()"></div>
}

func ExampleOnDrag() {
	el := Div(Attributes(OnDrag("handler()")))
	fmt.Println(el.String())
	// Output: <div ondrag="handler()"></div>
}

func ExampleOnDragEnd() {
	el := Div(Attributes(OnDragEnd("handler()")))
	fmt.Println(el.String())
	// Output: <div ondragend="handler()"></div>
}

func ExampleOnDragEnter() {
	el := Div(Attributes(OnDragEnter("handler()")))
	fmt.Println(el.String())
	// Output: <div ondragenter="handler()"></div>
}

func ExampleOnDragLeave() {
	el := Div(Attributes(OnDragLeave("handler()")))
	fmt.Println(el.String())
	// Output: <div ondragleave="handler()"></div>
}

func ExampleOnDragOver() {
	el := Div(Attributes(OnDragOver("handler()")))
	fmt.Println(el.String())
	// Output: <div ondragover="handler()"></div>
}

func ExampleOnDragStart() {
	el := Div(Attributes(OnDragStart("handler()")))
	fmt.Println(el.String())
	// Output: <div ondragstart="handler()"></div>
}

func ExampleOnDrop() {
	el := Div(Attributes(OnDrop("handler()")))
	fmt.Println(el.String())
	// Output: <div ondrop="handler()"></div>
}

func ExampleOnFocus() {
	el := Div(Attributes(OnFocus("handler()")))
	fmt.Println(el.String())
	// Output: <div onfocus="handler()"></div>
}

func ExampleOnInput() {
	el := Div(Attributes(OnInput("handler()")))
	fmt.Println(el.String())
	// Output: <div oninput="handler()"></div>
}

func ExampleOnInvalid() {
	el := Div(Attributes(OnInvalid("handler()")))
	fmt.Println(el.String())
	// Output: <div oninvalid="handler()"></div>
}

func ExampleOnKeyDown() {
	el := Div(Attributes(OnKeyDown("handler()")))
	fmt.Println(el.String())
	// Output: <div onkeydown="handler()"></div>
}

func ExampleOnKeyPress() {
	el := Div(Attributes(OnKeyPress("handler()")))
	fmt.Println(el.String())
	// Output: <div onkeypress="handler()"></div>
}

func ExampleOnKeyUp() {
	el := Div(Attributes(OnKeyUp("handler()")))
	fmt.Println(el.String())
	// Output: <div onkeyup="handler()"></div>
}

func ExampleOnLoad() {
	el := Div(Attributes(OnLoad("handler()")))
	fmt.Println(el.String())
	// Output: <div onload="handler()"></div>
}

func ExampleOnMouseDown() {
	el := Div(Attributes(OnMouseDown("handler()")))
	fmt.Println(el.String())
	// Output: <div onmousedown="handler()"></div>
}

func ExampleOnMouseEnter() {
	el := Div(Attributes(OnMouseEnter("handler()")))
	fmt.Println(el.String())
	// Output: <div onmouseenter="handler()"></div>
}

func ExampleOnMouseLeave() {
	el := Div(Attributes(OnMouseLeave("handler()")))
	fmt.Println(el.String())
	// Output: <div onmouseleave="handler()"></div>
}

func ExampleOnMouseMove() {
	el := Div(Attributes(OnMouseMove("handler()")))
	fmt.Println(el.String())
	// Output: <div onmousemove="handler()"></div>
}

func ExampleOnMouseOut() {
	el := Div(Attributes(OnMouseOut("handler()")))
	fmt.Println(el.String())
	// Output: <div onmouseout="handler()"></div>
}

func ExampleOnMouseOver() {
	el := Div(Attributes(OnMouseOver("handler()")))
	fmt.Println(el.String())
	// Output: <div onmouseover="handler()"></div>
}

func ExampleOnMouseUp() {
	el := Div(Attributes(OnMouseUp("handler()")))
	fmt.Println(el.String())
	// Output: <div onmouseup="handler()"></div>
}

func ExampleOnPaste() {
	el := Div(Attributes(OnPaste("handler()")))
	fmt.Println(el.String())
	// Output: <div onpaste="handler()"></div>
}

func ExampleOnResize() {
	el := Div(Attributes(OnResize("handler()")))
	fmt.Println(el.String())
	// Output: <div onresize="handler()"></div>
}

func ExampleOnScroll() {
	el := Div(Attributes(OnScroll("handler()")))
	fmt.Println(el.String())
	// Output: <div onscroll="handler()"></div>
}

func ExampleOnScrollEnd() {
	el := Div(Attributes(OnScrollEnd("handler()")))
	fmt.Println(el.String())
	// Output: <div onscrollend="handler()"></div>
}

func ExampleOnSecurityPolicyViolation() {
	el := Div(Attributes(OnSecurityPolicyViolation("handler()")))
	fmt.Println(el.String())
	// Output: <div onsecuritypolicyviolation="handler()"></div>
}

func ExampleOnSelect() {
	el := Div(Attributes(OnSelect("handler()")))
	fmt.Println(el.String())
	// Output: <div onselect="handler()"></div>
}

func ExampleOnToggle() {
	el := Details(Attributes(OnToggle("handler()")))
	fmt.Println(el.String())
	// Output: <details ontoggle="handler()"></details>
}

func ExampleOnWheel() {
	el := Div(Attributes(OnWheel("handler()")))
	fmt.Println(el.String())
	// Output: <div onwheel="handler()"></div>
}
