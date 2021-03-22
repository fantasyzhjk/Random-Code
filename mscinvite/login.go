package main

import (
	"flag"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var secure = flag.Bool("secure", true, "cookie是否禁用http")

func checkLogin(c *gin.Context) {
	if user, ok := c.GetPostForm("user"); ok {
		// user login
		login(c, user)
		return
	}
	// check login status
	tk, err := c.Cookie("invite_tk")
	if err != nil {
		c.HTML(http.StatusUnauthorized, "login.html", nil)
		c.Abort()
		return
	}
	// parse jwt token
	token, err := jwt.ParseWithClaims(tk, &JwtClaims{}, func(*jwt.Token) (interface{}, error) {
		return []byte(*tokenKey), nil
	})
	if err != nil {
		logrus.WithError(err).Error("Parse jwt token error")
		c.HTML(http.StatusUnauthorized, "login.html", nil)
		c.Abort()
		return
	}
	// check token valid, read claims
	if claims, ok := token.Claims.(*JwtClaims); ok && token.Valid {
		c.Set("name", claims.Name)
	} else {
		c.HTML(http.StatusUnauthorized, "login.html", nil)
		c.Abort()
		return
	}
}

func login(c *gin.Context, username string) {
	// pull user data form database
	var inviter Invitor
	if db.Where("Name = ?", username).First(&inviter).RecordNotFound() {
		logrus.WithField("user", username).Debug("Login fail: user not found")
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	// Check password
	if !inviter.Password.Confirm([]byte(c.PostForm("password"))) {
		logrus.WithField("user", username).Debug("Login fail: password confirm fail")
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	// generate jwt token
	now := time.Now()
	claims := JwtClaims{
		Name: inviter.Name,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  now.Unix(),
			ExpiresAt: now.Add(time.Duration(*cookieAge) * time.Second).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	tk, err := token.SignedString([]byte(*tokenKey))
	if err != nil {
		logrus.WithError(err).Error("Signed jwt token error")
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.SetCookie("invite_tk", tk, *cookieAge, *cookiePath, *cookieDomain, *secure, true)
	c.Set("name", claims.Name)
}

func logout(c *gin.Context) {
	c.SetCookie("invite_tk", "", -1, *cookiePath, *cookieDomain, *secure, true)
	c.AbortWithStatus(http.StatusNoContent)
}

func changePswd(c *gin.Context) {
	username := c.GetString("name")
	newPswd := c.PostForm("new_password")
	if newPswd == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "password is empty"})
		return
	}
	var inviter Invitor
	if db.Where("Name = ?", username).First(&inviter).RecordNotFound() {
		logrus.WithField("user", username).Debug("Confirm fail: user not found")
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	// Check password
	if !inviter.Password.Confirm([]byte(c.PostForm("password"))) {
		logrus.WithField("user", username).Debug("Confirm fail: password confirm fail")
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	// Set password
	inviter.Password = NewPassword([]byte(newPswd))
	if err := db.Save(&inviter).Error; err != nil {
		logrus.WithField("user", username).WithError(err).Error("Change password fail")
		c.AbortWithStatus(http.StatusInternalServerError)
	}
}
