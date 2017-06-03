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
  <tr>
    <th>
      <a href="http://localhost:8000/tracks?sort=Title">Title</a>
    </th>
    <th>
      <a href="http://localhost:8000/tracks?sort=Artist">Artist</a>
    </th>
    <th>
      <a href="http://localhost:8000/tracks?sort=Album">Album</a>
    </th>
    <th>
      <a href="http://localhost:8000/tracks?sort=Year">Year</a>
    </th>
    <th>
      <a href="http://localhost:8000/tracks?sort=Length">Length</a>
    </th>
  </tr>
  {{range .Tracks}}
  <tr>
    <td>{{.Title}}</td>
    <td>{{.Artist}}</td>
    <td>{{.Album}}</td>
    <td>{{.Year}}</td>
    <td>{{.Length}}</td>
  </tr>
  {{end}}
</table>
