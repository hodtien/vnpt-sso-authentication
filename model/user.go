package model

type (
	// TokenReqBody - TokenReqBody
	TokenReqBody struct {
		RefreshToken string `json:"refresh_token" validate:"required"`
	}

	// UpdateManyProfileRequest - UpdateManyProfileRequest
	UpdateManyProfileRequest struct {
		FilterName  string      `json:"filter_name" validate:"required"`
		FilterValue []string    `json:"filter_value" validate:"required"`
		UpdateName  string      `json:"update_name" validate:"required"`
		UpdateValue interface{} `json:"update_value" validate:"required"`
	}

	// UserProfile - UserProfile
	UserProfile struct {
		UserID      string                 `bson:"user_id" json:"user_id"`
		Username    string                 `bson:"username" json:"username" validate:"required,min=4,max=20"`
		Password    string                 `bson:"password" json:"password" validate:"required,min=8"`
		Firstname   string                 `bson:"first_name" json:"first_name" validate:"min=2,max=30"`
		Lastname    string                 `bson:"last_name" json:"last_name" validate:"min=2,max=30"`
		Email       string                 `bson:"email" json:"email" validate:"required,email,min=2,max=320"`
		PhoneNumber string                 `bson:"phone_number" json:"phone_number" validate:"number,min=9,max=11"`
		Address     string                 `bson:"address" json:"address,omitempty" validate:"omitempty,min=0,max=50"`
		DateOfBirth string                 `bson:"date_of_birth" json:"date_of_birth,omitempty" validate:"omitempty,number,max=10"`
		Status      string                 `bson:"status" json:"status,omitempty" validate:"number,max=20"`
		CheckExist  []string               `bson:"check_exist" json:"check_exist" validate:"required"`
		UserInfo    map[string]interface{} `bson:"user_info" json:"user_info"`
	}

	// UserRegisterIdentify - UserRegisterIdentify
	UserRegisterIdentify struct {
		ID          string                 `bson:"_id" json:"id"`
		UserID      string                 `bson:"user_id" json:"user_id"`
		Username    string                 `bson:"username" json:"username" validate:"required,min=4,max=20"`
		Password    string                 `bson:"password" json:"password,omitempty" validate:"required,min=8"`
		Firstname   string                 `bson:"first_name" json:"first_name" validate:"min=2,max=30"`
		Lastname    string                 `bson:"last_name" json:"last_name" validate:"min=2,max=30"`
		Email       string                 `bson:"email" json:"email" validate:"required,email,min=2,max=320"`
		PhoneNumber string                 `bson:"phone_number" json:"phone_number" validate:"required,number,min=9,max=11"`
		Address     string                 `bson:"address" json:"address" validate:"min=0,max=50"`
		DateOfBirth string                 `bson:"date_of_birth" json:"date_of_birth" validate:"number,max=10"`
		Status      string                 `bson:"status" json:"status" validate:"number,max=20"`
		CheckExist  []string               `bson:"check_exist" json:"check_exist" validate:"required"`
		UserInfo    map[string]interface{} `bson:"user_info" json:"user_info"`
	}

	// ResetPasswordByEmail - ResetPasswordByEmail
	ResetPasswordByEmail struct {
		Email              string `bson:"email" json:"email" validate:"required,email,min=2,max=320"`
		VerifyCode         string `json:"verify_code" validate:"required"`
		NewPassword        string `json:"new_password" validate:"required,min=8"`
		ConfirmNewPassword string `json:"confirm_new_password" validate:"required,min=8"`
	}

	// ResetPasswordByPhoneNumber - ResetPasswordByPhoneNumber
	ResetPasswordByPhoneNumber struct {
		PhoneNumber        string `json:"phone_number" validate:"required"`
		VerifyCode         string `json:"verify_code" validate:"required"`
		NewPassword        string `json:"new_password" validate:"required,min=8"`
		ConfirmNewPassword string `json:"confirm_new_password" validate:"required,min=8"`
	}

	// ConfirmResetPasswordByLink - ConfirmResetPasswordByLink
	ConfirmResetPasswordByLink struct {
		Email              string `json:"email,omitempty" validate:"omitempty,email,min=2,max=320"`
		PhoneNumber        string `json:"phone_number,omitempty" validate:"omitempty,number,min=9,max=11"`
		Token              string `json:"token" validate:"required"`
		NewPassword        string `json:"new_password" validate:"required,min=8"`
		ConfirmNewPassword string `json:"confirm_new_password" validate:"required,min=8"`
	}

	// ResetPasswordCommonTokenValidate - ResetPasswordCommonTokenValidate
	ResetPasswordCommonTokenValidate struct {
		Email       string `json:"email,omitempty" validate:"omitempty,email,min=2,max=320"`
		PhoneNumber string `json:"phone_number,omitempty" validate:"omitempty,number,min=9,max=11"`
		Token       string `json:"token" validate:"required"`
	}

	// UserIDList - UserIDList
	UserIDList struct {
		List []string `json:"user_id_list" validate:"required"`
	}

	// UpdateUserProfileModel - UpdateUserProfileModel
	UpdateUserProfileModel struct {
		Firstname   string                 `bson:"first_name" json:"first_name,omitempty" validate:"omitempty,min=2,max=30"`
		Lastname    string                 `bson:"last_name" json:"last_name,omitempty" validate:"omitempty,min=2,max=30"`
		Email       string                 `bson:"email" json:"email,omitempty" validate:"omitempty,email,min=2,max=320"`
		PhoneNumber string                 `bson:"phone_number" json:"phone_number,omitempty" validate:"omitempty,number,min=9,max=11"`
		Address     string                 `bson:"address" json:"address,omitempty" validate:"omitempty,min=0,max=50"`
		DateOfBirth string                 `bson:"date_of_birth" json:"date_of_birth,omitempty" validate:"omitempty,number,max=10"`
		Status      string                 `bson:"status" json:"status,omitempty" validate:"omitempty,number,max=20"`
		CheckExist  []string               `bson:"check_exist" json:"check_exist" validate:"required"`
		UserInfo    map[string]interface{} `bson:"user_info" json:"user_info"`
	}

	// LinkSSOUserModel - LinkSSOUserModel
	LinkSSOUserModel struct {
		UserID  string                 `json:"user_id" validate:"required"`
		Profile map[string]interface{} `json:"profile" validate:"required"`
	}

	// SearchUserModel - SearchUserModel
	SearchUserModel struct {
		Key   string `json:"key" validate:"required"`
		Limit int    `json:"limit" validate:"required,min=1"`
		Page  int    `json:"page" validate:"required,min=1"`
		Role  string `json:"role"`
	}
)
