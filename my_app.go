package main
import (
  express "github.com/DronRathore/goexpress"
  request "github.com/DronRathore/goexpress/request"
  response "github.com/DronRathore/goexpress/response"
)

func main (){
  var app = express.Express()

  app.Get("/", func(req *request.Request, res *response.Response, next func()){
    res.Write("Hello World")
    // you can skip closing connection
  })

  app.Get("/hello", func(req *request.Request, res *response.Response, next func()){
	   res.SendFile("./views/hello.html", false)
  })

  app.Start("8080")
}
