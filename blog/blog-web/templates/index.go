package template

var IndexBody = `
	{{ range $index, $element := .Posts }}
	<div class="entry" style="text-align: left">
		<h1>
			<a href="post/{{ $element.Slug }}">{{ $element.Title }}</a>
		</h1>
		<div id="content">
			{{ $element.Content }}
		</div>
	</div>
	<br />
	<br />
	{{ end }}
	{{ if not .Posts }}

	<h1>No posts here yet.</h1>

	{{ end }}
`
