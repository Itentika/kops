// Code generated by protoc-gen-goext. DO NOT EDIT.

package iam

import (
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

type CreateIamTokenRequest_Identity = isCreateIamTokenRequest_Identity

func (m *CreateIamTokenRequest) SetIdentity(v CreateIamTokenRequest_Identity) {
	m.Identity = v
}

func (m *CreateIamTokenRequest) SetYandexPassportOauthToken(v string) {
	m.Identity = &CreateIamTokenRequest_YandexPassportOauthToken{
		YandexPassportOauthToken: v,
	}
}

func (m *CreateIamTokenRequest) SetJwt(v string) {
	m.Identity = &CreateIamTokenRequest_Jwt{
		Jwt: v,
	}
}

func (m *CreateIamTokenResponse) SetIamToken(v string) {
	m.IamToken = v
}

func (m *CreateIamTokenResponse) SetExpiresAt(v *timestamppb.Timestamp) {
	m.ExpiresAt = v
}

func (m *CreateIamTokenForServiceAccountRequest) SetServiceAccountId(v string) {
	m.ServiceAccountId = v
}