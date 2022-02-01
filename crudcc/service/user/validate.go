package user

import "regexp"

func validateid(id int) bool {
	return id >= 1
}

func validatemail(email string) bool {
	emailreg := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailreg.MatchString(email)

}

func validp(phone string) bool {
	phonereg := regexp.MustCompile(`^[+]*[(]{0,1}[0-9]{1,4}[)]{0,1}[-\s\./0-9]*$`)
	return phonereg.MatchString(phone)
}
