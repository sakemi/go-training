<style>
table {
	border-collapse: collapse;
}
td {
	border: solid 1px;
	padding: 0.5em;
}
</style>
<table>

  {{range .}}
  <tr>
    <td>{{.Item}}</td>
    <td>{{.Price}}</td>
  </tr>
  {{end}}
</table>
