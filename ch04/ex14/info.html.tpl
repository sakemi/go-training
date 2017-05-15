<!DOCTYPE html>
<html>
  {{with .Issues}}
  <h1>
    Issue List
  </h1>
  <h2>
    {{.TotalCount}} issues
  </h2>
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
      <td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
      <td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
    </tr>
    {{end}}
  </table>
  {{end}}
  {{with .Users}}
  <h1>
    User List
  </h1>
  <h2>
    {{.TotalCount}} users
  </h2>
  <table>
    <tr style='text-align: left'>
      <th>Name</th>
    </tr>
    {{range .Items}}
    <tr>
      <td><a href='{{.HTMLURL}}'>{{.Login}}</a></td>
    </tr>
    {{end}}
  </table>
  {{end}}
  {{with .Milestones}}
  <h1>
    Milestone List
  </h1>
  <table>
    <tr style='text-align: left'>
      <th>#</th>
      <th>State</th>
      <th>Title</th>
    </tr>
    {{range .Milestones}}
    <tr>
      <td>{{.Number}}</td>
      <td>{{.State}}</td>
      <td>{{.Title}}</td>
    </tr>
    {{end}}
  </table>
  {{end}}
</html>
