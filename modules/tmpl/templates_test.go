// Copyright 2015-2017 Geofrey Ernest <geofreyernest@live.com>. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package tmpl

import (
	"html/template"
	"strings"
	"testing"
)

func TestNew(t *testing.T) {
	cfg := &Config{
		Name:         "base",
		IncludesDirs: []string{"test/base", "test/tmpl"},
	}
	tp, err := New(cfg)
	if err != nil {
		t.Errorf("loading templates %v", err)
	}
	rst, err := tp.Render("test/tmpl/hello.tmpl", nil)
	if err != nil {
		t.Error(err)
	}
	if !strings.Contains(rst, "Nyumbani") {
		t.Errorf("expected %s to contain Nyumbani", rst)
	}
}

func TestFuncs(t *testing.T) {

	// script
	s := script("test/tmpl/script.html")
	if !strings.Contains(string(s), "script") {
		t.Errorf("expected scripts template got %s", s)
	}
	serr := script("/desk/auth/date_picker")
	if !strings.Contains(string(serr), "not found") {
		t.Errorf(" expected not found got %s", serr)
	}

	// dashed
	if dashed("mjoba_sili") != "mjoba-sili" {
		t.Errorf("expected mjoba-sili got %s", dashed("mjomba_sili"))
	}

	// label
	myLabel := "my_lady"
	if label(myLabel) != "my lady" {
		t.Errorf("expected my lady got %s", label(myLabel))
	}

	// toHtml
	htm := "<home>"
	if toHTML(htm) != template.HTML(htm) {
		t.Errorf("expected <hom> got %s", toHTML(htm))
	}

	// toMarkdown
	msg := "_hello_"
	md := toMarkdown(msg)
	msgHtm := "<p><em>hello</em></p>"
	if strings.TrimSpace(string(md)) != msgHtm {
		t.Error(md)
	}

	// translate
	if translate("en", "home-btn") != "Home" {
		t.Errorf("expected Home got %s", translate("ne", "home-btn"))
	}
	if translate("sw", "home-btn") != "Nyumbani" {
		t.Errorf("expected Nyumbani got %s", translate("ne", "home-btn"))
	}
}
