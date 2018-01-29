package data

import (
	"time"
)

type Thread struct {
	Id        int
	Uuid      string
	Topic     string
	UserId    int
	CreatedAt time.Time
}

func Threads() (threads []Thread, err error) {
	rows, err := Db.Query("SELECT id, uuid, topic, user_id, created_at FROM threads ORDER BY created_at DESC")
	if err != nil {
		return
	}

	for rows.Next() {
		thr := Thread{}
		err = rows.Scan(&thr.Id, &thr.Uuid, &thr.Topic, &thr.UserId, &thr.CreatedAt)
		if err != nil {
			return
		}
		threads = append(threads, thr)
	}
	rows.Close()
	return
}

func (thread *Thread) NumReplies() (count int) {
	rows, err := Db.Query("SELECT count(*) FROM posts WHERE thread_id=$1", thread.Id)
	if err != nil {
		return
	}
	for rows.Next() {
		err = rows.Scan(&count)
		if err != nil {
			return
		}
	}
	rows.Close()

	return
}
