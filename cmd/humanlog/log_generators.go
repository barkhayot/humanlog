package main

import (
	"fmt"
	"time"

	"golang.org/x/exp/rand"
)

func generateLogfmtLog(now time.Time, src rand.Source) string {
	return fmt.Sprintf(
		"time=%s level=%s msg=%q user=%s org=%s",
		now.Format(time.RFC3339),
		randel(src, []string{"INFO", "DEBUG", "WARN", "ERROR"}),
		randel(src, nouns)+" "+randel(src, adjectives),
		genString(src, false),
		genString(src, false),
	)
}

func generateJSONLog(now time.Time, src rand.Source) string {
	return fmt.Sprintf(
		`{"time":"%s","level":"%s","message":"%s","user":"%s","org":"%s"}`,
		now.Format(time.RFC3339),
		randel(src, []string{"INFO", "DEBUG", "WARN", "ERROR"}),
		randel(src, nouns)+" "+randel(src, adjectives),
		genString(src, false),
		genString(src, false),
	)
}

func generateOtelLog(now time.Time, src rand.Source) string {
	return fmt.Sprintf(
		`{"time":"%s","severity":"%s","body":"%s","attributes":{"user":"%s","org":"%s"}}`,
		now.Format(time.RFC3339Nano),
		randel(src, []string{"INFO", "DEBUG", "WARN", "ERROR"}),
		randel(src, nouns)+" "+randel(src, adjectives),
		genString(src, false),
		genString(src, false),
	)
}
