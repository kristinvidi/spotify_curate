package apptype

import "fmt"

type AccessToken string

func (a *AccessToken) HeaderValue() string {
	return fmt.Sprintf("Bearer %s", *a)
}
