package base

import (
	"testing"

	"github.com/voicera/tester/assert"
)

func AnswerTheUltimateQuestion() (int, error) {
	return 30, nil
}

func TestDeepThought(t *testing.T) {
	answer, err := AnswerTheUltimateQuestion()
	if assert.For(t).ThatActualError(err).IsNil().Passed() {
		assert.For(t).ThatActual(answer).Equals(42)
	}
}
