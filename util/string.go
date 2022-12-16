package util

import (
	"errors"
)

func CharFinder(input string, array string) error {

	for i := 0; i < len(input); i++ {
		var decline int = len(array) - 1
		var incline int = 0

		for decline >= incline {

			if input[i] == array[decline] {
				return nil
			}
			if input[i] == array[incline] {
				return nil
			}
			decline--
			incline++

		}

	}

	return errors.New("error!!!")
}
