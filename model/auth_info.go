package model

type (
	// AuthenticationInfo - AuthenticationInfo
	AuthenticationInfo struct {
		UserID   string     `json:"user_id" validate:"required"`
		APIKey   string     `json:"api_key" validate:"required"`
		AuthInfo []AuthInfo `json:"auth_info" validate:"required"`
	}

	// AuthInfo - AuthInfo
	AuthInfo struct {
		BucketID   string `json:"bucket_id" validate:"required"`
		Permission string `json:"permission" validate:"required"`
	}

	// DataLogin - DataLogin
	DataLogin struct {
		Username string `json:"username" validate:"required,min=4,max=20"`
		Password string `json:"password" validate:"required,min=8"`
	}

	// CheckUserExist - CheckUserExist
	CheckUserExist struct {
		Username string `json:"username" validate:"required,min=4,max=20"`
	}

	// CheckEmailExist - CheckEmailExist
	CheckEmailExist struct {
		Email string `bson:"email" json:"email" validate:"required,email,min=2,max=320"`
	}

	// CheckPhoneExist - CheckPhoneExist
	CheckPhoneExist struct {
		PhoneNumber string `bson:"phone_number" json:"phone_number" validate:"required,number,min=9,max=11"`
	}

	// UpdatePasswordReq - UpdatePasswordReq
	UpdatePasswordReq struct {
		OldPassword        string `json:"old_password,omitempty" validate:"min=8"`
		NewPassword        string `json:"new_password,omitempty" validate:"min=8"`
		ConfirmNewPassword string `json:"confirm_new_password,omitempty" validate:"min=8"`
	}
)
