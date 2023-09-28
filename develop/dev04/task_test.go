package main

import (
	"reflect"
	"testing"
)

func TestNoAnograms(t *testing.T) {
	cases := []struct {
		s    []string
		want map[string][]string
	}{
		{
			s:    nil,
			want: map[string][]string{},
		},
		{
			s:    []string{"один", "два"},
			want: map[string][]string{},
		},
	}

	for _, c := range cases {
		if got := getAnagrams(c.s); !reflect.DeepEqual(got, c.want) {
			t.Errorf("getAnagrams(%v)= %v, want %v", c.s, got, c.want)
		}
	}
}

func TestAnogramsLowerCase(t *testing.T) {
	cases := []struct {
		s    []string
		want map[string][]string
	}{
		{
			s: []string{"строка", "пятак", "пятка", "тяпка", "листок", "сорока", "слиток", "столик"},
			want: map[string][]string{
				"пятак":  {"пятак", "пятка", "тяпка"},
				"листок": {"листок", "слиток", "столик"},
			},
		},
	}

	for _, c := range cases {
		if got := getAnagrams(c.s); !reflect.DeepEqual(got, c.want) {
			t.Errorf("getAnagrams(%v)= %v, want %v", c.s, got, c.want)
		}
	}
}
func TestAnagramsBothCase(t *testing.T) {
	cases := []struct {
		s    []string
		want map[string][]string
	}{
		{
			s: []string{"строка", "ПяТак", "пЯткА", "ТЯПка", "лИСток", "сОвОка", "слИТОк", "сТОлик"},
			want: map[string][]string{
				"пятак":  {"пятак", "пятка", "тяпка"},
				"листок": {"листок", "слиток", "столик"},
			},
		},
	}

	for _, c := range cases {
		if got := getAnagrams(c.s); !reflect.DeepEqual(got, c.want) {
			t.Errorf("extractAnagrams(%v) = %v, want %v", c.s, got, c.want)
		}
	}
}
