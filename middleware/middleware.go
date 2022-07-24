package middleware

import (
	"9bany/leaky-bucket/leaky"
	"9bany/leaky-bucket/rule"
	"crypto/md5"
	"fmt"
	"math/big"
	"net/http"

	"github.com/gin-gonic/gin"
)

var leakyCollector = leaky.NewCollector(true)

func Middleware(ctx *gin.Context) {
	userType := ctx.GetHeader("user-type")

	if userType == "" {
		userType = "gen-user"
	}

	rule := rule.GetBucket(userType)

	count := leakyCollector.Add(GetClientIndentifire(ctx), rule.Amout, (1<<20)*10, 10)

	if count == 0 {
		ctx.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
			"message": "Try again after sometime!",
		})
		return
	}
	ctx.Next()

}

func GetClientIndentifire(ctx *gin.Context) string {
	ip := ctx.ClientIP()
	url := ctx.Request.URL.Path
	data := fmt.Sprintf("%s-%s", ip, url)
	h := md5.Sum([]byte(data))
	hash := new(big.Int).SetBytes(h[:]).Text(62)
	return hash
}
