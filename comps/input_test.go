package comps

import (
	"fmt"

	"github.com/KrischanCS/roundly"
	. "github.com/KrischanCS/roundly/attribute"
)

func ExampleIptText() {
	textInput := IptText(InputAttributes{
		Label: "Username",
		Name:  "username",
		Id:    "ipt-username",
	})

	fmt.Println(textInput.StringPretty())

	// Output:
	//
	// <label for="ipt-username">
	// 	Username
	// </label>
	// <input
	// 	type="text"
	// 	name="username"
	// 	id="ipt-username"
	// >
}

func ExampleIptPassword() {
	passwordInput := IptPassword(InputAttributes{
		Label: "Password",
		Name:  "password",
		Id:    "ipt-password",
	})

	fmt.Println(passwordInput.StringPretty())

	// Output:
	//
	// <label for="ipt-password">
	// 	Password
	// </label>
	// <input
	// 	type="password"
	// 	name="password"
	// 	id="ipt-password"
	// >
}

func ExampleIptSubmit() {
	passwordInput := IptSubmit("Send Form")

	fmt.Println(passwordInput.StringPretty())

	// Output:
	//
	// <input type="submit" value="Send Form">
}

func ExampleIptReset() {
	passwordInput := IptReset("Reset Form")

	fmt.Println(passwordInput.StringPretty())

	// Output:
	//
	// <input type="reset" value="Reset Form">
}

func ExampleIptRadioGroup() {
	radioGroupInput := IptRadioGroup("Language",
		IptRadioButtonOption{
			Label:           "Go",
			Id:              "rb-lang-go",
			Value:           "Go",
			OtherAttributes: []roundly.Attribute{Checked()},
		},
		IptRadioButtonOption{
			Label: "Java",
			Id:    "rb-lang-java",
			Value: "Java",
		},
		IptRadioButtonOption{
			Label: "Rust",
			Id:    "rb-lang-rust",
			Value: "Rust",
		},
	)

	fmt.Println(radioGroupInput.StringPretty())

	// Output:
	//
	// <input
	// 	type="radio"
	// 	name="Language"
	// 	id="rb-lang-go"
	// 	value="Go"
	// 	checked
	// >
	// <label for="rb-lang-go">
	// 	Go
	// </label>
	// <input
	// 	type="radio"
	// 	name="Language"
	// 	id="rb-lang-java"
	// 	value="Java"
	// >
	// <label for="rb-lang-java">
	// 	Java
	// </label>
	// <input
	// 	type="radio"
	// 	name="Language"
	// 	id="rb-lang-rust"
	// 	value="Rust"
	// >
	// <label for="rb-lang-rust">
	// 	Rust
	// </label>
}

func ExampleIptCheckBox() {
	checkboxInput := IptCheckBox(InputAttributes{
		Label: "Checked?",
		Name:  "checkbox",
		Id:    "ipt-checkbox",
	})

	fmt.Println(checkboxInput.StringPretty())

	// Output:
	//
	// <input
	// 	type="checkbox"
	// 	name="checkbox"
	// 	id="ipt-checkbox"
	// >
	// <label for="ipt-checkbox">
	// 	Checked?
	// </label>
}

func ExampleIptButton() {
	buttonInput := IptButton("Click Me!", "alert('Clicked!')")

	fmt.Println(buttonInput.StringPretty())

	// Output:
	//
	// <input
	// 	type="button"
	// 	value="Click Me!"
	// 	onclick="alert('Clicked!')"
	// >
}

func ExampleIptColor() {
	colorInput := IptColor(InputAttributes{
		Label: "Favorite Color",
		Name:  "favorite-color",
		Id:    "ipt-favorite-color",
		Value: "#22dd88",
	})

	fmt.Println(colorInput.StringPretty())

	// Output:
	//
	// <label for="ipt-favorite-color">
	// 	Favorite Color
	// </label>
	// <input
	// 	type="color"
	// 	name="favorite-color"
	// 	id="ipt-favorite-color"
	// 	value="#22dd88"
	// >
}

func ExampleIptDate() {
	colorInput := IptDate(InputAttributes{
		Label: "Scheduled Date",
		Name:  "scheduled-date",
		Id:    "ipt-scheduled-date",
		Value: "2025-08-30",
	})

	fmt.Println(colorInput.StringPretty())

	// Output:
	//
	// <label for="ipt-scheduled-date">
	// 	Scheduled Date
	// </label>
	// <input
	// 	type="date"
	// 	name="scheduled-date"
	// 	id="ipt-scheduled-date"
	// 	value="2025-08-30"
	// >
}

func ExampleIptDateTimeLocal() {
	colorInput := IptDateTimeLocal(InputAttributes{
		Label: "The Moment",
		Name:  "the-moment",
		Id:    "ipt-the-moment",
		Value: "2025-08-30T12:22",
	})

	fmt.Println(colorInput.StringPretty())

	// Output:
	//
	// <label for="ipt-the-moment">
	// 	The Moment
	// </label>
	// <input
	// 	type="datetime-local"
	// 	name="the-moment"
	// 	id="ipt-the-moment"
	// 	value="2025-08-30T12:22"
	// >
}

func ExampleIptEmail() {
	colorInput := IptEmail(InputAttributes{
		Label: "E-Mail Address",
		Name:  "e-mail-address",
		Id:    "ipt-e-mail-address",
		Value: "someone@example.org",
	})

	fmt.Println(colorInput.StringPretty())

	// Output:
	//
	// <label for="ipt-e-mail-address">
	// 	E-Mail Address
	// </label>
	// <input
	// 	type="email"
	// 	name="e-mail-address"
	// 	id="ipt-e-mail-address"
	// 	value="someone@example.org"
	// >
}

func ExampleIptImage() {
	imageInput := IptImage(ImageAttributes{
		Label:  "Click on the center:",
		Id:     "ipt-image",
		Src:    "/some-image.jpg",
		Alt:    "alt text",
		Width:  31,
		Height: 31,
	})

	fmt.Println(imageInput.StringPretty())

	// Output:
	//
	// <div style="display:flex; flex-direction:column;">
	// 	<label id="ipt-image">
	// 		Click on the center:
	// 	</label>
	// 	<input
	// 		type="image"
	// 		id="ipt-image"
	// 		src="/some-image.jpg"
	// 		alt="alt text"
	// 		width="31"
	// 		height="31"
	// 	>
	// </div>
}

func ExampleIptFile() {
	fileInput := IptFile(InputAttributes{
		Label: "Select File",
		Name:  "file",
		Id:    "ipt-file",
	})

	fmt.Println(fileInput.StringPretty())

	// Output:
	//
	// <label for="ipt-file">
	// 	Select File
	// </label>
	// <input
	// 	type="file"
	// 	name="file"
	// 	id="ipt-file"
	// >
}

func ExampleIptHidden() {
	hiddenInput := IptHidden(
		"hidden-input",
		"this is sent with the form, but neither editable nor visible to the user")

	fmt.Println(hiddenInput.StringPretty())

	// Output:
	//
	// <input
	// 	type="hidden"
	// 	name="hidden-input"
	// 	value="this is sent with the form, but neither editable nor visible to the user"
	// >
}

func ExampleIptMonth() {
	monthInput := IptMonth(InputAttributes{
		Label: "The Month",
		Name:  "the-month",
		Id:    "ipt-the-month",
		Value: "2025-08",
	})

	fmt.Println(monthInput.StringPretty())

	// Output:
	//
	// <label for="ipt-the-month">
	// 	The Month
	// </label>
	// <input
	// 	type="month"
	// 	name="the-month"
	// 	id="ipt-the-month"
	// 	value="2025-08"
	// >
}

func ExampleIptNumber() {
	numberInput := IptNumber(
		InputAttributes{
			Label: "The Number",
			Name:  "the-number",
			Id:    "ipt-the-number",
			Value: "45",
		},
		Attributes(
			Min("0"),
			Max("50"),
			Step(5),
		),
	)

	fmt.Println(numberInput.StringPretty())

	// Output:
	//
	// <label for="ipt-the-number">
	// 	The Number
	// </label>
	// <input
	// 	type="number"
	// 	name="the-number"
	// 	id="ipt-the-number"
	// 	value="45"
	// 	min="0"
	// 	max="50"
	// 	step="5"
	// >
}

func ExampleIptRange() {
	rangeInput := IptRange(
		InputAttributes{
			Label: "The Range",
			Name:  "the-range",
			Id:    "ipt-the-range",
			Value: "45",
		},
		"0", "70",
		Attributes(
			Step(5),
		),
	)

	fmt.Println(rangeInput.StringPretty())

	// Output:
	//
	// <label for="ipt-the-range">
	// 	The Range
	// </label>
	// <input
	// 	type="range"
	// 	name="the-range"
	// 	id="ipt-the-range"
	// 	value="45"
	// 	min="0"
	// 	max="70"
	// 	step="5"
	// >
}

func ExampleIptSearch() {
	textInput := IptSearch(InputAttributes{
		Label:       "Search",
		Name:        "search-query",
		Id:          "ipt-search",
		Placeholder: "Search placeholder",
	})

	fmt.Println(textInput.StringPretty())

	// Output:
	//
	// <label for="ipt-search">
	// 	Search
	// </label>
	// <input
	// 	type="search"
	// 	name="search-query"
	// 	id="ipt-search"
	// 	placeholder="Search placeholder"
	// >
}

func ExampleIptTel() {
	telInput := IptTel(InputAttributes{
		Label:       "Number",
		Name:        "number",
		Id:          "ipt-number",
		Placeholder: "555-1337",
	})

	fmt.Println(telInput.StringPretty())

	// Output:
	//
	// <label for="ipt-number">
	// 	Number
	// </label>
	// <input
	// 	type="tel"
	// 	name="number"
	// 	id="ipt-number"
	// 	placeholder="555-1337"
	// >
}

func ExampleIptTime() {
	timeInput := IptTime(InputAttributes{
		Label: "Time",
		Name:  "time",
		Id:    "ipt-time",
		Value: "12:22",
	})

	fmt.Println(timeInput.StringPretty())

	// Output:
	//
	// <label for="ipt-time">
	// 	Time
	// </label>
	// <input
	// 	type="time"
	// 	name="time"
	// 	id="ipt-time"
	// 	value="12:22"
	// >
}

func ExampleIptUrl() {
	urlInput := IptUrl(InputAttributes{
		Label:       "Url",
		Name:        "url",
		Id:          "ipt-url",
		Placeholder: "https://example.org",
	})

	fmt.Println(urlInput.StringPretty())

	// Output:
	//
	// <label for="ipt-url">
	// 	Url
	// </label>
	// <input
	// 	type="url"
	// 	name="url"
	// 	id="ipt-url"
	// 	placeholder="https://example.org"
	// >
}

func ExampleIptWeek() {
	weekInput := IptWeek(InputAttributes{
		Label:       "Week",
		Name:        "week",
		Id:          "ipt-week",
		Placeholder: "2025-W35",
	})

	fmt.Println(weekInput.StringPretty())

	// Output:
	//
	// <label for="ipt-week">
	// 	Week
	// </label>
	// <input
	// 	type="week"
	// 	name="week"
	// 	id="ipt-week"
	// 	placeholder="2025-W35"
	// >
}

func ExampleForm() {
	form := Form("#", nil,
		IptText(InputAttributes{
			Label:       "Username",
			Name:        "username",
			Id:          "ipt-username",
			Placeholder: "Username",
		}),
		IptPassword(InputAttributes{
			Label:       "Password",
			Name:        "password",
			Id:          "ipt-password",
			Placeholder: "Password",
		}),
		IptRadioGroup(
			"user-type",
			IptRadioButtonOption{
				Label: "Developer",
				Value: "developer",
				Id:    "user-type-opt-developer",
			},
			IptRadioButtonOption{
				Label: "Product Owner",
				Value: "product-owner",
				Id:    "user-type-opt-product-owner",
			},
			IptRadioButtonOption{
				Label: "Architect",
				Value: "product-architect",
				Id:    "user-type-opt-product-architect",
			},
			IptRadioButtonOption{
				Label: "Scrum Master",
				Value: "scrum-master",
				Id:    "user-type-opt-scrum-master",
			},
		),
		IptReset("Reset"),
		IptSubmit("Send"),
	)

	fmt.Println(form.StringPretty())

	// Output:
	//
	// <form action="#" style="display:grid; grid-template-columns: fit-content(1%) 1fr">
	// 	<label for="ipt-username">
	// 		Username
	// 	</label>
	// 	<input
	// 		type="text"
	// 		name="username"
	// 		id="ipt-username"
	// 		placeholder="Username"
	// 	>
	// 	<label for="ipt-password">
	// 		Password
	// 	</label>
	// 	<input
	// 		type="password"
	// 		name="password"
	// 		id="ipt-password"
	// 		placeholder="Password"
	// 	>
	// 	<input
	// 		type="radio"
	// 		name="user-type"
	// 		id="user-type-opt-developer"
	// 		value="developer"
	// 	>
	// 	<label for="user-type-opt-developer">
	// 		Developer
	// 	</label>
	// 	<input
	// 		type="radio"
	// 		name="user-type"
	// 		id="user-type-opt-product-owner"
	// 		value="product-owner"
	// 	>
	// 	<label for="user-type-opt-product-owner">
	// 		Product Owner
	// 	</label>
	// 	<input
	// 		type="radio"
	// 		name="user-type"
	// 		id="user-type-opt-product-architect"
	// 		value="product-architect"
	// 	>
	// 	<label for="user-type-opt-product-architect">
	// 		Architect
	// 	</label>
	// 	<input
	// 		type="radio"
	// 		name="user-type"
	// 		id="user-type-opt-scrum-master"
	// 		value="scrum-master"
	// 	>
	// 	<label for="user-type-opt-scrum-master">
	// 		Scrum Master
	// 	</label>
	// 	<input type="reset" value="Reset">
	// 	<input type="submit" value="Send">
	// </form>
}
