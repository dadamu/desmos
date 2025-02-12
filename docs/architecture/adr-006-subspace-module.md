# ADR 006: Subspace module

## Changelog

- December 15th, 2021: Initial draft;
- December 16th, 2021: First review;
- January 06th, 2022: Second review;
- January 13th, 2021: Third review;
- January 14th, 2021: Fourth review;
- January 17th, 2021: Fifth review.

## Status

PROPOSED

## Abstract
This ADR defines the `x/subspaces` module which allows users to create and manage the representation of different social networks inside which contents will be created.  

## Context
In order to support building independent social networks, it is vital for Desmos to expose a mechanism that allows to replicate as much as possible the current stace of social networks in which each platform has different economic model, set of administrators, Terms of Services and ways of enforcing such terms.  

## Decision
We will create a module named `subspaces` which provided users the ability to create independent spaces inside Desmos, each one having its own administrators, ToS and tokenomics.

### Types 
Subspaces must always have an account that is elected as the _owner_ and should act as the final decision maker and accountable user. Additionally, each subspace can have the following data:

* a _human-readable name_ and a _description_, in order to allow users to easily identify the scope of such subspace;
* a set of _connections_ with external applications, allowing users to verify the validity of such subspace and avoid any possible fakes;
* an _ACL_ (access-control list) defining the permissions that different kind of users have inside the subspace itself (e.g. administrators, moderators, etc.).

#### Subspace

```go
// Subspace contains all the data related to a Desmos subspace
type Subspace struct {
  // Unique identifier of this subspace
  ID uint64

  // Human-readable name of the subspace
  Name string
  
  // Optional description of this subspace
  Description string
  
  // Address of the user that owns the subspace
  Owner  string

  // Address of the subspace creator
  Creator string
  
  // Represents the account that is associated with the subspace and
  // should be used to connect external applications to verify this subspace
  Treasury string

  // the creation time of the subspace
  CreationTime time.Time
}
```

#### ACL
In order to easily implement an ACL, we will use a simple set of keys made as follows: 
```
ACLPrefix + Subspace ID + <Group name | User Address> -> ACL Value
```

The `ACL Value` will be a simple binary value allowing us to perform bitwise operations to combine the following different permissions: 

```go
const (
  // Identifies users that can create content inside the subspace
  PermissionWrite           = 0b000001
  
  // Allows users to moderate contents of other users (e.g. deleting it) 
  PermissionModerateContent = 0b000010
  
  // Allows to add a link for this subspace
  PermissionAddLink         = 0b000100
  
  // Allows to change the information of the subspace
  PermissionChangeInfo      = 0b001000
  
  // Allows to set other users' permissions (except PermissionSetPermissions). 
  // This includes managing user groups and the associated permissions
  PermissionSetPermissions  = 0b010000
)
```

> **Note**:  
> Only the `Owner` account will be able to grant other users the `PermissionSetPermissions`

Using this kind of permissions will allow us to easily set permissions and check whether a user has a permission or not: 
```go
userPermissions := PermissionWrite | PermissionAddLink | PermissionChangeInfo

canWrite := (userPermissions & PermissionWrite) == PermissionWrite  // True
canModerateContent := (userPermissions & PermissionModerateContent) == PermissionModerateContent // False
```

##### Group permissions 
In order to simplify the handling of multiple users' permissions, we SHOULD allow subspace admins to set group-wide permissions. 
Each group will be represented by its own name, which can be defined by the admins themselves, and each group will have the same importance as others.
We reserve the group name `Others` to identify all the users that are not part of any other group (i.e. those users who are not registered inside a subspace).

While checking a user permission to do something, the following actions will be performed: 
1. get the permissions set for that specific user;
2. get all the permissions for all the groups the user is part of; 
3. compute the resulting permission associated to the user using the `OR` operator.

In order to properly store groups and members information, the following store keys will be used: 
``` 
// Store the belonging of a group to a subspace
SubspaceGroupPrefix + Subspace ID + Group name -> 0x01

// Store the belonging of a user to a specific subspace group
UserGroupPrefix + Subspace ID + Group name + User Address -> 0x01 
```

In both cases, the `0x01` value is used only as a placeholder to make it possible for the key to exist.

These keys will allow us to iterate over all the users that are part of a group as well as all the groups inside a subspace. This will allow clients to easily get all the groups of a subspace and the users that are part of such groups. On the other hand, it will make it harder to get all the groups that a user is part of. The decision to prioritize the first instead of the latter is made to make sure that permission checking is sufficiently fast: since subspace will have a limited amount of groups, it will be quite inexpensive to iterate over all of them and check if a given user is part of each group.

### `Msg` Service
We will allow the following operations to be performed.

**Subspace administration**
* Create a subspace
* Edit a subspace

**Content management**
* Delete contents that do not respect the ToS

**Groups management**
* Create a new group
* Delete a group
* Set group permissions

**Users management**
* Add a user to a group
* Remove a user from a group
* Set user permissions


```protobuf
// Msg defines subspaces Msg service.
service Msg {

  // CreateSubspace allows to create a subspace
  rpc CreateSubspace(MsgCreateSubspace) returns (MsgCreateSubspaceResponse);

  // EditSubspace allows to edit a subspace
  rpc EditSubspace(MsgEditSubspace) returns (MsgEditSubspaceResponse);
  
  // CreateUserGroup allows to create a new user group
  rpc CreateUserGroup(MsgCreateUserGroup) returns (MsgCreateUserGroupResponse);
  
  // DeleteUserGroup allows to delete an existing user group
  rpc DeleteUserGroup(MsgDeleteUserGroup) returns (MsgDeleteUserGroupResponse);
  
  // AddUserToUserGroup allows to add a specific user to a specific user group
  rpc AddUserToUserGroup(MsgAddUserToUserGroup) returns (MsgAddUserToUserGroupResponse);

  // RemoveUserFromUserGroup allows to remove a specific user from a specific user group
  rpc RemoveUserFromUserGroup(MsgRemoveUserFromUserGroup) returns (MsgRemoveUserFromUserGroupResponse); 
  
  // SetPermissions allows to set the permissions of a user or user group
  rpc SetPermissions(MsgSetPermissions) returns (MsgSetPermissionsResponse);
}

message MsgCreateSubspace {
  string name = 1;
  string description = 2;
  string owner = 3;
  string treasury = 4;
  string creator = 5;
}

message MsgCreateSubspaceResponse {
  uint64 subspace_id = 1;
}

message MsgEditSubspace {
  uint64 id = 1;
  string name = 2;
  string description = 3;
  string owner = 4;
  string treasury = 5;
  string signer = 6;
}

message MsgEditSubspaceResponse {}

message MsgCreateUserGroup {
  uint64 subspace_id = 1;
  string group_name = 2;
  bytes default_permissions = 3;
  string signer = 4;
}

message MsgCreateUserGroupResponse {}

message MsgDeleteUserGroup {
  uint64 subspace_id = 1;
  string group_name = 2;
  string signer = 3;
}

message MsgDeleteUserGroupResponse {}

message MsgAddUserToUserGroup { 
  uint64 subspace_id = 1;
  string user = 2;
  string group_name = 3;
  string signer = 4;
}

message MsgAddUserToUserGroupResponse {}

message MsgRemoveUserFromUserGroup {
  uint64 subspace_id = 1;
  string user = 2; 
  string group_name = 3; 
  string signer = 4;
}

message MsgRemoveUserFromUserGroupResponse {}

message MsgSetPermissions {
  uint64 subspace_id = 1;
  string target = 2;
  bytes permissions = 3;
  string signer = 4;
}

message MsgSetPermissionsResponse {}
```

## Consequences
### Positive

* Users will be able to create their own subspace representing a social network inside Desmos
* The ACL implementation proposed is generic enough to allow future permissions to be implemented without much work to be done
* The link-based verification system proposed is generic enough to allow any creator to verify their subspace relying on the already existing `x/profiles` module without the need of new code 

### Negative

### Neutral

## References
- Extend the concept of subspaces: https://github.com/desmos-labs/desmos/discussions/375