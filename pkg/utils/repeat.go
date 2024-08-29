package utils

import "time"

func DoWithTries(fn func() error, tries int, delay time.Duration) error {
	for tries > 0 {
		if err := fn(); err != nil {
			time.Sleep(delay)
			tries--

			continue
		}
		tries = 0
		return nil
	}
	return nil
}
