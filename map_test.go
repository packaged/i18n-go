package translate

import "testing"

func TestMap(t *testing.T) {
	m := NewMap()
	m.Add("hello_world_b10a", "Hello World")
	m.Add("hello_name_a696", "Hello {name}")

	if val := m.Translate("hello_world_b10a"); val != "Hello World" {
		t.Error("Basic translation failed, got " + val)
	}

	if val := m.TranslateWith("hello_name_a696", Args("name", "John")); val != "Hello John" {
		t.Error("Arg translation failed, got " + val)
	}
}
