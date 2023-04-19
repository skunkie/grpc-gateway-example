# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [api.proto](#api-proto)
    - [Api](#pb-Api)
  
- [token.proto](#token-proto)
    - [CreateTokenReq](#pb-CreateTokenReq)
    - [Token](#pb-Token)
  
- [users.proto](#users-proto)
    - [CreateUserReq](#pb-CreateUserReq)
    - [Id](#pb-Id)
    - [Ids](#pb-Ids)
    - [UpdateUserReq](#pb-UpdateUserReq)
    - [User](#pb-User)
    - [Users](#pb-Users)
  
- [Scalar Value Types](#scalar-value-types)



<a name="api-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## api.proto


 

 

 


<a name="pb-Api"></a>

### Api


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| CreateToken | [CreateTokenReq](#pb-CreateTokenReq) | [Token](#pb-Token) | Authorization |
| ListUsers | [Ids](#pb-Ids) | [Users](#pb-Users) | List Users |
| CreateUser | [CreateUserReq](#pb-CreateUserReq) | [User](#pb-User) | Create User |
| GetUser | [Id](#pb-Id) | [User](#pb-User) | Get User |
| UpdateUser | [UpdateUserReq](#pb-UpdateUserReq) | [User](#pb-User) | Update User Info |
| DeleteUser | [Id](#pb-Id) | [.google.protobuf.Empty](#google-protobuf-Empty) | Delete User |

 



<a name="token-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## token.proto



<a name="pb-CreateTokenReq"></a>

### CreateTokenReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| username | [string](#string) |  |  |
| password | [string](#string) |  |  |






<a name="pb-Token"></a>

### Token



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| token_string | [string](#string) |  |  |





 

 

 

 



<a name="users-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## users.proto



<a name="pb-CreateUserReq"></a>

### CreateUserReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| email | [string](#string) |  |  |






<a name="pb-Id"></a>

### Id



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="pb-Ids"></a>

### Ids



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) | repeated |  |






<a name="pb-UpdateUserReq"></a>

### UpdateUserReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| name | [string](#string) | optional |  |
| email | [string](#string) | optional |  |






<a name="pb-User"></a>

### User



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| name | [string](#string) |  |  |
| email | [string](#string) |  |  |






<a name="pb-Users"></a>

### Users



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| users | [User](#pb-User) | repeated |  |





 

 

 

 



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

