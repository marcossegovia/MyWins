package main

import "testing"

func TestExample(t *testing.T) {
        if testing.Short() {
                t.Skip("skipping test in short mode.")
        }
}
