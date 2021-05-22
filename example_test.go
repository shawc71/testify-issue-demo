package testify_issue_demo

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

type structWithTime struct {
	someTime   *time.Time
	someString string
}

func TestDemo(t *testing.T) {

	t.Run("same tzs", func(t *testing.T) {
		expectedTime, err := time.Parse("2006-01-02", "2020-05-22")
		assert.NoError(t, err)
		expected := &structWithTime{
			someTime:   &expectedTime,
			someString: "some val",
		}

		actualTime, err := time.Parse("2006-01-02", "2020-05-21")
		actual := &structWithTime{
			someTime:   &actualTime,
			someString: "some val",
		}
		assert.Equal(t, expected, actual)
	})

	t.Run("different tzs", func(t *testing.T) {
		expectedTime, err := time.Parse("2006-01-02", "2020-05-22")
		assert.NoError(t, err)
		expected := &structWithTime{
			someTime:   &expectedTime,
			someString: "some val",
		}

		loc, err := time.LoadLocation("America/New_York")
		assert.NoError(t, err)
		actualTime, err := time.ParseInLocation("2006-01-02", "2020-05-21", loc)
		actual := &structWithTime{
			someTime:   &actualTime,
			someString: "some val",
		}
		assert.Equal(t, expected, actual)
	})
}
