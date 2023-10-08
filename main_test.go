package main

import (
	"testing"
)

func TestFormatPath(t *testing.T) {
	type testCase struct {
		title    string
		input    string
		home     string
		expected string
	}

	for _, c := range []testCase{
		{
			title:    "Absolute path in home dir",
			input:    "/home/test/projects/projectA",
			home:     "/home/test",
			expected: "~/p/projectA",
		},
		{
			title:    "Absolute path not in home dir",
			input:    "/opt/test/projects/projectA",
			home:     "/home/test",
			expected: "/o/t/p/projectA",
		},
		{
			title:    "Relative path",
			input:    "../projects/important/projectA",
			home:     "/home/test",
			expected: "../p/i/projectA",
		},
		{
			title:    "Relative path which include home path",
			input:    "../home/test/important/projectA",
			home:     "/home/test",
			expected: "../h/t/i/projectA",
		},
		{
			title:    "Single element path",
			input:    "test",
			home:     "/home/test",
			expected: "test",
		},
		{
			title:    "Just root",
			input:    "/",
			home:     "/home/test",
			expected: "/",
		},
	} {
		testCase := c

		t.Run(testCase.title, func(t *testing.T) {
			t.Setenv("HOME", testCase.home)

			result := formatPath(testCase.input)

			if result != testCase.expected {
				t.Logf("unexpected result: got %s, expected: %s", result, testCase.expected)
				t.Fail()
			}
		})

	}
}
