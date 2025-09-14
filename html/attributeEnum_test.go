package html

import (
	"fmt"
)

// AutoCapitalize
func ExampleAutoCapitalizeCharacters() {
	el := Input(Attributes(TypeText(), AutoCapitalizeCharacters()))
	fmt.Println(el.String())
	// Output: <input type="text" autocapitalize="characters">
}

func ExampleAutoCapitalizeNone() {
	el := Input(Attributes(TypeText(), AutoCapitalizeNone()))
	fmt.Println(el.String())
	// Output: <input type="text" autocapitalize="none">
}

func ExampleAutoCapitalizeOff() {
	el := Input(Attributes(TypeText(), AutoCapitalizeOff()))
	fmt.Println(el.String())
	// Output: <input type="text" autocapitalize="off">
}

func ExampleAutoCapitalizeOn() {
	el := Input(Attributes(TypeText(), AutoCapitalizeOn()))
	fmt.Println(el.String())
	// Output: <input type="text" autocapitalize="on">
}

func ExampleAutoCapitalizeSentences() {
	el := Input(Attributes(TypeText(), AutoCapitalizeSentences()))
	fmt.Println(el.String())
	// Output: <input type="text" autocapitalize="sentences">
}

func ExampleAutoCapitalizeWords() {
	el := Input(Attributes(TypeText(), AutoCapitalizeWords()))
	fmt.Println(el.String())
	// Output: <input type="text" autocapitalize="words">
}

// AutoComplete
func ExampleAutoCompleteOff() {
	el := Input(Attributes(TypeText(), AutoCompleteOff()))
	fmt.Println(el.String())
	// Output: <input type="text" autocomplete="off">
}

func ExampleAutoCompleteOn() {
	el := Input(Attributes(TypeText(), AutoCompleteOn()))
	fmt.Println(el.String())
	// Output: <input type="text" autocomplete="on">
}

// AutoCorrect
func ExampleAutoCorrectOff() {
	el := Input(Attributes(TypeText(), AutoCorrectOff()))
	fmt.Println(el.String())
	// Output: <input type="text" autocorrect="off">
}

func ExampleAutoCorrectOn() {
	el := Input(Attributes(TypeText(), AutoCorrectOn()))
	fmt.Println(el.String())
	// Output: <input type="text" autocorrect="on">
}

// CharSet
func ExampleCharSetUtf8() {
	el := Meta(Attributes(CharSetUtf8()))
	fmt.Println(el.String())
	// Output: <meta charset="utf-8">
}

// Dialog closedby
func ExampleClosedByAny() {
	el := Dialog(Attributes(ClosedByAny()))
	fmt.Println(el.String())
	// Output: <dialog closedby="any"></dialog>
}

func ExampleClosedByCloserequest() {
	el := Dialog(Attributes(ClosedByCloserequest()))
	fmt.Println(el.String())
	// Output: <dialog closedby="closerequest"></dialog>
}

func ExampleClosedByNone() {
	el := Dialog(Attributes(ClosedByNone()))
	fmt.Println(el.String())
	// Output: <dialog closedby="none"></dialog>
}

// Color space for input color
func ExampleColorSpaceDisplayP3() {
	el := Input(Attributes(TypeColor(), ColorSpaceDisplayP3()))
	fmt.Println(el.String())
	// Output: <input type="color" colorspace="display-p3">
}

func ExampleColorSpaceLimitedSrgb() {
	el := Input(Attributes(TypeColor(), ColorSpaceLimitedSrgb()))
	fmt.Println(el.String())
	// Output: <input type="color" colorspace="limited-srgb">
}

// ContentEditable
func ExampleContentEditableTrue() {
	el := Div(ContentEditableTrue())
	fmt.Println(el.String())
	// Output: <div contenteditable="true"></div>
}

func ExampleContentEditablePlaintextOnly() {
	el := Div(ContentEditablePlaintextOnly())
	fmt.Println(el.String())
	// Output: <div contenteditable="plaintext-only"></div>
}

func ExampleContentEditableFalse() {
	el := Div(ContentEditableFalse())
	fmt.Println(el.String())
	// Output: <div contenteditable="false"></div>
}

// CrossOrigin
func ExampleCrossOriginAnonymous() {
	el := Img(Attributes(CrossOriginAnonymous(), Src("image.png")))
	fmt.Println(el.String())
	// Output: <img crossorigin="anonymous" src="image.png">
}

func ExampleCrossOriginUseCredentials() {
	el := Img(Attributes(CrossOriginUseCredentials(), Src("image.png")))
	fmt.Println(el.String())
	// Output: <img crossorigin="use-credentials" src="image.png">
}

// Decoding
func ExampleDecodingAsync() {
	el := Img(Attributes(DecodingAsync(), Src("image.png")))
	fmt.Println(el.String())
	// Output: <img decoding="async" src="image.png">
}

func ExampleDecodingAuto() {
	el := Img(Attributes(DecodingAuto(), Src("image.png")))
	fmt.Println(el.String())
	// Output: <img decoding="auto" src="image.png">
}

func ExampleDecodingSync() {
	el := Img(Attributes(DecodingSync(), Src("image.png")))
	fmt.Println(el.String())
	// Output: <img decoding="sync" src="image.png">
}

// Dir
func ExampleDirAuto() {
	el := Div(Attributes(DirAuto()))
	fmt.Println(el.String())
	// Output: <div dir="auto"></div>
}

func ExampleDirLtr() {
	el := Div(Attributes(DirLtr()))
	fmt.Println(el.String())
	// Output: <div dir="ltr"></div>
}

func ExampleDirRtl() {
	el := Div(Attributes(DirRtl()))
	fmt.Println(el.String())
	// Output: <div dir="rtl"></div>
}

// Draggable
func ExampleDraggableTrue() {
	el := Div(Attributes(DraggableTrue()))
	fmt.Println(el.String())
	// Output: <div draggable="true"></div>
}

func ExampleDraggableFalse() {
	el := Div(Attributes(DraggableFalse()))
	fmt.Println(el.String())
	// Output: <div draggable="false"></div>
}

// Enctype on form
func ExampleEncTypeApplicationXWwwFormUrlencoded() {
	el := Form(Attributes(EncTypeApplicationXWwwFormUrlencoded()))
	fmt.Println(el.String())
	// Output: <form enctype="application/x-www-form-urlencoded"></form>
}

func ExampleEncTypeMultipartFormData() {
	el := Form(Attributes(EncTypeMultipartFormData()))
	fmt.Println(el.String())
	// Output: <form enctype="multipart/form-data"></form>
}

func ExampleEncTypeTextPlain() {
	el := Form(Attributes(EncTypeTextPlain()))
	fmt.Println(el.String())
	// Output: <form enctype="text/plain"></form>
}

// EnterKeyHint
func ExampleEnterKeyHintDone() {
	el := Input(Attributes(TypeText(), EnterKeyHintDone()))
	fmt.Println(el.String())
	// Output: <input type="text" enterkeyhint="done">
}

func ExampleEnterKeyHintEnter() {
	el := Input(Attributes(TypeText(), EnterKeyHintEnter()))
	fmt.Println(el.String())
	// Output: <input type="text" enterkeyhint="enter">
}

func ExampleEnterKeyHintGo() {
	el := Input(Attributes(TypeText(), EnterKeyHintGo()))
	fmt.Println(el.String())
	// Output: <input type="text" enterkeyhint="go">
}

func ExampleEnterKeyHintNext() {
	el := Input(Attributes(TypeText(), EnterKeyHintNext()))
	fmt.Println(el.String())
	// Output: <input type="text" enterkeyhint="next">
}

func ExampleEnterKeyHintPrevious() {
	el := Input(Attributes(TypeText(), EnterKeyHintPrevious()))
	fmt.Println(el.String())
	// Output: <input type="text" enterkeyhint="previous">
}

func ExampleEnterKeyHintSearch() {
	el := Input(Attributes(TypeText(), EnterKeyHintSearch()))
	fmt.Println(el.String())
	// Output: <input type="text" enterkeyhint="search">
}

func ExampleEnterKeyHintSend() {
	el := Input(Attributes(TypeText(), EnterKeyHintSend()))
	fmt.Println(el.String())
	// Output: <input type="text" enterkeyhint="send">
}

// FetchPriority
func ExampleFetchPriorityAuto() {
	el := Link(Attributes(HRef("/app.js"), Rel("preload"), FetchPriorityAuto()))
	fmt.Println(el.String())
	// Output: <link href="/app.js" rel="preload" fetchpriority="auto">
}

func ExampleFetchPriorityHigh() {
	el := Link(Attributes(HRef("/app.js"), Rel("preload"), FetchPriorityHigh()))
	fmt.Println(el.String())
	// Output: <link href="/app.js" rel="preload" fetchpriority="high">
}

func ExampleFetchPriorityLow() {
	el := Link(Attributes(HRef("/app.js"), Rel("preload"), FetchPriorityLow()))
	fmt.Println(el.String())
	// Output: <link href="/app.js" rel="preload" fetchpriority="low">
}

// Form-associated overrides
func ExampleFormEnctypeApplicationXWwwFormUrlencoded() {
	el := Button(Attributes(FormEnctypeApplicationXWwwFormUrlencoded()))
	fmt.Println(el.String())
	// Output: <button formenctype="application/x-www-form-urlencoded"></button>
}

func ExampleFormEnctypeMultipartFormData() {
	el := Button(Attributes(FormEnctypeMultipartFormData()))
	fmt.Println(el.String())
	// Output: <button formenctype="multipart/form-data"></button>
}

func ExampleFormEnctypeTextPlain() {
	el := Button(Attributes(FormEnctypeTextPlain()))
	fmt.Println(el.String())
	// Output: <button formenctype="text/plain"></button>
}

func ExampleFormMethodDialog() {
	el := Button(Attributes(FormMethodDialog()))
	fmt.Println(el.String())
	// Output: <button formmethod="dialog"></button>
}

func ExampleFormMethodGET() {
	el := Button(Attributes(FormMethodGET()))
	fmt.Println(el.String())
	// Output: <button formmethod="GET"></button>
}

func ExampleFormMethodPOST() {
	el := Button(Attributes(FormMethodPOST()))
	fmt.Println(el.String())
	// Output: <button formmethod="POST"></button>
}

// Hidden
func ExampleHiddenEmpty() {
	el := Div(Attributes(HiddenEmpty()))
	fmt.Println(el.String())
	// Output: <div></div>
}

func ExampleHiddenHidden() {
	el := Div(Attributes(HiddenHidden()))
	fmt.Println(el.String())
	// Output: <div hidden="hidden"></div>
}

func ExampleHiddenUntilFound() {
	el := Div(Attributes(HiddenUntilFound()))
	fmt.Println(el.String())
	// Output: <div hidden="until-found"></div>
}

// http-equiv
func ExampleHttpEquivContentSecurityPolicy() {
	el := Meta(Attributes(HttpEquivContentSecurityPolicy()))
	fmt.Println(el.String())
	// Output: <meta http-equiv="content-security-policy">
}

func ExampleHttpEquivContentType() {
	el := Meta(Attributes(HttpEquivContentType()))
	fmt.Println(el.String())
	// Output: <meta http-equiv="content-type">
}

func ExampleHttpEquivDefaultStyle() {
	el := Meta(Attributes(HttpEquivDefaultStyle()))
	fmt.Println(el.String())
	// Output: <meta http-equiv="default-style">
}

func ExampleHttpEquivRefresh() {
	el := Meta(Attributes(HttpEquivRefresh()))
	fmt.Println(el.String())
	// Output: <meta http-equiv="refresh">
}

func ExampleHttpEquivXUaCompatible() {
	el := Meta(Attributes(HttpEquivXUaCompatible()))
	fmt.Println(el.String())
	// Output: <meta http-equiv="x-ua-compatible">
}

// InputMode
func ExampleInputModeDecimal() {
	el := Input(Attributes(TypeText(), InputModeDecimal()))
	fmt.Println(el.String())
	// Output: <input type="text" inputmode="decimal">
}

func ExampleInputModeEmail() {
	el := Input(Attributes(TypeText(), InputModeEmail()))
	fmt.Println(el.String())
	// Output: <input type="text" inputmode="email">
}

func ExampleInputModeNone() {
	el := Input(Attributes(TypeText(), InputModeNone()))
	fmt.Println(el.String())
	// Output: <input type="text" inputmode="none">
}

func ExampleInputModeNumeric() {
	el := Input(Attributes(TypeText(), InputModeNumeric()))
	fmt.Println(el.String())
	// Output: <input type="text" inputmode="numeric">
}

func ExampleInputModeSearch() {
	el := Input(Attributes(TypeText(), InputModeSearch()))
	fmt.Println(el.String())
	// Output: <input type="text" inputmode="search">
}

func ExampleInputModeTel() {
	el := Input(Attributes(TypeText(), InputModeTel()))
	fmt.Println(el.String())
	// Output: <input type="text" inputmode="tel">
}

func ExampleInputModeText() {
	el := Input(Attributes(TypeText(), InputModeText()))
	fmt.Println(el.String())
	// Output: <input type="text" inputmode="text">
}

func ExampleInputModeUrl() {
	el := Input(Attributes(TypeText(), InputModeUrl()))
	fmt.Println(el.String())
	// Output: <input type="text" inputmode="url">
}

// Kind
func ExampleKindCaptions() {
	el := Track(Attributes(KindCaptions()))
	fmt.Println(el.String())
	// Output: <track kind="captions">
}

func ExampleKindChapters() {
	el := Track(Attributes(KindChapters()))
	fmt.Println(el.String())
	// Output: <track kind="chapters">
}

func ExampleKindDescriptions() {
	el := Track(Attributes(KindDescriptions()))
	fmt.Println(el.String())
	// Output: <track kind="descriptions">
}

func ExampleKindMetadata() {
	el := Track(Attributes(KindMetadata()))
	fmt.Println(el.String())
	// Output: <track kind="metadata">
}

func ExampleKindSubtitles() {
	el := Track(Attributes(KindSubtitles()))
	fmt.Println(el.String())
	// Output: <track kind="subtitles">
}

// Loading
func ExampleLoadingEager() {
	el := Img(Attributes(LoadingEager(), Src("image.png")))
	fmt.Println(el.String())
	// Output: <img loading="eager" src="image.png">
}

func ExampleLoadingLazy() {
	el := Img(Attributes(LoadingLazy(), Src("image.png")))
	fmt.Println(el.String())
	// Output: <img loading="lazy" src="image.png">
}

// Method on form
func ExampleMethodDialog() {
	el := Form(Attributes(MethodDialog()))
	fmt.Println(el.String())
	// Output: <form method="dialog"></form>
}

func ExampleMethodGET() {
	el := Form(Attributes(MethodGET()))
	fmt.Println(el.String())
	// Output: <form method="GET"></form>
}

func ExampleMethodPOST() {
	el := Form(Attributes(MethodPOST()))
	fmt.Println(el.String())
	// Output: <form method="POST"></form>
}

// Popover
func ExamplePopOverAuto() {
	el := Div(Attributes(PopOverAuto()))
	fmt.Println(el.String())
	// Output: <div popover="auto"></div>
}

func ExamplePopOverHint() {
	el := Div(Attributes(PopOverHint()))
	fmt.Println(el.String())
	// Output: <div popover="hint"></div>
}

func ExamplePopOverManual() {
	el := Div(Attributes(PopOverManual()))
	fmt.Println(el.String())
	// Output: <div popover="manual"></div>
}

// PopoverTargetAction
func ExamplePopOverTargetActionHide() {
	el := Button(Attributes(PopOverTargetActionHide()))
	fmt.Println(el.String())
	// Output: <button popovertargetaction="hide"></button>
}

func ExamplePopOverTargetActionShow() {
	el := Button(Attributes(PopOverTargetActionShow()))
	fmt.Println(el.String())
	// Output: <button popovertargetaction="show"></button>
}

func ExamplePopOverTargetActionToggle() {
	el := Button(Attributes(PopOverTargetActionToggle()))
	fmt.Println(el.String())
	// Output: <button popovertargetaction="toggle"></button>
}

// Preload
func ExamplePreLoadAuto() {
	el := Video(Attributes(PreLoadAuto()))
	fmt.Println(el.String())
	// Output: <video preload="auto"></video>
}

func ExamplePreLoadMetadata() {
	el := Video(Attributes(PreLoadMetadata()))
	fmt.Println(el.String())
	// Output: <video preload="metadata"></video>
}

func ExamplePreLoadNone() {
	el := Video(Attributes(PreLoadNone()))
	fmt.Println(el.String())
	// Output: <video preload="none"></video>
}

// Scope
func ExampleScopeCol() {
	el := Th(Attributes(ScopeCol()))
	fmt.Println(el.String())
	// Output: <th scope="col"></th>
}

func ExampleScopeColgroup() {
	el := Th(Attributes(ScopeColgroup()))
	fmt.Println(el.String())
	// Output: <th scope="colgroup"></th>
}

func ExampleScopeRow() {
	el := Th(Attributes(ScopeRow()))
	fmt.Println(el.String())
	// Output: <th scope="row"></th>
}

func ExampleScopeRowgroup() {
	el := Th(Attributes(ScopeRowgroup()))
	fmt.Println(el.String())
	// Output: <th scope="rowgroup"></th>
}

// ShadowRootMode
func ExampleShadowRootModeClosed() {
	el := Template(Attributes(ShadowRootModeClosed()))
	fmt.Println(el.String())
	// Output: <template shadowrootmode="closed"></template>
}

func ExampleShadowRootModeOpen() {
	el := Template(Attributes(ShadowRootModeOpen()))
	fmt.Println(el.String())
	// Output: <template shadowrootmode="open"></template>
}

// Shape
func ExampleShapeCircle() {
	el := Area(Attributes(ShapeCircle()))
	fmt.Println(el.String())
	// Output: <area shape="circle">
}

func ExampleShapeDefault() {
	el := Area(Attributes(ShapeDefault()))
	fmt.Println(el.String())
	// Output: <area shape="default">
}

func ExampleShapePoly() {
	el := Area(Attributes(ShapePoly()))
	fmt.Println(el.String())
	// Output: <area shape="poly">
}

func ExampleShapeRect() {
	el := Area(Attributes(ShapeRect()))
	fmt.Println(el.String())
	// Output: <area shape="rect">
}

// SpellCheck
func ExampleSpellCheckEmpty() {
	el := Div(Attributes(SpellCheckEmpty()))
	fmt.Println(el.String())
	// Output: <div></div>
}

func ExampleSpellCheckFalse() {
	el := Div(Attributes(SpellCheckFalse()))
	fmt.Println(el.String())
	// Output: <div spellcheck="false"></div>
}

func ExampleSpellCheckTrue() {
	// TODO this makes no sense at the moment, have to think about how to handle it. Currently attributes with empty string are not printed at all

	el := Div(Attributes(SpellCheckTrue()))
	fmt.Println(el.String())
	// Output: <div spellcheck="true"></div>
}

// Translate
func ExampleTranslateNo() {
	el := Div(Attributes(TranslateNo()))
	fmt.Println(el.String())
	// Output: <div translate="no"></div>
}

func ExampleTranslateYes() {
	el := Div(Attributes(TranslateYes()))
	fmt.Println(el.String())
	// Output: <div translate="yes"></div>
}

// Type for ol
func ExampleTypeAlpha() {
	el := Ol(Attributes(TypeAlpha()))
	fmt.Println(el.String())
	// Output: <ol type="A"></ol>
}

func ExampleTypeAlphaLower() {
	el := Ol(Attributes(TypeAlphaLower()))
	fmt.Println(el.String())
	// Output: <ol type="a"></ol>
}

func ExampleTypeRoman() {
	el := Ol(Attributes(TypeRoman()))
	fmt.Println(el.String())
	// Output: <ol type="I"></ol>
}

func ExampleTypeRomanLower() {
	el := Ol(Attributes(TypeRomanLower()))
	fmt.Println(el.String())
	// Output: <ol type="i"></ol>
}

// Type for button/input
func ExampleTypeButton() {
	el := Button(Attributes(TypeButton()))
	fmt.Println(el.String())
	// Output: <button type="button"></button>
}

func ExampleTypeCheckbox() {
	el := Input(Attributes(TypeCheckbox()))
	fmt.Println(el.String())
	// Output: <input type="checkbox">
}

func ExampleTypeColor() {
	el := Input(Attributes(TypeColor()))
	fmt.Println(el.String())
	// Output: <input type="color">
}

func ExampleTypeDate() {
	el := Input(Attributes(TypeDate()))
	fmt.Println(el.String())
	// Output: <input type="date">
}

func ExampleTypeDatetimeLocal() {
	el := Input(Attributes(TypeDatetimeLocal()))
	fmt.Println(el.String())
	// Output: <input type="datetime-local">
}

func ExampleTypeEmail() {
	el := Input(Attributes(TypeEmail()))
	fmt.Println(el.String())
	// Output: <input type="email">
}

func ExampleTypeFile() {
	el := Input(Attributes(TypeFile()))
	fmt.Println(el.String())
	// Output: <input type="file">
}

func ExampleTypeHidden() {
	el := Input(Attributes(TypeHidden()))
	fmt.Println(el.String())
	// Output: <input type="hidden">
}

func ExampleTypeImage() {
	el := Input(Attributes(TypeImage()))
	fmt.Println(el.String())
	// Output: <input type="image">
}

func ExampleTypeMonth() {
	el := Input(Attributes(TypeMonth()))
	fmt.Println(el.String())
	// Output: <input type="month">
}

func ExampleTypeNumber() {
	el := Input(Attributes(TypeNumber()))
	fmt.Println(el.String())
	// Output: <input type="number">
}

func ExampleTypeNumeric() {
	el := Ol(Attributes(TypeNumeric()))
	fmt.Println(el.String())
	// Output: <ol type="1"></ol>
}

func ExampleTypePassword() {
	el := Input(Attributes(TypePassword()))
	fmt.Println(el.String())
	// Output: <input type="password">
}

func ExampleTypeRadio() {
	el := Input(Attributes(TypeRadio()))
	fmt.Println(el.String())
	// Output: <input type="radio">
}

func ExampleTypeRange() {
	el := Input(Attributes(TypeRange()))
	fmt.Println(el.String())
	// Output: <input type="range">
}

func ExampleTypeReset() {
	el := Input(Attributes(TypeReset()))
	fmt.Println(el.String())
	// Output: <input type="reset">
}

func ExampleTypeSearch() {
	el := Input(Attributes(TypeSearch()))
	fmt.Println(el.String())
	// Output: <input type="search">
}

func ExampleTypeSubmit() {
	el := Input(Attributes(TypeSubmit()))
	fmt.Println(el.String())
	// Output: <input type="submit">
}

func ExampleTypeTel() {
	el := Input(Attributes(TypeTel()))
	fmt.Println(el.String())
	// Output: <input type="tel">
}

func ExampleTypeText() {
	el := Input(Attributes(TypeText()))
	fmt.Println(el.String())
	// Output: <input type="text">
}

func ExampleTypeTime() {
	el := Input(Attributes(TypeTime()))
	fmt.Println(el.String())
	// Output: <input type="time">
}

func ExampleTypeUrl() {
	el := Input(Attributes(TypeUrl()))
	fmt.Println(el.String())
	// Output: <input type="url">
}

func ExampleTypeWeek() {
	el := Input(Attributes(TypeWeek()))
	fmt.Println(el.String())
	// Output: <input type="week">
}

// Wrap
func ExampleWrapHard() {
	el := Textarea(Attributes(WrapHard()))
	fmt.Println(el.String())
	// Output: <textarea wrap="hard"></textarea>
}

func ExampleWrapSoft() {
	el := Textarea(Attributes(WrapSoft()))
	fmt.Println(el.String())
	// Output: <textarea wrap="soft"></textarea>
}

// WritingSuggestions
func ExampleWritingSuggestionsEmpty() {
	// TODO this makes no sense at the moment, have to think about how to handle it. Currently attributes with empty string are not printed at all

	el := Div(Attributes(WritingSuggestionsEmpty()))
	fmt.Println(el.String())
	// Output: <div></div>
}

func ExampleWritingSuggestionsFalse() {
	el := Div(Attributes(WritingSuggestionsFalse()))
	fmt.Println(el.String())
	// Output: <div writingsuggestions="false"></div>
}

func ExampleWritingSuggestionsTrue() {
	el := Div(Attributes(WritingSuggestionsTrue()))
	fmt.Println(el.String())
	// Output: <div writingsuggestions="true"></div>
}
