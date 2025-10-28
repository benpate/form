package widget

import "github.com/benpate/html"

func loadingIcon(b *html.Builder) {
	b.Container("svg").
		Class("spin").
		Style("height:1.5em", "width:1.5em").
		Attr("xmlns", "http://www.w3.org/2000/svg").
		Attr("fill", "currentColor").
		Attr("viewBox", "0 0 16 16")

	b.Container("path").
		Attr("fill-rule", "evenodd").
		Attr("d", "M8 3a5 5 0 1 0 4.546 2.914.5.5 0 0 1 .908-.417A6 6 0 1 1 8 2z").
		Close()

	b.Container("path").
		Attr("d", "M8 4.466V.534a.25.25 0 0 1 .41-.192l2.36 1.966c.12.1.12.284 0 .384L8.41 4.658A.25.25 0 0 1 8 4.466").
		Close()

	b.Close()
}

func locateIcon(b *html.Builder) {
	b.Container("svg").
		Style("height:1.5em", "width:1.5em").
		Attr("xmlns", "http://www.w3.org/2000/svg").
		Attr("fill", "currentColor").
		Attr("viewBox", "0 0 16 16")

	b.Container("path").
		Attr("d", "M8.5.5a.5.5 0 0 0-1 0v.518A7 7 0 0 0 1.018 7.5H.5a.5.5 0 0 0 0 1h.518A7 7 0 0 0 7.5 14.982v.518a.5.5 0 0 0 1 0v-.518A7 7 0 0 0 14.982 8.5h.518a.5.5 0 0 0 0-1h-.518A7 7 0 0 0 8.5 1.018zm-6.48 7A6 6 0 0 1 7.5 2.02v.48a.5.5 0 0 0 1 0v-.48a6 6 0 0 1 5.48 5.48h-.48a.5.5 0 0 0 0 1h.48a6 6 0 0 1-5.48 5.48v-.48a.5.5 0 0 0-1 0v.48A6 6 0 0 1 2.02 8.5h.48a.5.5 0 0 0 0-1zM8 10a2 2 0 1 0 0-4 2 2 0 0 0 0 4").
		Close()

	b.Close()
}
