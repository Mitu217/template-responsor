<!DOCTYPE html>
<html>
<body>
    <div>
        <button onclick="location.href='/admin/create'">Add</button>
        <h4>ResponseList</h4>
        {{ if eq .len 0 }}
            <p>no data</p>
        {{ else }}
            <table border="1" cellspacing="0" cellpadding="5" bordercolor="#333333">
                <tr>
                    <th>Name</th>
                    <th>Route</th>
                    <th>RequestMethod</th>
                    <th>ContentType</th>
                    <th>Data</th>
                </tr>
                {{range $i, $res := .ress}}
                    <tr>
                        <td>{{ $res.Name }}</td>
                        <td>{{ $res.Route }}</td>
                        <td>{{ $res.RequestMethod }}</td>
                        <td>{{ $res.ContentType }}</td>
                        <td>{{ $res.Data }}</td>
                    </tr>
                {{end}}
            </table>
        {{ end }}
    </div>
</body>
</html>