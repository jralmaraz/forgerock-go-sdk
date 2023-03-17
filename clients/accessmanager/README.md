# Access Manager Client

Client to authenticate to Access Manager API and set the `iplanetDirectoryPro` header required for sub-sequent requests.

To be able to manage resources on an access managet instance, authentication flow follows the steps on [Using the Session Token After Authentication](https://backstage.forgerock.com/docs/am/7/authentication-guide/rest-using-ssotokens.html).

```shell
curl \
 --insecure \
 --request POST \
 --header "Content-Type: application/json" \
 --header "X-OpenAM-Username: amadmin" \
 --header "X-OpenAM-Password: ***" \
 --header "Accept-API-Version: resource=2.0" \
 --data "{}" \
 "https://dev.example.com/am/json/realms/root/authenticate"

#Response:

{"tokenId":"P5d-syX7evW11eW_s8hWt2KfKk0.*AAJTSQACMDIAAlNLABxuaFBFUW5iditOUWRONjZIZzdHUVFGUFFSWVU9AAR0eXBlAANDVFMAAlMxAAIwMQ..*","successUrl":"/am/console","realm":"/"}
```

```shell
curl -k \
--header "iPlanetDirectoryPro: $value_from_tokenId" \
--header  "Accept-API-Version: resource=1.0, protocol=2.1" \
https://dev.example.com/am/json/global-config/realms?_queryFilter=true


#Response:

{"result":[{"_id":"Lw","_rev":"442134239","parentPath":null,"active":true,"name":"/","aliases":["dev.example.com","am-config","am"]},{"_id":"L3Rlc3Q","_rev":"1983000027","parentPath":"/","active":true,"name":"test","aliases":[]}],"resultCount":2,"pagedResultsCookie":null,"totalPagedResultsPolicy":"NONE","totalPagedResults":-1,"remainingPagedResults":-1}%
```