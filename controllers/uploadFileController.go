package controllers

import (
	"DataCertPlatform/models"
	"crypto/md5"
	sha2562 "crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/astaxie/beego"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

type UploadFileController struct {
	beego.Controller
}

/*
该post方法用于处理用户在客户端提交的文件
 */
func (u *UploadFileController) Post() {

	phone := u.Ctx.Request.PostFormValue("phone")
	title := u.Ctx.Request.PostFormValue("upload_title")
	fmt.Println("电子数据标签：",title)
	//用户上传的文件
	file,header,err := u.GetFile("ouyang")
	//defer file.Close() //如果u.GetFile有错误，file里面就没有值，这时候会报空指针错误：invaild memory or nil pointer dere...,所以放到下面
	if err != nil {//解析客户端提交的文件出现错误
		u.Ctx.WriteString("抱歉，文件解析失败，请重试")
		return
	}

	defer file.Close()
	//使用io包提供的方法保存文件
	saveFilePath := "static/upload/" + header.Filename
	saveFile,err := os.OpenFile(saveFilePath,os.O_CREATE|os.O_RDWR,777)
	if err !=nil {
		u.Ctx.WriteString("抱歉，电子数据认证失败，请重试")
		return
	}

	_,err = io.Copy(saveFile,file)
	if err != nil {
		u.Ctx.WriteString("抱歉，电子数据又认证失败，请重新尝试一下吧")
		return
	}

	//2、计算文件的SHA256值
	hash256 := sha2562.New()
	fileBytes,_ := ioutil.ReadAll(file)
	hash256.Write(fileBytes)
	hashBytes := hash256.Sum(nil)
	fmt.Println("存证的hash值为：",hex.EncodeToString(hashBytes))

	//先查询id
	user1,err := models.User{Phone:phone}.QueryUserByPhone()
	if err != nil {
		u.Ctx.WriteString("抱歉，电子数据认证失败，请等会儿再试")
	}

	//把上传的文件作为记录保存到数据库中
	//1、计算md5值
	md5Hash := md5.New()
	fileMd5Bytes,err := ioutil.ReadAll(saveFile)
	md5Hash.Write(fileMd5Bytes)
	bytes := md5Hash.Sum(nil)
	record := models.UploadRecord{
		UserId:   user1.Id,
		FileName: header.Filename,
		FileSize: header.Size,
		FileCert:hex.EncodeToString(bytes),
		FileTitle: title,
		CertTime: time.Now().Unix(),
	}
	//2、保存认证数据到数据库中
	_,err = record.SaveRedcord()
	if err != nil {
		u.Ctx.WriteString("抱歉，电子数据认证保存失败，请重视")
		fmt.Println("重视错误：",err.Error())
		return
	}

	//上传文件保存到数据库成功
	records,err := models.QueryRecordsUserId(user1.Id)
	if err != nil {
		u.Ctx.WriteString("抱歉，电子数据列表获取失败，请重试")
		return
	}

	u.Data["records"] = records
	u.TplName = "list_recorc.html"

}

/**
该post方法用于处理用户在客户端提交的认证文件
*/
func (u *UploadFileController) Post1() {
	//1、解析用户上传的数据
	//用户上传的自定义的标题
	title := u.Ctx.Request.PostFormValue("upload_title")

	//用户上传的文件
	file,header,err := u.GetFile("ouyang")
	defer file.Close()
	if err != nil {//解析客户端提交的文件出现错误
		u.Ctx.WriteString("抱歉，文件解析失败，请重试")
		return
	}
	fmt.Println("自定义的标题：",title)

	//限制文件类型
	//字符串切割,返回一个切片类型
	fileNameSlice := strings.Split(header.Filename,".")
	fileType := fileNameSlice[1]
	if fileType != "jpg" && fileType != "png" {
		u.Ctx.WriteString("抱歉，数据类型不正确，请重兴上传")
		return
	}

	//限制文件大小,
	//因为这个为硬代码，所以限制的大小可以写在app.conf中
	config := beego.AppConfig
	filesize,err := config.Int64("file_size")
	if err != nil {
		return
	}
	if header.Size / 1024 > filesize {
		u.Ctx.WriteString("抱歉，上传的文件过大，请重新尝试")
	}

	//到此，获得到了上传的文件

	fmt.Println("上传的文件名称：",header.Filename)
	fmt.Println("上传的文件大小：",header.Size)//返回的字节大小

	//perm:permission 权限
	//权限的组成：a+b+c
		//a:文件所有者对文件的操作权限   读4、写2、执行1
		//b：文件所有者所在组的用户的操作权限   读、写、执行
		//c：其他用户的操作权限	读、写、执行
		//eg: m文件，权限是651，651分别代表a，b，c
	saveDir := "static/upload"
	//1、先尝试打开文件
	_,err = os.Open(saveDir)
	if err != nil {
		//2、创建文件夹
		err = os.Mkdir(saveDir,777)
		if err != nil {
			fmt.Println(err.Error())
			u.Ctx.WriteString("抱歉，文件认证错误，请重新尝试")
			return
		}

	}
	//fmt.Println("打开的文件夹:",f.Name())

	saveName := saveDir + "/" + header.Filename
	fmt.Println("要保存的文件名：",saveName)
	//fromFile:文件
	//toFile：要保存的文件
	err = u.SaveToFile("ouyang",saveName)
	if err != nil {
		fmt.Println(err.Error())
		u.Ctx.WriteString("抱歉，文件认证失败，请重试...")
		return
	}

	fmt.Println("上传的文件:",file)
	u.Ctx.WriteString("成功")

}