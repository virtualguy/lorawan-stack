// Code generated by protoc-gen-go-json. DO NOT EDIT.
// versions:
// - protoc-gen-go-json v1.0.0
// - protoc             v3.9.1
// source: lorawan-stack/api/user.proto

package ttnpb

import (
	gogo "github.com/TheThingsIndustries/protoc-gen-go-json/gogo"
	jsonplugin "github.com/TheThingsIndustries/protoc-gen-go-json/jsonplugin"
)

// MarshalProtoJSON marshals the User message to JSON.
func (x *User) MarshalProtoJSON(s *jsonplugin.MarshalState) {
	if x == nil {
		s.WriteNil()
		return
	}
	s.WriteObjectStart()
	var wroteField bool
	if true { // (gogoproto.nullable) = false
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("ids")
		// NOTE: UserIdentifiers does not seem to implement MarshalProtoJSON.
		gogo.MarshalMessage(s, &x.UserIdentifiers)
	}
	if true { // (gogoproto.nullable) = false
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("created_at")
		s.WriteTime(x.CreatedAt)
	}
	if true { // (gogoproto.nullable) = false
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("updated_at")
		s.WriteTime(x.UpdatedAt)
	}
	if x.DeletedAt != nil || s.HasField("deleted_at") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("deleted_at")
		if x.DeletedAt == nil {
			s.WriteNil()
		} else {
			s.WriteTime(*x.DeletedAt)
		}
	}
	if x.Name != "" || s.HasField("name") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("name")
		s.WriteString(x.Name)
	}
	if x.Description != "" || s.HasField("description") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("description")
		s.WriteString(x.Description)
	}
	if x.Attributes != nil || s.HasField("attributes") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("attributes")
		s.WriteObjectStart()
		var wroteElement bool
		for k, v := range x.Attributes {
			s.WriteMoreIf(&wroteElement)
			s.WriteObjectStringField(k)
			s.WriteString(v)
		}
		s.WriteObjectEnd()
	}
	if len(x.ContactInfo) > 0 || s.HasField("contact_info") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("contact_info")
		s.WriteArrayStart()
		var wroteElement bool
		for _, element := range x.ContactInfo {
			s.WriteMoreIf(&wroteElement)
			element.MarshalProtoJSON(s.WithField("contact_info"))
		}
		s.WriteArrayEnd()
	}
	if x.PrimaryEmailAddress != "" || s.HasField("primary_email_address") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("primary_email_address")
		s.WriteString(x.PrimaryEmailAddress)
	}
	if x.PrimaryEmailAddressValidatedAt != nil || s.HasField("primary_email_address_validated_at") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("primary_email_address_validated_at")
		if x.PrimaryEmailAddressValidatedAt == nil {
			s.WriteNil()
		} else {
			s.WriteTime(*x.PrimaryEmailAddressValidatedAt)
		}
	}
	if x.Password != "" || s.HasField("password") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("password")
		s.WriteString(x.Password)
	}
	if x.PasswordUpdatedAt != nil || s.HasField("password_updated_at") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("password_updated_at")
		if x.PasswordUpdatedAt == nil {
			s.WriteNil()
		} else {
			s.WriteTime(*x.PasswordUpdatedAt)
		}
	}
	if x.RequirePasswordUpdate || s.HasField("require_password_update") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("require_password_update")
		s.WriteBool(x.RequirePasswordUpdate)
	}
	if x.State != 0 || s.HasField("state") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("state")
		x.State.MarshalProtoJSON(s)
	}
	if x.StateDescription != "" || s.HasField("state_description") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("state_description")
		s.WriteString(x.StateDescription)
	}
	if x.Admin || s.HasField("admin") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("admin")
		s.WriteBool(x.Admin)
	}
	if x.TemporaryPassword != "" || s.HasField("temporary_password") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("temporary_password")
		s.WriteString(x.TemporaryPassword)
	}
	if x.TemporaryPasswordCreatedAt != nil || s.HasField("temporary_password_created_at") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("temporary_password_created_at")
		if x.TemporaryPasswordCreatedAt == nil {
			s.WriteNil()
		} else {
			s.WriteTime(*x.TemporaryPasswordCreatedAt)
		}
	}
	if x.TemporaryPasswordExpiresAt != nil || s.HasField("temporary_password_expires_at") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("temporary_password_expires_at")
		if x.TemporaryPasswordExpiresAt == nil {
			s.WriteNil()
		} else {
			s.WriteTime(*x.TemporaryPasswordExpiresAt)
		}
	}
	if x.ProfilePicture != nil || s.HasField("profile_picture") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("profile_picture")
		// NOTE: Picture does not seem to implement MarshalProtoJSON.
		gogo.MarshalMessage(s, x.ProfilePicture)
	}
	s.WriteObjectEnd()
}

// UnmarshalProtoJSON unmarshals the User message from JSON.
func (x *User) UnmarshalProtoJSON(s *jsonplugin.UnmarshalState) {
	if s.ReadNil() {
		return
	}
	s.ReadObject(func(key string) {
		switch key {
		default:
			s.ReadAny() // ignore unknown field
		case "ids":
			s.AddField("ids")
			// NOTE: UserIdentifiers does not seem to implement UnmarshalProtoJSON.
			var v UserIdentifiers
			gogo.UnmarshalMessage(s, &v)
			x.UserIdentifiers = v
		case "created_at", "createdAt":
			s.AddField("created_at")
			v := s.ReadTime()
			if s.Err() != nil {
				return
			}
			x.CreatedAt = *v
		case "updated_at", "updatedAt":
			s.AddField("updated_at")
			v := s.ReadTime()
			if s.Err() != nil {
				return
			}
			x.UpdatedAt = *v
		case "deleted_at", "deletedAt":
			s.AddField("deleted_at")
			v := s.ReadTime()
			if s.Err() != nil {
				return
			}
			x.DeletedAt = v
		case "name":
			s.AddField("name")
			x.Name = s.ReadString()
		case "description":
			s.AddField("description")
			x.Description = s.ReadString()
		case "attributes":
			s.AddField("attributes")
			x.Attributes = make(map[string]string)
			s.ReadStringMap(func(key string) {
				x.Attributes[key] = s.ReadString()
			})
		case "contact_info", "contactInfo":
			s.AddField("contact_info")
			s.ReadArray(func() {
				if s.ReadNil() {
					x.ContactInfo = append(x.ContactInfo, nil)
					return
				}
				v := &ContactInfo{}
				v.UnmarshalProtoJSON(s.WithField("contact_info", false))
				if s.Err() != nil {
					return
				}
				x.ContactInfo = append(x.ContactInfo, v)
			})
		case "primary_email_address", "primaryEmailAddress":
			s.AddField("primary_email_address")
			x.PrimaryEmailAddress = s.ReadString()
		case "primary_email_address_validated_at", "primaryEmailAddressValidatedAt":
			s.AddField("primary_email_address_validated_at")
			v := s.ReadTime()
			if s.Err() != nil {
				return
			}
			x.PrimaryEmailAddressValidatedAt = v
		case "password":
			s.AddField("password")
			x.Password = s.ReadString()
		case "password_updated_at", "passwordUpdatedAt":
			s.AddField("password_updated_at")
			v := s.ReadTime()
			if s.Err() != nil {
				return
			}
			x.PasswordUpdatedAt = v
		case "require_password_update", "requirePasswordUpdate":
			s.AddField("require_password_update")
			x.RequirePasswordUpdate = s.ReadBool()
		case "state":
			s.AddField("state")
			x.State.UnmarshalProtoJSON(s)
		case "state_description", "stateDescription":
			s.AddField("state_description")
			x.StateDescription = s.ReadString()
		case "admin":
			s.AddField("admin")
			x.Admin = s.ReadBool()
		case "temporary_password", "temporaryPassword":
			s.AddField("temporary_password")
			x.TemporaryPassword = s.ReadString()
		case "temporary_password_created_at", "temporaryPasswordCreatedAt":
			s.AddField("temporary_password_created_at")
			v := s.ReadTime()
			if s.Err() != nil {
				return
			}
			x.TemporaryPasswordCreatedAt = v
		case "temporary_password_expires_at", "temporaryPasswordExpiresAt":
			s.AddField("temporary_password_expires_at")
			v := s.ReadTime()
			if s.Err() != nil {
				return
			}
			x.TemporaryPasswordExpiresAt = v
		case "profile_picture", "profilePicture":
			s.AddField("profile_picture")
			// NOTE: Picture does not seem to implement UnmarshalProtoJSON.
			var v Picture
			gogo.UnmarshalMessage(s, &v)
			x.ProfilePicture = &v
		}
	})
}

// MarshalProtoJSON marshals the Users message to JSON.
func (x *Users) MarshalProtoJSON(s *jsonplugin.MarshalState) {
	if x == nil {
		s.WriteNil()
		return
	}
	s.WriteObjectStart()
	var wroteField bool
	if len(x.Users) > 0 || s.HasField("users") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("users")
		s.WriteArrayStart()
		var wroteElement bool
		for _, element := range x.Users {
			s.WriteMoreIf(&wroteElement)
			element.MarshalProtoJSON(s.WithField("users"))
		}
		s.WriteArrayEnd()
	}
	s.WriteObjectEnd()
}

// UnmarshalProtoJSON unmarshals the Users message from JSON.
func (x *Users) UnmarshalProtoJSON(s *jsonplugin.UnmarshalState) {
	if s.ReadNil() {
		return
	}
	s.ReadObject(func(key string) {
		switch key {
		default:
			s.ReadAny() // ignore unknown field
		case "users":
			s.AddField("users")
			s.ReadArray(func() {
				if s.ReadNil() {
					x.Users = append(x.Users, nil)
					return
				}
				v := &User{}
				v.UnmarshalProtoJSON(s.WithField("users", false))
				if s.Err() != nil {
					return
				}
				x.Users = append(x.Users, v)
			})
		}
	})
}

// MarshalProtoJSON marshals the CreateUserRequest message to JSON.
func (x *CreateUserRequest) MarshalProtoJSON(s *jsonplugin.MarshalState) {
	if x == nil {
		s.WriteNil()
		return
	}
	s.WriteObjectStart()
	var wroteField bool
	if true { // (gogoproto.nullable) = false
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("user")
		x.User.MarshalProtoJSON(s.WithField("user"))
	}
	if x.InvitationToken != "" || s.HasField("invitation_token") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("invitation_token")
		s.WriteString(x.InvitationToken)
	}
	s.WriteObjectEnd()
}

// UnmarshalProtoJSON unmarshals the CreateUserRequest message from JSON.
func (x *CreateUserRequest) UnmarshalProtoJSON(s *jsonplugin.UnmarshalState) {
	if s.ReadNil() {
		return
	}
	s.ReadObject(func(key string) {
		switch key {
		default:
			s.ReadAny() // ignore unknown field
		case "user":
			if !s.ReadNil() {
				x.User.UnmarshalProtoJSON(s.WithField("user", true))
			}
		case "invitation_token", "invitationToken":
			s.AddField("invitation_token")
			x.InvitationToken = s.ReadString()
		}
	})
}

// MarshalProtoJSON marshals the UpdateUserRequest message to JSON.
func (x *UpdateUserRequest) MarshalProtoJSON(s *jsonplugin.MarshalState) {
	if x == nil {
		s.WriteNil()
		return
	}
	s.WriteObjectStart()
	var wroteField bool
	if true { // (gogoproto.nullable) = false
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("user")
		x.User.MarshalProtoJSON(s.WithField("user"))
	}
	if x.FieldMask != nil || s.HasField("field_mask") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("field_mask")
		if x.FieldMask == nil {
			s.WriteNil()
		} else {
			gogo.MarshalFieldMask(s, x.FieldMask)
		}
	}
	s.WriteObjectEnd()
}

// UnmarshalProtoJSON unmarshals the UpdateUserRequest message from JSON.
func (x *UpdateUserRequest) UnmarshalProtoJSON(s *jsonplugin.UnmarshalState) {
	if s.ReadNil() {
		return
	}
	s.ReadObject(func(key string) {
		switch key {
		default:
			s.ReadAny() // ignore unknown field
		case "user":
			if !s.ReadNil() {
				x.User.UnmarshalProtoJSON(s.WithField("user", true))
			}
		case "field_mask", "fieldMask":
			s.AddField("field_mask")
			v := gogo.UnmarshalFieldMask(s)
			if s.Err() != nil {
				return
			}
			x.FieldMask = v
		}
	})
}

// MarshalProtoJSON marshals the CreateUserAPIKeyRequest message to JSON.
func (x *CreateUserAPIKeyRequest) MarshalProtoJSON(s *jsonplugin.MarshalState) {
	if x == nil {
		s.WriteNil()
		return
	}
	s.WriteObjectStart()
	var wroteField bool
	if true { // (gogoproto.nullable) = false
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("user_ids")
		// NOTE: UserIdentifiers does not seem to implement MarshalProtoJSON.
		gogo.MarshalMessage(s, &x.UserIdentifiers)
	}
	if x.Name != "" || s.HasField("name") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("name")
		s.WriteString(x.Name)
	}
	if len(x.Rights) > 0 || s.HasField("rights") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("rights")
		s.WriteArrayStart()
		var wroteElement bool
		for _, element := range x.Rights {
			s.WriteMoreIf(&wroteElement)
			element.MarshalProtoJSON(s)
		}
		s.WriteArrayEnd()
	}
	if x.ExpiresAt != nil || s.HasField("expires_at") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("expires_at")
		if x.ExpiresAt == nil {
			s.WriteNil()
		} else {
			s.WriteTime(*x.ExpiresAt)
		}
	}
	s.WriteObjectEnd()
}

// UnmarshalProtoJSON unmarshals the CreateUserAPIKeyRequest message from JSON.
func (x *CreateUserAPIKeyRequest) UnmarshalProtoJSON(s *jsonplugin.UnmarshalState) {
	if s.ReadNil() {
		return
	}
	s.ReadObject(func(key string) {
		switch key {
		default:
			s.ReadAny() // ignore unknown field
		case "user_ids", "userIds":
			s.AddField("user_ids")
			// NOTE: UserIdentifiers does not seem to implement UnmarshalProtoJSON.
			var v UserIdentifiers
			gogo.UnmarshalMessage(s, &v)
			x.UserIdentifiers = v
		case "name":
			s.AddField("name")
			x.Name = s.ReadString()
		case "rights":
			s.AddField("rights")
			s.ReadArray(func() {
				var v Right
				v.UnmarshalProtoJSON(s)
				x.Rights = append(x.Rights, v)
			})
		case "expires_at", "expiresAt":
			s.AddField("expires_at")
			v := s.ReadTime()
			if s.Err() != nil {
				return
			}
			x.ExpiresAt = v
		}
	})
}

// MarshalProtoJSON marshals the UpdateUserAPIKeyRequest message to JSON.
func (x *UpdateUserAPIKeyRequest) MarshalProtoJSON(s *jsonplugin.MarshalState) {
	if x == nil {
		s.WriteNil()
		return
	}
	s.WriteObjectStart()
	var wroteField bool
	if true { // (gogoproto.nullable) = false
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("user_ids")
		// NOTE: UserIdentifiers does not seem to implement MarshalProtoJSON.
		gogo.MarshalMessage(s, &x.UserIdentifiers)
	}
	if true { // (gogoproto.nullable) = false
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("api_key")
		x.APIKey.MarshalProtoJSON(s.WithField("api_key"))
	}
	if x.FieldMask != nil || s.HasField("field_mask") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("field_mask")
		if x.FieldMask == nil {
			s.WriteNil()
		} else {
			gogo.MarshalFieldMask(s, x.FieldMask)
		}
	}
	s.WriteObjectEnd()
}

// UnmarshalProtoJSON unmarshals the UpdateUserAPIKeyRequest message from JSON.
func (x *UpdateUserAPIKeyRequest) UnmarshalProtoJSON(s *jsonplugin.UnmarshalState) {
	if s.ReadNil() {
		return
	}
	s.ReadObject(func(key string) {
		switch key {
		default:
			s.ReadAny() // ignore unknown field
		case "user_ids", "userIds":
			s.AddField("user_ids")
			// NOTE: UserIdentifiers does not seem to implement UnmarshalProtoJSON.
			var v UserIdentifiers
			gogo.UnmarshalMessage(s, &v)
			x.UserIdentifiers = v
		case "api_key", "apiKey":
			if !s.ReadNil() {
				x.APIKey.UnmarshalProtoJSON(s.WithField("api_key", true))
			}
		case "field_mask", "fieldMask":
			s.AddField("field_mask")
			v := gogo.UnmarshalFieldMask(s)
			if s.Err() != nil {
				return
			}
			x.FieldMask = v
		}
	})
}
