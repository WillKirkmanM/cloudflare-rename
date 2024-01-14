# cloudflare-rename
> Bulk Change DNS Records in Cloudflare

## Get Started
### Run the Program
Install the dependencies
```
$ go get
```

Run the program
```
go run main.go
```

### Get API Token
1. Head to https://dash.cloudflare.com/profile/api-tokens
2. Click `Create Token`
![User API Tokens](/assets/user-api-tokens.png "User API Tokens")
3. Use the Edit zone DNS template
![Create API Token](/assets/create-api-token.png "Create API Token")
4. Select the Zone (Domain) you would like to rename
![Zone Resources](/assets/zone-resources.png "Zone Resources")
5. Create the API Key
![Create It](/assets/create-it.png "Create It")
6. Copy it and paste it into the program
![Copy Token](/assets/copy-token.png "Copy Token")
If you have already created the API key but have made a mistake, you can always regenerate it by clicking `Roll`
![Roll API Key](/assets/roll-api-key.png "Roll API Key")
