package html // import "github.com/tdewolff/minify/html"

import "github.com/tdewolff/parse/html"

type traits uint8

const (
	rawTag traits = 1 << iota
	nonPhrasingTag
	booleanAttr
	caselessAttr
	urlAttr
)

var tagMap = map[html.Hash]traits{
	html.Address:    nonPhrasingTag,
	html.Article:    nonPhrasingTag,
	html.Aside:      nonPhrasingTag,
	html.Blockquote: nonPhrasingTag,
	html.Body:       nonPhrasingTag,
	html.Br:         nonPhrasingTag,
	html.Canvas:     nonPhrasingTag,
	html.Caption:    nonPhrasingTag,
	html.Code:       rawTag,
	html.Col:        nonPhrasingTag,
	html.Colgroup:   nonPhrasingTag,
	html.Dd:         nonPhrasingTag,
	html.Div:        nonPhrasingTag,
	html.Dl:         nonPhrasingTag,
	html.Dt:         nonPhrasingTag,
	html.Embed:      nonPhrasingTag,
	html.Fieldset:   nonPhrasingTag,
	html.Figcaption: nonPhrasingTag,
	html.Figure:     nonPhrasingTag,
	html.Footer:     nonPhrasingTag,
	html.Form:       nonPhrasingTag,
	html.H1:         nonPhrasingTag,
	html.H2:         nonPhrasingTag,
	html.H3:         nonPhrasingTag,
	html.H4:         nonPhrasingTag,
	html.H5:         nonPhrasingTag,
	html.H6:         nonPhrasingTag,
	html.Head:       nonPhrasingTag,
	html.Header:     nonPhrasingTag,
	html.Hgroup:     nonPhrasingTag,
	html.Hr:         nonPhrasingTag,
	html.Html:       nonPhrasingTag,
	html.Iframe:     rawTag,
	html.Li:         nonPhrasingTag,
	html.Main:       nonPhrasingTag,
	html.Math:       rawTag,
	html.Meta:       nonPhrasingTag,
	html.Nav:        nonPhrasingTag,
	html.Noscript:   nonPhrasingTag,
	html.Ol:         nonPhrasingTag,
	html.Output:     nonPhrasingTag,
	html.P:          nonPhrasingTag,
	html.Pre:        rawTag | nonPhrasingTag,
	html.Progress:   nonPhrasingTag,
	html.Script:     rawTag,
	html.Section:    nonPhrasingTag,
	html.Style:      rawTag | nonPhrasingTag,
	html.Svg:        rawTag,
	html.Table:      nonPhrasingTag,
	html.Tbody:      nonPhrasingTag,
	html.Td:         nonPhrasingTag,
	html.Textarea:   rawTag,
	html.Tfoot:      nonPhrasingTag,
	html.Th:         nonPhrasingTag,
	html.Thead:      nonPhrasingTag,
	html.Title:      nonPhrasingTag,
	html.Tr:         nonPhrasingTag,
	html.Ul:         nonPhrasingTag,
	html.Video:      nonPhrasingTag,
}

var attrMap = map[html.Hash]traits{
	html.Accept:          caselessAttr,
	html.Accept_Charset:  caselessAttr,
	html.Action:          urlAttr,
	html.Align:           caselessAttr,
	html.Alink:           caselessAttr,
	html.Allowfullscreen: booleanAttr,
	html.Async:           booleanAttr,
	html.Autofocus:       booleanAttr,
	html.Autoplay:        booleanAttr,
	html.Axis:            caselessAttr,
	html.Background:      urlAttr,
	html.Bgcolor:         caselessAttr,
	html.Charset:         caselessAttr,
	html.Checked:         booleanAttr,
	html.Cite:            urlAttr,
	html.Classid:         urlAttr,
	html.Clear:           caselessAttr,
	html.Codebase:        urlAttr,
	html.Codetype:        caselessAttr,
	html.Color:           caselessAttr,
	html.Compact:         booleanAttr,
	html.Controls:        booleanAttr,
	html.Data:            urlAttr,
	html.Declare:         booleanAttr,
	html.Default:         booleanAttr,
	html.DefaultChecked:  booleanAttr,
	html.DefaultMuted:    booleanAttr,
	html.DefaultSelected: booleanAttr,
	html.Defer:           booleanAttr,
	html.Dir:             caselessAttr,
	html.Disabled:        booleanAttr,
	html.Draggable:       booleanAttr,
	html.Enabled:         booleanAttr,
	html.Enctype:         caselessAttr,
	html.Face:            caselessAttr,
	html.Formaction:      urlAttr,
	html.Formnovalidate:  booleanAttr,
	html.Frame:           caselessAttr,
	html.Hidden:          booleanAttr,
	html.Href:            urlAttr,
	html.Hreflang:        caselessAttr,
	html.Http_Equiv:      caselessAttr,
	html.Icon:            urlAttr,
	html.Inert:           booleanAttr,
	html.Ismap:           booleanAttr,
	html.Itemscope:       booleanAttr,
	html.Lang:            caselessAttr,
	html.Language:        caselessAttr,
	html.Link:            caselessAttr,
	html.Longdesc:        urlAttr,
	html.Manifest:        urlAttr,
	html.Media:           caselessAttr,
	html.Method:          caselessAttr,
	html.Multiple:        booleanAttr,
	html.Muted:           booleanAttr,
	html.Nohref:          booleanAttr,
	html.Noresize:        booleanAttr,
	html.Noshade:         booleanAttr,
	html.Novalidate:      booleanAttr,
	html.Nowrap:          booleanAttr,
	html.Open:            booleanAttr,
	html.Pauseonexit:     booleanAttr,
	html.Poster:          urlAttr,
	html.Profile:         urlAttr,
	html.Readonly:        booleanAttr,
	html.Rel:             caselessAttr,
	html.Required:        booleanAttr,
	html.Rev:             caselessAttr,
	html.Reversed:        booleanAttr,
	html.Rules:           caselessAttr,
	html.Scope:           caselessAttr,
	html.Scoped:          booleanAttr,
	html.Scrolling:       caselessAttr,
	html.Seamless:        booleanAttr,
	html.Selected:        booleanAttr,
	html.Shape:           caselessAttr,
	html.Sortable:        booleanAttr,
	html.Spellcheck:      booleanAttr,
	html.Src:             urlAttr,
	html.Target:          caselessAttr,
	html.Text:            caselessAttr,
	html.Translate:       booleanAttr,
	html.Truespeed:       booleanAttr,
	html.Type:            caselessAttr,
	html.Typemustmatch:   booleanAttr,
	html.Undeterminate:   booleanAttr,
	html.Usemap:          urlAttr,
	html.Valign:          caselessAttr,
	html.Valuetype:       caselessAttr,
	html.Vlink:           caselessAttr,
	html.Visible:         booleanAttr,
	html.Xmlns:           urlAttr,
}