<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>User List</title>
    <link rel="stylesheet" href="https://cdn.datatables.net/1.10.22/css/jquery.dataTables.min.css" />
    <!-- jQuery library file -->
    <script type="text/javascript" src="https://code.jquery.com/jquery-3.5.1.js">
    </script>
    <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bootstrap@5.0.0-beta1/dist/js/bootstrap.bundle.min.js">
    <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bootstrap@5.0.0-beta1/dist/css/bootstrap.min.css">
    <!-- Datatable plugin JS library file -->
    <script type="text/javascript" src="https://cdn.datatables.net/1.10.22/js/jquery.dataTables.min.js">
    </script>
    <style>
        table,
        th,
        td {
            border: 1px solid black;
            border-collapse: collapse;
        }
    </style>
</head>

<body>
    {{ if .message }}
    <div class="alert alert-success" role="alert">
        <p style="text-align: center;">{{ .message }}</p>
    </div>
    {{ end }}
    <h1>User List</h1>
    <div>
        <table id="users">
            <thead>
                <tr>
                    <th>SR</th>
                    <th>Name</th>
                    <th>Last Name</th>
                    <th>Email</th>
                    <th>Age</th>
                    <th>Country</th>
                    <th>Action</th>
                </tr>
            </thead>
            <tbody>
                {{range .users}}
                <tr>
                    <td>{{.Id}}</td>
                    <td>{{.FirstName}}</td>
                    <td>{{.LastName}}</td>
                    <td>{{.Email}}</td>
                    <td>{{.Age}}</td>
                    <td>{{.Country}}</td>
                    <td><a href=""></a></td>
                </tr>
                {{end}}
            </tbody>
        </table>
    </div>
</body>
<script>
    /* Initialization of datatable */
    $(document).ready(function() {
        $('#users').DataTable();
    });
</script>
</html>