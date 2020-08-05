# Post Service

The posts service stores posts

## Usage

### Create a post

```
micro call go.micro.service.posts Posts.Save '{"post":{"id":"1","title":"How to Micro","content":"Simply put, Micro is awesome."}}'
micro call go.micro.service.posts Posts.Save '{"post":{"id":"2","title":"Fresh posts are fresh","content":"This post is fresher than the How to Micro one"}}'
```

### Query posts

```
micro call go.micro.service.posts Posts.Query '{}'
micro call go.micro.service.posts Posts.Query '{"slug":"how-to-micro"}'
micro call go.micro.service.posts Posts.Query '{"offset": 10, "limit": 10}'
```

### Delete posts

```
micro call go.micro.service.posts Posts.Delete '{"offset": 10, "limit": 10}'
```
