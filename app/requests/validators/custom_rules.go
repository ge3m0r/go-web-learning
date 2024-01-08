package validators

import (
	"errors"
	"fmt"
	"golearning/pkg/database"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/thedevsaddam/govalidator"
)

func init(){
	govalidator.AddCustomRule("not_exists", func(field, rule, message string, value interface{}) error {
		rng := strings.Split(strings.TrimPrefix(rule, "not_exists:"),",")

		tableName := rng[0]
		dbFiled := rng[1]

		var exceptID string
		if len(rng) > 2{
			exceptID = rng[2]
		}

		requestValue := value.(string)

		query := database.DB.Table(tableName).Where(dbFiled+" =?", requestValue)
		if len(exceptID) > 0 {
            query.Where("id != ?", exceptID)
        }

        // 查询数据库
        var count int64
        query.Count(&count)

        // 验证不通过，数据库能找到对应的数据
        if count != 0 {
            // 如果有自定义错误消息的话
            if message != "" {
                return errors.New(message)
            }
            // 默认的错误消息
            return fmt.Errorf("%v 已被占用", requestValue)
        }
        // 验证通过
        return nil
	})

	govalidator.AddCustomRule("max_cn", func(field, rule, message string, value interface{}) error {
		valLength := utf8.RuneCountInString(value.(string))
		l, _ := strconv.Atoi(strings.TrimPrefix(rule, "max_cn"))

		if valLength > l{
			if message != ""{
				return errors.New(message)
			}
			return fmt.Errorf("长度不能超过 %d 个字", l)
		}
		return nil
	})

	govalidator.AddCustomRule("min_cn", func(field, rule, message string, value interface{}) error {
		valLength := utf8.RuneCountInString(value.(string))
		l, _ := strconv.Atoi(strings.TrimPrefix(rule, "min_cn"))

		if valLength < l{
			if message != ""{
				return errors.New(message)
			}
			return fmt.Errorf("长度需大于 %d 个字", l)
		}
		return nil
	})

	govalidator.AddCustomRule("exists", func(field string, rule string, message string, value interface{}) error {
        rng := strings.Split(strings.TrimPrefix(rule, "exists:"), ",")

        // 第一个参数，表名称，如 categories
        tableName := rng[0]
        // 第二个参数，字段名称，如 id
        dbFiled := rng[1]

        // 用户请求过来的数据
        requestValue := value.(string)

        // 查询数据库
        var count int64
        database.DB.Table(tableName).Where(dbFiled+" = ?", requestValue).Count(&count)
        // 验证不通过，数据不存在
        if count == 0 {
            // 如果有自定义错误消息的话，使用自定义消息
            if message != "" {
                return errors.New(message)
            }
            return fmt.Errorf("%v 不存在", requestValue)
        }
        return nil
    })
}