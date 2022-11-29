

# gin_catering_backEnd 说明文档

## 默认配置

| 配置内容 | 参数 | 说明 |
| -------- | ---- | ---- |
| ip地址   | 未定 | 待定 |
| 端口号   | 8401 | 暂定 |

## 文件目录

| 包名       | 功能       | 说明 |
| ---------- | ---------- | ---- |
| common     | 数据库连接 |      |
| controller | 请求配置   |      |
| model      | 数据模型   |      |
| util       | 工具       |      |

# 请求

## 用户请求



### 1./user_login（登录请求）

#### 前端：

请求类型（GET）

| 名称 | 数据类型 | 说明 |
| ---- | -------- | ---- |
| 账号 | String   |      |
| 密码 | int      |      |

```json
{
    "user_name":"admin",
    "user_password":123456,
}
```

#### 数据库：

查询是否拥有该管理员

| 管理员编号 | 账号 | 密码 | 授权等级 |
| ---------- | ---- | ---- | -------- |
|            |      |      |          |

#### 返回数据：

```json
{
    "code":200,
    "user_level":1,
    "msg":"登录成功",
}
```



### 2./user_register

#### 前端：

请求类型（POST）

| 名称     | 数据类型 | 说明          |
| -------- | -------- | ------------- |
| 账号     | String   |               |
| 密码     | int      |               |
| 权限等级 | int      | 暂定为int类型 |

```json
{
    "user_name":"admin",
    "user_password":123456,
    "user_level":1,
}
```

#### 数据库：

插入数据

| 管理员编号 | 账号 | 密码 | 授权等级 |
| ---------- | ---- | ---- | -------- |
|            |      |      |          |

#### 返回数据：

```json
{
    "code":200,
    "user_name":"admin",
    "user_password":123456,
    "user_level":1,
    "msg":"注册成功",
}
```



### 3./user_changePassword

#### 前端：

请求类型（POST）

| 名称     | 数据类型 | 说明          |
| -------- | -------- | ------------- |
| 账号     | String   |               |
| 密码     | int      |               |
| 新密码   | int      | 修改的新密码  |
| 权限等级 | int      | 暂定为int类型 |

```json
{
    "user_name":"admin",
    "user_password":123456,
    "user_level":1,
}
```

#### 数据库：

修改密码数据

| 管理员编号 | 账号 | 密码 | 授权等级 |
| ---------- | ---- | ---- | -------- |
|            |      |      |          |

#### 返回数据：

```json
{
    "code":200,
    "user_name":"admin",
    "user_password":123456,
    "user_level":1,
    "msg":"注册成功",
}
```

### 

### 4./user_delete（可舍弃）

#### 前端：

请求类型（POST）

| 名称     | 数据类型 | 说明          |
| -------- | -------- | ------------- |
| 账号     | String   |               |
| 密码     | int      |               |
| 权限等级 | int      | 暂定为int类型 |

```json
{
    "user_name":"admin",
    "user_password":123456,
    "user_changePassword":234567,
    "user_level":1,
}
```

#### 数据库：

删除数据

| 管理员编号 | 账号 | 密码 | 授权等级 |
| ---------- | ---- | ---- | -------- |
|            |      |      |          |

#### 返回数据：

```json
{
    "code":200,
    "user_name":"admin",
    "user_changePassword":234567,
    "user_level":1,
    "msg":"修改成功",
}
```



## 订单请求

### 1./show_order

#### 前端：

请求类型（POST）

| 名称     | 数据类型 | 说明 |
| -------- | -------- | ---- |
| 权限等级 | int      | 暂定 |

```json
{
    "user_level":1,
}
```

#### 数据库：

插入数据

| 管理员编号 | 账号 | 密码 | 授权等级 |
| ---------- | ---- | ---- | -------- |
|            |      |      |          |

#### 返回数据：

```json
{
    {
    "order_number":40001,
    "customer_number":90001,
    "customer_name":"李一",
    "customer_age":18,
    "customer_sex":"男",
    "customer_contact":"12345678910",
    "customer_address":"福建省福州市马尾区登龙路阳光学院"
    "menu_number":30001,
    "order_time":"2020.1.1",
    "order_price":200.00
    "order_rating":"好评"
	}
,
	{
    "order_number":40002,
    "customer_number":90002,
    "customer_name":"李二",
    "customer_age":18,
    "customer_sex":"男",
    "customer_contact":"12345678910",
    "customer_address":"福建省福州市马尾区登龙路阳光学院"
    "menu_number":30002,
    "order_time":"2020.1.1",
    "order_price":300.00
    "order_rating":"差评"
	}
}
```

