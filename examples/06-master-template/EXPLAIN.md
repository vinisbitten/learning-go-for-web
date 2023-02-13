# Explaination

This is the last example of the series. In this example we will use a master template to avoid repeating the same HTML code in all the templates.

## Master template

- The master template called **layout.html** will contain the HTML code that is **common** to all the templates.
- All the other templates (index, list, view) will **extend** the master template, for each individual case.
- The master template will have a **placeholder** for the content of the page. In our case:

```html
{{ content "body" . }}
```

- The other templates will **fill** the placeholder with their own content. In our case:

```html
{{ define "body" }}
    ... content of the page ...
{{ end }}
```

## using the master template

- It's acctually very simple. We just have to call both "templates" in the handler function:

```go
// ParseFiles() --> parse all the files and return a template

// old code --> calling only the index template
// Used in example 4 --> in this case is a full html page
t := template.Must(template.ParseFiles("template/index.html"))
if err := t.ExecuteTemplate(w, "index.html", page); err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
}

// new code --> calling the master template and the index template
// Used in this example --> layout and index combined
t := template.Must(template.ParseFiles("template/layout.html", "template/index.html"))
if err := t.ExecuteTemplate(w, "layout.html", page); err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
}
```
