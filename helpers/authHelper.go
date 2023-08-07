package helper

import(
	"errors"
	"github.com/gin-gonic/gin"
)

// Functions that check user type and return error or nil error

func CheckUserType(c *gin.Context, role string) (err error) {
	userType := c.GetString("user_type")
	err = nil
	if userType != role {
		err = errors.New("You must be an administrator to access this!")
		return err
	}
	return err
}

func MatchUserTypeToUid(c *gin.Context, userId string) (err error) {
	userType := c.GetString("user_type")
	uId := c.GetString("uid")
	err = nil

	if userType =="USER" && uId != userId {
		err = errors.New("You must be an administrator to access this!")
		return err
	}
	err = CheckUserType(c, userType)
	return err
}

