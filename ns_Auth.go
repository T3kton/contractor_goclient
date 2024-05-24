// Package contractor - (version: "2.0") - Automatically generated by cinp-codegen from /api/v1/Auth/ at 2024-05-24T15:42:51.144790
package contractor

import (
	"context"
	cinp "github.com/cinp/go"
	"reflect"
)

// AuthUser - Model User(/api/v1/Auth/User)
/*

 */
type AuthUser struct {
	cinp.BaseObject
	cinp cinp.CInPClient `json:"-"`
}

// AuthUserNew - Make a new object of Model User
func (service *Contractor) AuthUserNew() *AuthUser {
	return &AuthUser{cinp: service.cinp}
}

// AuthUserCallLogin calls login
func (service *Contractor) AuthUserCallLogin(ctx context.Context, Username string, Password string) (string, error) {
	args := map[string]interface{}{
		"username": Username,
		"password": Password,
	}
	uri := "/api/v1/Auth/User(login)"

	result := ""

	if err := service.cinp.Call(ctx, uri, &args, &result); err != nil {
		return "", err
	}

	return result, nil
}

// AuthUserCallLogout calls logout
func (service *Contractor) AuthUserCallLogout(ctx context.Context) error {
	args := map[string]interface{}{}
	uri := "/api/v1/Auth/User(logout)"

	result := ""

	if err := service.cinp.Call(ctx, uri, &args, &result); err != nil {
		return err
	}

	return nil
}

// AuthUserCallWhoami calls whoami
func (service *Contractor) AuthUserCallWhoami(ctx context.Context) (string, error) {
	args := map[string]interface{}{}
	uri := "/api/v1/Auth/User(whoami)"

	result := ""

	if err := service.cinp.Call(ctx, uri, &args, &result); err != nil {
		return "", err
	}

	return result, nil
}

// AuthUserCallChangePassword calls changePassword
func (service *Contractor) AuthUserCallChangePassword(ctx context.Context, Password string) error {
	args := map[string]interface{}{
		"password": Password,
	}
	uri := "/api/v1/Auth/User(changePassword)"

	result := ""

	if err := service.cinp.Call(ctx, uri, &args, &result); err != nil {
		return err
	}

	return nil
}
func registerAuth(cinp cinp.CInPClient) {
	cinp.RegisterType("/api/v1/Auth/User", reflect.TypeOf((*AuthUser)(nil)).Elem())
}
