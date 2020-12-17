package Services

import (
	"elearn100/Model/Admin"
	"elearn100/Pkg/e"
	"elearn100/Pkg/util"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"regexp"
	"strconv"
)

//登录验证
func Login(c *gin.Context) (code int, err string) {
	isOk, tokVal := e.GetVal("token")
	if isOk {
		return e.SUCCESS, tokVal
	}
	c.Request.Body = e.GetBody(c)
	loginName := e.Trim(c.PostForm("uname"))
	pwd := e.Trim(c.PostForm("pword"))

	valid := validation.Validation{}
	valid.Required(loginName, "login_name").Message("用户名不能为空")
	valid.Required(pwd, "pwd").Message("密码不能为空")
	if !valid.HasErrors() {
		err, uuid := Admin.GetUserInfo(loginName, pwd)
		if err != nil {
			return e.ERROR, err.Error()
		}
		token := util.GetSignContent(c)
		e.SetCookie(c, uuid, 10800)
		e.SetVal("loginuid", string(uuid))
		SaveUserInfo(uuid)
		return e.SUCCESS, token
	}
	return ViewErr(valid)
}

// @Summer 添加/编辑用户
func AddUser(c *gin.Context) (code int, err string) {
	c.Request.Body = e.GetBody(c)
	nickName := c.PostForm("nick_name")
	id := com.StrTo(c.PostForm("id")).MustInt64()
	loginName := com.StrTo(c.PostForm("login_name")).String()
	email := com.StrTo(c.PostForm("email")).String()
	pwd := com.StrTo(c.PostForm("pwd")).String()
	statues := com.StrTo(c.PostForm("status")).MustInt64()
	tel := e.Trim(com.StrTo(c.PostForm("tel")).String())

	valid := validation.Validation{}
	valid.Required(nickName, "nick_name").Message("昵称不能为空")
	valid.Required(loginName, "login_name").Message("账号不能为空")
	valid.Required(email, "email").Message("邮箱不能为空")
	valid.Required(tel, "tel").Message("手机号码不能为空")
	valid.Required(statues, "statues").Message("状态必选")

	if id < 1 {
		valid.Required(pwd, "pwd").Message("密码不能为空")
	}

	data := make(map[string]interface{})

	if !valid.HasErrors() {
		reg := regexp.MustCompile(`^1{1}\d{10}$`)
		if !reg.MatchString(tel) || len(tel) < 11 {
			return e.ERROR, "手机号码格式不正确"
		}
		data["nick_name"] = nickName
		data["login_name"] = loginName
		data["email"] = email
		if id < 1 {
			data["pwd"] = Admin.Md5Pwd(pwd)
		} else {
			if len(pwd) > 1 {
				data["pwd"] = Admin.Md5Pwd(pwd)
			}
		}
		data["statues"] = statues
		data["tel"] = tel
		var isOk bool
		if id < 1 { //编辑
			validLogin := Admin.ExistsByLoginName(loginName)
			if validLogin {
				return e.ERROR, "账号已存在，填写新的账号"
			}
			validTel := Admin.ExistsByTel(tel)
			if validTel {
				return e.ERROR, "手机号码已存在，填写新的手机号码"
			}
			if !validTel && !validTel {
				isOk = Admin.AddUser(data)
			}

		} else {
			user := Admin.Find(id)
			if user.Tel != tel {
				validTel := Admin.ExistsByTel(tel)
				if validTel {
					return e.ERROR, "手机号码已存在，填写新的手机号码"
				}
			}
			isOk = Admin.EditUser(id, data)
		}

		if isOk {
			return e.SUCCESS, "操作成功"
		}
		return e.ERROR, "操作失败"
	}
	return ViewErr(valid)
}

// @Summer 修改用户信息和密码
func EditUser(c *gin.Context) (code int, err string) {
	c.Request.Body = e.GetBody(c)
	id := com.StrTo(c.PostForm("id")).MustInt64()
	email := com.StrTo(c.PostForm("email")).String()
	pwd := com.StrTo(c.PostForm("pwd")).String()
	newpwd := com.StrTo(c.PostForm("newpwd")).String()
	tel := e.Trim(com.StrTo(c.PostForm("tel")).String())
	act := e.Trim(com.StrTo(c.PostForm("act")).String())

	valid := validation.Validation{}
	valid.Required(id, "id").Message("操作失败")
	valid.Min(id, 0, "id").Message("操作失败")
	fmt.Println("lalla:", id)
	if act == "user" {
		valid.Required(email, "email").Message("邮箱不能为空")
		valid.Required(tel, "tel").Message("手机号码不能为空")
	} else {
		valid.Required(pwd, "pwd").Message("原始密码不能为空")
		valid.Required(newpwd, "newpwd").Message("新密码不能为空")
	}

	data := make(map[string]interface{})
	if !valid.HasErrors() {
		user := Admin.Find(id)
		if act == "user" {
			reg := regexp.MustCompile(`^1{1}\d{10}$`)
			if !reg.MatchString(tel) || len(tel) < 11 {
				return e.ERROR, "手机号码格式不正确"
			}

			data["tel"] = tel
			data["email"] = email

			if user.Tel != tel {
				validTel := Admin.ExistsByTel(tel)
				if validTel {
					return e.ERROR, "手机号码已存在，填写新的手机号码"
				}
			}
		} else {
			if user.Pwd != Admin.Md5Pwd(pwd) {
				return e.ERROR, "原始密码不正确"
			}
			data["pwd"] = Admin.Md5Pwd(newpwd)
		}
		isOk := Admin.EditUser(id, data)
		msg := ""
		if isOk {
			if act == "user" {
				SaveUserInfo(int(id))
				msg = "用户信息"
			} else {
				if LogOut(c) {
					c.Header("Cache-Control", "no-cache,no-store")
				}
				msg = "修改密码"
			}
			return e.SUCCESS, msg
		}
		return e.ERROR, "操作失败"
	}
	return ViewErr(valid)
}

// @Summer 获取当个用户信息
func GetUser(c *gin.Context) (code int, err string, con interface{}) {
	c.Request.Body = e.GetBody(c)
	id := com.StrTo(c.PostForm("id")).MustInt64()
	var data interface{}
	if id < 1 {
		return e.ERROR, "ID必须大于0", data
	}

	data = Admin.Find(id)
	return e.SUCCESS, "操作成功", data
}

// 管理员信息存储redis
func GetUserById(c *gin.Context) (con string) {
	getUUID, uOk := c.Request.Cookie("uuid")
	if uOk != nil {
		fmt.Println("没有获取到uuid")
		return
	}

	getUid := getUUID.Value
	if len(getUid) == 0 {
		fmt.Println("uuid不正确")
		return
	}

	uuid, _ := strconv.Atoi(getUid)
	if isOk, val := e.GetVal("user_" + getUid); isOk {
		fmt.Sprintf("读取缓存:'%v\n", val)
		return val
	}
	return SaveUserInfo(uuid)
}

// @Summer 解析错误原因
func ViewErr(valid validation.Validation) (code int, err string) {
	for _, err := range valid.Errors {
		return e.ERROR, err.Message
	}
	return e.SUCCESS, "操作成功"
}

//管理员登录存储缓存中
func SaveUserInfo(uuid int) (con string) {
	fmt.Println("设置用户缓存")
	userInfo := Admin.Find(int64(uuid))
	userResult, _ := json.Marshal(userInfo)
	con = string(userResult)
	e.SetVal("user_"+strconv.Itoa(uuid), con)
	return
}

// @Summer退出
func LogOut(c *gin.Context) bool {
	isOk, token := e.GetVal("token")

	if isOk {
		isOK, uuid := e.GetUUID(c)
		if isOK {
			e.SetCookie(c, uuid, -1)
		}
		fmt.Println("shanctoken")
		e.DelVal("token")
		if uuid > 0 {
			e.DelVal("user_" + strconv.Itoa(uuid))
		}
		return true
	}
	fmt.Println("token:", token)
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
		admins = Admin.Find(int64(uid))
	}
	return nil, admins
}
func GetLastUser() (id int) {
	return Admin.GetLastUserId().ID
}
