package template

import (
	"fmt"
	"github.com/rasa/shortme/conf"
	"github.com/rasa/shortme/static"
	"html/template"
	"io/ioutil"
	"log"
)

const (
	staticDir = "/static"
	cssDir    = "/css"
	jsDir     = "/js"
	cssFile   = "<link href=\"%s\" rel=\"stylesheet\"></link>\n"
	jsFile    = "<script src=\"%s\"></script>"
	// see http://www.webdevout.net/articles/escaping-style-and-script-data
	cssInline = "\n<style type=\"text/css\">\n<!--/*--><![CDATA[/*><!-- */\n/* %s */\n%s\n/*]]>*/-->\n</style>\n"
	jsInline  = "\n<script type=\"text/javascript\">\n<!--//--><![CDATA[//><!-- \n/* %s */\n%s\n//--><!]]>\n</script>\n"
)

var inline = true

type Index struct {
	Title                            string
	ShortURL                         string
	Bootstrap_min_css                template.HTML
	Ie10_viewport_bug_workaround_css template.HTML
	Shortme_css                      template.HTML

	Bootstrap_min_js                template.HTML
	Html5shiv_min_js                template.HTML
	Ie10_viewport_bug_workaround_js template.HTML
	Ie_emulation_modes_warning_js   template.HTML
	// Ie8_responsive_file_warning_js template.HTML
	Jquery_min_js        template.HTML
	Jquery_qrcode_min_js template.HTML
	Respond_min_js       template.HTML
	Shortme_js           template.HTML
}

func include(file string, fileFmt string, inlineFmt string) template.HTML {
	if !inline {
		return template.HTML(fmt.Sprintf(fileFmt, staticDir+file))
	}

	fh, err := static.Assets.Open(file)
	if err != nil {
		log.Fatalf("Failed to open %v: %v", file, err)
	}
	defer fh.Close()
	data, err := ioutil.ReadAll(fh)
	if err != nil {
		log.Fatalf("Failed to read %v: %v", file, err)
	}
	return template.HTML(fmt.Sprintf(inlineFmt, staticDir+file, string(data)))
}

func css(file string) template.HTML {
	return include(cssDir+"/"+file, cssFile, cssInline)
}

func js(file string) template.HTML {
	return include(jsDir+"/"+file, jsFile, jsInline)
}

var Vars Index

func Init() {
	Vars.Title = conf.Conf.Common.Title
	Vars.ShortURL = conf.Conf.Common.ShortURL
	Vars.Bootstrap_min_css = css("bootstrap.min.css")
	Vars.Ie10_viewport_bug_workaround_css = css("ie10-viewport-bug-workaround.css")
	Vars.Shortme_css = css("shortme.css")

	Vars.Bootstrap_min_js = js("bootstrap.min.js")
	Vars.Html5shiv_min_js = js("html5shiv.min.js")
	Vars.Ie10_viewport_bug_workaround_js = js("ie10-viewport-bug-workaround.js")
	Vars.Ie_emulation_modes_warning_js = js("ie-emulation-modes-warning.js")
	// Vars.Ie8_responsive_file_warning_js = js("ie8-responsive-file-warning.js")
	Vars.Jquery_min_js = js("jquery.min.js")
	Vars.Jquery_qrcode_min_js = js("jquery.qrcode.min.js")
	Vars.Respond_min_js = js("respond.min.js")
	Vars.Shortme_js = js("shortme.js")
}

/*
<link href="/static/css/shortme.css" rel="stylesheet">
<link href="https://cdnjs.cloudflare.com/ajax/libs/twitter-bootstrap/4.1.3/css/bootstrap.min.css" integrity="sha384-MCw98/SFnGE8fJT3GXwEOngsV7Zt27NXFoaoApmYm81iuXoPkFOJwJ8ERdknLPMO" crossorigin="anonymous" rel="stylesheet">
<link href="https://maxcdn.bootstrapcdn.com/css/ie10-viewport-bug-workaround.css" integrity="sha384-sPZtnf00OKtohgCjYPiHYhx4vOa4gD182/f6+z1N+LzXdE2e5bENObUq17/FrRFS" crossorigin="anonymous" rel="stylesheet">
<script src="/static/js/shortme.js"></script>
<script src="https://cdn.jsdelivr.net/npm/bootstrap-offcanvas@1.0.0/ie-emulation-modes-warning.js" integrity="sha256-yEQpKbHsFQ0y3701lo5beguiJuTGFD+Nb3Gq2k3+wOo=" crossorigin="anonymous"></script>
<script src="https://cdnjs.cloudflare.com/ajax/libs/html5shiv/3.7.3/html5shiv.min.js" integrity="sha384-qFIkRsVO/J5orlMvxK1sgAt2FXT67og+NyFTITYzvbIP1IJavVEKZM7YWczXkwpB" crossorigin="anonymous"></script>
<script src="https://cdnjs.cloudflare.com/ajax/libs/jquery.qrcode/1.0/jquery.qrcode.min.js" integrity="sha256-9MzwK2kJKBmsJFdccXoIDDtsbWFh8bjYK/C7UjB1Ay0=" crossorigin="anonymous"></script>
<script src="https://cdnjs.cloudflare.com/ajax/libs/jquery/3.3.1/jquery.min.js" integrity="sha384-tsQFqpEReu7ZLhBV2VZlAu7zcOV+rXbYlF2cqB8txI/8aZajjp4Bqd+V6D5IgvKT" crossorigin="anonymous"></script>
<script src="https://cdnjs.cloudflare.com/ajax/libs/respond.js/1.4.2/respond.min.js" integrity="sha256-g6iAfvZp+nDQ2TdTR/VVKJf3bGro4ub5fvWSWVRi2NE=" crossorigin="anonymous"></script>
<script src="https://cdnjs.cloudflare.com/ajax/libs/twitter-bootstrap/4.1.3/js/bootstrap.min.js" integrity="sha384-ChfqqxuZUCnJSK3+MXmPNIyE6ZbWh2IMqE241rYiqJxyMiZ6OW/JmZQ5stwEULTy" crossorigin="anonymous"></script>
<script src="https://maxcdn.bootstrapcdn.com/js/ie10-viewport-bug-workaround.js" integrity="sha384-EZKKO3vHj6CHKQPIi5+Ubzvx7GjCAfgb/28vGjgly8qKb2DMq7V5D2o//Bjp9z03" crossorigin="anonymous"></script>
*/
