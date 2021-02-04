---
title: files Micro.js
servicename: files
labels: 
- Micro.js
---

## Micro.js


### Files List
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->
List files by their project and optionally a path.
```html
<script src="https://web.m3o.com/assets/micro.js"></script>
<script type="text/javascript">
  document.addEventListener("DOMContentLoaded", function (event) {
    // Login is only required for endpoints doing authorization
    Micro.requireLogin(function () {
      Micro.post(
        "/files/Files/List",
        "micro",
        {
          "path": "Defaults to '/', ie. lists all files in a project.. Supply path if of a folder if you want to list. file inside that folder. eg. '/docs'",
          "project": "Project, required for listing."
},
        function (data) {
          console.log("Success.");
        }
      );
    });
  });
</script>
```


### Files Save
<!-- We use the request body description here as endpoint descriptions are not
being lifted correctly from the proto by the openapi spec generator -->
The save endpoint lets you batch save text files.
```html
<script src="https://web.m3o.com/assets/micro.js"></script>
<script type="text/javascript">
  document.addEventListener("DOMContentLoaded", function (event) {
    // Login is only required for endpoints doing authorization
    Micro.requireLogin(function () {
      Micro.post(
        "/files/Files/Save",
        "micro",
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
},
        function (data) {
          console.log("Success.");
        }
      );
    });
  });
</script>
```


