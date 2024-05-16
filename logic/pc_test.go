package logic

import (
	"context"
	"testing"
)

func TestPcModel(t *testing.T) {
	m := NewPcModel(context.Background(), 50, 1, 10, 10)
	m.Run()
}

func TestPrintInTurn2(t *testing.T) {
	print_in_turn3("春江潮水连海平，海上明月共潮生", 100, 20)
}
