# telegram-get-id

Bot send id: Channel, Private Channel, User or Group

---------

Example:

![](img/demo.png)

-----

Run:

```bash
export MOD=GET_UPDATES
export BOT_TOKEN=BOK_TOKEN

go run main.go
```


```bash
docker build . -t telegram-get-id
docker run --env=MOD=GET_UPDATES --env=BOT_TOKEN=BOK_TOKEN telegram-get-id 
```
