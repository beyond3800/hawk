##  Storage

Hawk provides storage system

Where Hawk storage command can be use in the terminal

```bash
    hawk storage
```
It creates storage folder where files are being stored
Storage has Put, Get, Delete, Exists, Url
```go

    Put is use to save file in storage folder or sub folder
	app.Post("/upload", func(c *hawk.Context) {

    file, header, err := c.OpenFile("cover")
	
    if err != nil {
        c.JSON(400, hawk.H{
            "error": err.Error(),
        })
        return
    }
    defer file.Close()

    this creates a sub folder in storage name uploads
    err = storage.Default().Put(
        "uploads/"+header.Filename,
        file,
    )
    

    if err != nil {
        c.JSON(500, hawk.H{
            "error": err.Error(),
        })
        return
    }

    c.JSON(200, hawk.H{
        "message": "Uploaded successfully",
    })

})
```