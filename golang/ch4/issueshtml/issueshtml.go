package issueshtml

const Templ = `
<h1>issues</h1>
<table>
<tr style='text-align: left'>
	<th>#</th>
	<th>State</th>
	<th>User</th>
	<th>Title</th>
</tr>
{{range .Items}}
<tr>
	<td><a href='{{.HTMLURL}}'>{{.Number}}</a></td>
	<td>{{.State}}</td>
	<td>{{.User.Login}}</td>
	<td>{{.Title}}</td>
</tr>
{{end}}
</table>
`
