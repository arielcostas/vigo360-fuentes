/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */
package templates

import (
	"bytes"
	"embed"
	"html/template"
	"io"
)

//go:embed html/*
var raw embed.FS

var funcs = template.FuncMap{
	"sino": func(cond bool) string {
		if cond {
			return "Sí"
		}
		return "No"
	},
}

var t = func() *template.Template {
	t := template.New("")

	entries, _ := raw.ReadDir("html")
	for _, de := range entries {
		filename := de.Name()
		contents, _ := raw.ReadFile("html/" + filename)

		_, err := t.New(filename).Funcs(funcs).Parse(string(contents))
		if err != nil {
			panic(err)
		}
	}

	return t
}()

func Render(w io.Writer, name string, params any) error {
	var temporalWriter bytes.Buffer
	err := t.ExecuteTemplate(&temporalWriter, name, params)
	if err != nil {
		return err
	}
	w.Write(temporalWriter.Bytes())
	return nil
}
