package dao

import (
	"TikTok/config"
	"github.com/dutchcoders/goftp"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var Db *gorm.DB

func Init() {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second,  // 慢 SQL 阈值
			LogLevel:      logger.Error, // Log level
			Colorful:      true,         // 彩色打印
		},
	)
	var err error
	dsn := "douyin:zjqxy@tcp(43.138.25.60:3306)/douyin?charset=utf8mb4&parseTime=True&loc=Local"
	//想要正确的处理time.Time,需要带上 parseTime 参数，
	//要支持完整的UTF-8编码，需要将 charset=utf8 更改为 charset=utf8mb4
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		log.Panicln("err:", err.Error())
	}

}

var MyFTP *goftp.FTP

func InitFTP() {
	//获取到ftp的链接
	var err error
	MyFTP, err = goftp.Connect(config.ConConfig)
	if err != nil {
		log.Printf("获取到FTP链接失败！！！")
	}
	log.Printf("获取到FTP链接成功%v：", MyFTP)
	//登录
	err = MyFTP.Login(config.FtpUser, config.FtpPsw)
	if err != nil {
		log.Printf("FTP登录失败！！！")
	}
	log.Printf("FTP登录成功！！！")
}
