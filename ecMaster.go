package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
	"fmt"
	_ "github.com/lib/pq"
	"database/sql"
	"log"
	"platformOps-EC/models"
	"platformOps-EC/services"
	"time"
	"crypto/md5"
	"io"
	"strconv"
	"os"
)

type Env struct {
	db *sql.DB
}


func main() {

	db, err := models.NewDB()
	services.SetSearchPath(db, "baseline")
	if err != nil {
		log.Panic(err)
	}

	app := iris.New()
	app.Logger().SetLevel("debug")
	app.RegisterView(iris.HTML("./templates", ".html"))

	// Optionally, add two built'n handlers
	// that can recover from any http-relative panics
	// and log the requests to the terminal.
	app.Use(recover.New())
	app.Use(logger.New())

	// Method:   GET
	// Resource: http://localhost:8080
	app.Handle("GET", "/", func(ctx iris.Context) {
		ctx.HTML("<h1>Welcome to EC Master</h1>")
	})

	// Method:   GET
	// Resource: http://localhost:8080/manifest/{baselineId}
	app.Get("/manifest/{baselineId}", func(ctx iris.Context)  {
		id, _ := ctx.Params().GetInt("baselineId")

		ctx.JSON(services.GetManifestByBaselineId(db, id))
	})

	// Method:   GET
	// Resource: http://localhost:8080/hello
	app.Get("/hello/{name}", func(ctx iris.Context) {

		ctx.JSON(iris.Map{"message": fmt.Sprintf("Hello %s!", ctx.Params().Get("name"))})
	})

	// Method:   GET
	// Resource: http://localhost:8080/ecResult/{ecResultId}
	app.Get("/ecResult/{ecResultId}", func(ctx iris.Context)  {
		id, _ := ctx.Params().GetInt("ecResultId")

		ctx.WriteString(services.GetECResultById(db, id))
	})

	// Method POST: http://localhost:8080/ecResults
	app.Post("/ecResults", func(ctx iris.Context) {
		var ecResult []models.ECResult
		ctx.ReadJSON(&ecResult)
		saveId := services.SaveECResult(db, ecResult)
		if saveId == 0 {
			ctx.JSON(iris.Map{"status":"failed","id": fmt.Sprintf("%d", saveId)})
		}
		ctx.JSON(iris.Map{"status":"success","id": fmt.Sprintf("%d", saveId)})
	})

	// Serve the upload_form.html to the client.
	app.Get("/upload", func(ctx iris.Context) {
		// create a token (optionally).

		now := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(now, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))

		// render the form with the token for any use you'd like.
		// ctx.ViewData("", token)
		// or add second argument to the `View` method.
		// Token will be passed as {{.}} in the template.
		ctx.View("upload_form.html", token)
	})

	// Handle the post request from the upload_form.html to the server
	app.Post("/upload", func(ctx iris.Context) {
		// iris.LimitRequestBodySize(32 <<20) as middleware to a route
		// or use ctx.SetMaxRequestBodySize(32 << 20)
		// to limit the whole request body size,
		//
		// or let the configuration option at app.Run for global setting
		// for POST/PUT methods, including uploads of course.

		// Get the file from the request.
		file, info, err := ctx.FormFile("uploadfile")

		if err != nil {
			ctx.StatusCode(iris.StatusInternalServerError)
			ctx.HTML("Error while uploading: <b>" + err.Error() + "</b>")
			return
		}

		defer file.Close()
		fname := info.Filename

		// Create a file with the same name
		// assuming that you have a folder named 'uploads'
		out, err := os.OpenFile("./uploads/"+fname,
			os.O_WRONLY|os.O_CREATE, 0666)

		if err != nil {
			ctx.StatusCode(iris.StatusInternalServerError)
			ctx.HTML("Error while uploading: <b>" + err.Error() + "</b>")
			return
		}
		defer out.Close()

		io.Copy(out, file)

		baseline, controls := services.LoadFromExcel(out.Name())
		var manifest []models.ECManifest

		fmt.Println("Converting to Json object")

		for _, c := range controls {

			m := models.ECManifest{ReqId: c.ReqId, Title: c.Category,
				Baseline: baseline.Name}
			manifest = append(manifest, m)

		}
		ctx.WriteString(models.ToJson(manifest))
	})

	// http://localhost:8080
	// http://localhost:8080/ping
	// http://localhost:8080/hello
	app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))
}



