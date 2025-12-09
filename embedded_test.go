package i18n

import (
	"embed"
	"testing"
)

//go:embed testdata/*
var embeddedTestFS embed.FS

func TestEmbeddedJSON(t *testing.T) {
	tr, err := NewEmbeddedFile(embeddedTestFS, "testdata/en.json")
	if err != nil {
		t.Fatalf("failed to load embedded json: %v", err)
	}

	if val := tr.Translate("hello_world_b10a"); val != "Hello World" {
		t.Error("Basic JSON translation failed, got " + val)
	}

	if val := tr.TranslateWith("hello_name_a696", Args("name", "John")); val != "Hello John" {
		t.Error("Arg JSON translation failed, got " + val)
	}
}

func TestEmbeddedYAML(t *testing.T) {
	tr, err := NewEmbeddedFile(embeddedTestFS, "testdata/en.yaml")
	if err != nil {
		t.Fatalf("failed to load embedded yaml: %v", err)
	}

	if val := tr.Translate("hello_world_b10a"); val != "Hello World" {
		t.Error("Basic YAML translation failed, got " + val)
	}

	if val := tr.TranslateWith("hello_name_a696", Args("name", "Jane")); val != "Hello Jane" {
		t.Error("Arg YAML translation failed, got " + val)
	}

	if val := tr.TranslatePlural("qty_bag_s_full_19772c_17", 1, nil); val != "1 bag full" {
		t.Error("Plural YAML translation (singular) failed, got " + val)
	}
	if val := tr.TranslatePlural("qty_bag_s_full_19772c_17", 2, nil); val != "2 bags full" {
		t.Error("Plural YAML translation (plural) failed, got " + val)
	}

	if val := tr.TranslatePlural("qty_cat_s_c712", 1, nil); val != "1 cat" {
		t.Error("Plural YAML translation (singular) failed, got " + val)
	}
	if val := tr.TranslatePlural("qty_cat_s_c712", 2, nil); val != "2 cats" {
		t.Error("Plural YAML translation (plural) failed, got " + val)
	}
}
