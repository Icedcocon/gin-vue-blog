package utils

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"log"
	"math/rand"
	"time"

	"golang.org/x/crypto/bcrypt"
	"golang.org/x/crypto/scrypt"
)

type _encrypt struct{}

var Encryptor = new(_encrypt)

// 使用 scrypt 对密码进行加密生成一个哈希值
func (*_encrypt) ScryptHash(pasword string) string {
	const KeyLen = 10
	salt := []byte{12, 32, 4, 6, 66, 22, 222, 11} // 随便写

	// Key() 方法从密码、盐值和成本参数派生密钥，返回一个长度为 keyLen 的字节切片，可用作加密密钥。
	// 第3个参数    N 是 CPU/内存成本参数，必须是大于 1 的 2 的幂。
	// 第4、5个参数 r 和 p 必须满足 r * p < 2³⁰。不满足则函数返回一个空字节切片和一个错误。

	// 2017 年的交互式登录的推荐参数是 N=32768，r=8 和 p=1。
	// 随着内存延迟和 CPU 并行性的增加，应该增加参数 N、r 和 p；
	// 考虑将 N 设置为您可以在 100 毫秒内派生的最高 2 的幂。
	// 记住要获得一个好的随机盐值。
	hashPwd, err := scrypt.Key([]byte(pasword), salt, 1<<15, 8, 1, KeyLen)
	if err != nil {
		log.Fatal("加密失败: ", err)
	}

	// 对密码进行 base64 编码
	return base64.StdEncoding.EncodeToString(hashPwd)
}

// 使用 scrypt 对比 明文密码 和 数据库中哈希值
func (c *_encrypt) ScryptCheck(password, hash string) bool {
	return c.ScryptHash(password) == hash
}

// 使用 bcrypt 对密码进行加密生成一个哈希值
func (*_encrypt) BcryptHash(password string) string {
	// GenerateFromPassword函数返回给定成本的密码的bcrypt哈希值。
	// 如果给定的成本小于MinCost，则代价将被设置为DefaultCost。
	// 使用此包中定义的CompareHashAndPassword将返回的哈希密码与其明文版本进行比较。
	// GenerateFromPassword不接受超过72个字节的密码，这是bcrypt将要操作的最长密码。
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes)
}

// 使用 bcrypt 对比 明文密码 和 数据库中哈希值
func (*_encrypt) BcryptCheck(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// MD5 加密
func (*_encrypt) MD5(str string, b ...byte) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(b))
}

// 验证码
func (*_encrypt) ValidateCode() string {
	return fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
}

// TODO
func UUID() string {
	return ""
}
