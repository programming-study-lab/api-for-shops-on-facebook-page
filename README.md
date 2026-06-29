

## ข้อมูล Facebook page
***Request Body:***
```
curl -X POST http://localhost:5000/api/v1/send-message \
  -H "Content-Type: application/json" \
  -d '{
        "recipient":{
            "id": "<facebook_id>"
        },
        "messaging_type": "RESPONSE",
        "message":{
            "text": "Hello World"
        }
    }'
```
***Response:***
```
{
    "name": "<facebook_page_name>",
    "id": "<facebook_page_id>"
}
```
## ดูรายการ conversion id ของเพจ
***Request Body:***
```
curl -i -X GET "http://localhost:5000/api/v1/conversations
    ?platform=messages
```
***Response:***
```
{
  "messages": {
    "data": [
      {
        "id": "Message ID-1",
        "created_time": "UNIX-TIMESTAMP-MOST-RECENT-MESSAGE"
      },
      {
        "id": "Message ID-2",
        "created_time": "UNIX-TIMESTAMP"
      },
      {
        "id": "Message ID-3",
        "created_time": "UNIX-TIMESTAMP"
      },
...
    ]
  },
  "id": "Conversation ID",
}
```
## ดูเนื้อหาแชท จาก conversion id
***Request Body:***
```
curl -i -X GET "http://localhost:5000/api/v1/<conversion_id>/message
    ?fields=message,from,created_time
```
***Response:***
```
{
  "data": [
    {
      "message": "ราคาเท่าไหร่ครับ", // ข้อความที่ลูกค้าพิมพ์มา
      "from": {
        "name": "ชื่อลูกค้า",
        "id": "<PSID_ของลูกค้า>"
      },
      "created_time": "2026-06-28T10:30:00+0000",
      "id": "m_mid.14586453..."
    },
    {
      "message": "สวัสดีครับ ยินดีต้อนรับครับ", // ข้อความที่เพจตอบกลับ
      "from": {
        "name": "ชื่อเพจของคุณ",
        "id": "<PAGE_ID>"
      },
      "created_time": "2026-06-28T10:31:00+0000",
      "id": "m_mid.14586454..."
    }
  ]
}

```
## ส่งข้อความถึงลูกค้าของเพจ
***Request Body:***
```
curl -X POST http://localhost:5000/api/v1/send-message \
  -H "Content-Type: application/json" \
  -d '{
        "recipient":{
            "id": "<facebook_id>"
        },
        "messaging_type": "RESPONSE",
        "message":{
            "text": "Hello World"
        }
    }'
```
***Response:***
```
{
    "recipient_id": <facebook_id>,
    "message_id": "<message_id>"
}
```
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