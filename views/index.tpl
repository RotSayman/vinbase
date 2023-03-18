<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>Vinbase</title>
    <link rel="shortcut icon" href="../static/img/favicon.ico" type="image/x-icon">
    <link href="https://cdn.jsdelivr.net/npm/bootstrap@5.3.0-alpha1/dist/css/bootstrap.min.css" rel="stylesheet" integrity="sha384-GLhlTQ8iRABdZLl6O3oVMWSktQOp6b7In1Zl3/Jr59b6EGGoI1aFkw7cmDA6j6gD" crossorigin="anonymous">
    <link href="../static/css/grid.css" rel="stylesheet">
  </head>
  <body>
    <main>
        <div class="container">
            <h1><a href="/" style="text-decoration: none; color: #008B8B; ">Vinbase</a></h1>
             <div class="d-grid gap-2 d-md-flex justify-content-md-end">
                <form action="/" method="get" class="col-12 col-lg-auto mb-3 mb-lg-0 me-lg-3">
                    <input type="search" name="search" class="form-control" placeholder="Search..." aria-label="Search", value="{{.Search}}">
                    </form>
                <a class="btn btn-lg btn-primary" href="/Add" role="button">Add VIN</a>
            </div>  
            <table class="row mb-3 text-center">
                <tr class="row mb-3 text-center">
                    <th class="col-4 themed-grid-col">Name</th>
                    <th class="col-4 themed-grid-col">Vin</th>
                </tr>
                {{range .vins}}
                <tr class="row mb-3 text-center">
                    <td class="col-4 themed-grid-col">{{.Name}}</td>
                    <td class="col-4 themed-grid-col">{{.VinID}}</td>
                    <td class="col-2">
                        <a class="btn btn-primary" href="/Edit/{{.VinID}}" role="button">&#x270e</a>
                        <a class="btn btn-primary" href="/Delete/{{.VinID}}" role="button">X</span></a>

                    </td>
                </tr>
                {{end}}
            </table>
           
         </div>
    </main>
    <script src="/static/js/reload.min.js"></script>
    <script src="https://cdn.jsdelivr.net/npm/bootstrap@5.3.0-alpha1/dist/js/bootstrap.bundle.min.js" integrity="sha384-w76AqPfDkMBDXo30jS1Sgez6pr3x5MlQ1ZAGC+nuZB+EYdgRZgiwxhTBTkF7CXvN" crossorigin="anonymous"></script>
  </body>
</html>

