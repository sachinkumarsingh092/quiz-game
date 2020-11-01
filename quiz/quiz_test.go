package quiz

import (
	"testing"
	"time"
)

func TestRunQuiz(t *testing.T) {
	var tests = []struct {
		question string
		answer   string
	}{
		{"1+1", "2"},
		{"3+2", "5"},
		{"10+11", "21"},
		{"", ""},
	}

	timer := time.NewTimer(time.Duration(2) * time.Second).C
	done := make(chan string)
	var ans int
	var err error
	allDone := make(chan bool)
	for _, test := range tests {
		go func() {
			ans, err = runQuiz(test.question, test.answer, done, timer)
			allDone <- true
		}()
		done <- test.answer
		<-allDone

		t.Logf("want:%v, get:%v\n", 1, ans)

		if err != nil {
			t.Error(err)
		}

		if ans != 1 {
			t.Errorf("want:%v, get:%v\n", 1, ans)
		}
	}

}
