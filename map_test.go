package i18n

import "testing"

func TestMap(t *testing.T) {
	m := NewMap()
	m.Add("hello_world_b10a", "Hello World")
	m.Add("hello_name_a696", "Hello {name}")
	m.AddPlural("bag_count_0d3e", "{qty} bag", "{qty} bags")

	if val := m.Translate("hello_world_b10a"); val != "Hello World" {
		t.Error("Basic translation failed, got " + val)
	}

	if val := m.TranslateWith("hello_name_a696", Args("name", "John")); val != "Hello John" {
		t.Error("Arg translation failed, got " + val)
	}

	if val := m.TranslatePlural("bag_count_0d3e", 1, nil); val != "1 bag" {
		t.Error("Plural translation (singular) failed, got " + val)
	}
	if val := m.TranslatePlural("bag_count_0d3e", 3, nil); val != "3 bags" {
		t.Error("Plural translation (plural) failed, got " + val)
	}
}
