然后main函数里gin的路由，下面这个是我的

	r := gin.Default()
	HTMLTemplates = loadTemplates()
	r.SetHTMLTemplate(HTMLTemplates)

	r.GET("/admin", admin)
	r.POST("/admin", admin)

	authorized := r.Group("/", checkLogin)
	authorized.GET("/invite", invite)
	authorized.POST("/invite", invite)
	authorized.POST("/checkName", checkName)
	authorized.POST("/mail", mail)
	authorized.POST("/preview", preview)
	authorized.POST("/changePassword", changePswd)
	authorized.POST("/logout", logout)


	Inviter结构体的定义

	type Invitor struct {
		gorm.Model
		Name string
		Password
		Remarks string
	}

	type Password struct {
		PasswordMAC  []byte
		PasswordSalt []byte
	}
	
	顺便把Password类的两个方法也发给你
	func NewPassword(password []byte) (p Password) {
		p.PasswordSalt = make([]byte, 32)
		if _, err := rand.Read(p.PasswordSalt); err != nil {
			logrus.WithError(err).Panic("Generate random number error")
		}
		mac := hmac.New(sha256.New, p.PasswordSalt)
		mac.Write(password)
		p.PasswordMAC = mac.Sum(nil)
		return
	}
	
	func (p *Password) Confirm(password []byte) bool {
		mac := hmac.New(sha256.New, p.PasswordSalt)
		mac.Write(password)
		expectedMAC := mac.Sum(nil)
		return hmac.Equal(p.PasswordMAC, expectedMAC)
	}
	
	添加和删除用户的方法

	case "add":
				// 添加inviter
				var i Invitor
				i.Name = c.PostForm("name")
				i.Remarks = c.PostForm("remarks")
				i.Password = NewPassword([]byte(c.PostForm("password")))
				if err := db.Create(&i).Error; err != nil {
					logrus.WithError(err).Error("Add new account to database error")
					c.AbortWithStatus(http.StatusInternalServerError)
					return
				}
				logrus.WithField("name", i.Name).WithField("remarks", i.Remarks).Info("Add new account")
	
			case "del":
				// 删除inviter
				id, err := strconv.ParseUint(c.PostForm("id"), 10, 0)
				if err != nil {
					logrus.WithError(err).Error("can't parse target ID")
					c.AbortWithStatus(http.StatusBadRequest)
					return
				}
				logrus.WithField("id", id).Info("Delete account")
				db.Delete(&Invitor{Model: gorm.Model{ID: uint(id)}})
			