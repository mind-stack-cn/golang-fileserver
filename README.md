# GoLangFileServer

### build
docker build -t oceanwu/golang-fileserver .

### Usage
docker run -d -p 8088:8088 -v $(pwd)/data:/go/src/github.com/mind-stack-cn/golang-fileserver/data oceanwu/golang-fileserver

## test page
http://127.0.0.1:8088/test

## upload File
```
POST http://127.0.0.1:8088/?fileType=image

fileType=file|image|audio|video 
```

## Sample Return Json
```json
Image
{
    "header": {
        "code": 1000,
        "description": "success"
    },
    "data": [
        {
            "uri": "/828c8279/62ea/4829/bb16/0ad7bee0e5c6.png",
            "size": 141913,
            "fileType": "image",
            "width": 848,
            "height": 878
        },
        {
            "uri": "/97f9882f/81cf/4ad4/9bcc/b80bf2d4205f.png",
            "size": 171262,
            "fileType": "image",
            "width": 784,
            "height": 818
        }
    ]
}

Audio
{
    "header": {
        "code": 1000,
        "description": "success"
    },
    "data": [
        {
            "uri": "/0efc42f4/7d76/4771/9717/49ad46ac1f6a.mp3",
            "size": 1673125,
            "fileType": "audio",
            "duration": 192.339592
        }
    ]
}

Video
{
    "header": {
        "code": 1000,
        "description": "success"
    },
    "data": [
        {
            "uri": "/d6165916/2724/4822/ba83/689e6984583c.mp4",
            "size": 424407,
            "fileType": "video",
            "duration": 10.024,
            "thumbnail": {
                "uri": "/d6165916/2724/4822/ba83/689e6984583c.jpg",
                "size": 18109,
                "fileType": "image",
                "width": 640,
                "height": 360
            }
        }
    ]
}

common File
{
    "header": {
        "code": 1000,
        "description": "success"
    },
    "data": [
        {
            "uri": "/8cd1f889/fd43/49ba/8016/40f366327bc4.word",
            "size": 45,
            "fileType": "file"
        }
    ]
}
```

## Image Thumbnail
```
Query Parameter : width && height
example: http://127.0.0.1:8088/828c8279/62ea/4829/bb16/0ad7bee0e5c6.png?width=800&height=800
```

