package template

var IndexBody = `
	{{ range $index, $element := .Posts }}

		<h1>
			{{ $element.Title }}
		</h1>
		<div id="content">
			{{ $element.Content }}
		</div>

	{{ end }}
`
