package printalphabet

import "ft"

func printalphabet() err {
	for i := 'A'; i <= 'Z'; i++ {
		err = ft.PrintRune(i)
    }
	return err
}