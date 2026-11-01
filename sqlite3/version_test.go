package sqlite3

import "testing"

func TestBundledSQLiteVersion(t *testing.T) {
	if got := Version(); got != "3.54.1" {
		t.Fatalf("Version() = %q, want %q", got, "3.54.1")
	}
}

func TestUpdateDeleteLimitStillEnabled(t *testing.T) {
	conn, err := Open(":memory:")
	if err != nil {
		t.Fatalf("Open: %v", err)
	}
	defer conn.Close()

	if err := conn.Exec(`CREATE TABLE student(name TEXT, age INTEGER)`); err != nil {
		t.Fatalf("create table: %v", err)
	}
	if err := conn.Exec(`INSERT INTO student VALUES ('Bob', 18), ('Jane', 19)`); err != nil {
		t.Fatalf("insert students: %v", err)
	}
	if err := conn.Exec(`DELETE FROM student LIMIT 1`); err != nil {
		t.Fatalf("DELETE with LIMIT should be enabled: %v", err)
	}

	stmt, err := conn.Prepare(`SELECT COUNT(*) FROM student`)
	if err != nil {
		t.Fatalf("prepare count: %v", err)
	}
	defer stmt.Close()
	hasRow, err := stmt.Step()
	if err != nil {
		t.Fatalf("step count: %v", err)
	}
	if !hasRow {
		t.Fatal("count query returned no row")
	}
	var count int
	if err := stmt.Scan(&count); err != nil {
		t.Fatalf("scan count: %v", err)
	}
	if count != 1 {
		t.Fatalf("count = %d, want 1", count)
	}
}
