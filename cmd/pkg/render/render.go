package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

// RenderTemplate renders tesmplats using html/template
func RenderTemplate(w http.ResponseWriter, tmpl string) {
	// create a template cache
	tc, err := createTemplateCache()
	if err != nil {
		log.Fatal(err)
	}

	// get requested template from cache
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal(err)
	}

	// extra error check
	buf := new(bytes.Buffer)
	err = t.Execute(buf, nil)
	if err != nil {
		log.Println(err)
	}

	// render the template
	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}
}

func createTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	// get all files *.page.tmpl
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}
	//range through files
	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}
		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}
		myCache[name] = ts
	}
	return myCache, nil
}


// This is the basic way to do the same as above:

// var tc = make(map[string]*template.Template)

// func RenderTemplateBasic(w http.ResponseWriter, t string) {
// 	var tmpl *template.Template
// 	var err error

// 	// check to see if we already have the template in cache
// 	_, inMap := tc[t]
// 	if !inMap {
// 		// need to create template
// 		log.Println("creating template and adding to cache")
// 		err = createTemplateCacheBasic(t)
// 		if err != nil {
// 			log.Println(err)
// 		}
// 	} else {
// 		// the template is in cache
// 		log.Println("using cached template")
// 	}
// 	tmpl = tc[t]
// 	err = tmpl.Execute(w, nil)
// 	if err != nil {
// 		log.Println(err)
// 	}
// }

// func createTemplateCacheBasic(t string) error {
// 	templates := []string{
// 		fmt.Sprintf("./templates/%s", t),
// 		"./templates/base.layout.tmpl",
// 	}
// 	tmpl, err := template.ParseFiles(templates...)
// 	if err != nil {
// 		return err
// 	}
// 	tc[t] = tmpl
// 	return nil
// }
