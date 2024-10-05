package utils

import "regexp"

func ValidatePassword(password string) bool {
    var passwordRegex = `^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[@$!%*?&])[A-Za-z\d@$!%*?&]{8,}$`
    re := regexp.MustCompile(passwordRegex)
    return re.MatchString(password)
}