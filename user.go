// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package jamesburvelocallaghaniiicitibankdemobusinessinc

import (
	"context"
	"net/http"
	"slices"
	"time"

	"github.com/jocall3/1231-go/internal/apijson"
	"github.com/jocall3/1231-go/internal/param"
	"github.com/jocall3/1231-go/internal/requestconfig"
	"github.com/jocall3/1231-go/option"
)

// UserService contains methods and other services that help with interacting with
// the 1231 API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUserService] method instead.
type UserService struct {
	Options       []option.RequestOption
	PasswordReset *UserPasswordResetService
	Me            *UserMeService
}

// NewUserService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewUserService(opts ...option.RequestOption) (r *UserService) {
	r = &UserService{}
	r.Options = opts
	r.PasswordReset = NewUserPasswordResetService(opts...)
	r.Me = NewUserMeService(opts...)
	return
}

// Authenticates a user and creates a secure session, returning access tokens. May
// require MFA depending on user settings.
func (r *UserService) Login(ctx context.Context, body UserLoginParams, opts ...option.RequestOption) (res *UserLoginResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "users/login"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Registers a new user account with , initiating the onboarding process. Requires
// basic user details.
func (r *UserService) Register(ctx context.Context, body UserRegisterParams, opts ...option.RequestOption) (res *User, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "users/register"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type Address struct {
	City    string      `json:"city"`
	Country string      `json:"country"`
	State   string      `json:"state"`
	Street  string      `json:"street"`
	Zip     string      `json:"zip"`
	JSON    addressJSON `json:"-"`
}

// addressJSON contains the JSON metadata for the struct [Address]
type addressJSON struct {
	City        apijson.Field
	Country     apijson.Field
	State       apijson.Field
	Street      apijson.Field
	Zip         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Address) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressJSON) RawJSON() string {
	return r.raw
}

type AddressParam struct {
	City    param.Field[string] `json:"city"`
	Country param.Field[string] `json:"country"`
	State   param.Field[string] `json:"state"`
	Street  param.Field[string] `json:"street"`
	Zip     param.Field[string] `json:"zip"`
}

func (r AddressParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type User struct {
	ID                string             `json:"id,required"`
	Email             string             `json:"email,required" format:"email"`
	IdentityVerified  bool               `json:"identityVerified,required"`
	Name              string             `json:"name,required"`
	Address           Address            `json:"address"`
	AIPersona         string             `json:"aiPersona"`
	DateOfBirth       time.Time          `json:"dateOfBirth" format:"date"`
	GamificationLevel int64              `json:"gamificationLevel"`
	LoyaltyPoints     int64              `json:"loyaltyPoints"`
	LoyaltyTier       string             `json:"loyaltyTier"`
	Phone             string             `json:"phone"`
	Preferences       UserPreferences    `json:"preferences"`
	SecurityStatus    UserSecurityStatus `json:"securityStatus"`
	JSON              userJSON           `json:"-"`
}

// userJSON contains the JSON metadata for the struct [User]
type userJSON struct {
	ID                apijson.Field
	Email             apijson.Field
	IdentityVerified  apijson.Field
	Name              apijson.Field
	Address           apijson.Field
	AIPersona         apijson.Field
	DateOfBirth       apijson.Field
	GamificationLevel apijson.Field
	LoyaltyPoints     apijson.Field
	LoyaltyTier       apijson.Field
	Phone             apijson.Field
	Preferences       apijson.Field
	SecurityStatus    apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *User) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userJSON) RawJSON() string {
	return r.raw
}

type UserSecurityStatus struct {
	BiometricsEnrolled bool                   `json:"biometricsEnrolled"`
	LastLogin          time.Time              `json:"lastLogin" format:"date-time"`
	LastLoginIP        string                 `json:"lastLoginIp"`
	TwoFactorEnabled   bool                   `json:"twoFactorEnabled"`
	JSON               userSecurityStatusJSON `json:"-"`
}

// userSecurityStatusJSON contains the JSON metadata for the struct
// [UserSecurityStatus]
type userSecurityStatusJSON struct {
	BiometricsEnrolled apijson.Field
	LastLogin          apijson.Field
	LastLoginIP        apijson.Field
	TwoFactorEnabled   apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *UserSecurityStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userSecurityStatusJSON) RawJSON() string {
	return r.raw
}

type UserLoginResponse struct {
	AccessToken  string                `json:"accessToken,required"`
	ExpiresIn    int64                 `json:"expiresIn,required"`
	RefreshToken string                `json:"refreshToken,required"`
	TokenType    string                `json:"tokenType,required"`
	JSON         userLoginResponseJSON `json:"-"`
}

// userLoginResponseJSON contains the JSON metadata for the struct
// [UserLoginResponse]
type userLoginResponseJSON struct {
	AccessToken  apijson.Field
	ExpiresIn    apijson.Field
	RefreshToken apijson.Field
	TokenType    apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *UserLoginResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userLoginResponseJSON) RawJSON() string {
	return r.raw
}

type UserLoginParams struct {
	Email    param.Field[string] `json:"email,required" format:"email"`
	Password param.Field[string] `json:"password,required" format:"password"`
	MfaCode  param.Field[string] `json:"mfaCode"`
}

func (r UserLoginParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type UserRegisterParams struct {
	Email    param.Field[string] `json:"email,required" format:"email"`
	Name     param.Field[string] `json:"name,required"`
	Password param.Field[string] `json:"password,required" format:"password"`
	Phone    param.Field[string] `json:"phone"`
}

func (r UserRegisterParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
