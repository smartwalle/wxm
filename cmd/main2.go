package main

//import (
//	"github.com/gin-gonic/gin"
//	"github.com/smartwalle/wxm"
//	"net/http"
//)
//
//func main() {
//	var client = wxm.New("wx7262a2f023e9aef8", "133850581b156b304ed23b30766aee90")
//
//	var s = gin.Default()
//	s.GET("/MP_verify_VgtkH8dpacKCg867.txt", func(c *gin.Context) {
//		c.Writer.WriteString("VgtkH8dpacKCg867")
//	})
//
//	s.GET("/auth", func(c *gin.Context) {
//		c.Redirect(http.StatusTemporaryRedirect, client.GetAuthorizeURL("http://fxx.scyoule.com/", wxm.AuthScopeBase, "hehehe"))
//	})
//
//	s.GET("/", func(c *gin.Context) {
//		c.Request.ParseForm()
//		var code = c.Request.Form.Get("code")
//		var token, err = client.GetAccessToken(code)
//		if err != nil {
//			c.Writer.WriteString(err.Error())
//			return
//		}
//		c.Writer.WriteString(token.OpenId)
//	})
//
//	s.Run(":9800")
//}
