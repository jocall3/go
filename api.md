# Users

Params Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#AddressParam">AddressParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#Address">Address</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#User">User</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#UserLoginResponse">UserLoginResponse</a>

Methods:

- <code title="post /users/login">client.Users.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#UserService.Login">Login</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#UserLoginParams">UserLoginParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#UserLoginResponse">UserLoginResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /users/register">client.Users.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#UserService.Register">Register</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#UserRegisterParams">UserRegisterParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#User">User</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## PasswordReset

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#UserPasswordResetConfirmResponse">UserPasswordResetConfirmResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#UserPasswordResetInitiateResponse">UserPasswordResetInitiateResponse</a>

Methods:

- <code title="post /users/password-reset/confirm">client.Users.PasswordReset.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#UserPasswordResetService.Confirm">Confirm</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#UserPasswordResetConfirmParams">UserPasswordResetConfirmParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#UserPasswordResetConfirmResponse">UserPasswordResetConfirmResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /users/password-reset/initiate">client.Users.PasswordReset.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#UserPasswordResetService.Initiate">Initiate</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#UserPasswordResetInitiateParams">UserPasswordResetInitiateParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#UserPasswordResetInitiateResponse">UserPasswordResetInitiateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Me

Methods:

- <code title="get /users/me">client.Users.Me.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#UserMeService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#User">User</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /users/me">client.Users.Me.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#UserMeService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#UserMeUpdateParams">UserMeUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#User">User</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Preferences

Params Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#UserPreferencesParam">UserPreferencesParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#UserPreferences">UserPreferences</a>

Methods:

- <code title="get /users/me/preferences">client.Users.Me.Preferences.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#UserMePreferenceService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#UserPreferences">UserPreferences</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /users/me/preferences">client.Users.Me.Preferences.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#UserMePreferenceService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#UserMePreferenceUpdateParams">UserMePreferenceUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#UserPreferences">UserPreferences</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Devices

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#Device">Device</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#PaginatedList">PaginatedList</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#UserMeDeviceListResponse">UserMeDeviceListResponse</a>

Methods:

- <code title="get /users/me/devices">client.Users.Me.Devices.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#UserMeDeviceService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#UserMeDeviceListParams">UserMeDeviceListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#UserMeDeviceListResponse">UserMeDeviceListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /users/me/devices/{deviceId}">client.Users.Me.Devices.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#UserMeDeviceService.Deregister">Deregister</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, deviceID <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="post /users/me/devices">client.Users.Me.Devices.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#UserMeDeviceService.Register">Register</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#UserMeDeviceRegisterParams">UserMeDeviceRegisterParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#Device">Device</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Biometrics

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#BiometricStatus">BiometricStatus</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#UserMeBiometricVerifyResponse">UserMeBiometricVerifyResponse</a>

Methods:

- <code title="delete /users/me/biometrics">client.Users.Me.Biometrics.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#UserMeBiometricService.Deregister">Deregister</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="post /users/me/biometrics/enroll">client.Users.Me.Biometrics.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#UserMeBiometricService.Enroll">Enroll</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#UserMeBiometricEnrollParams">UserMeBiometricEnrollParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#BiometricStatus">BiometricStatus</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /users/me/biometrics">client.Users.Me.Biometrics.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#UserMeBiometricService.Status">Status</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#BiometricStatus">BiometricStatus</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /users/me/biometrics/verify">client.Users.Me.Biometrics.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#UserMeBiometricService.Verify">Verify</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#UserMeBiometricVerifyParams">UserMeBiometricVerifyParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#UserMeBiometricVerifyResponse">UserMeBiometricVerifyResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Accounts

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#LinkedAccount">LinkedAccount</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#AccountLinkResponse">AccountLinkResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#AccountGetDetailsResponse">AccountGetDetailsResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#AccountGetMeResponse">AccountGetMeResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#AccountGetStatementsResponse">AccountGetStatementsResponse</a>

Methods:

- <code title="post /accounts/link">client.Accounts.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#AccountService.Link">Link</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#AccountLinkParams">AccountLinkParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#AccountLinkResponse">AccountLinkResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /accounts/{accountId}/details">client.Accounts.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#AccountService.GetDetails">GetDetails</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#AccountGetDetailsResponse">AccountGetDetailsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /accounts/me">client.Accounts.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#AccountService.GetMe">GetMe</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#AccountGetMeParams">AccountGetMeParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#AccountGetMeResponse">AccountGetMeResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /accounts/{accountId}/statements">client.Accounts.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#AccountService.GetStatements">GetStatements</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#AccountGetStatementsParams">AccountGetStatementsParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#AccountGetStatementsResponse">AccountGetStatementsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Transactions

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#AccountTransactionGetPendingResponse">AccountTransactionGetPendingResponse</a>

Methods:

- <code title="get /accounts/{accountId}/transactions/pending">client.Accounts.Transactions.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#AccountTransactionService.GetPending">GetPending</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#AccountTransactionGetPendingParams">AccountTransactionGetPendingParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#AccountTransactionGetPendingResponse">AccountTransactionGetPendingResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## OverdraftSettings

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#OverdraftSettings">OverdraftSettings</a>

Methods:

- <code title="get /accounts/{accountId}/overdraft-settings">client.Accounts.OverdraftSettings.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#AccountOverdraftSettingService.GetOverdraftSettings">GetOverdraftSettings</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#OverdraftSettings">OverdraftSettings</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /accounts/{accountId}/overdraft-settings">client.Accounts.OverdraftSettings.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#AccountOverdraftSettingService.UpdateOverdraftSettings">UpdateOverdraftSettings</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#AccountOverdraftSettingUpdateOverdraftSettingsParams">AccountOverdraftSettingUpdateOverdraftSettingsParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#OverdraftSettings">OverdraftSettings</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Transactions

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#PaginatedTransactions">PaginatedTransactions</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#Transaction">Transaction</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#TransactionDisputeResponse">TransactionDisputeResponse</a>

Methods:

- <code title="get /transactions/{transactionId}">client.Transactions.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#TransactionService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, transactionID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#Transaction">Transaction</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /transactions">client.Transactions.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#TransactionService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#TransactionListParams">TransactionListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#PaginatedTransactions">PaginatedTransactions</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /transactions/{transactionId}/categorize">client.Transactions.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#TransactionService.Categorize">Categorize</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, transactionID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#TransactionCategorizeParams">TransactionCategorizeParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#Transaction">Transaction</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /transactions/{transactionId}/dispute">client.Transactions.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#TransactionService.Dispute">Dispute</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, transactionID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#TransactionDisputeParams">TransactionDisputeParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#TransactionDisputeResponse">TransactionDisputeResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /transactions/{transactionId}/notes">client.Transactions.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#TransactionService.UpdateNotes">UpdateNotes</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, transactionID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#TransactionUpdateNotesParams">TransactionUpdateNotesParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#Transaction">Transaction</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Recurring

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#RecurringTransaction">RecurringTransaction</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#TransactionRecurringListResponse">TransactionRecurringListResponse</a>

Methods:

- <code title="post /transactions/recurring">client.Transactions.Recurring.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#TransactionRecurringService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#TransactionRecurringNewParams">TransactionRecurringNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#RecurringTransaction">RecurringTransaction</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /transactions/recurring">client.Transactions.Recurring.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#TransactionRecurringService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#TransactionRecurringListParams">TransactionRecurringListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#TransactionRecurringListResponse">TransactionRecurringListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Insights

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#AIInsight">AIInsight</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#TransactionInsightGetSpendingTrendsResponse">TransactionInsightGetSpendingTrendsResponse</a>

Methods:

- <code title="get /transactions/insights/spending-trends">client.Transactions.Insights.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#TransactionInsightService.GetSpendingTrends">GetSpendingTrends</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#TransactionInsightGetSpendingTrendsResponse">TransactionInsightGetSpendingTrendsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Budgets

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#Budget">Budget</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#BudgetListResponse">BudgetListResponse</a>

Methods:

- <code title="post /budgets">client.Budgets.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#BudgetService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#BudgetNewParams">BudgetNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#Budget">Budget</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /budgets/{budgetId}">client.Budgets.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#BudgetService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, budgetID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#Budget">Budget</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /budgets/{budgetId}">client.Budgets.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#BudgetService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, budgetID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#BudgetUpdateParams">BudgetUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#Budget">Budget</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /budgets">client.Budgets.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#BudgetService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#BudgetListParams">BudgetListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#BudgetListResponse">BudgetListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /budgets/{budgetId}">client.Budgets.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#BudgetService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, budgetID <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

# Investments

## Portfolios

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#InvestmentPortfolio">InvestmentPortfolio</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#InvestmentPortfolioListResponse">InvestmentPortfolioListResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#InvestmentPortfolioRebalanceResponse">InvestmentPortfolioRebalanceResponse</a>

Methods:

- <code title="post /investments/portfolios">client.Investments.Portfolios.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#InvestmentPortfolioService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#InvestmentPortfolioNewParams">InvestmentPortfolioNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#InvestmentPortfolio">InvestmentPortfolio</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /investments/portfolios/{portfolioId}">client.Investments.Portfolios.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#InvestmentPortfolioService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, portfolioID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#InvestmentPortfolio">InvestmentPortfolio</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /investments/portfolios/{portfolioId}">client.Investments.Portfolios.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#InvestmentPortfolioService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, portfolioID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#InvestmentPortfolioUpdateParams">InvestmentPortfolioUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#InvestmentPortfolio">InvestmentPortfolio</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /investments/portfolios">client.Investments.Portfolios.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#InvestmentPortfolioService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#InvestmentPortfolioListParams">InvestmentPortfolioListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#InvestmentPortfolioListResponse">InvestmentPortfolioListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /investments/portfolios/{portfolioId}/rebalance">client.Investments.Portfolios.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#InvestmentPortfolioService.Rebalance">Rebalance</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, portfolioID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#InvestmentPortfolioRebalanceParams">InvestmentPortfolioRebalanceParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#InvestmentPortfolioRebalanceResponse">InvestmentPortfolioRebalanceResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Assets

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#InvestmentAssetSearchResponse">InvestmentAssetSearchResponse</a>

Methods:

- <code title="get /investments/assets/search">client.Investments.Assets.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#InvestmentAssetService.Search">Search</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#InvestmentAssetSearchParams">InvestmentAssetSearchParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#InvestmentAssetSearchResponse">InvestmentAssetSearchResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# AI

## Advisor

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#AIAdvisorListToolsResponse">AIAdvisorListToolsResponse</a>

Methods:

- <code title="get /ai/advisor/tools">client.AI.Advisor.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#AIAdvisorService.ListTools">ListTools</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#AIAdvisorListToolsParams">AIAdvisorListToolsParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#AIAdvisorListToolsResponse">AIAdvisorListToolsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Chat

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#AIAdvisorChatGetHistoryResponse">AIAdvisorChatGetHistoryResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#AIAdvisorChatSendMessageResponse">AIAdvisorChatSendMessageResponse</a>

Methods:

- <code title="get /ai/advisor/chat/history">client.AI.Advisor.Chat.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#AIAdvisorChatService.GetHistory">GetHistory</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#AIAdvisorChatGetHistoryParams">AIAdvisorChatGetHistoryParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#AIAdvisorChatGetHistoryResponse">AIAdvisorChatGetHistoryResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /ai/advisor/chat">client.AI.Advisor.Chat.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#AIAdvisorChatService.SendMessage">SendMessage</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#AIAdvisorChatSendMessageParams">AIAdvisorChatSendMessageParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#AIAdvisorChatSendMessageResponse">AIAdvisorChatSendMessageResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Oracle

### Simulate

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#AdvancedSimulationResponse">AdvancedSimulationResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#SimulationResponse">SimulationResponse</a>

Methods:

- <code title="post /ai/oracle/simulate/advanced">client.AI.Oracle.Simulate.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#AIOracleSimulateService.RunAdvanced">RunAdvanced</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#AIOracleSimulateRunAdvancedParams">AIOracleSimulateRunAdvancedParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#AdvancedSimulationResponse">AdvancedSimulationResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /ai/oracle/simulate">client.AI.Oracle.Simulate.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#AIOracleSimulateService.RunStandard">RunStandard</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#AIOracleSimulateRunStandardParams">AIOracleSimulateRunStandardParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#SimulationResponse">SimulationResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Simulations

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#AIOracleSimulationGetResponse">AIOracleSimulationGetResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#AIOracleSimulationListResponse">AIOracleSimulationListResponse</a>

Methods:

- <code title="get /ai/oracle/simulations/{simulationId}">client.AI.Oracle.Simulations.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#AIOracleSimulationService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, simulationID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#AIOracleSimulationGetResponse">AIOracleSimulationGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /ai/oracle/simulations">client.AI.Oracle.Simulations.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#AIOracleSimulationService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#AIOracleSimulationListParams">AIOracleSimulationListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#AIOracleSimulationListResponse">AIOracleSimulationListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /ai/oracle/simulations/{simulationId}">client.AI.Oracle.Simulations.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#AIOracleSimulationService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, simulationID <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

## Incubator

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#AIIncubatorListPitchesResponse">AIIncubatorListPitchesResponse</a>

Methods:

- <code title="get /ai/incubator/pitches">client.AI.Incubator.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#AIIncubatorService.ListPitches">ListPitches</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#AIIncubatorListPitchesParams">AIIncubatorListPitchesParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#AIIncubatorListPitchesResponse">AIIncubatorListPitchesResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Pitch

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#QuantumWeaverState">QuantumWeaverState</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#AIIncubatorPitchGetDetailsResponse">AIIncubatorPitchGetDetailsResponse</a>

Methods:

- <code title="get /ai/incubator/pitch/{pitchId}/details">client.AI.Incubator.Pitch.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#AIIncubatorPitchService.GetDetails">GetDetails</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, pitchID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#AIIncubatorPitchGetDetailsResponse">AIIncubatorPitchGetDetailsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /ai/incubator/pitch">client.AI.Incubator.Pitch.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#AIIncubatorPitchService.Submit">Submit</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#AIIncubatorPitchSubmitParams">AIIncubatorPitchSubmitParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#QuantumWeaverState">QuantumWeaverState</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /ai/incubator/pitch/{pitchId}/feedback">client.AI.Incubator.Pitch.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#AIIncubatorPitchService.SubmitFeedback">SubmitFeedback</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, pitchID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#AIIncubatorPitchSubmitFeedbackParams">AIIncubatorPitchSubmitFeedbackParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#QuantumWeaverState">QuantumWeaverState</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Ads

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#VideoOperationStatus">VideoOperationStatus</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#AIAdListGeneratedResponse">AIAdListGeneratedResponse</a>

Methods:

- <code title="get /ai/ads">client.AI.Ads.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#AIAdService.ListGenerated">ListGenerated</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#AIAdListGeneratedParams">AIAdListGeneratedParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#AIAdListGeneratedResponse">AIAdListGeneratedResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /ai/ads/operations/{operationId}">client.AI.Ads.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#AIAdService.GetStatus">GetStatus</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, operationID interface{}) (<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#VideoOperationStatus">VideoOperationStatus</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Generate

Params Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#GenerateVideoParam">GenerateVideoParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#AIAdGenerateAdvancedResponse">AIAdGenerateAdvancedResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#AIAdGenerateStandardResponse">AIAdGenerateStandardResponse</a>

Methods:

- <code title="post /ai/ads/generate/advanced">client.AI.Ads.Generate.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#AIAdGenerateService.Advanced">Advanced</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#AIAdGenerateAdvancedParams">AIAdGenerateAdvancedParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#AIAdGenerateAdvancedResponse">AIAdGenerateAdvancedResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /ai/ads/generate">client.AI.Ads.Generate.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#AIAdGenerateService.Standard">Standard</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#AIAdGenerateStandardParams">AIAdGenerateStandardParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#AIAdGenerateStandardResponse">AIAdGenerateStandardResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Corporate

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#CorporatePerformSanctionScreeningResponse">CorporatePerformSanctionScreeningResponse</a>

Methods:

- <code title="post /corporate/sanction-screening">client.Corporate.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#CorporateService.PerformSanctionScreening">PerformSanctionScreening</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#CorporatePerformSanctionScreeningParams">CorporatePerformSanctionScreeningParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#CorporatePerformSanctionScreeningResponse">CorporatePerformSanctionScreeningResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Cards

Params Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#CorporateCardControlsParam">CorporateCardControlsParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#CorporateCard">CorporateCard</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#CorporateCardControls">CorporateCardControls</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#CorporateCardListResponse">CorporateCardListResponse</a>

Methods:

- <code title="get /corporate/cards">client.Corporate.Cards.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#CorporateCardService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#CorporateCardListParams">CorporateCardListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#CorporateCardListResponse">CorporateCardListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /corporate/cards/virtual">client.Corporate.Cards.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#CorporateCardService.NewVirtual">NewVirtual</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#CorporateCardNewVirtualParams">CorporateCardNewVirtualParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#CorporateCard">CorporateCard</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /corporate/cards/{cardId}/freeze">client.Corporate.Cards.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#CorporateCardService.Freeze">Freeze</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, cardID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#CorporateCardFreezeParams">CorporateCardFreezeParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#CorporateCard">CorporateCard</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /corporate/cards/{cardId}/transactions">client.Corporate.Cards.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#CorporateCardService.ListTransactions">ListTransactions</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, cardID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#CorporateCardListTransactionsParams">CorporateCardListTransactionsParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#PaginatedTransactions">PaginatedTransactions</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /corporate/cards/{cardId}/controls">client.Corporate.Cards.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#CorporateCardService.UpdateControls">UpdateControls</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, cardID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#CorporateCardUpdateControlsParams">CorporateCardUpdateControlsParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#CorporateCard">CorporateCard</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Anomalies

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#FinancialAnomaly">FinancialAnomaly</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#CorporateAnomalyListResponse">CorporateAnomalyListResponse</a>

Methods:

- <code title="get /corporate/anomalies">client.Corporate.Anomalies.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#CorporateAnomalyService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#CorporateAnomalyListParams">CorporateAnomalyListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#CorporateAnomalyListResponse">CorporateAnomalyListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /corporate/anomalies/{anomalyId}/status">client.Corporate.Anomalies.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#CorporateAnomalyService.UpdateStatus">UpdateStatus</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, anomalyID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#CorporateAnomalyUpdateStatusParams">CorporateAnomalyUpdateStatusParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#FinancialAnomaly">FinancialAnomaly</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Compliance

### Audits

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#CorporateComplianceAuditRequestResponse">CorporateComplianceAuditRequestResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#CorporateComplianceAuditGetReportResponse">CorporateComplianceAuditGetReportResponse</a>

Methods:

- <code title="post /corporate/compliance/audits">client.Corporate.Compliance.Audits.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#CorporateComplianceAuditService.Request">Request</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#CorporateComplianceAuditRequestParams">CorporateComplianceAuditRequestParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#CorporateComplianceAuditRequestResponse">CorporateComplianceAuditRequestResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /corporate/compliance/audits/{auditId}/report">client.Corporate.Compliance.Audits.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#CorporateComplianceAuditService.GetReport">GetReport</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, auditID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#CorporateComplianceAuditGetReportResponse">CorporateComplianceAuditGetReportResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Treasury

### CashFlow

## Risk

### Fraud

#### Rules

# Web3

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#Web3GetNFTsResponse">Web3GetNFTsResponse</a>

Methods:

- <code title="get /web3/nfts">client.Web3.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#Web3Service.GetNFTs">GetNFTs</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#Web3GetNFTsParams">Web3GetNFTsParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#Web3GetNFTsResponse">Web3GetNFTsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Wallets

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#CryptoWalletConnection">CryptoWalletConnection</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#Web3WalletListResponse">Web3WalletListResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#Web3WalletGetBalancesResponse">Web3WalletGetBalancesResponse</a>

Methods:

- <code title="get /web3/wallets">client.Web3.Wallets.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#Web3WalletService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#Web3WalletListParams">Web3WalletListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#Web3WalletListResponse">Web3WalletListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /web3/wallets">client.Web3.Wallets.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#Web3WalletService.Connect">Connect</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#Web3WalletConnectParams">Web3WalletConnectParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#CryptoWalletConnection">CryptoWalletConnection</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /web3/wallets/{walletId}/balances">client.Web3.Wallets.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#Web3WalletService.GetBalances">GetBalances</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, walletID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#Web3WalletGetBalancesParams">Web3WalletGetBalancesParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#Web3WalletGetBalancesResponse">Web3WalletGetBalancesResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Transactions

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#Web3TransactionInitiateTransferResponse">Web3TransactionInitiateTransferResponse</a>

Methods:

- <code title="post /web3/transactions/initiate">client.Web3.Transactions.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#Web3TransactionService.InitiateTransfer">InitiateTransfer</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#Web3TransactionInitiateTransferParams">Web3TransactionInitiateTransferParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#Web3TransactionInitiateTransferResponse">Web3TransactionInitiateTransferResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Payments

## International

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#InternationalPaymentStatus">InternationalPaymentStatus</a>

Methods:

- <code title="post /payments/international/initiate">client.Payments.International.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#PaymentInternationalService.Initiate">Initiate</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#PaymentInternationalInitiateParams">PaymentInternationalInitiateParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#InternationalPaymentStatus">InternationalPaymentStatus</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /payments/international/{paymentId}/status">client.Payments.International.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#PaymentInternationalService.GetStatus">GetStatus</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, paymentID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#InternationalPaymentStatus">InternationalPaymentStatus</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Fx

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#PaymentFxConvertResponse">PaymentFxConvertResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#PaymentFxGetRatesResponse">PaymentFxGetRatesResponse</a>

Methods:

- <code title="post /payments/fx/convert">client.Payments.Fx.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#PaymentFxService.Convert">Convert</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#PaymentFxConvertParams">PaymentFxConvertParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#PaymentFxConvertResponse">PaymentFxConvertResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /payments/fx/rates">client.Payments.Fx.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#PaymentFxService.GetRates">GetRates</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#PaymentFxGetRatesParams">PaymentFxGetRatesParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#PaymentFxGetRatesResponse">PaymentFxGetRatesResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Sustainability

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#SustainabilityPurchaseCarbonOffsetsResponse">SustainabilityPurchaseCarbonOffsetsResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#SustainabilityGetCarbonFootprintResponse">SustainabilityGetCarbonFootprintResponse</a>

Methods:

- <code title="post /sustainability/carbon-offsets">client.Sustainability.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#SustainabilityService.PurchaseCarbonOffsets">PurchaseCarbonOffsets</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#SustainabilityPurchaseCarbonOffsetsParams">SustainabilityPurchaseCarbonOffsetsParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#SustainabilityPurchaseCarbonOffsetsResponse">SustainabilityPurchaseCarbonOffsetsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /sustainability/carbon-footprint">client.Sustainability.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#SustainabilityService.GetCarbonFootprint">GetCarbonFootprint</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#SustainabilityGetCarbonFootprintResponse">SustainabilityGetCarbonFootprintResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Investments

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#SustainabilityInvestmentAnalyzeImpactResponse">SustainabilityInvestmentAnalyzeImpactResponse</a>

Methods:

- <code title="get /sustainability/investments/impact">client.Sustainability.Investments.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#SustainabilityInvestmentService.AnalyzeImpact">AnalyzeImpact</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#SustainabilityInvestmentAnalyzeImpactResponse">SustainabilityInvestmentAnalyzeImpactResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Lending

## Applications

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#LoanApplicationStatus">LoanApplicationStatus</a>

Methods:

- <code title="get /lending/applications/{applicationId}">client.Lending.Applications.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#LendingApplicationService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, applicationID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#LoanApplicationStatus">LoanApplicationStatus</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /lending/applications">client.Lending.Applications.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#LendingApplicationService.Submit">Submit</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#LendingApplicationSubmitParams">LendingApplicationSubmitParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#LoanApplicationStatus">LoanApplicationStatus</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Offers

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#LoanOffer">LoanOffer</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#LendingOfferListPreApprovedResponse">LendingOfferListPreApprovedResponse</a>

Methods:

- <code title="get /lending/offers/pre-approved">client.Lending.Offers.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#LendingOfferService.ListPreApproved">ListPreApproved</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#LendingOfferListPreApprovedParams">LendingOfferListPreApprovedParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#LendingOfferListPreApprovedResponse">LendingOfferListPreApprovedResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Developers

## Webhooks

## APIKeys

# Identity

## KYC

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#KYCStatus">KYCStatus</a>

Methods:

- <code title="get /identity/kyc/status">client.Identity.KYC.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#IdentityKYCService.GetStatus">GetStatus</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#KYCStatus">KYCStatus</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Goals

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#FinancialGoal">FinancialGoal</a>

Methods:

- <code title="post /goals">client.Goals.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#GoalService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#GoalNewParams">GoalNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go">jamesburvelocallaghaniiicitibankdemobusinessinc</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/1231-go#FinancialGoal">FinancialGoal</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Notifications

## Settings

# Marketplace

## Products

## Offers
