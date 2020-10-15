# Post Service

The posts service stores posts

## Usage

### Create a post

```sh
micro posts save --id=1 --title="How to Micro" --content="Simply put, Micro is awesome."
micro posts save --id=2 --title="Fresh posts are fresh" --content="This post is fresher than the How to Micro one"
```

### Create a post with tags

```sh
micro posts save --id=3 --title="How to do epic things with Micro" --content="Everything is awesome." --tags=a,b
# or
micro posts save --id=3 --title="How to do epic things with Micro" --content="Everything is awesome." --tags=a --tags=b
```

### Query posts

```sh
micro posts query
micro posts query --slug=how-to-micro
micro posts query --offset=10 --limit=10
```

### Delete posts

```
micro call posts Posts.Delete '{"offset": 10, "limit": 10}'
```
