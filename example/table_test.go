package example

import (
	"fmt"

	"github.com/KrischanCS/roundly"
	. "github.com/KrischanCS/roundly/attribute"
	. "github.com/KrischanCS/roundly/element"
	. "github.com/KrischanCS/roundly/logic"
	. "github.com/KrischanCS/roundly/text"
)

func ExampleTable() {
	type Mascot struct {
		Language string
		Kind     string
		Name     string
	}

	table := Table(
		Attributes(
			Id("mascot-table"),
			Class("fancy-table", "mascots"),
		),
		Thead(nil,
			Tr(nil,
				ThText("Language"),
				ThText("Kind"),
				ThText("Name"),
			),
		),
		Tbody(nil,
			Range([]Mascot{
				{"Go", "Blue Gopher", "The Go Gopher"},
				{"Rust", "Orange Crab", "Ferris"},
				{"Gleam", "Pink Starfish", "Lucy"},
				{"Zig", "Ziguana", "Ziggy & Zero"},
				{"Java", "?", "Duke"},
			}, func(i int, mascot Mascot) roundly.Element {
				return Tr(nil,
					Th(ScopeRow(), Text(mascot.Language)),
					TdText(mascot.Kind),
					TdText(mascot.Name),
				)
			}),
		),
	)

	w := roundly.NewWriter()

	err := table.RenderElementWithOptions(w, &roundly.RenderOptions{Pretty: true})
	if err != nil {
		panic(err)
	}

	fmt.Println(w.String())

	// Output:
	//
	// <table id="mascot-table" class="fancy-table mascots">
	// 	<thead>
	// 		<tr>
	// 			<th>
	// 				Language
	// 			</th>
	// 			<th>
	// 				Kind
	// 			</th>
	// 			<th>
	// 				Name
	// 			</th>
	// 		</tr>
	// 	</thead>
	// 	<tbody>
	// 		<tr>
	// 			<th scope="row">
	// 				Go
	// 			</th>
	// 			<td>
	// 				Blue Gopher
	// 			</td>
	// 			<td>
	// 				The Go Gopher
	// 			</td>
	// 		</tr>
	// 		<tr>
	// 			<th scope="row">
	// 				Rust
	// 			</th>
	// 			<td>
	// 				Orange Crab
	// 			</td>
	// 			<td>
	// 				Ferris
	// 			</td>
	// 		</tr>
	// 		<tr>
	// 			<th scope="row">
	// 				Gleam
	// 			</th>
	// 			<td>
	// 				Pink Starfish
	// 			</td>
	// 			<td>
	// 				Lucy
	// 			</td>
	// 		</tr>
	// 		<tr>
	// 			<th scope="row">
	// 				Zig
	// 			</th>
	// 			<td>
	// 				Ziguana
	// 			</td>
	// 			<td>
	// 				Ziggy &amp; Zero
	// 			</td>
	// 		</tr>
	// 		<tr>
	// 			<th scope="row">
	// 				Java
	// 			</th>
	// 			<td>
	// 				?
	// 			</td>
	// 			<td>
	// 				Duke
	// 			</td>
	// 		</tr>
	// 	</tbody>
	// </table>
}
