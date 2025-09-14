package html

import (
	"fmt"
)

func ExampleAbbrAttribute() {
 el := Th(AbbrAttribute("FYI"))
	fmt.Println(el.String())
	// Output: <th abbr="FYI"></th>
}

func ExampleAcceptCharset() {
 el := Form(AcceptCharset("utf-8"))
	fmt.Println(el.String())
	// Output: <form accept-charset="utf-8"></form>
}

func ExampleAction() {
 el := Form(Action("/submit"))
	fmt.Println(el.String())
	// Output: <form action="/submit"></form>
}

func ExampleAllow() {
 el := Iframe(Allow("fullscreen *"))
	fmt.Println(el.String())
	// Output: <iframe allow="fullscreen *"></iframe>
}

func ExampleAlt() {
	el := Img(Attributes(Src("pic.png"), Alt("Profile")))
	fmt.Println(el.String())
	// Output: <img src="pic.png" alt="Profile">
}

func ExampleAs() {
 el := Link(As("script"))
	fmt.Println(el.String())
	// Output: <link as="script">
}

func ExampleAutoComplete_text() {
 el := Form(AutoComplete("one-time-code"))
	fmt.Println(el.String())
	// Output: <form autocomplete="one-time-code"></form>
}

func ExampleCiteAttribute() {
 el := Blockquote(CiteAttribute("/quote-source"))
	fmt.Println(el.String())
	// Output: <blockquote cite="/quote-source"></blockquote>
}

func ExampleColor() {
	el := Link(Attributes(HRef("/app.css"), Color("#ff0000")))
	fmt.Println(el.String())
	// Output: <link href="/app.css" color="#ff0000">
}

func ExampleCommand() {
	el := Div(Attributes(Command("doit")))
	fmt.Println(el.String())
	// Output: <div command="doit"></div>
}

func ExampleCommandfor() {
	el := Div(Attributes(Commandfor("elem1")))
	fmt.Println(el.String())
	// Output: <div commandfor="elem1"></div>
}

func ExampleContent() {
 el := Meta(Content("text/html; charset=utf-8"))
	fmt.Println(el.String())
	// Output: <meta content="text/html; charset=utf-8">
}

func ExampleDataAttribute() {
 el := Object(DataAttribute("/movie.mp4"))
	fmt.Println(el.String())
	// Output: <object data="/movie.mp4"></object>
}

func ExampleDateTime() {
 el := Time(DateTime("2025-09-14"))
	fmt.Println(el.String())
	// Output: <time datetime="2025-09-14"></time>
}

func ExampleDirName() {
	el := Input(Attributes(TypeText(), DirName("comment.dir")))
	fmt.Println(el.String())
	// Output: <input type="text" dirname="comment.dir">
}

func ExampleDownload() {
 el := A(Attributes(HRef("/file.txt"), Download("file.txt")))
	fmt.Println(el.String())
	// Output: <a href="/file.txt" download="file.txt"></a>
}

func ExampleFor() {
 el := Label(For("input1"))
	fmt.Println(el.String())
	// Output: <label for="input1"></label>
}

func ExampleFormAttribute() {
	el := Input(Attributes(TypeText(), FormAttribute("myform")))
	fmt.Println(el.String())
	// Output: <input type="text" form="myform">
}

func ExampleFormAction() {
 el := Button(FormAction("/go"))
	fmt.Println(el.String())
	// Output: <button formaction="/go"></button>
}

func ExampleFormTarget() {
 el := Button(FormTarget("_blank"))
	fmt.Println(el.String())
	// Output: <button formtarget="_blank"></button>
}

func ExampleHRef() {
 el := A(HRef("#top"), Text("Top"))
	fmt.Println(el.String())
	// Output: <a href="#top">Top</a>
}

func ExampleHRefLang() {
	el := Link(Attributes(HRef("/app.js"), HRefLang("en")))
	fmt.Println(el.String())
	// Output: <link href="/app.js" hreflang="en">
}

func ExampleId() {
	el := Div(Id("main"))
	fmt.Println(el.String())
	// Output: <div id="main"></div>
}

func ExampleImageSizes() {
	el := Link(Attributes(Rel("preload"), ImageSizes("(max-width: 600px) 100vw, 600px")))
	fmt.Println(el.String())
	// Output: <link rel="preload" imagesizes="(max-width: 600px) 100vw, 600px">
}

func ExampleIntegrity() {
 el := Script(Integrity("sha256-xyz"))
	fmt.Println(el.String())
	// Output: <script integrity="sha256-xyz"></script>
}

func ExampleIs() {
 el := Div(Is("x-my-element"))
	fmt.Println(el.String())
	// Output: <div is="x-my-element"></div>
}

func ExampleItemId() {
 el := Div(ItemId("item-1"))
	fmt.Println(el.String())
	// Output: <div itemid="item-1"></div>
}

func ExampleLabelAttribute() {
 el := Track(LabelAttribute("English"))
	fmt.Println(el.String())
	// Output: <track label="English">
}

func ExampleLang() {
 el := Div(Lang("en"))
	fmt.Println(el.String())
	// Output: <div lang="en"></div>
}

func ExampleList() {
	el := Input(Attributes(TypeText(), List("datalist-id")))
	fmt.Println(el.String())
	// Output: <input type="text" list="datalist-id">
}

func ExampleMax_text() {
	el := Input(Attributes(TypeText(), Max("xyz")))
	fmt.Println(el.String())
	// Output: <input type="text" max="xyz">
}

func ExampleMedia() {
 el := Link(Media("(min-width: 800px)"))
	fmt.Println(el.String())
	// Output: <link media="(min-width: 800px)">
}

func ExampleMin_text() {
	el := Input(Attributes(TypeText(), Min("aaa")))
	fmt.Println(el.String())
	// Output: <input type="text" min="aaa">
}

func ExampleName() {
	el := Input(Attributes(TypeText(), Name("username")))
	fmt.Println(el.String())
	// Output: <input type="text" name="username">
}

func ExampleNonce() {
 el := Script(Nonce("abc123"))
	fmt.Println(el.String())
	// Output: <script nonce="abc123"></script>
}

func ExamplePattern() {
	el := Input(Attributes(TypeText(), Pattern("[A-Za-z]+")))
	fmt.Println(el.String())
	// Output: <input type="text" pattern="[A-Za-z]+">
}

func ExamplePlaceHolder() {
	el := Input(Attributes(TypeText(), PlaceHolder("Your name")))
	fmt.Println(el.String())
	// Output: <input type="text" placeholder="Your name">
}

func ExamplePopOverTarget() {
 el := Button(PopOverTarget("menu1"))
	fmt.Println(el.String())
	// Output: <button popovertarget="menu1"></button>
}

func ExamplePoster() {
 el := Video(Poster("preview.jpg"))
	fmt.Println(el.String())
	// Output: <video poster="preview.jpg"></video>
}

func ExampleReferrerPolicy() {
 el := Link(ReferrerPolicy("no-referrer"))
	fmt.Println(el.String())
	// Output: <link referrerpolicy="no-referrer">
}

func ExampleSizes_text() {
 el := Link(Sizes("(max-width: 600px) 100vw, 600px"))
	fmt.Println(el.String())
	// Output: <link sizes="(max-width: 600px) 100vw, 600px">
}

func ExampleSlotAttribute() {
 el := Div(SlotAttribute("slot1"))
	fmt.Println(el.String())
	// Output: <div slot="slot1"></div>
}

func ExampleSrc() {
	el := Img(Attributes(Src("pic.png"), Alt("Profile")))
	fmt.Println(el.String())
	// Output: <img src="pic.png" alt="Profile">
}

func ExampleSrcDoc() {
 el := Iframe(SrcDoc("<p>Hi</p>"))
	fmt.Println(el.String())
	// Output: <iframe srcdoc="<p>Hi</p>"></iframe>
}

func ExampleSrcLang() {
 el := Track(SrcLang("en"))
	fmt.Println(el.String())
	// Output: <track srclang="en">
}

func ExampleStyleAttribute() {
 el := Div(StyleAttribute("color:red"))
	fmt.Println(el.String())
	// Output: <div style="color:red"></div>
}

func ExampleTarget() {
	el := A(Attributes(HRef("/page"), Target("_blank")))
	fmt.Println(el.String())
	// Output: <a href="/page" target="_blank"></a>
}

func ExampleTitleAttribute() {
 el := Span(TitleAttribute("Hello"), Text("Hover"))
	fmt.Println(el.String())
	// Output: <span title="Hello">Hover</span>
}

func ExampleType_text() {
	el := Script(Attributes(Type("module")))
	fmt.Println(el.String())
	// Output: <script type="module"></script>
}

func ExampleUseMap() {
	el := Img(Attributes(UseMap("#m")))
	fmt.Println(el.String())
	// Output: <img usemap="#m">
}

func ExampleValue_text() {
	el := Input(Attributes(TypeText(), Value("abc")))
	fmt.Println(el.String())
	// Output: <input type="text" value="abc">
}
