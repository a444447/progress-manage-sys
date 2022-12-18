package middware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"net/http"
	"progress-manage-system/utils"
	"progress-manage-system/utils/errmsg"
	"strings"
	"time"
)

// 用于加密jwt的特殊签名字符串
var JwtKey = []byte(utils.JwtKey)

// 生成token的参数
type Claims struct {
	IdentityID string `json:"identityID"`
	PassWord   string `json:"password"`
	jwt.StandardClaims
}

// GenerateToken 生成token
func GenerateToken(identityID, password string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour) //过期时间设置为3小时

	claims := Claims{
		IdentityID: identityID,
		PassWord:   password,
		StandardClaims: jwt.StandardClaims{
			//过期时间
			ExpiresAt: expireTime.Unix(),
			//TOKEN发行人
			Issuer: "progress_sys",
		},
	}

	reqClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := reqClaims.SignedString(JwtKey)
	if err != nil {
		return "", errors.Wrapf(err, "error->GenerateToken:")
	}
	return token, nil
}

// ParseToken 解析token
func ParseToken(token string) (*Claims, error) {
	//用于解析鉴权的声明，方法内部主要是具体的解码和校验的过程，最终返回*Token
	parsedToken, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})
	if key, ok := parsedToken.Claims.(*Claims); ok && parsedToken.Valid {
		return key, nil
	} else {
		return nil, err
	}
}

// JwtToken jwt中间件
func JwtToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenHeader := c.Request.Header.Get("Authorization")
		var code int
		if tokenHeader == "" {
			code = errmsg.ErrorTokenNotExist
			c.JSON(http.StatusOK, gin.H{
				"code":    code,
				"message": errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}
		//jwt中规定，authorization需要以bearer开头
		//jwt可以放在cookie中进行返回，但是这样不容易跨域；因此放在header的authorization字段中可以更好处理跨域问题
		checkToken := strings.SplitN(tokenHeader, " ", 2)
		if len(checkToken) != 2 && checkToken[0] != "Bearer" {
			code = errmsg.ErrorTokenType
			c.JSON(http.StatusOK, gin.H{
				"code":    code,
				"message": errmsg.GetErrMsg(code),
			})
		}
		//解析token
		key, _ := ParseToken(checkToken[1])
		if key == nil {
			code = errmsg.ErrorTokenWrong
			c.JSON(http.StatusOK, gin.H{
				"code":     code,
				"messsage": errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}
		//检验token是否过期
		if time.Now().Unix() > key.ExpiresAt {
			code = errmsg.ErrorTokenRuntime
			c.JSON(http.StatusOK, gin.H{
				"code":    code,
				"message": errmsg.GetErrMsg(code),
			})
			c.Abort()
			return
		}
		c.Set("user", key)
		c.Next()
	}
}
