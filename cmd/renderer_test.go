package cmd

import (
	"strconv"
	"strings"
	"testing"
)

func TestCreateData(t *testing.T) {
	width, height := 256, 256
	data := CreateData(width, height)
	var expected strings.Builder

	expected.WriteString("P3\n" + strconv.Itoa(width) + " " + strconv.Itoa(height) + "\n255\n")
	formatLen := len(expected.String())
	t.Run("正しいフォーマットが出力されること", func(t *testing.T) {
		if !strings.HasPrefix(data, expected.String()) {
			t.Errorf("expected format %v", expected.String())
			t.Errorf("result of format %v", data[:formatLen])
		}
	})

	t.Run("各色は0以上255以下であること", func(t *testing.T) {
		image := data[formatLen:]

		for _, row := range strings.Split(image, "\n") {
			for _, pixel := range strings.Fields(row) {
				color, err := strconv.Atoi(pixel)
				if err != nil {
					t.Error(err)
				}
				if color < 0 || color > 255 {
					t.Errorf("Color is greater than or equal to 0 and less than or equal to 255.")
					t.Errorf("result: %d", color)
				}
			}
		}
	})
}
