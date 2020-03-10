package todo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

// テストデータが保管されているファイルパスです。
const testDataPath = "testdata/fixture.json.golden"

// テスト名の構造体です。
type TestName struct {
	Name string `golden:"name"`
}

// テスト用のTodoの構造体です。
type TestTodo struct {
	Todo string `golden:"todo"`
	Done bool   `golden:"done"`
}

// 期待されるTodoリストの構造体です。
type Expecteds []TestTodo

// Add関数のテストデータを格納する構造体です。
type TestCaseAdd struct {
	TestName
	TestTodo
	Expected string `golden:"expected"`
}

// IsDone関数とIsNotDone関数のテストデータを格納する構造体です。
type TestCaseIsDones struct {
	TestName
	Done     bool `golden:"done"`
	Expected bool `golden:"expected"`
}

// List関数とDoneList関数のExpectedデータを格納する構造体です。
type TestCaseLists struct {
	TestName
	Expecteds `golden:"expecteds"`
}

// Done関数とDel関数のテストデータを格納する構造体です。
type TestCaseDoneDel struct {
	TestName
	TestTodo
	Expecteds `golden:"expecteds"`
}

// テスト対象の関数のテストケースを設定した構造体です。
type Fixture struct {
	Add       []TestCaseAdd     `golden:"Todo"`
	IsDone    []TestCaseIsDones `golden:"IsDone"`
	IsNotDone []TestCaseIsDones `golden:"IsNotDone"`
	List      []TestCaseLists   `golden:"List"`
	Done      []TestCaseDoneDel `golden:"Done"`
	DoneList  []TestCaseLists   `golden:"DoneList"`
	Del       []TestCaseDoneDel `golden:"Del"`
}

// ファイルパスからJSONファイルを読み込んだ後にデコードを行い、type Fixture を返します。
func getDecodeJSON(t *testing.T, path string) (fixture []Fixture) {
	t.Helper()
	// Read JSON file
	d, err := ioutil.ReadFile(path)
	if err != nil {
		t.Fatalf("ioutil.ReadAll(%s) got unexpected error %#v", testDataPath, err)
	}
	// Decode JSON file
	if err := json.Unmarshal(d, &fixture); err != nil {
		t.Fatalf("json.Unmarshal() got unexpected error %#v", err)
	}
	return
}

// ----------------------------------------------------------------------------------------
// Test Code
// ----------------------------------------------------------------------------------------

func TestMain(m *testing.M) {
	fmt.Printf("-------- before test --------\n\n")
	code := m.Run()
	fmt.Printf("\n-------- after test ---------\n")
	os.Exit(code)
}

func TestIsDone(t *testing.T) {
	// Read test data from JSON file.
	fixture := getDecodeJSON(t, testDataPath)

	// Executes a test based on the read test data.
	for idx, key := range fixture[0].IsDone {
		t.Logf("%v:%v / pattern: %v / expected: %v\n", idx, key.TestName.Name, key.Done, key.Expected)
		todo := Todo{Done: key.Done}
		d := isDone(todo)
		if key.Expected != d {
			t.Errorf("Error: Not equal\npattern : %d\nexpected: %v\nactual  : %v\nTest    : %v", idx, key.Expected, d, "TestIsDone")
		}
	}
}

func TestIsNotDone(t *testing.T) {
	// Read test data from JSON file.
	fixture := getDecodeJSON(t, testDataPath)

	// Executes a test based on the read test data.
	for idx, key := range fixture[0].IsNotDone {
		t.Logf("%v:%v / pattern: %v / expected: %v\n", idx, key.TestName.Name, key.Done, key.Expected)
		todo := Todo{Done: key.Done}
		d := isNotDone(todo)
		if key.Expected != d {
			t.Errorf("Error: Not equal\npattern : %d\nexpected: %v\nactual  : %v\nTest    : %v", idx, key.Expected, d, "TestIsNotDone")
		}
	}
}

func TestTodos(t *testing.T) {
	// Read test data from JSON file.
	fixture := getDecodeJSON(t, testDataPath)

	// Todoリストに未完了のTodoがない場合にerrorになるかのサブテスト
	t.Run("NotList", func(t *testing.T) {
		name := "NotTodoList"
		p, e := "nil", "Not nil"
		t.Logf("%v:%v / pattern: %v / expected: %v\n", 0, name, p, e)
		if _, err := List(); err == nil {
			t.Errorf("Error: Not equal\npattern : %v\nexpected: %v\nactual  : %v\nTest    : %v", 0, e, "nil", name)
		}
	})

	// Todoリストに完了済みのTodoがない場合にerrorになるかのサブテスト
	t.Run("NotDoneList", func(t *testing.T) {
		name := "NotTodoDoneList"
		p, e := "nil", "Not nil"
		t.Logf("%v:%v / pattern: %v / expected: %v\n", 0, name, p, e)
		if _, err := List(); err == nil {
			t.Errorf("Error: Not equal\npattern : %v\nexpected: %v\nactual  : %v\nTest    : %v", 0, e, "nil", name)
		}
	})

	// Add Todo
	t.Run("Add", func(t *testing.T) {
		testAdd(t, fixture)
	})

	// List Todo
	t.Run("List", func(t *testing.T) {
		n := fixture[0].List[0].TestName.Name
		e := fixture[0].List[0].Expecteds
		testList(t, n, e)
	})

	// Done Todo & List Todo
	t.Run("Done & List", func(t *testing.T) {
		len := len(fixture[0].Done)
		testDone(t, fixture, len)
		n := fixture[0].Done[len-1].TestName.Name
		e := fixture[0].Done[len-1].Expecteds
		testList(t, n, e)
	})

	// Done List Todo
	t.Run("DoneList", func(t *testing.T) {
		n := fixture[0].DoneList[0].TestName.Name
		e := fixture[0].DoneList[0].Expecteds
		testDoneList(t, n, e)
	})

	// Delete Todo
	t.Run("Del Todo", func(t *testing.T) {
		len := len(fixture[0].Del)
		testDel(t, fixture, len)
		n := fixture[0].Del[len-2].TestName.Name
		e := fixture[0].Del[len-2].Expecteds
		testList(t, n, e)
		n = fixture[0].Del[len-1].TestName.Name
		e = fixture[0].Del[len-1].Expecteds
		testDoneList(t, n, e)
	})
}

// Todoを追加するサブテストです。
func testAdd(t *testing.T, fixture []Fixture) {
	t.Helper()
	// Executes a test based on the read test data.
	for idx, key := range fixture[0].Add {
		t.Logf("%v:%v / pattern: %v / expected: %v\n", idx, key.TestName.Name, key.TestTodo.Todo, key.Expected)
		Add(key.Todo)
		if key.Expected != Todos[idx].Todo {
			t.Errorf("Error: Not equal\npattern : %d\nexpected: %v\nactual  : %v\nTest    : %v", idx, key.Expected, Todos[idx].Todo, "testAdd")
		}
	}
}

// 未完了のTodoリストの結果をテストするサブテストです。
func testList(t *testing.T, testName string, todos Expecteds) {
	t.Helper()
	result, _ := List()
	for idx, key := range todos {
		actT := result[idx]
		expT := key.Todo
		t.Logf("%v:%v / expected: %v / actual: %v\n", idx, testName, expT, actT)
		if expT != actT {
			t.Errorf("Error: Not equal\npattern : %d\nexpected: %v\nactual  : %v\nTest    : %v", idx, expT, actT, testName)
		}
	}
}

// ToDoを完了させるサブテストです。
func testDone(t *testing.T, fixture []Fixture, length int) {
	t.Helper()
	// Complete Todo tasks
	for idx, key := range fixture[0].Done {
		t.Logf("%v:%v / actual: %v\n", idx, key.TestName.Name, key.Todo)
		if !Done(key.Todo) {
			t.Errorf("Error: Not equal\npattern : %d\nexpected: %v\nactual  : %v\nTest    : %v", idx, true, false, key.TestName)
		}
		if idx == length-2 {
			break
		}
	}
}

// 完了済みのTodoリストの結果をテストするサブテストです。
func testDoneList(t *testing.T, testName string, todos Expecteds) {
	t.Helper()
	result, _ := DoneList()
	for idx, key := range todos {
		actT := result[idx]
		expT := key.Todo
		t.Logf("%v:%v / expected: %v / actual: %v\n", idx, testName, expT, actT)
		if expT != actT {
			t.Errorf("Error: Not equal\npattern : %d\nexpected: %v\nactual  : %v\nTest    : %v", idx, expT, actT, testName)
		}
	}
}

// タスクをTodoから削除するサブテストです。
func testDel(t *testing.T, fixture []Fixture, length int) {
	t.Helper()
	for idx, key := range fixture[0].Del {
		t.Logf("%v:%v / actual: %v\n", idx, key.TestName.Name, key.Todo)
		if !Del(key.Todo) {
			t.Errorf("Error: Not equal\npattern : %d\nexpected: %v\nactual  : %v\nTest    : %v", idx, true, false, key.TestName)
		}
		if idx == length-3 {
			break
		}
	}
}
