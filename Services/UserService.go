package Services

import (
	"elearn100/Model/Admin"
	"elearn100/Pkg/e"
	"elearn100/Pkg/util"
	"github.com/astaxie/beego/validation"
	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"strconv"
)

// @Desc登录验证
func Login(c *gin.Context) (int, string) {
	c.Request.Body = e.GetBody(c)
	conn := e.PoolConnect()
	defer conn.Close()

	if tokVal, isOk := redis.String(conn.Do("get", e.Token)); isOk == nil {
		return e.SUCCESS, tokVal
	}

	loginName := e.Trim(c.PostForm("uname"))
	pwd := e.Trim(c.PostForm("pword"))

	if code, err := validLogin(loginName, pwd); code == e.ERROR {
		return code, err
	}

	err, uuid := Admin.GetUserInfo(loginName, pwd)
	if uuid < 1 {
		return e.ERROR, err.Error()
	}
	token := util.GetSignContent(c)
	SaveUserInfo(token)
	return e.SUCCESS, token
}

// @Desc 登录校验
func validLogin(loginName, pwd string) (int, string) {
	valid := validation.Validation{}
	valid.Required(loginName, "login_name").Message("用户名不能为空")
	valid.Required(pwd, "pwd").Message("密码不能为空")
	if !valid.HasErrors() {
		return e.ReSuccess()
	}
	return e.ViewErr(valid)
}

// @Desc 添加/编辑用户
func AddUser(c *gin.Context) (code int, err string) {
	c.Request.Body = e.GetBody(c)
	nickName := c.PostForm("nick_name")
	id := com.StrTo(c.PostForm("id")).MustInt()
	loginName := com.StrTo(c.PostForm("login_name")).String()
	email := com.StrTo(c.PostForm("email")).String()
	pwd := com.StrTo(c.PostForm("pwd")).String()
	statues := com.StrTo(c.PostForm("statues")).MustInt64()
	tel := e.Trim(com.StrTo(c.PostForm("tel")).String())

	if code, msg := validUsers(nickName, loginName, email, tel, pwd, statues); code == e.ERROR {
		return code, msg
	}

	if !e.CheckPhone(tel) {
		return e.ERROR, "手机号码格式不正确"
	}

	var user Admin.SysAdminUser
	user.NickName = nickName
	user.LoginName = loginName
	user.Email = email
	if id < 1 {
		user.Pwd = Admin.Md5Pwd(pwd)
	}
	user.Statues = statues
	user.Tel = tel
	user.DepartmentId = 1
	user.Avatar = "#"
	user.CityId = "10000"
	user.PositionId = "1"

	var isOk bool
	if id < 1 { //编辑
		if Admin.ExistsByLoginName(loginName) {
			return e.ERROR, "账号已存在，填写新的账号"
		}

		if Admin.ExistsByTel(tel) {
			return e.ERROR, "手机号码已存在，填写新的手机号码"
		}
		isOk = Admin.AddUser(user)
	} else {
		userInfo := Admin.GetAdmin(id)
		if userInfo.Tel != tel {
			if Admin.ExistsByTel(tel) {
				return e.ERROR, "手机号码已存在，填写新的手机号码"
			}
		}
		if userInfo.Pwd == pwd {
			user.Pwd = pwd
		} else {
			user.Pwd = Admin.Md5Pwd(pwd)
		}
		isOk = Admin.EditUser(id, user)
	}

	if isOk {
		return e.ReSuccess()
	}
	return e.ReError()
}

// @Desc 添加用户验证
func validUsers(nickName, loginName, email, tel, pwd string, statues int64) (int, string) {
	valid := validation.Validation{}
	valid.Required(nickName, "nick_name").Message("昵称不能为空")
	valid.Required(loginName, "login_name").Message("账号不能为空")
	valid.Required(email, "email").Message("邮箱不能为空")
	valid.Required(tel, "tel").Message("手机号码不能为空")
	valid.Required(statues, "statues").Message("状态必选")
	valid.Required(pwd, "pwd").Message("密码不能为空")
	if !valid.HasErrors() {
		return e.ReSuccess()
	}
	return e.ViewErr(valid)
}

// @Summer 修改用户信息和密码
func EditUser(c *gin.Context) (code int, err string) {
	c.Request.Body = e.GetBody(c)
	id := com.StrTo(c.PostForm("id")).MustInt()
	email := com.StrTo(c.PostForm("email")).String()
	pwd := com.StrTo(c.PostForm("pwd")).String()
	newPwd := com.StrTo(c.PostForm("newpwd")).String()
	tel := e.Trim(com.StrTo(c.PostForm("tel")).String())
	act := e.Trim(com.StrTo(c.PostForm("act")).String())

	if code, msg := validUser(id, act, email, tel, pwd, newPwd); code == e.ERROR {
		return code, msg
	}

	user := Admin.GetAdmin(id)
	var userInfo Admin.SysAdminUser
	if act == "user" {
		if !e.CheckPhone(tel) {
			return e.ERROR, "手机号码格式不正确"
		}

		userInfo.Tel = tel
		userInfo.Email = email

		if user.Tel != tel {
			if Admin.ExistsByTel(tel) {
				return e.ERROR, "手机号码已存在，填写新的手机号码"
			}
		}
	} else {
		if user.Pwd != Admin.Md5Pwd(pwd) {
			return e.ERROR, "原始密码不正确"
		}
		userInfo.Pwd = Admin.Md5Pwd(newPwd)
	}

	msg := ""
	if Admin.EditUser(id, userInfo) {
		if act == "user" {
			conn := e.PoolConnect()
			defer conn.Close()

			conn.Do("expire", e.Token, -1)
			token := util.GetSignContent(c)
			SaveUserInfo(token)
			msg = "用户信息"
		} else {
			if LogOut() {
				c.Header("Cache-Control", "no-cache,no-store")
			}
			msg = "修改密码"
		}
		return e.SUCCESS, msg
	}
	return e.ReError()
}

// @Desc 数据验证
func validUser(id int, act, email, tel, pwd, newPwd string) (int, string) {
	valid := validation.Validation{}
	valid.Required(id, "id").Message("操作失败")
	valid.Min(id, 0, "id").Message("操作失败")
	if act == "user" {
		valid.Required(email, "email").Message("邮箱不能为空")
		valid.Required(tel, "tel").Message("手机号码不能为空")
	} else {
		valid.Required(pwd, "pwd").Message("原始密码不能为空")
		valid.Required(newPwd, "newpwd").Message("新密码不能为空")
	}
	if !valid.HasErrors() {
		return e.ReSuccess()
	}
	return e.ViewErr(valid)
}

// @Summer 获取当个用户信息
func GetUser(c *gin.Context) (int, string, interface{}) {
	c.Request.Body = e.GetBody(c)
	id := com.StrTo(c.PostForm("id")).MustInt()
	var data interface{}
	if id < 1 {
		return e.ERROR, "ID必须大于0", data
	}

	data = Admin.GetAdmin(id)
	return e.SUCCESS, "操作成功", data
}

//管理员登录存储缓存中
func SaveUserInfo(tokenStr string) {
	conn := e.PoolConnect()
	defer conn.Close()

	if _, err := conn.Do("set", e.Token, tokenStr); err == nil {
		conn.Do("expire", e.Token, e.VALIDTime)
	}
}

// @Desc退出
func LogOut() bool {
	conn := e.PoolConnect()
	defer conn.Close()

	_, err := redis.Bool(conn.Do("expire", e.Token, -1))
	if err == nil {
		return true
	}
	return false
}

func DetailsUser(c *gin.Context) (err error, admins Admin.SysAdminUser) {
	uuid, uOk := c.Request.Cookie("uuid")
	if uOk != nil {
		return err, Admin.SysAdminUser{}
	} else {
		uid, err := strconv.Atoi(uuid.Value)
		if err != nil {
			return err, Admin.SysAdminUser{}
		}
		admins = Admin.GetAdmin(uid)
	}
	return nil, admins
}

func GetLastUser() (id int) {
	return Admin.GetLastUserId().ID
}
