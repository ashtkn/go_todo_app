package store

import (
	"context"
	"testing"

	"github.com/ashtkn/go_todo_app/clock"
	"github.com/ashtkn/go_todo_app/entity"
	"github.com/ashtkn/go_todo_app/testutil"
	"github.com/google/go-cmp/cmp"
)

func TestRepository_ListTask(t *testing.T) {
	ctx := context.Background()

	tx, err := testutil.OpenDBForTest(t).BeginTxx(ctx, nil)
	// テストケースが完了したらロールバックする
	t.Cleanup(func() { _ = tx.Rollback() })
	if err != nil {
		t.Fatal(err)
	}

	wants := prepareTasks(ctx, t, tx)

	sut := &Repository{}
	gots, err := sut.ListTasks(ctx, tx)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if d := cmp.Diff(gots, wants); len(d) != 0 {
		t.Errorf("differs: (-got +want)\n%s", d)
	}
}

func prepareTasks(ctx context.Context, t *testing.T, con Execer) entity.Tasks {
	t.Helper()

	if _, err := con.ExecContext(ctx, `DELETE FROM task;`); err != nil {
		t.Logf("failed to initialize task: %v", err)
	}

	c := clock.FixedClocker{}

	wants := entity.Tasks{
		&entity.Task{
			Title:    "want task 1",
			Status:   "todo",
			Created:  c.Now(),
			Modified: c.Now(),
		},
		&entity.Task{
			Title:    "want task 2",
			Status:   "todo",
			Created:  c.Now(),
			Modified: c.Now(),
		},
		&entity.Task{
			Title:    "want task 3",
			Status:   "done",
			Created:  c.Now(),
			Modified: c.Now(),
		},
	}

	results, err := con.ExecContext(ctx,
		`INSERT INTO task (title, status, created, modified)
			VALUES
				(?, ?, ?, ?),
				(?, ?, ?, ?),
				(?, ?, ?, ?);`,
		wants[0].Title, wants[0].Status, wants[0].Created, wants[0].Modified,
		wants[1].Title, wants[1].Status, wants[1].Created, wants[1].Modified,
		wants[2].Title, wants[2].Status, wants[2].Created, wants[2].Modified,
	)

	if err != nil {
		t.Fatal(err)
	}

	// MySQLの場合はLastInsertId()で発行されたIDのうち「最も小さなID」を取得できる
	id, err := results.LastInsertId()
	if err != nil {
		t.Fatal(err)
	}

	wants[0].ID = entity.TaskID(id)
	wants[1].ID = entity.TaskID(id + 1)
	wants[2].ID = entity.TaskID(id + 2)

	return wants
}
