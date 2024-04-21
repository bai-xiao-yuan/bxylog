# BxyLog
## 支持功能
- 日志等级控制
- 输出到文件和控制台
- 日志颜色控制
- 自定义日志前缀
- 日志文件按条件切割
  - 日期
  - 文件大小
  - 时间间隔
- 是否携带日志时间  
以上功能皆可通过修改log对象的config配置数据即可实现。

## 日志等级控制
通过修改config对象的Level属性即可控制不同等级的日志数据输出：
日志等级有：
lInfo
lDebug
lWarn
lError
lPanic
五个等级
```
	l := NewLog(&conf.Config{
            Level: conf.LDebug,
          })
```
## 日志输出控制
默认仅输出到控制台，如果需要输出到文件，则需要将config的OutTarget配置为file,同时需要指定输出的路径文件名，如果路径不存在则会自动创建该文件夹：
```
	l := NewLog(&conf.Config{
		Level:     conf.LDebug,
		OutTarget: conf.File,
		FileName:  "log/t1.log",
	})
```

## 日志颜色控制
bxylog默认输出不带颜色，如果需要带颜色输出日志数据则需要将config的Color属性设置为true
```
	l := NewLog(&conf.Config{
		Level:     conf.LDebug,
		OutTarget: conf.File,
		FileName:  "log/t1.log",
		Color:     true,
	})
```

## 自定义前缀
设置Prefix字段，则日志数据会自动增加该前缀字符串
```
	l := NewLog(&conf.Config{
		Level:     conf.LDebug,
		OutTarget: conf.File,
		FileName:  "log/t1.log",
		Color:     true,
		Prefix:    "TestLogger: ",
	})
```

### 携带日志时间
设置timeFlag为true则日志数据会携带当前时间：
```
	l := NewLog(&conf.Config{
		Level:     conf.LDebug,
		OutTarget: conf.File,
		FileName:  "log/t1.log",
		Color:     true,
		Prefix:    "TestLogger: ",
		TimeFlag:  true,
	})
```

### 日志切割方式
日志切割可按日志文件大小（kb）,时间间隔（时），日期间隔(天)进行日志文件切割，需要同时设置
FileSliceType（日志文件切割方式），SliceProp(根据方式填写参数，对应切割方式的单位)
```
	l := NewLog(&conf.Config{
		Level:         conf.LDebug,
		OutTarget:     conf.File,
		FileName:      "log/t1.log",
		Color:         true,
		Prefix:        "TestLogger: ",
		TimeFlag:      true,
		FileSliceType: conf.FileSize,
		SliceProp:     1024,
	})
```
