package main

func auth(n, p string) (bool) {
	users := map[string]string{
		"boris"	: "verysecure",
		"john"	: "notverysecure",
	}

	for name, pass := range users {
		if n == name && p == pass {
			return true
		}
		continue
	}
	return false
}