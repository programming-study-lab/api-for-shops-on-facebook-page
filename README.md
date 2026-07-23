

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

# -------------------------------------------------------
# ไม่มีในระบบ
## หมายเหตุ
- แอปพลิเคชันยังอยู่ในโหมดพัฒนา (Development Mode) จะไม่สามารถเป็นแพร่เป็น Public ได้ เช่น การโพสต์ Content
## สร้างโพสต์ในเพจ
***Request Body:***
```
curl -X POST "http://localhost:5000/api/feed" \
     -H "Content-Type: application/json" \
     -d '{
           "message":"your_message_text",
           "access_token":"page_access_token",
         }'
```
***Response:***
```
{
  "id": "page_post_id"
}
```
##  ดูรายการโพสต์ในเพจ
***Request Body:***
```
curl -i -X GET "http://localhost:5000/api/feed"
```
***Response:***
```
{
  "data": [
    {
      "created_time": "2020-03-25T17:33:34+0000",
      "message": "Hello World!",
      "id": "<content_id>"
    },
...
  ]
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