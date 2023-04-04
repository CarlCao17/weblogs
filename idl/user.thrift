namespace go github.com.weblogs.user

include "./base.thrift"
include "./model.thrift"

struct CreateUserRequest {
    1: required string UserName
    2: required string Password
    3: required string Email
    4: optional string FirstName
    5: optional string LastName
}

struct CreateUserResponse {
    1: required i64 UserID
    2: required model.UserRole Role
    255: base.BaseResp BaseResp
}

struct DeleteUserRequest {
    1: required i64 UserID
}

struct DeleteUserResponse {
    255: base.BaseResp BaseResp
}

struct SignInRequest {
    1: required string UserName
    2: required string Password
}

struct SignInResponse {
    1: required i64 UserID
    255: base.BaseResp BaseResp
}

struct SignOutRequest {
    1: required i64 UserID
}

struct SignOutResponse {
    255: base.BaseResp BaseResp
}

service UserService {
    CreateUserResponse CreateUser(1: CreateUserRequest req)
    DeleteUserResponse DeleteUser(1: DeleteUserRequest req)
    SignInResponse SignIn(1: SignInRequest req)
    SignOutResponse SignOut(1: SignOutRequest req)
}