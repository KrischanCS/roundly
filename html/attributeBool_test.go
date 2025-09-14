package html

import (
	"fmt"
)

func ExampleAllowFullScreen() {
	el := Iframe(AllowFullScreen())
	fmt.Println(el.String())
	// Output: <iframe allowfullscreen></iframe>
}

func ExampleAlpha() {
	el := Input(Attributes(TypeColor(), Alpha()))
	fmt.Println(el.String())
	// Output: <input type="color" alpha>
}

func ExampleAsync() {
	el := Script(Async(), Text("// some script"))
	fmt.Println(el.String())
	// Output: <script async>// some script</script>
}

func ExampleAutoFocus() {
	el := Input(Attributes(TypeText(), AutoFocus()))
	fmt.Println(el.String())
	// Output: <input type="text" autofocus>
}

func ExampleAutoPlay() {
	el := Video(Attributes(AutoPlay(), Src("video.mp4")))
	fmt.Println(el.String())
	// Output: <video autoplay src="video.mp4"></video>
}

func ExampleChecked() {
	el := Input(Attributes(TypeCheckbox(), Checked()))
	fmt.Println(el.String())
	// Output: <input type="checkbox" checked>
}

func ExampleControls() {
	el := Audio(Attributes(Controls(), Src("audio.mp3")))
	fmt.Println(el.String())
	// Output: <audio controls src="audio.mp3"></audio>
}

func ExampleDefault() {
	el := Track(Attributes(Default(), Src("captions.vtt")))
	fmt.Println(el.String())
	// Output: <track default src="captions.vtt">
}

func ExampleDefer() {
	el := Script(Defer(), Text("// some script"))
	fmt.Println(el.String())
	// Output: <script defer>// some script</script>
}

func ExampleDisabled() {
	el := Input(Attributes(TypeText(), Disabled()))
	fmt.Println(el.String())
	// Output: <input type="text" disabled>
}

func ExampleFormNoValidate() {
	el := Input(Attributes(TypeSubmit(), FormNoValidate()))
	fmt.Println(el.String())
	// Output: <input type="submit" formnovalidate>
}

func ExampleInert() {
	el := Div(Inert())
	fmt.Println(el.String())
	// Output: <div inert></div>
}

func ExampleIsMap() {
	el := Img(Attributes(Src("image.png"), IsMap()))
	fmt.Println(el.String())
	// Output: <img src="image.png" ismap>
}

func ExampleItemScope() {
	el := Div(ItemScope())
	fmt.Println(el.String())
	// Output: <div itemscope></div>
}

func ExampleLoop() {
	el := Video(Attributes(Loop(), Src("video.mp4")))
	fmt.Println(el.String())
	// Output: <video loop src="video.mp4"></video>
}

func ExampleMultiple() {
	el := Select(Multiple())
	fmt.Println(el.String())
	// Output: <select multiple></select>
}

func ExampleMuted() {
	el := Audio(Attributes(Muted(), Src("audio.mp3")))
	fmt.Println(el.String())
	// Output: <audio muted src="audio.mp3"></audio>
}

func ExampleNoModule() {
	el := Script(NoModule(), Text("// some script"))
	fmt.Println(el.String())
	// Output: <script nomodule>// some script</script>
}

func ExampleNoValidate() {
	el := Form(NoValidate())
	fmt.Println(el.String())
	// Output: <form novalidate></form>
}

func ExampleOpen() {
	el := Details(Open())
	fmt.Println(el.String())
	// Output: <details open></details>
}

func ExamplePlaysInline() {
	el := Video(Attributes(PlaysInline(), Src("video.mp4")))
	fmt.Println(el.String())
	// Output: <video playsinline src="video.mp4"></video>
}

func ExampleReadOnly() {
	el := Textarea(Attributes(ReadOnly()))
	fmt.Println(el.String())
	// Output: <textarea readonly></textarea>
}

func ExampleRequired() {
	el := Input(Attributes(TypeText(), Required()))
	fmt.Println(el.String())
	// Output: <input type="text" required>
}

func ExampleReversed() {
	el := Ol(Reversed())
	fmt.Println(el.String())
	// Output: <ol reversed></ol>
}

func ExampleSelected() {
	el := Option(Selected(), Text("One"))
	fmt.Println(el.String())
	// Output: <option selected>One</option>
}

func ExampleShadowRootClonable() {
	el := Template(ShadowRootClonable())
	fmt.Println(el.String())
	// Output: <template shadowrootclonable></template>
}

func ExampleShadowRootCustomElementRegistry() {
	el := Template(ShadowRootCustomElementRegistry())
	fmt.Println(el.String())
	// Output: <template shadowrootcustomelementregistry></template>
}

func ExampleShadowRootDelegatesFocus() {
	el := Template(ShadowRootDelegatesFocus())
	fmt.Println(el.String())
	// Output: <template shadowrootdelegatesfocus></template>
}

func ExampleShadowRootSerializable() {
	el := Template(ShadowRootSerializable())
	fmt.Println(el.String())
	// Output: <template shadowrootserializable></template>
}
