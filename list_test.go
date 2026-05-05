package i18n

import (
	"testing"

	"golang.org/x/text/language"
)

func TestLocalizedList(t *testing.T) {
	tests := []struct {
		name  string
		tag   language.Tag
		items []string
		want  string
	}{
		{
			name:  "empty",
			tag:   language.AmericanEnglish,
			items: nil,
			want:  "",
		},
		{
			name:  "one item",
			tag:   language.AmericanEnglish,
			items: []string{"LINK1"},
			want:  "LINK1",
		},
		{
			name:  "two items english",
			tag:   language.AmericanEnglish,
			items: []string{"LINK1", "LINK2"},
			want:  "LINK1 and LINK2",
		},
		{
			name:  "four items english",
			tag:   language.AmericanEnglish,
			items: []string{"LINK1", "LINK2", "LINK3", "LINK4"},
			want:  "LINK1, LINK2, LINK3 and LINK4",
		},
		{
			name:  "four items french",
			tag:   language.French,
			items: []string{"LIEN1", "LIEN2", "LIEN3", "LIEN4"},
			want:  "LIEN1, LIEN2, LIEN3 et LIEN4",
		},
		{
			name:  "four items chinese",
			tag:   language.SimplifiedChinese,
			items: []string{"链接1", "链接2", "链接3", "链接4"},
			want:  "链接1、链接2、链接3和链接4",
		},
	}

	for _, tt := range tests {
		if got := LocalizedList(tt.tag, tt.items); got != tt.want {
			t.Errorf("%s: got %q want %q", tt.name, got, tt.want)
		}
	}
}

func TestTranslatedList(t *testing.T) {
	m := NewMap()
	m.Add("link_1", "LINK1")
	m.Add("link_2", "LINK2")
	m.Add("link_3", "LINK3")
	m.Add("link_4", "LINK4")

	keys := []string{"link_1", "link_2", "link_3", "link_4"}
	if got := TranslatedList(language.AmericanEnglish, m, keys); got != "LINK1, LINK2, LINK3 and LINK4" {
		t.Fatalf("translated list failed, got %q", got)
	}
}

func TestCatalogTranslateList(t *testing.T) {
	m := NewMap()
	m.Add("link_1", "LINK1")
	m.Add("link_2", "LINK2")
	m.Add("link_3", "LINK3")
	m.Add("link_4", "LINK4")

	c := NewCatalog(language.AmericanEnglish, m)
	keys := []string{"link_1", "link_2", "link_3", "link_4"}

	if got := c.TranslateList(language.AmericanEnglish, keys); got != "LINK1, LINK2, LINK3 and LINK4" {
		t.Fatalf("catalog translated list failed, got %q", got)
	}
}
