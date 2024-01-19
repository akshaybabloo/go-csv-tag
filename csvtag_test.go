package csvtag

import (
	"fmt"
	"testing"
)

func TestDumpAndLoad(t *testing.T) {
	str, err := DumpToString(tabTest)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	tabT := []test{}
	err = LoadFromString(str, &tabT)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if tabT[0].ID != tabTest[0].ID ||
		tabT[0].Name != tabTest[0].Name ||
		tabT[0].Num != tabTest[0].Num {
		t.Fail()
	}
}

func TestDumpAndLoadTagKey(t *testing.T) {
	str, err := DumpToString(tabTestTagKey, CsvOptions{TagKey: "json"})
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	tabT := []testTagKey{}
	err = LoadFromString(str, &tabT, CsvOptions{TagKey: "json"})
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if tabT[0].ID != tabTestTagKey[0].ID ||
		tabT[0].Name != tabTestTagKey[0].Name ||
		tabT[0].Num != tabTestTagKey[0].Num {
		t.Fail()
	}
}
