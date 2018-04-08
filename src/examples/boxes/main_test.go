package main

import (
	"testing"
)

func TestBoxesMain(t *testing.T) {
	t.Run("Main application configuration", func(t *testing.T) {
		win, err := createWindow()
		one := win.GetChildAt(0)

		if err != nil {
			t.Error(err)
		}
		if win == nil {
			t.Error("Expected win to be returned from createWindow")
		}
		if win.GetChildCount() < 1 {
			t.Errorf("Expected at least 1 child on window, but got %d", win.GetChildCount())
		}
		if one == nil {
			t.Errorf("Expected at least one child")
		}
	})
}
