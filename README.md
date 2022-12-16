

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

## 用户请求（即将修改，先不动）



### 1./user_login（登录请求）

#### 前端：

type:

| 名称 | 数据类型 | 说明             |
| ---- | -------- | ---------------- |
| 种类 | string   | 功能待定，可不传 |

data:

| 名称 | 数据类型 | 说明 |
| ---- | -------- | ---- |
| 账号 | String   |      |
| 密码 | int      |      |

```json
{
    "type":"post",
    "data":[
        {
            "user_name":"admin",
            "user_password":123456,
        }
    ]
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
    "data":[
        {
            "user_id":90001,
            "user_name":"admin",
            "user_password":"123456",
            "user_level":1,
            "msg":"登录成功",
        }
    ]
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
    "type":"post",
    "data":[
        {
            "user_name":"admin",
            "user_password":123456,
            "user_level":1,
        }
    ]
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
    "data":[
        {
            "user_id":90001,
            "user_name":"admin",
            "user_password":"123456",
            "user_level":1,
            "msg":"登录成功",
        }
    ]
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
    "type":"post",
    "data":[
        {
            "user_name":"admin",
            "user_password":123456,
            "user_newassword":123456,
            "user_level":1,
        }
    ]
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
    "data":[
        {
            "user_id":90001,
            "user_name":"admin",
            "user_password":123456,
            "user_level":1,
            "msg":"注册成功",
        }
    ]
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
    "type":"post",
    "data":[
        {
            "user_name":"admin",
            "user_password":123456,
            "user_changePassword":234567,
            "user_level":1,
        }
    ]
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
    "data":[
        {
            "user_id":90001,
            "user_name":"admin",
            "user_changePassword":234567,
            "user_level":1,
            "msg":"修改成功",
        }
    ]
}
```

### 

## 订单请求

### 1./order_show

#### 前端：

请求类型（get），无需json数据

#### 返回数据：

```json
{
  "code": 200,
  "data": [
    {
      "order_no": 1,
      "order_customerName": "杜詩涵",
      "order_address": "房山区岳琉路30号",
      "order_time": "2022-12-16 13:13:52",
      "order_age": 35,
      "order_gender": "女",
      "order_post": ""
    },
    {
      "order_no": 2,
      "order_customerName": "方宇宁",
      "order_address": "天河区大信商圈大信南路475号",
      "order_time": "2022-12-16 13:13:31",
      "order_age": 45,
      "order_gender": "女",
      "order_post": ""
    },
    {
      "order_no": 3,
      "order_customerName": "许詩涵",
      "order_address": "白云区小坪东路178号",
      "order_time": "2022-12-16 13:13:32",
      "order_age": 29,
      "order_gender": "男",
      "order_post": ""
    },
    {
      "order_no": 4,
      "order_customerName": "邹岚",
      "order_address": "东泰五街154号",
      "order_time": "2022-12-16 13:13:33",
      "order_age": 27,
      "order_gender": "女",
      "order_post": ""
    },
    {
      "order_no": 5,
      "order_customerName": "袁云熙",
      "order_address": "东城区东单王府井东街353号",
      "order_time": "2022-12-16 13:13:33",
      "order_age": 30,
      "order_gender": "女",
      "order_post": ""
    },
    {
      "order_no": 6,
      "order_customerName": "于杰宏",
      "order_address": "锦江区人民南路四段184号",
      "order_time": "2022-12-16 13:13:34",
      "order_age": 27,
      "order_gender": "女",
      "order_post": ""
    },
    {
      "order_no": 7,
      "order_customerName": "韦云熙",
      "order_address": "罗湖区蔡屋围深南东路517号",
      "order_time": "2022-12-16 13:13:34",
      "order_age": 30,
      "order_gender": "女",
      "order_post": ""
    },
    {
      "order_no": 8,
      "order_customerName": "阎詩涵",
      "order_address": "天河区大信商圈大信南路329号",
      "order_time": "2022-12-16 13:13:35",
      "order_age": 37,
      "order_gender": "女",
      "order_post": ""
    },
    {
      "order_no": 9,
      "order_customerName": "王璐",
      "order_address": "浦东新区健祥路57号",
      "order_time": "2022-12-16 13:13:35",
      "order_age": 28,
      "order_gender": "男",
      "order_post": ""
    },
    {
      "order_no": 10,
      "order_customerName": "尹杰宏",
      "order_address": "房山区岳琉路220号",
      "order_time": "2022-12-16 13:13:36",
      "order_age": 18,
      "order_gender": "女",
      "order_post": ""
    },
    {
      "order_no": 11,
      "order_customerName": "夏晓明",
      "order_address": "延庆区028县道717号",
      "order_time": "2022-12-16 13:13:36",
      "order_age": 44,
      "order_gender": "女",
      "order_post": ""
    },
    {
      "order_no": 12,
      "order_customerName": "杜詩涵",
      "order_address": "房山区岳琉路30号",
      "order_time": "2022-12-16 13:13:37",
      "order_age": 35,
      "order_gender": "女",
      "order_post": ""
    },
    {
      "order_no": 13,
      "order_customerName": "孔宇宁",
      "order_address": "天河区天河路520号",
      "order_time": "2022-12-16 13:13:37",
      "order_age": 33,
      "order_gender": "男",
      "order_post": ""
    },
    {
      "order_no": 14,
      "order_customerName": "范云熙",
      "order_address": "成华区双庆路602号",
      "order_time": "2022-12-16 13:13:38",
      "order_age": 42,
      "order_gender": "女",
      "order_post": ""
    },
    {
      "order_no": 15,
      "order_customerName": "潘致远",
      "order_address": "黄浦区淮海中路364号",
      "order_time": "2022-12-16 13:14:13",
      "order_age": 44,
      "order_gender": "女",
      "order_post": ""
    }
  ]
}
```

## 员工请求

### 1./employee_chancePost

#### 前端：

post_id:

| 名称 | 数据类型 | 说明             |
| ---- | -------- | ---------------- |
| 种类 | string   | 功能待定，可不传 |

data:

| 名称     | 数据类型 | 说明          | 备注                                     |
| -------- | -------- | ------------- | ---------------------------------------- |
| 员工职位 | string   | employee_post | 只能以下四选一<br>经理，主厨，收银，职员 |
| 员工编号 | int      | employee_no   | 2                                        |

```json
{
    "post_id": 1,
    "data": [
        {
            "employee_post":"主厨"
            "employee_no": 1
        }
    ]
}
```

#### 返回数据

```json
{
  "code": 200,
  "data": [
    {
      "dish_name": "主厨",
      "dish_no": 1
    }
  ],
  "msg": "修改成功"
}
```

### 2./employee_delete

#### 前端：

| 名称     | 数据类型 | 说明        |
| -------- | -------- | ----------- |
| 员工编号 | int      | employee_no |

```json
{
    "post_id": 1,
    "data": [
        {
            "employee_no":2
        }
    ]
}
```

#### 返回数据:

```json
{
  "code": 200,
  "msg": "删除成功"
}
```



### 3./employee_add

#### 前端：

post_id:

| 名称 | 数据类型 | 说明             |
| ---- | -------- | ---------------- |
| 种类 | string   | 功能待定，可不传 |

data:

| 名称           | 数据类型 | 说明 | 备注 |
| :------------- | -------- | ---- | ---- |
| 太多了，先如图 |          |      |      |

```json
{
    "post_id": 1,
    "data": [
        {
            "employee_name":"王大猛",
            "employee_age": 30,
            "employee_gender":"男",
            "employee_enterTime":"2022-12-03",
            "employee_address":"福建省福州市马尾区阳光学院",
            "employee_post":"主厨",
            "employee_salary":2000.00,
            "employee_power":2
        }
    ]
}
```

#### 返回数据：（bug：返回的id号不是更新的id）

```json
{
  "code": 422,
  "data": {
    "post_id": "",
    "data": [
      {
        "employee_no": 0,
        "employee_name": "王大猛",
        "employee_age": 30,
        "employee_gender": "男",
        "employee_enterTime": "2022-12-03",
        "employee_address": "福建省福州市马尾区阳光学院",
        "employee_post": "主厨",
        "employee_salary": 2000,
        "employee_power": 2
      }
    ]
  },
  "msg": "新增成功"
}
```



## 菜单请求：

### 1./Dish_show

#### 前端：

发送get请求

#### 返回数据：

```json
{
  "code": 200,
  "data": [
    {
      "dish_no": 1,
      "dish_name": "西红柿炒鸡蛋",
      "dish_price": 15,
      "dish_taste": "甜"
    },
    {
      "dish_no": 2,
      "dish_name": "荔枝肉",
      "dish_price": 20,
      "dish_taste": "酸"
    },
    {
      "dish_no": 3,
      "dish_name": "肉末茄子",
      "dish_price": 18,
      "dish_taste": "咸"
    },
    {
      "dish_no": 4,
      "dish_name": "麻婆豆腐",
      "dish_price": 25,
      "dish_taste": "辣"
    }
  ]
}
```

### 2./dish_chanceName

#### 前端：

post_id:

| 名称 | 数据类型 | 说明             |
| ---- | -------- | ---------------- |
| 种类 | string   | 功能待定，可不传 |

data:

| 名称     | 数据类型 | 说明    |
| -------- | -------- | ------- |
| 菜单编号 | int      | dish_no |

```json
{
    "post_id": 1,
    "data": [
        {
            "dish_no": 1,
            "dish_name": "大西瓜"
        }
    ]
}
```

#### 返回数据：

```json
{
  "code": 200,
  "data": [
    {
      "dish_name": "大西瓜",
      "dish_no": 1
    }
  ],
  "msg": "修改成功"
}
```

### 3./dish_delete（有bug）

#### 前端：

```json
```

### 4./dish_add

#### 前端：

post_id:

| 名称 | 数据类型 | 说明             |
| ---- | -------- | ---------------- |
| 种类 | string   | 功能待定，可不传 |

data:

| 名称     | 数据类型 | 说明       |
| -------- | -------- | ---------- |
| 菜单名称 | string   | dish_name  |
| 菜单价格 | int      | dish_price |
| 菜单味道 | string   | dish_taste |

```json
{
    "post_id": 1,
    "data": [
        {
            "dish_name": "大西瓜",
            "dish_price": 20,
            "dish_taste": "甜"
        }
    ]
}
```

#### 返回数据：

```json
{
  "code": 422,
  "data": [
    {
      "dish_name": "大西瓜",
      "dish_price": 20,
      "dish_taste": "甜"
    }
  ],
  "msg": "新增成功"
}
```

### 
