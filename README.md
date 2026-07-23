## วิธี run
```
> go run cmd/api/main.go
```
## API มี 2 version
```
path v1: base-url/api/v1
path v2: base-url/api/v2
```
## ทดสอสเบื้องต้น v1
***Request Body:***
```
curl -i -X GET http://localhost:5000/api/v1/test
```
***Response:***
```
{
    "data": "[{}]",
    "message": "success",
    "status": true
}
```
## ทดสอสเบื้องต้น v2
***Request Body:***
```
curl -i -X GET http://localhost:5000/api/v2/test
```
***Response:***
```
{
    "message": "success",
    "status": true,
    "time": "<time>"
}
```
## ข้อมูล
```
- โปรแกรม version 1 อยู่ "นอก" โฟลเดอร์ internal
- โปรแกรม version 2 อยู่ "ใน" โฟลเดอร์ internal
```
# base-url/api/v1
## ข้อมูล Facebook page
***Request Body:***
```
curl -X GET http://localhost:5000/api/v1/get-info
```
***Response:***
```
{
  "status": true,
  "message": "success",
  "data": {
    "name": "<facebook_page_name>",
    "id": "<facebook_page_id>"
  }
}
```
## ดูรายการ conversation id ของเพจ
***Request Body:***
```
curl -i -X GET http://localhost:5000/api/v1/get-conversations
```
***Response:***
```
{
  "status": true,
  "message": "success",
  "data": [
        {
          "id": "conversation_id",
          "participants": {
            "data": [
              {
                "email": "<id>@facebook.com",
                "id": "facebook_id",
                "name": "<customer>"
              },
              {
                "email": "<id>@facebook.com",
                "id": "facebook_id",
                "name": "<facebook_page>"
              }
            ]
          },
          "updated_time": "<updated_time>"
        }
        ...
        {
          "id": "conversation_id",
          ....
        }
      ]
}
```
## ดูเนื้อหาแชท จาก conversation id
***Request Body:***
```
curl -i -X GET "http://localhost:5000/api/v1/get-message/<conversation_id>
```
***Response:***
```
{
  "status": true,
  "message": "success",
  "data": [
      {
        "attachments": {
          "data": [
            {
              "id": "<facebook id>",
              "image_data": {
                "height": "<number>",
                "image_type": "<number>",
                "max_height": "<number>",
                "max_width": "<number>",
                "preview_url": "preview_url",
                "render_as_sticker": "boolean: true, false",
                "url": "<image_path>",
                "width": "<number>"
              },
              "mime_type": "image/jpeg",
              "name": "<image_name>",
              "size": "<image_size>"
            }
          ],
          "paging": {
            "cursors": {
              "after": "MAZDZD",
              "before": "MAZDZD"
            },
            "next": "<api หน้าถัดไปของข้อมูล>"
          }
        },
        "created_time": "<time>",
        "from": {
          "email": "<facebook id>@facebook.com",
          "id": "<facebook id>",
          "name": "ชื่อเพจ (ผู้ส่งข้อความ)"
        },
        "id": "message_id",
        "message": "<ข้อความ>",
        "to": {
          "data": [
            {
              "email": "<facebook id>@facebook.com",
              "id": "<facebook id>",
              "name": "<ชื่อ facebook (ผู้รับข้อความ)>"
            }
          ]
        }
      },
    ]
}

```
## ส่งข้อความถึงลูกค้าของเพจ
***Request Body:***
```
curl -X POST http://localhost:5000/api/v1/send-message \
  -H "Content-Type: application/json" \
  -d '{
        "recipientId":"<facebook_id>",
        "messageText": "<ข้อความ>",
        "mediaType": "<ส่งข้อความไม่ใส่ก็ได้>"
    }'
```
***Response:***
```
{
  "status": true,
  "message": "success",
  "data": {
    "recipient_id": <facebook_id>,
    "message_id": "<message_id>"
  }
}
```
## ส่งรูปภาพถึงลูกค้าของเพจ
***Request Body:***
```
curl -X POST "http://localhost:5000/api/v1/send-message" \
  -F "mediaType=<ประเภทของข้อมูลที่จะส่ง เช่น image, video, message>" \
  -F "recipientId=<facebook_id>" \
  -F "filedata=@/Users/<username>/<Pictures>/<beach.png>"
```
***Response:***
```
{
  "status": true,
  "message": "success",
  "data": {
    "recipientId": <facebook_id>,
    "messageId": "<message_id>",
    "attachmentId": "<attachmentId>"
  }
}
```
# base-url/api/v2
## หมายเหตุ
- แอปพลิเคชันยังอยู่ในโหมดพัฒนา (Development Mode) จะไม่สามารถเผยแพร่เป็น Public ได้ เช่น การโพสต์ Content
## สร้างโพสต์ในเพจ
***Request Body:***
```
curl -X POST "http://localhost:5000/api/v2/feed" \
     -H "Content-Type: application/json" \
     -d '{
           "message":"<your_message_text>",
           "published":<boolean>,
           "privacy": {
            "value": "EVERYONE"
           }
         }'
```
***Response:***
```
{
  "id": "<feed_id>"
}
```
##  ดูรายการโพสต์ในเพจ
***Request Body:***
```
curl -X PATCH "http://localhost:5000/api/v2/feed/<feed_id>" \
     -H "Content-Type: application/json" \
     -d '{
           "message":"<your_message_text>"
         }'
```
***Response:***
```
{
  "status": true,
  "message": "success",
  "data": {
    "Success": <boolean>
  }
}
```
##  แก้ไขโพสต์ในเพจ
***Request Body:***
```
curl -i -X PATCH "http://localhost:5000/api/v2/feed/<feed_id>"
```
***Response:***
```
{
  "status": true,
  "message": "success",
  "data": {
    "data": [
      {
        "created_time": "2020-03-25T17:33:34+0000",
        "message": "Hello World!",
        "id": "<content_id>"
      },
    ...
  ]
  }
}
```

## สิทธิ์การอนุญาตและฟีเจอร์ของ API Facebook pages
```
pages_show_list – แสดงเพจที่ผู้ใช้จัดการ
pages_read_engagement – อ่านเนื้อหาที่โพสต์ลงในเพจ
pages_manage_posts – เผยแพร่และกำหนดเวลาเนื้อหา
pages_manage_engagement – ควบคุมความคิดเห็น ลบโพสต์
pages_read_user_content – อ่านเนื้อหาที่ผู้ใช้สร้างบนเพจ
pages_manage_metadata – จัดการการตั้งค่าเพจ
pages_manage_ads – สร้างและจัดการโฆษณาสำหรับเพจ
pages_manage_cta – ดูและอัพเดตปุ่มกระตุ้นให้ดำเนินการ
pages_messaging – จัดการและส่งข้อความในนามของเพจ
business_management – จัดการสินทรัพย์ธุรกิจที่เกี่ยวข้องกับเพจ
```

# Directory Tree
```
├───cmd
│   └───api
├───configs
├───controllers
│   └───facebook-controller
├───internal
│   ├───app
│   │   └───comment_app
│   ├───common
│   │   └───config
│   ├───infrastructure
│   │   └───server
│   ├───module
│   │   ├───comment_facebook
│   │   │   ├───adapter
│   │   │   │   ├───comment_facebook_adapter
│   │   │   │   └───data_transfer_object
│   │   │   ├───delivery
│   │   │   │   └───http
│   │   │   ├───domain
│   │   │   └───usecase
│   │   ├───conversation_facebook
│   │   │   ├───adapter
│   │   │   │   ├───conversation_facebook_adapter
│   │   │   │   └───data_transfer_object
│   │   │   ├───delivery
│   │   │   │   └───http
│   │   │   ├───domain
│   │   │   └───usecase
│   │   ├───feed_facebook
│   │   │   ├───adapter
│   │   │   │   ├───data_transfer_object
│   │   │   │   └───feed_facebook_adapter.go
│   │   │   ├───delivery
│   │   │   │   └───http
│   │   │   ├───domain
│   │   │   └───usecase
│   │   └───infrastructure
│   │       └───response
│   │           └───http
│   ├───public
│   │   └───asset
│   │       └───image
│   └───router
│       └───v2
├───middlewares
├───models
│   └───facebook-model
├───public
│   └───assets
│       └───images
├───routers
├───services
│   └───facebook-service
└───test
```