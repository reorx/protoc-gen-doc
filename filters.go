package gendoc

import (
	"fmt"
	"html/template"
	"regexp"
	"strings"
)

var (
	paraPattern         = regexp.MustCompile(`(\n|\r|\r\n)\s*`)
	spacePattern        = regexp.MustCompile("( )+")
	multiNewlinePattern = regexp.MustCompile(`(\r\n|\r|\n){2,}`)
)

// PFilter splits the content by new lines and wraps each one in a <p> tag.
func PFilter(content string) template.HTML {
	paragraphs := paraPattern.Split(content, -1)
	return template.HTML(fmt.Sprintf("<p>%s</p>", strings.Join(paragraphs, "</p><p>")))
}

// ParaFilter splits the content by new lines and wraps each one in a <para> tag.
func ParaFilter(content string) string {
	paragraphs := paraPattern.Split(content, -1)
	return fmt.Sprintf("<para>%s</para>", strings.Join(paragraphs, "</para><para>"))
}

// NoBrFilter removes single CR and LF from content.
func NoBrFilter(content string) string {
	normalized := strings.Replace(content, "\r\n", "\n", -1)
	paragraphs := multiNewlinePattern.Split(normalized, -1)
	for i, p := range paragraphs {
		withoutCR := strings.Replace(p, "\r", " ", -1)
		withoutLF := strings.Replace(withoutCR, "\n", " ", -1)
		paragraphs[i] = spacePattern.ReplaceAllString(withoutLF, " ")
	}
	return strings.Join(paragraphs, "\n\n")
}

// Anchor formats content as an anchor string, replace . with -
func Anchor(content string) string {
	return strings.Replace(content, ".", "-", -1)
}

var injectTagRe = regexp.MustCompile(`@inject_tag: \x60.+\x60`)

// RefineComment formats comment, put <br/> before @inject_tag
func RefineComment(content string) string {
	// handle @inject_tag, assume only one @inject_tag
	injectTagPrefix := strings.HasPrefix(content, "@inject_tag")
	content = injectTagRe.ReplaceAllStringFunc(content, func(v string) string {
		v = fmt.Sprintf("<code>%s</code>", v)
		if injectTagPrefix {
			return v + "<br>"
		}
		return "<br>" + v
	})

	// trim
	content = strings.Trim(content, "\n ")
	content = strings.TrimSuffix(content, "<br>")
	return content
}
