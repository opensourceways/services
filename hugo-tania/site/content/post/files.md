---
title: files
servicename: files
labels: 
- Readme
---
# Files Service

Save and list text files by project and path.
## cURL


### Files List
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->
List files by their project and optionally a path.
```shell
> curl 'https://api.m3o.com/files/Files/List' \
  -H 'micro-namespace: $yourNamespace' \
  -H 'authorization: Bearer $yourToken' \
  -d {
  "path": "Defaults to '/', ie. lists all files in a project.. Supply path if of a folder if you want to list. file inside that folder. eg. '/docs'",
  "project": "Project, required for listing."
};
# Response
{
  "files": [
    {
      "created": 1,
      "file_contents": "File contents. Empty for directories.",
      "is_directory": true,
      "name": "Name of folder or file.",
      "path": "Path. Default is '/', ie. top level",
      "project": "A custom string for namespacing purposes. eg. files-of-mywebsite.com",
      "updated": 1
    }
  ]
}
```


### Files Save
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->
The save endpoint lets you batch save text files.
```shell
> curl 'https://api.m3o.com/files/Files/Save' \
  -H 'micro-namespace: $yourNamespace' \
  -H 'authorization: Bearer $yourToken' \
  -d {
  "files": [
    {
      "created": 1,
      "file_contents": "File contents. Empty for directories.",
      "is_directory": true,
      "name": "Name of folder or file.",
      "path": "Path. Default is '/', ie. top level",
      "project": "A custom string for namespacing purposes. eg. files-of-mywebsite.com",
      "updated": 1
    }
  ]
};
# Response
{}
```


