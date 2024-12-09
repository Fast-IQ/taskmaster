package taskmaster

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDayOfWeek_String(t *testing.T) {
	d := Monday
	assert.EqualValues(t, "Monday", d.String())
}

func TestMonth_String(t *testing.T) {
	m := April
	assert.EqualValues(t, "April", m.String())

	m = AllMonths
	assert.EqualValues(t, "All months", m.String())

}

func TestWeek_String(t *testing.T) {
	w := Second
	assert.EqualValues(t, "Second", w.String())
}
