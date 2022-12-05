package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"reflect"
	"strconv"
	"strings"
)

// ini配置文件解析器

// MysqlConfig MYSQL配置结构体
type MysqlConfig struct {
	Host     string `ini:"host"`
	Port     int    `ini:"port"`
	User     string `ini:"user"`
	Password string `ini:"password"`
	Test     bool   `ini:"test"`
}

// RedisConfig Redis配置结构体
type RedisConfig struct {
	Host     string `ini:"host"`
	Port     int    `ini:"port"`
	Password string `ini:"password"`
	Database int    `ini:"database"`
}

// Config 配置文件结构体
type Config struct {
	MysqlConfig `ini:"mysql"`
	RedisConfig `ini:"redis"`
}

func loadIni(fileName string, data interface{}) (err error) {
	// 0. 参数的校验
	// 0.1 传进来的data参数必须是指针类型(因为需要在函数中对其赋值)
	t := reflect.TypeOf(data)
	// fmt.Println(t, t.Kind())
	if t.Kind() != reflect.Ptr {
		err = errors.New("data param should be a pointer") // 新创建一个错误
		return
	}
	// 0.2 传进来的data参数必须是结构体类型指针（因为配置文件中有各种键值对需要赋值给结构体的字段）
	if t.Elem().Kind() != reflect.Struct {
		err = errors.New("data param should be a struct pointer") // 新创建一个错误
		return
	}

	// 1. 读文件得到字节类型数据
	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		return
	}
	// string(b)   // 将字节类型的文件内容转换为字符串
	lineSlice := strings.Split(string(b), "\n")
	// fmt.Printf("%#v\n", lineSlice) // []string{"# redis config", "[mysql]", "host=192.168.1.18", "port=3306", "user=root", "password=Admin@123", "", "# redis config", "[redis]", "host=127.0.0.1", "port=6379", "password=rootAd2", "database=0", "", ""}

	var structName string
	// 2. 一行一行的读数据
	for idx, line := range lineSlice {
		// 去掉字符串首位的空格
		line = strings.TrimSpace(line)
		// 如果是空行就跳过
		if len(line) == 0 {
			continue
		}
		// 2.1 如果是注释就忽略
		if strings.HasPrefix(line, ";") || strings.HasPrefix(line, "#") {
			continue
		}
		// 2.2 如果是[开头的就表示是节(section)
		if strings.HasPrefix(line, "[") {
			// 2.2.1 判断是否以[开头，并以]结尾
			if line[0] != '[' || line[len(line)-1] != ']' {
				err = fmt.Errorf("line:%d syntax error", idx+1)
				return
			}
			// 2.2.2 把这一行首尾的[]去掉，取到中间的内容把首尾空格去掉拿到内容
			sectionName := strings.TrimSpace(line[1 : len(line)-1])
			if len(sectionName) == 0 {
				err = fmt.Errorf("line:%d syntax error", idx+1)
				return
			}
			// 根据字符串sectionName去data里面根据反射找到对应的结构体
			for i := 0; i < t.Elem().NumField(); i++ {
				field := t.Elem().Field(i)
				if sectionName == field.Tag.Get("ini") {
					// 说明找到了对应的嵌套结构体，把字段名记下来
					structName = field.Name
					fmt.Printf("找到%s对应的嵌套结构体%s\n", sectionName, structName)
					// break
				}
			}

		} else {
			// 2.3 如果不是[开头的就是=分割的键值对
			// 2.3.1 以等号分割这一行，等号左边是key，等号右边是value
			if strings.Index(line, "=") == -1 || strings.HasPrefix(line, "=") {
				err = fmt.Errorf("line:%d syntax error", idx+1)
				return
			}
			index := strings.Index(line, "=")          // 拿到=的索引
			key := strings.TrimSpace(line[:index])     // 拿到key
			value := strings.TrimSpace(line[index+1:]) // 拿到value
			// 2.3.2 根据structName 去 data 里面把对应的嵌套结构体取出来
			v := reflect.ValueOf(data)
			sValue := v.Elem().FieldByName(structName) // 拿到嵌套结构体的值信息
			sType := sValue.Type()                     // 拿到嵌套结构体的类型信息

			if sType.Kind() != reflect.Struct {
				err = fmt.Errorf("data 中的%s应该是一个结构体", structName)
				return
			}
			var fieldName string
			var fileType reflect.StructField
			// 2.3.3 遍历嵌套结构体的每一个字段，判断tag是不是等于这个key
			for i := 0; i < sValue.NumField(); i++ {
				filed := sType.Field(i) // tag 信息是存储在类型信息中的
				fileType = filed
				if filed.Tag.Get("ini") == key {
					// 找到对应的字段
					fieldName = filed.Name
					break
				}
			}
			// 2.3.3 如果key = tag，给这个字段赋值
			// 2.3.3.1 根据fieldName去取出这个字段
			if len(fieldName) == 0 {
				// 在结构体中找不到对应的字段
				continue
			}
			fileObj := sValue.FieldByName(fieldName)

			// 2.3.3.2 对其赋值
			switch fileType.Type.Kind() {
			case reflect.String:
				fileObj.SetString(value)
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				var valueInt int64
				valueInt, err = strconv.ParseInt(value, 10, 64)
				if err != nil {
					err = fmt.Errorf("line:%d value type error", idx+1)
					return
				}
				fileObj.SetInt(valueInt)
			case reflect.Bool:
				var valueBool bool
				valueBool, err = strconv.ParseBool(value)
				if err != nil {
					err = fmt.Errorf("line:%d value type error", idx+1)
					return
				}
				fileObj.SetBool(valueBool)
			case reflect.Float32, reflect.Float64:
				var valueFloat float64
				valueFloat, err = strconv.ParseFloat(value, 64)
				if err != nil {
					err = fmt.Errorf("line:%d value type error", idx+1)
					return
				}
				fileObj.SetFloat(valueFloat)
			}
		}
	}
	return
}

func main() {
	// 声明一个Config类型的变量
	var cfg Config
	// 调用loadIni方法
	err := loadIni("./conf.ini", &cfg)
	if err != nil {
		fmt.Printf("load ini failed, err: %v\n", err)
		return
	}
	fmt.Println(cfg)
	fmt.Printf("%#v\n", cfg)
	fmt.Println(cfg.MysqlConfig.Host)     // 192.168.1.18
	fmt.Println(cfg.MysqlConfig.Port)     // 3306
	fmt.Println(cfg.MysqlConfig.User)     // root
	fmt.Println(cfg.MysqlConfig.Password) // Admin@123
}
