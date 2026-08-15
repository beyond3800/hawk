# Auth


Hawk provides authentification system

Where Hawk authentification command that can use in the terminal

```bash
    hawk auth
```
It generates a secret key which is used by hawk to authentificate 


```.env
    APP_SECRET_KEY=4567890lkadzxfcvbytrewAZXCVBN
    APP_ISSUER=Hawk
```

The APP_SECRET_KEY is not meant to be edited because hawk uses it to authentificate user
