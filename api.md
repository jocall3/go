# Users

Params Types:

- <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#AddressParam">AddressParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#Address">Address</a>
- <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#User">User</a>
- <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#UserLoginResponse">UserLoginResponse</a>

Methods:

- <code title="post /users/login">client.Users.<a href="https://pkg.go.dev/github.com/jocall3/cli#UserService.Login">Login</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#UserLoginParams">UserLoginParams</a>) (<a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#UserLoginResponse">UserLoginResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /users/register">client.Users.<a href="https://pkg.go.dev/github.com/jocall3/cli#UserService.Register">Register</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#UserRegisterParams">UserRegisterParams</a>) (<a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#User">User</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## PasswordReset

Response Types:

- <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#UserPasswordResetConfirmResponse">UserPasswordResetConfirmResponse</a>
- <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#UserPasswordResetInitiateResponse">UserPasswordResetInitiateResponse</a>

Methods:

- <code title="post /users/password-reset/confirm">client.Users.PasswordReset.<a href="https://pkg.go.dev/github.com/jocall3/cli#UserPasswordResetService.Confirm">Confirm</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#UserPasswordResetConfirmParams">UserPasswordResetConfirmParams</a>) (<a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#UserPasswordResetConfirmResponse">UserPasswordResetConfirmResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /users/password-reset/initiate">client.Users.PasswordReset.<a href="https://pkg.go.dev/github.com/jocall3/cli#UserPasswordResetService.Initiate">Initiate</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#UserPasswordResetInitiateParams">UserPasswordResetInitiateParams</a>) (<a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#UserPasswordResetInitiateResponse">UserPasswordResetInitiateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Me

Methods:

- <code title="get /users/me">client.Users.Me.<a href="https://pkg.go.dev/github.com/jocall3/cli#UserMeService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#User">User</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /users/me">client.Users.Me.<a href="https://pkg.go.dev/github.com/jocall3/cli#UserMeService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#UserMeUpdateParams">UserMeUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#User">User</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Preferences

Params Types:

- <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#UserPreferencesParam">UserPreferencesParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#UserPreferences">UserPreferences</a>

Methods:

- <code title="get /users/me/preferences">client.Users.Me.Preferences.<a href="https://pkg.go.dev/github.com/jocall3/cli#UserMePreferenceService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#UserPreferences">UserPreferences</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /users/me/preferences">client.Users.Me.Preferences.<a href="https://pkg.go.dev/github.com/jocall3/cli#UserMePreferenceService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#UserMePreferenceUpdateParams">UserMePreferenceUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#UserPreferences">UserPreferences</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Devices

Response Types:

- <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#Device">Device</a>
- <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#PaginatedList">PaginatedList</a>
- <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#UserMeDeviceListResponse">UserMeDeviceListResponse</a>

Methods:

- <code title="get /users/me/devices">client.Users.Me.Devices.<a href="https://pkg.go.dev/github.com/jocall3/cli#UserMeDeviceService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#UserMeDeviceListParams">UserMeDeviceListParams</a>) (<a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#UserMeDeviceListResponse">UserMeDeviceListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /users/me/devices/{deviceId}">client.Users.Me.Devices.<a href="https://pkg.go.dev/github.com/jocall3/cli#UserMeDeviceService.Deregister">Deregister</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, deviceID interface{}) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="post /users/me/devices">client.Users.Me.Devices.<a href="https://pkg.go.dev/github.com/jocall3/cli#UserMeDeviceService.Register">Register</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#UserMeDeviceRegisterParams">UserMeDeviceRegisterParams</a>) (<a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#Device">Device</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Biometrics

Response Types:

- <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#BiometricStatus">BiometricStatus</a>
- <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#UserMeBiometricVerifyResponse">UserMeBiometricVerifyResponse</a>

Methods:

- <code title="delete /users/me/biometrics">client.Users.Me.Biometrics.<a href="https://pkg.go.dev/github.com/jocall3/cli#UserMeBiometricService.Deregister">Deregister</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="post /users/me/biometrics/enroll">client.Users.Me.Biometrics.<a href="https://pkg.go.dev/github.com/jocall3/cli#UserMeBiometricService.Enroll">Enroll</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#UserMeBiometricEnrollParams">UserMeBiometricEnrollParams</a>) (<a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#BiometricStatus">BiometricStatus</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /users/me/biometrics">client.Users.Me.Biometrics.<a href="https://pkg.go.dev/github.com/jocall3/cli#UserMeBiometricService.Status">Status</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#BiometricStatus">BiometricStatus</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /users/me/biometrics/verify">client.Users.Me.Biometrics.<a href="https://pkg.go.dev/github.com/jocall3/cli#UserMeBiometricService.Verify">Verify</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#UserMeBiometricVerifyParams">UserMeBiometricVerifyParams</a>) (<a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#UserMeBiometricVerifyResponse">UserMeBiometricVerifyResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Accounts

Response Types:

- <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#LinkedAccount">LinkedAccount</a>
- <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#AccountLinkResponse">AccountLinkResponse</a>
- <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#AccountGetDetailsResponse">AccountGetDetailsResponse</a>
- <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#AccountGetMeResponse">AccountGetMeResponse</a>
- <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#AccountGetStatementsResponse">AccountGetStatementsResponse</a>

Methods:

- <code title="post /accounts/link">client.Accounts.<a href="https://pkg.go.dev/github.com/jocall3/cli#AccountService.Link">Link</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#AccountLinkParams">AccountLinkParams</a>) (<a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#AccountLinkResponse">AccountLinkResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /accounts/{accountId}/details">client.Accounts.<a href="https://pkg.go.dev/github.com/jocall3/cli#AccountService.GetDetails">GetDetails</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID interface{}) (<a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#AccountGetDetailsResponse">AccountGetDetailsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /accounts/me">client.Accounts.<a href="https://pkg.go.dev/github.com/jocall3/cli#AccountService.GetMe">GetMe</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#AccountGetMeParams">AccountGetMeParams</a>) (<a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#AccountGetMeResponse">AccountGetMeResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /accounts/{accountId}/statements">client.Accounts.<a href="https://pkg.go.dev/github.com/jocall3/cli#AccountService.GetStatements">GetStatements</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID interface{}, query <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#AccountGetStatementsParams">AccountGetStatementsParams</a>) (<a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#AccountGetStatementsResponse">AccountGetStatementsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Transactions

Response Types:

- <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#AccountTransactionGetPendingResponse">AccountTransactionGetPendingResponse</a>

Methods:

- <code title="get /accounts/{accountId}/transactions/pending">client.Accounts.Transactions.<a href="https://pkg.go.dev/github.com/jocall3/cli#AccountTransactionService.GetPending">GetPending</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID interface{}, query <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#AccountTransactionGetPendingParams">AccountTransactionGetPendingParams</a>) (<a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#AccountTransactionGetPendingResponse">AccountTransactionGetPendingResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## OverdraftSettings

Response Types:

- <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#OverdraftSettings">OverdraftSettings</a>

Methods:

- <code title="get /accounts/{accountId}/overdraft-settings">client.Accounts.OverdraftSettings.<a href="https://pkg.go.dev/github.com/jocall3/cli#AccountOverdraftSettingService.GetOverdraftSettings">GetOverdraftSettings</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID interface{}) (<a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#OverdraftSettings">OverdraftSettings</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /accounts/{accountId}/overdraft-settings">client.Accounts.OverdraftSettings.<a href="https://pkg.go.dev/github.com/jocall3/cli#AccountOverdraftSettingService.UpdateOverdraftSettings">UpdateOverdraftSettings</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, accountID interface{}, body <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#AccountOverdraftSettingUpdateOverdraftSettingsParams">AccountOverdraftSettingUpdateOverdraftSettingsParams</a>) (<a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#OverdraftSettings">OverdraftSettings</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Transactions

Response Types:

- <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#PaginatedTransactions">PaginatedTransactions</a>
- <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#Transaction">Transaction</a>
- <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#TransactionDisputeResponse">TransactionDisputeResponse</a>

Methods:

- <code title="get /transactions/{transactionId}">client.Transactions.<a href="https://pkg.go.dev/github.com/jocall3/cli#TransactionService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, transactionID interface{}) (<a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#Transaction">Transaction</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /transactions">client.Transactions.<a href="https://pkg.go.dev/github.com/jocall3/cli#TransactionService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#TransactionListParams">TransactionListParams</a>) (<a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#PaginatedTransactions">PaginatedTransactions</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /transactions/{transactionId}/categorize">client.Transactions.<a href="https://pkg.go.dev/github.com/jocall3/cli#TransactionService.Categorize">Categorize</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, transactionID interface{}, body <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#TransactionCategorizeParams">TransactionCategorizeParams</a>) (<a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#Transaction">Transaction</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /transactions/{transactionId}/dispute">client.Transactions.<a href="https://pkg.go.dev/github.com/jocall3/cli#TransactionService.Dispute">Dispute</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, transactionID interface{}, body <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#TransactionDisputeParams">TransactionDisputeParams</a>) (<a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#TransactionDisputeResponse">TransactionDisputeResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /transactions/{transactionId}/notes">client.Transactions.<a href="https://pkg.go.dev/github.com/jocall3/cli#TransactionService.UpdateNotes">UpdateNotes</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, transactionID interface{}, body <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#TransactionUpdateNotesParams">TransactionUpdateNotesParams</a>) (<a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#Transaction">Transaction</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Recurring

Response Types:

- <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#RecurringTransaction">RecurringTransaction</a>
- <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#TransactionRecurringListResponse">TransactionRecurringListResponse</a>

Methods:

- <code title="post /transactions/recurring">client.Transactions.Recurring.<a href="https://pkg.go.dev/github.com/jocall3/cli#TransactionRecurringService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#TransactionRecurringNewParams">TransactionRecurringNewParams</a>) (<a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#RecurringTransaction">RecurringTransaction</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /transactions/recurring">client.Transactions.Recurring.<a href="https://pkg.go.dev/github.com/jocall3/cli#TransactionRecurringService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#TransactionRecurringListParams">TransactionRecurringListParams</a>) (<a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#TransactionRecurringListResponse">TransactionRecurringListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Insights

Response Types:

- <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#AIInsight">AIInsight</a>
- <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#TransactionInsightGetSpendingTrendsResponse">TransactionInsightGetSpendingTrendsResponse</a>

Methods:

- <code title="get /transactions/insights/spending-trends">client.Transactions.Insights.<a href="https://pkg.go.dev/github.com/jocall3/cli#TransactionInsightService.GetSpendingTrends">GetSpendingTrends</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#TransactionInsightGetSpendingTrendsResponse">TransactionInsightGetSpendingTrendsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Budgets

Response Types:

- <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#Budget">Budget</a>
- <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#BudgetListResponse">BudgetListResponse</a>

Methods:

- <code title="post /budgets">client.Budgets.<a href="https://pkg.go.dev/github.com/jocall3/cli#BudgetService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#BudgetNewParams">BudgetNewParams</a>) (<a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#Budget">Budget</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /budgets/{budgetId}">client.Budgets.<a href="https://pkg.go.dev/github.com/jocall3/cli#BudgetService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, budgetID interface{}) (<a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#Budget">Budget</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /budgets/{budgetId}">client.Budgets.<a href="https://pkg.go.dev/github.com/jocall3/cli#BudgetService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, budgetID interface{}, body <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#BudgetUpdateParams">BudgetUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#Budget">Budget</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /budgets">client.Budgets.<a href="https://pkg.go.dev/github.com/jocall3/cli#BudgetService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#BudgetListParams">BudgetListParams</a>) (<a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#BudgetListResponse">BudgetListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /budgets/{budgetId}">client.Budgets.<a href="https://pkg.go.dev/github.com/jocall3/cli#BudgetService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, budgetID interface{}) <a href="https://pkg.go.dev/builtin#error">error</a></code>

# Investments

## Portfolios

Response Types:

- <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#InvestmentPortfolio">InvestmentPortfolio</a>
- <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#InvestmentPortfolioListResponse">InvestmentPortfolioListResponse</a>
- <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#InvestmentPortfolioRebalanceResponse">InvestmentPortfolioRebalanceResponse</a>

Methods:

- <code title="post /investments/portfolios">client.Investments.Portfolios.<a href="https://pkg.go.dev/github.com/jocall3/cli#InvestmentPortfolioService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#InvestmentPortfolioNewParams">InvestmentPortfolioNewParams</a>) (<a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#InvestmentPortfolio">InvestmentPortfolio</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /investments/portfolios/{portfolioId}">client.Investments.Portfolios.<a href="https://pkg.go.dev/github.com/jocall3/cli#InvestmentPortfolioService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, portfolioID interface{}) (<a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#InvestmentPortfolio">InvestmentPortfolio</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /investments/portfolios/{portfolioId}">client.Investments.Portfolios.<a href="https://pkg.go.dev/github.com/jocall3/cli#InvestmentPortfolioService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, portfolioID interface{}, body <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#InvestmentPortfolioUpdateParams">InvestmentPortfolioUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#InvestmentPortfolio">InvestmentPortfolio</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /investments/portfolios">client.Investments.Portfolios.<a href="https://pkg.go.dev/github.com/jocall3/cli#InvestmentPortfolioService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#InvestmentPortfolioListParams">InvestmentPortfolioListParams</a>) (<a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#InvestmentPortfolioListResponse">InvestmentPortfolioListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /investments/portfolios/{portfolioId}/rebalance">client.Investments.Portfolios.<a href="https://pkg.go.dev/github.com/jocall3/cli#InvestmentPortfolioService.Rebalance">Rebalance</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, portfolioID interface{}, body <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#InvestmentPortfolioRebalanceParams">InvestmentPortfolioRebalanceParams</a>) (<a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#InvestmentPortfolioRebalanceResponse">InvestmentPortfolioRebalanceResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Assets

Response Types:

- <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#InvestmentAssetSearchResponse">InvestmentAssetSearchResponse</a>

Methods:

- <code title="get /investments/assets/search">client.Investments.Assets.<a href="https://pkg.go.dev/github.com/jocall3/cli#InvestmentAssetService.Search">Search</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#InvestmentAssetSearchParams">InvestmentAssetSearchParams</a>) (<a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#InvestmentAssetSearchResponse">InvestmentAssetSearchResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# AI

## Advisor

Response Types:

- <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#AIAdvisorListToolsResponse">AIAdvisorListToolsResponse</a>

Methods:

- <code title="get /ai/advisor/tools">client.AI.Advisor.<a href="https://pkg.go.dev/github.com/jocall3/cli#AIAdvisorService.ListTools">ListTools</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#AIAdvisorListToolsParams">AIAdvisorListToolsParams</a>) (<a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#AIAdvisorListToolsResponse">AIAdvisorListToolsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Chat

Response Types:

- <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#AIAdvisorChatGetHistoryResponse">AIAdvisorChatGetHistoryResponse</a>
- <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#AIAdvisorChatSendMessageResponse">AIAdvisorChatSendMessageResponse</a>

Methods:

- <code title="get /ai/advisor/chat/history">client.AI.Advisor.Chat.<a href="https://pkg.go.dev/github.com/jocall3/cli#AIAdvisorChatService.GetHistory">GetHistory</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#AIAdvisorChatGetHistoryParams">AIAdvisorChatGetHistoryParams</a>) (<a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#AIAdvisorChatGetHistoryResponse">AIAdvisorChatGetHistoryResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /ai/advisor/chat">client.AI.Advisor.Chat.<a href="https://pkg.go.dev/github.com/jocall3/cli#AIAdvisorChatService.SendMessage">SendMessage</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#AIAdvisorChatSendMessageParams">AIAdvisorChatSendMessageParams</a>) (<a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#AIAdvisorChatSendMessageResponse">AIAdvisorChatSendMessageResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Oracle

### Simulate

Response Types:

- <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#AdvancedSimulationResponse">AdvancedSimulationResponse</a>
- <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#SimulationResponse">SimulationResponse</a>

Methods:

- <code title="post /ai/oracle/simulate/advanced">client.AI.Oracle.Simulate.<a href="https://pkg.go.dev/github.com/jocall3/cli#AIOracleSimulateService.RunAdvanced">RunAdvanced</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#AIOracleSimulateRunAdvancedParams">AIOracleSimulateRunAdvancedParams</a>) (<a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#AdvancedSimulationResponse">AdvancedSimulationResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /ai/oracle/simulate">client.AI.Oracle.Simulate.<a href="https://pkg.go.dev/github.com/jocall3/cli#AIOracleSimulateService.RunStandard">RunStandard</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#AIOracleSimulateRunStandardParams">AIOracleSimulateRunStandardParams</a>) (<a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#SimulationResponse">SimulationResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Simulations

Response Types:

- <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#AIOracleSimulationGetResponse">AIOracleSimulationGetResponse</a>
- <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#AIOracleSimulationListResponse">AIOracleSimulationListResponse</a>

Methods:

- <code title="get /ai/oracle/simulations/{simulationId}">client.AI.Oracle.Simulations.<a href="https://pkg.go.dev/github.com/jocall3/cli#AIOracleSimulationService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, simulationID interface{}) (<a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#AIOracleSimulationGetResponse">AIOracleSimulationGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /ai/oracle/simulations">client.AI.Oracle.Simulations.<a href="https://pkg.go.dev/github.com/jocall3/cli#AIOracleSimulationService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#AIOracleSimulationListParams">AIOracleSimulationListParams</a>) (<a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#AIOracleSimulationListResponse">AIOracleSimulationListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /ai/oracle/simulations/{simulationId}">client.AI.Oracle.Simulations.<a href="https://pkg.go.dev/github.com/jocall3/cli#AIOracleSimulationService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, simulationID interface{}) <a href="https://pkg.go.dev/builtin#error">error</a></code>

## Incubator

Response Types:

- <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#AIIncubatorListPitchesResponse">AIIncubatorListPitchesResponse</a>

Methods:

- <code title="get /ai/incubator/pitches">client.AI.Incubator.<a href="https://pkg.go.dev/github.com/jocall3/cli#AIIncubatorService.ListPitches">ListPitches</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#AIIncubatorListPitchesParams">AIIncubatorListPitchesParams</a>) (<a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#AIIncubatorListPitchesResponse">AIIncubatorListPitchesResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Pitch

Response Types:

- <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#QuantumWeaverState">QuantumWeaverState</a>
- <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#AIIncubatorPitchGetDetailsResponse">AIIncubatorPitchGetDetailsResponse</a>

Methods:

- <code title="get /ai/incubator/pitch/{pitchId}/details">client.AI.Incubator.Pitch.<a href="https://pkg.go.dev/github.com/jocall3/cli#AIIncubatorPitchService.GetDetails">GetDetails</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, pitchID interface{}) (<a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#AIIncubatorPitchGetDetailsResponse">AIIncubatorPitchGetDetailsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /ai/incubator/pitch">client.AI.Incubator.Pitch.<a href="https://pkg.go.dev/github.com/jocall3/cli#AIIncubatorPitchService.Submit">Submit</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#AIIncubatorPitchSubmitParams">AIIncubatorPitchSubmitParams</a>) (<a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#QuantumWeaverState">QuantumWeaverState</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /ai/incubator/pitch/{pitchId}/feedback">client.AI.Incubator.Pitch.<a href="https://pkg.go.dev/github.com/jocall3/cli#AIIncubatorPitchService.SubmitFeedback">SubmitFeedback</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, pitchID interface{}, body <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#AIIncubatorPitchSubmitFeedbackParams">AIIncubatorPitchSubmitFeedbackParams</a>) (<a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#QuantumWeaverState">QuantumWeaverState</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Ads

Response Types:

- <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#VideoOperationStatus">VideoOperationStatus</a>
- <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#AIAdListGeneratedResponse">AIAdListGeneratedResponse</a>

Methods:

- <code title="get /ai/ads">client.AI.Ads.<a href="https://pkg.go.dev/github.com/jocall3/cli#AIAdService.ListGenerated">ListGenerated</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#AIAdListGeneratedParams">AIAdListGeneratedParams</a>) (<a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#AIAdListGeneratedResponse">AIAdListGeneratedResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /ai/ads/operations/{operationId}">client.AI.Ads.<a href="https://pkg.go.dev/github.com/jocall3/cli#AIAdService.GetStatus">GetStatus</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, operationID interface{}) (<a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#VideoOperationStatus">VideoOperationStatus</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Generate

Params Types:

- <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#GenerateVideoParam">GenerateVideoParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#AIAdGenerateAdvancedResponse">AIAdGenerateAdvancedResponse</a>
- <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#AIAdGenerateStandardResponse">AIAdGenerateStandardResponse</a>

Methods:

- <code title="post /ai/ads/generate/advanced">client.AI.Ads.Generate.<a href="https://pkg.go.dev/github.com/jocall3/cli#AIAdGenerateService.Advanced">Advanced</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#AIAdGenerateAdvancedParams">AIAdGenerateAdvancedParams</a>) (<a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#AIAdGenerateAdvancedResponse">AIAdGenerateAdvancedResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /ai/ads/generate">client.AI.Ads.Generate.<a href="https://pkg.go.dev/github.com/jocall3/cli#AIAdGenerateService.Standard">Standard</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#AIAdGenerateStandardParams">AIAdGenerateStandardParams</a>) (<a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#AIAdGenerateStandardResponse">AIAdGenerateStandardResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Corporate

Response Types:

- <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#CorporatePerformSanctionScreeningResponse">CorporatePerformSanctionScreeningResponse</a>

Methods:

- <code title="post /corporate/sanction-screening">client.Corporate.<a href="https://pkg.go.dev/github.com/jocall3/cli#CorporateService.PerformSanctionScreening">PerformSanctionScreening</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#CorporatePerformSanctionScreeningParams">CorporatePerformSanctionScreeningParams</a>) (<a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#CorporatePerformSanctionScreeningResponse">CorporatePerformSanctionScreeningResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Cards

Params Types:

- <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#CorporateCardControlsParam">CorporateCardControlsParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#CorporateCard">CorporateCard</a>
- <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#CorporateCardControls">CorporateCardControls</a>
- <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#CorporateCardListResponse">CorporateCardListResponse</a>

Methods:

- <code title="get /corporate/cards">client.Corporate.Cards.<a href="https://pkg.go.dev/github.com/jocall3/cli#CorporateCardService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#CorporateCardListParams">CorporateCardListParams</a>) (<a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#CorporateCardListResponse">CorporateCardListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /corporate/cards/virtual">client.Corporate.Cards.<a href="https://pkg.go.dev/github.com/jocall3/cli#CorporateCardService.NewVirtual">NewVirtual</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#CorporateCardNewVirtualParams">CorporateCardNewVirtualParams</a>) (<a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#CorporateCard">CorporateCard</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /corporate/cards/{cardId}/freeze">client.Corporate.Cards.<a href="https://pkg.go.dev/github.com/jocall3/cli#CorporateCardService.Freeze">Freeze</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, cardID interface{}, body <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#CorporateCardFreezeParams">CorporateCardFreezeParams</a>) (<a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#CorporateCard">CorporateCard</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /corporate/cards/{cardId}/transactions">client.Corporate.Cards.<a href="https://pkg.go.dev/github.com/jocall3/cli#CorporateCardService.ListTransactions">ListTransactions</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, cardID interface{}, query <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#CorporateCardListTransactionsParams">CorporateCardListTransactionsParams</a>) (<a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#PaginatedTransactions">PaginatedTransactions</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /corporate/cards/{cardId}/controls">client.Corporate.Cards.<a href="https://pkg.go.dev/github.com/jocall3/cli#CorporateCardService.UpdateControls">UpdateControls</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, cardID interface{}, body <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#CorporateCardUpdateControlsParams">CorporateCardUpdateControlsParams</a>) (<a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#CorporateCard">CorporateCard</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Anomalies

Response Types:

- <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#FinancialAnomaly">FinancialAnomaly</a>
- <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#CorporateAnomalyListResponse">CorporateAnomalyListResponse</a>

Methods:

- <code title="get /corporate/anomalies">client.Corporate.Anomalies.<a href="https://pkg.go.dev/github.com/jocall3/cli#CorporateAnomalyService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#CorporateAnomalyListParams">CorporateAnomalyListParams</a>) (<a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#CorporateAnomalyListResponse">CorporateAnomalyListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /corporate/anomalies/{anomalyId}/status">client.Corporate.Anomalies.<a href="https://pkg.go.dev/github.com/jocall3/cli#CorporateAnomalyService.UpdateStatus">UpdateStatus</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, anomalyID interface{}, body <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#CorporateAnomalyUpdateStatusParams">CorporateAnomalyUpdateStatusParams</a>) (<a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#FinancialAnomaly">FinancialAnomaly</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Compliance

### Audits

Response Types:

- <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#CorporateComplianceAuditRequestResponse">CorporateComplianceAuditRequestResponse</a>
- <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#CorporateComplianceAuditGetReportResponse">CorporateComplianceAuditGetReportResponse</a>

Methods:

- <code title="post /corporate/compliance/audits">client.Corporate.Compliance.Audits.<a href="https://pkg.go.dev/github.com/jocall3/cli#CorporateComplianceAuditService.Request">Request</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#CorporateComplianceAuditRequestParams">CorporateComplianceAuditRequestParams</a>) (<a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#CorporateComplianceAuditRequestResponse">CorporateComplianceAuditRequestResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /corporate/compliance/audits/{auditId}/report">client.Corporate.Compliance.Audits.<a href="https://pkg.go.dev/github.com/jocall3/cli#CorporateComplianceAuditService.GetReport">GetReport</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, auditID interface{}) (<a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#CorporateComplianceAuditGetReportResponse">CorporateComplianceAuditGetReportResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Treasury

Response Types:

- <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#CorporateTreasuryGetLiquidityPositionsResponse">CorporateTreasuryGetLiquidityPositionsResponse</a>

Methods:

- <code title="get /corporate/treasury/liquidity-positions">client.Corporate.Treasury.<a href="https://pkg.go.dev/github.com/jocall3/cli#CorporateTreasuryService.GetLiquidityPositions">GetLiquidityPositions</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#CorporateTreasuryGetLiquidityPositionsResponse">CorporateTreasuryGetLiquidityPositionsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### CashFlow

Response Types:

- <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#CorporateTreasuryCashFlowGetForecastResponse">CorporateTreasuryCashFlowGetForecastResponse</a>

Methods:

- <code title="get /corporate/treasury/cash-flow/forecast">client.Corporate.Treasury.CashFlow.<a href="https://pkg.go.dev/github.com/jocall3/cli#CorporateTreasuryCashFlowService.GetForecast">GetForecast</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#CorporateTreasuryCashFlowGetForecastParams">CorporateTreasuryCashFlowGetForecastParams</a>) (<a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#CorporateTreasuryCashFlowGetForecastResponse">CorporateTreasuryCashFlowGetForecastResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Risk

### Fraud

#### Rules

Params Types:

- <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#FraudRuleActionParam">FraudRuleActionParam</a>
- <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#FraudRuleCriteriaParam">FraudRuleCriteriaParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#FraudRule">FraudRule</a>
- <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#FraudRuleAction">FraudRuleAction</a>
- <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#FraudRuleCriteria">FraudRuleCriteria</a>
- <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#CorporateRiskFraudRuleListResponse">CorporateRiskFraudRuleListResponse</a>

Methods:

- <code title="post /corporate/risk/fraud/rules">client.Corporate.Risk.Fraud.Rules.<a href="https://pkg.go.dev/github.com/jocall3/cli#CorporateRiskFraudRuleService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#CorporateRiskFraudRuleNewParams">CorporateRiskFraudRuleNewParams</a>) (<a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#FraudRule">FraudRule</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /corporate/risk/fraud/rules/{ruleId}">client.Corporate.Risk.Fraud.Rules.<a href="https://pkg.go.dev/github.com/jocall3/cli#CorporateRiskFraudRuleService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, ruleID interface{}, body <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#CorporateRiskFraudRuleUpdateParams">CorporateRiskFraudRuleUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#FraudRule">FraudRule</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /corporate/risk/fraud/rules">client.Corporate.Risk.Fraud.Rules.<a href="https://pkg.go.dev/github.com/jocall3/cli#CorporateRiskFraudRuleService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#CorporateRiskFraudRuleListParams">CorporateRiskFraudRuleListParams</a>) (<a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#CorporateRiskFraudRuleListResponse">CorporateRiskFraudRuleListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /corporate/risk/fraud/rules/{ruleId}">client.Corporate.Risk.Fraud.Rules.<a href="https://pkg.go.dev/github.com/jocall3/cli#CorporateRiskFraudRuleService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, ruleID interface{}) <a href="https://pkg.go.dev/builtin#error">error</a></code>

# Web3

Response Types:

- <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#Web3GetNFTsResponse">Web3GetNFTsResponse</a>

Methods:

- <code title="get /web3/nfts">client.Web3.<a href="https://pkg.go.dev/github.com/jocall3/cli#Web3Service.GetNFTs">GetNFTs</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#Web3GetNFTsParams">Web3GetNFTsParams</a>) (<a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#Web3GetNFTsResponse">Web3GetNFTsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Wallets

Response Types:

- <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#CryptoWalletConnection">CryptoWalletConnection</a>
- <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#Web3WalletListResponse">Web3WalletListResponse</a>
- <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#Web3WalletGetBalancesResponse">Web3WalletGetBalancesResponse</a>

Methods:

- <code title="get /web3/wallets">client.Web3.Wallets.<a href="https://pkg.go.dev/github.com/jocall3/cli#Web3WalletService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#Web3WalletListParams">Web3WalletListParams</a>) (<a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#Web3WalletListResponse">Web3WalletListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /web3/wallets">client.Web3.Wallets.<a href="https://pkg.go.dev/github.com/jocall3/cli#Web3WalletService.Connect">Connect</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#Web3WalletConnectParams">Web3WalletConnectParams</a>) (<a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#CryptoWalletConnection">CryptoWalletConnection</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /web3/wallets/{walletId}/balances">client.Web3.Wallets.<a href="https://pkg.go.dev/github.com/jocall3/cli#Web3WalletService.GetBalances">GetBalances</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, walletID interface{}, query <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#Web3WalletGetBalancesParams">Web3WalletGetBalancesParams</a>) (<a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#Web3WalletGetBalancesResponse">Web3WalletGetBalancesResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Transactions

Response Types:

- <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#Web3TransactionInitiateTransferResponse">Web3TransactionInitiateTransferResponse</a>

Methods:

- <code title="post /web3/transactions/initiate">client.Web3.Transactions.<a href="https://pkg.go.dev/github.com/jocall3/cli#Web3TransactionService.InitiateTransfer">InitiateTransfer</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#Web3TransactionInitiateTransferParams">Web3TransactionInitiateTransferParams</a>) (<a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#Web3TransactionInitiateTransferResponse">Web3TransactionInitiateTransferResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Payments

## International

Response Types:

- <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#InternationalPaymentStatus">InternationalPaymentStatus</a>

Methods:

- <code title="post /payments/international/initiate">client.Payments.International.<a href="https://pkg.go.dev/github.com/jocall3/cli#PaymentInternationalService.Initiate">Initiate</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#PaymentInternationalInitiateParams">PaymentInternationalInitiateParams</a>) (<a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#InternationalPaymentStatus">InternationalPaymentStatus</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /payments/international/{paymentId}/status">client.Payments.International.<a href="https://pkg.go.dev/github.com/jocall3/cli#PaymentInternationalService.GetStatus">GetStatus</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, paymentID interface{}) (<a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#InternationalPaymentStatus">InternationalPaymentStatus</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Fx

Response Types:

- <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#PaymentFxConvertResponse">PaymentFxConvertResponse</a>
- <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#PaymentFxGetRatesResponse">PaymentFxGetRatesResponse</a>

Methods:

- <code title="post /payments/fx/convert">client.Payments.Fx.<a href="https://pkg.go.dev/github.com/jocall3/cli#PaymentFxService.Convert">Convert</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#PaymentFxConvertParams">PaymentFxConvertParams</a>) (<a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#PaymentFxConvertResponse">PaymentFxConvertResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /payments/fx/rates">client.Payments.Fx.<a href="https://pkg.go.dev/github.com/jocall3/cli#PaymentFxService.GetRates">GetRates</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#PaymentFxGetRatesParams">PaymentFxGetRatesParams</a>) (<a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#PaymentFxGetRatesResponse">PaymentFxGetRatesResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Sustainability

Response Types:

- <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#SustainabilityPurchaseCarbonOffsetsResponse">SustainabilityPurchaseCarbonOffsetsResponse</a>
- <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#SustainabilityGetCarbonFootprintResponse">SustainabilityGetCarbonFootprintResponse</a>

Methods:

- <code title="post /sustainability/carbon-offsets">client.Sustainability.<a href="https://pkg.go.dev/github.com/jocall3/cli#SustainabilityService.PurchaseCarbonOffsets">PurchaseCarbonOffsets</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#SustainabilityPurchaseCarbonOffsetsParams">SustainabilityPurchaseCarbonOffsetsParams</a>) (<a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#SustainabilityPurchaseCarbonOffsetsResponse">SustainabilityPurchaseCarbonOffsetsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /sustainability/carbon-footprint">client.Sustainability.<a href="https://pkg.go.dev/github.com/jocall3/cli#SustainabilityService.GetCarbonFootprint">GetCarbonFootprint</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#SustainabilityGetCarbonFootprintResponse">SustainabilityGetCarbonFootprintResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Investments

Response Types:

- <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#SustainabilityInvestmentAnalyzeImpactResponse">SustainabilityInvestmentAnalyzeImpactResponse</a>

Methods:

- <code title="get /sustainability/investments/impact">client.Sustainability.Investments.<a href="https://pkg.go.dev/github.com/jocall3/cli#SustainabilityInvestmentService.AnalyzeImpact">AnalyzeImpact</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#SustainabilityInvestmentAnalyzeImpactResponse">SustainabilityInvestmentAnalyzeImpactResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Lending

## Applications

Response Types:

- <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#LoanApplicationStatus">LoanApplicationStatus</a>

Methods:

- <code title="get /lending/applications/{applicationId}">client.Lending.Applications.<a href="https://pkg.go.dev/github.com/jocall3/cli#LendingApplicationService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, applicationID interface{}) (<a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#LoanApplicationStatus">LoanApplicationStatus</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /lending/applications">client.Lending.Applications.<a href="https://pkg.go.dev/github.com/jocall3/cli#LendingApplicationService.Submit">Submit</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#LendingApplicationSubmitParams">LendingApplicationSubmitParams</a>) (<a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#LoanApplicationStatus">LoanApplicationStatus</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Offers

Response Types:

- <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#LoanOffer">LoanOffer</a>
- <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#LendingOfferListPreApprovedResponse">LendingOfferListPreApprovedResponse</a>

Methods:

- <code title="get /lending/offers/pre-approved">client.Lending.Offers.<a href="https://pkg.go.dev/github.com/jocall3/cli#LendingOfferService.ListPreApproved">ListPreApproved</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#LendingOfferListPreApprovedParams">LendingOfferListPreApprovedParams</a>) (<a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#LendingOfferListPreApprovedResponse">LendingOfferListPreApprovedResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Developers

## Webhooks

Response Types:

- <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#WebhookSubscription">WebhookSubscription</a>
- <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#DeveloperWebhookListResponse">DeveloperWebhookListResponse</a>

Methods:

- <code title="post /developers/webhooks">client.Developers.Webhooks.<a href="https://pkg.go.dev/github.com/jocall3/cli#DeveloperWebhookService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#DeveloperWebhookNewParams">DeveloperWebhookNewParams</a>) (<a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#WebhookSubscription">WebhookSubscription</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /developers/webhooks/{subscriptionId}">client.Developers.Webhooks.<a href="https://pkg.go.dev/github.com/jocall3/cli#DeveloperWebhookService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, subscriptionID interface{}, body <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#DeveloperWebhookUpdateParams">DeveloperWebhookUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#WebhookSubscription">WebhookSubscription</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /developers/webhooks">client.Developers.Webhooks.<a href="https://pkg.go.dev/github.com/jocall3/cli#DeveloperWebhookService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#DeveloperWebhookListParams">DeveloperWebhookListParams</a>) (<a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#DeveloperWebhookListResponse">DeveloperWebhookListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /developers/webhooks/{subscriptionId}">client.Developers.Webhooks.<a href="https://pkg.go.dev/github.com/jocall3/cli#DeveloperWebhookService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, subscriptionID interface{}) <a href="https://pkg.go.dev/builtin#error">error</a></code>

## APIKeys

Response Types:

- <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#APIKey">APIKey</a>
- <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#DeveloperAPIKeyListResponse">DeveloperAPIKeyListResponse</a>

Methods:

- <code title="post /developers/api-keys">client.Developers.APIKeys.<a href="https://pkg.go.dev/github.com/jocall3/cli#DeveloperAPIKeyService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#DeveloperAPIKeyNewParams">DeveloperAPIKeyNewParams</a>) (<a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#APIKey">APIKey</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /developers/api-keys">client.Developers.APIKeys.<a href="https://pkg.go.dev/github.com/jocall3/cli#DeveloperAPIKeyService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#DeveloperAPIKeyListParams">DeveloperAPIKeyListParams</a>) (<a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#DeveloperAPIKeyListResponse">DeveloperAPIKeyListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /developers/api-keys/{keyId}">client.Developers.APIKeys.<a href="https://pkg.go.dev/github.com/jocall3/cli#DeveloperAPIKeyService.Revoke">Revoke</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, keyID interface{}) <a href="https://pkg.go.dev/builtin#error">error</a></code>

# Identity

## KYC

Response Types:

- <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#KYCStatus">KYCStatus</a>

Methods:

- <code title="get /identity/kyc/status">client.Identity.KYC.<a href="https://pkg.go.dev/github.com/jocall3/cli#IdentityKYCService.GetStatus">GetStatus</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#KYCStatus">KYCStatus</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /identity/kyc/submit">client.Identity.KYC.<a href="https://pkg.go.dev/github.com/jocall3/cli#IdentityKYCService.Submit">Submit</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#IdentityKYCSubmitParams">IdentityKYCSubmitParams</a>) (<a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#KYCStatus">KYCStatus</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Goals

Response Types:

- <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#FinancialGoal">FinancialGoal</a>
- <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#GoalListResponse">GoalListResponse</a>

Methods:

- <code title="post /goals">client.Goals.<a href="https://pkg.go.dev/github.com/jocall3/cli#GoalService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#GoalNewParams">GoalNewParams</a>) (<a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#FinancialGoal">FinancialGoal</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /goals/{goalId}">client.Goals.<a href="https://pkg.go.dev/github.com/jocall3/cli#GoalService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, goalID interface{}) (<a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#FinancialGoal">FinancialGoal</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /goals/{goalId}">client.Goals.<a href="https://pkg.go.dev/github.com/jocall3/cli#GoalService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, goalID interface{}, body <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#GoalUpdateParams">GoalUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#FinancialGoal">FinancialGoal</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /goals">client.Goals.<a href="https://pkg.go.dev/github.com/jocall3/cli#GoalService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#GoalListParams">GoalListParams</a>) (<a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#GoalListResponse">GoalListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /goals/{goalId}">client.Goals.<a href="https://pkg.go.dev/github.com/jocall3/cli#GoalService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, goalID interface{}) <a href="https://pkg.go.dev/builtin#error">error</a></code>

# Notifications

Response Types:

- <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#Notification">Notification</a>
- <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#NotificationListUserNotificationsResponse">NotificationListUserNotificationsResponse</a>

Methods:

- <code title="get /notifications/me">client.Notifications.<a href="https://pkg.go.dev/github.com/jocall3/cli#NotificationService.ListUserNotifications">ListUserNotifications</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#NotificationListUserNotificationsParams">NotificationListUserNotificationsParams</a>) (<a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#NotificationListUserNotificationsResponse">NotificationListUserNotificationsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /notifications/{notificationId}/mark-read">client.Notifications.<a href="https://pkg.go.dev/github.com/jocall3/cli#NotificationService.MarkAsRead">MarkAsRead</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, notificationID interface{}) (<a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#Notification">Notification</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Settings

Response Types:

- <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#NotificationSettings">NotificationSettings</a>

Methods:

- <code title="get /notifications/settings">client.Notifications.Settings.<a href="https://pkg.go.dev/github.com/jocall3/cli#NotificationSettingService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#NotificationSettings">NotificationSettings</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /notifications/settings">client.Notifications.Settings.<a href="https://pkg.go.dev/github.com/jocall3/cli#NotificationSettingService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#NotificationSettingUpdateParams">NotificationSettingUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#NotificationSettings">NotificationSettings</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Marketplace

## Products

Response Types:

- <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#MarketplaceProductListResponse">MarketplaceProductListResponse</a>
- <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#MarketplaceProductSimulateImpactResponse">MarketplaceProductSimulateImpactResponse</a>

Methods:

- <code title="get /marketplace/products">client.Marketplace.Products.<a href="https://pkg.go.dev/github.com/jocall3/cli#MarketplaceProductService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#MarketplaceProductListParams">MarketplaceProductListParams</a>) (<a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#MarketplaceProductListResponse">MarketplaceProductListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /marketplace/products/{productId}/impact-simulate">client.Marketplace.Products.<a href="https://pkg.go.dev/github.com/jocall3/cli#MarketplaceProductService.SimulateImpact">SimulateImpact</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, productID interface{}, body <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#MarketplaceProductSimulateImpactParams">MarketplaceProductSimulateImpactParams</a>) (<a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#MarketplaceProductSimulateImpactResponse">MarketplaceProductSimulateImpactResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Offers

Response Types:

- <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#MarketplaceOfferRedeemResponse">MarketplaceOfferRedeemResponse</a>

Methods:

- <code title="post /marketplace/offers/{offerId}/redeem">client.Marketplace.Offers.<a href="https://pkg.go.dev/github.com/jocall3/cli#MarketplaceOfferService.Redeem">Redeem</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, offerID interface{}, body <a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#MarketplaceOfferRedeemParams">MarketplaceOfferRedeemParams</a>) (<a href="https://pkg.go.dev/github.com/jocall3/cli">jocall3</a>.<a href="https://pkg.go.dev/github.com/jocall3/cli#MarketplaceOfferRedeemResponse">MarketplaceOfferRedeemResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
