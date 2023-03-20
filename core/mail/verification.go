package mail

import (
	"GoAPIfy/core/math"
	"GoAPIfy/model"
	"GoAPIfy/service/appService"
	"fmt"
)

// VerificationMail sends a verification email to the specified user's email address,
// containing a link to verify their account. The function generates a random verification
// token, saves it to the database, and constructs a verification link using the token
// and the provided URL as the domain name. The email message includes the verification
// link and a subject line indicating that the user should verify their account. The
// function returns an error if the email fails to send, or if there is an error creating
// the verification token.
//
// Parameters:
// - s: an instance of appService.AppService which is used to interact with the
// database and Meilisearch server.
// - userData: an instance of model.User representing the user to whom the
// verification email should be sent.
// - url: a string representing the domain name and path to be used when constructing
// the verification link. The URL should include the scheme (http:// or https://),
// domain name, and path to the verification endpoint, separated by forward slashes (/).
// For example, "https://example.com/verify". The user's email and the randomly-generated
// verification token will be appended to the URL path as path parameters. The URL that
// generated will be formatted like "https://example.com/verify?email=example@mail.com&token=someRandomString"
//
// Returns:
// - error: an error if the email fails to send, or if there is an error creating the verification token.
func VerificationMail(s appService.AppService, userData model.User, url string) error {
	// Generate the verification token
	token := math.RandomString(32)
	if err := s.Model.Load(&model.EmailVerification{Token: token}).Save(); err != nil {
		return err
	}

	verifyLink := fmt.Sprintf("%s?email=%s&token=%s", url, userData.Email, token)

	// Construct the email message
	to := []string{userData.Email}
	subject := "Verify your account"
	body := fmt.Sprintf("Click the following link to verify your account: %s", verifyLink)

	// Send the email message
	if err := SendMail(to, subject, body, "text"); err != nil {
		return err
	}

	return nil
}
