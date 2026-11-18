package rotator

import (
	"testing"
	"time"
)

func TestClosedPartitionCutoff(t *testing.T) {
	loc := time.FixedZone("UTC+3", 3*3600)
	cases := []struct {
		now      time.Time
		stepMin  int
		wantDate string
		wantTime string
	}{
		{time.Date(2026, 11, 18, 15, 40, 0, 0, loc), 60, "20261118", "1500"},
		{time.Date(2026, 11, 18, 15, 40, 0, 0, loc), 120, "20261118", "1400"},
		{time.Date(2026, 11, 18, 15, 40, 0, 0, loc), 720, "20261118", "1200"},
		{time.Date(2026, 11, 18, 0, 10, 0, 0, loc), 60, "20261118", "0000"},
		{time.Date(2026, 11, 18, 0, 0, 0, 0, loc), 60, "20261118", "0000"},
	}
	for _, tc := range cases {
		date, hm := closedPartitionCutoff(tc.now, tc.stepMin)
		if date != tc.wantDate || hm != tc.wantTime {
			t.Fatalf("now=%s step=%d: got %s %s, want %s %s", tc.now, tc.stepMin, date, hm, tc.wantDate, tc.wantTime)
		}
	}
}
