# Summary of {{.Title}}</h1>

Here is the gist of an [article]({{.URL}}) by `{{.Author}}` about {{.Topic}}:

> {{index .Quotes 0}}

Further more, the article continues:

> {{index .Quotes 1}}

And

> {{index .Quotes 2}}

## Topics in Article:

{{range .Phrases}}
* {{.}}
{{end}}