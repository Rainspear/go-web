package main

import (
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

type fileState struct{}

func (fileState) executeMain() {
	http.Handle("/", http.HandlerFunc(cat))
	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./assets"))))
	http.Handle("/image", http.HandlerFunc(serveDogImage))
	http.Handle("/test", http.HandlerFunc(test))
	http.ListenAndServe(":8089", nil)
}

func cat(w http.ResponseWriter, r *http.Request) {
	var s string
	if r.Method == http.MethodPost {
		// open
		f, h, err := r.FormFile("q")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer f.Close()

		fmt.Println("\nfile:", f, "\nheader:", h, "\nerr", err)

		// read
		bs, err := ioutil.ReadAll(f)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		s = string(bs)
	}

	w.Header().Add("Content-Type", "text/html; charset=utf-8;")
	io.WriteString(w, `
	<form method="POST" enctype="multipart/form-data">
	<input type="file" name="q">
	<input type="submit">
	</form>
	<br>`+s)
}

func serveDogImage(w http.ResponseWriter, r *http.Request) {
	status := ""
	if r.Method == http.MethodPost { // upload image

		r.ParseMultipartForm(1024 * 1024 * 10) // 10Mb maximum memory
		f, _, err := r.FormFile("q")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer f.Close()
		// fmt.Printf("Uploaded File: %+v\n", h.Filename)
		// fmt.Printf("File Size: %+v\n", h.Size)
		// fmt.Printf("MIME Header: %+v\n", h.Header)
		t := time.Now().Format(time.RFC3339) // 2009-11-10T23:00:00Z
		name := string([]byte(`./assets/dog ` + t + ".jpg"))
		tempFile, err := os.Create(name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer tempFile.Close()

		fileBytes, err := ioutil.ReadAll(f)
		if err != nil {
			fmt.Println(err)
		}
		_, err = tempFile.Write(fileBytes)
		if err != nil {
			status = "Write file to server failure\n"
		} else {
			status = "Successfully uploaded file\n"
		}
	}
	// find image name
	dirs, err := os.ReadDir("./assets")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	imgName := "shiba.jpg"
	for _, d := range dirs {
		if strings.Contains(d.Name(), "dog") {
			imgName = d.Name()
			break
		}
	}
	// response
	w.Header().Add("Content-Type", "text/html; charset=utf-8;")
	io.WriteString(w, `
	<form method="POST" enctype="multipart/form-data">
	<input type="file" name="q">
	<input type="submit">
	</form>
	<div>`+status+`</div>
	<img src="/static/`+imgName+"\""+` alt="dog" />
	<br>`)
}

type responseFile struct {
	Name  string
	IsDir bool
}

func test(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html; charset=utf-8;")
	dirs, err := os.ReadDir("./assets")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	var fltd []responseFile
	for _, d := range dirs {
		fltd = append(fltd, responseFile{
			d.Name(),
			d.IsDir(),
		})
	}

	fmt.Printf("test %+v \r\n", fltd)
	// io.WriteString(w, "Hello world")
	body := `
	<html lang="en">
		<head>
				<meta charset="UTF-8">
				<meta name="viewport" content="width=device-width, initial-scale=1">
				<title>Directory</title>
				<link href="https://cdn.jsdelivr.net/npm/bootstrap@5.2.1/dist/css/bootstrap.min.css" rel="stylesheet" integrity="sha384-iYQeCzEYFbKjA/T2uDLTpkwGzCiq6soy8tYaI1GyVh/UjpbCx/TYkiZhlZB6+fzT" crossorigin="anonymous">
		</head>
		<body>
			<h1>Directory Assets</h1>
				<ul>
				{{range .}}
					<li>
					{{.Name}}
						{{ if .IsDir }}
						<svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" class="bi bi-folder" viewBox="0 0 16 16">
							<path d="M.54 3.87.5 3a2 2 0 0 1 2-2h3.672a2 2 0 0 1 1.414.586l.828.828A2 2 0 0 0 9.828 3h3.982a2 2 0 0 1 1.992 2.181l-.637 7A2 2 0 0 1 13.174 14H2.826a2 2 0 0 1-1.991-1.819l-.637-7a1.99 1.99 0 0 1 .342-1.31zM2.19 4a1 1 0 0 0-.996 1.09l.637 7a1 1 0 0 0 .995.91h10.348a1 1 0 0 0 .995-.91l.637-7A1 1 0 0 0 13.81 4H2.19zm4.69-1.707A1 1 0 0 0 6.172 2H2.5a1 1 0 0 0-1 .981l.006.139C1.72 3.042 1.95 3 2.19 3h5.396l-.707-.707z"/>
						</svg>
						{{end}}
					</li>
				{{end}}
			</ul>
			<script src="https://cdn.jsdelivr.net/npm/@popperjs/core@2.11.6/dist/umd/popper.min.js" integrity="sha384-oBqDVmMz9ATKxIep9tiCxS/Z9fNfEXiDAYTujMAeBAsjFuCZSmKbSSUnQlmh/jp3" crossorigin="anonymous"></script>
			<script src="https://cdn.jsdelivr.net/npm/bootstrap@5.2.1/dist/js/bootstrap.min.js" integrity="sha384-7VPbUDkoPSGFnVtYi0QogXtr74QeVeeIs99Qfg5YCF+TidwNdjvaKZX19NZ/e6oz" crossorigin="anonymous"></script>
		</body>
	</html>
	`
	tpl := template.Must(template.New("test.gohtml").Parse(body))
	tpl.ExecuteTemplate(w, "test.gohtml", fltd)
}
