package main

import (
    "testing"
)

// Test function for Add
func TestAdd(t *testing.T) {
    // Test cases
    cases := []struct {
        a, b, expected int
    }{
        {1, 2, 3},
        {0, 0, 0},
        {-1, 1, 0},
        {5, -2, 3},
    }

    // Iterate through test cases
    for _, c := range cases {
        result := Add(c.a, c.b)
        if result != c.expected {
            t.Errorf("Add(%d, %d) == %d, expected %d", c.a, c.b, result, c.expected)
        }
    }
}
