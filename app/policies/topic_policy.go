package policies

import(
	"golearning/app/models/topic"
	"golearning/pkg/auth"

	"github.com/gin-gonic/gin"
)

func CanModifyTopic(c *gin.Context, _topic topic.Topic) bool{
	return auth.CurrentUID(c) == _topic.UserID
}