# Authorization

This app authorizes user credentials using the OAuth2.0 Authorization Code Flow. You can read more on this here: https://developer.spotify.com/documentation/web-api/tutorials/code-flow

## Step 1 - Unmarshal the config.toml to get required parameters
- Call config.GetEnvironmentVariables()

## Step 2 - Call Authorization Endpoint
- Call the authorization endpoint: https://accounts.spotify.com/authorize
- Build the URL by setting required params (see params and values in the Request User Authorization section here: https://developer.spotify.com/documentation/web-api/tutorials/code-flow).

## Step 3 - Authorization URL Opens in the Browser
- If it's the first time a user is logging in, they will have to approve the app. Otherwise this step is skipped.
- The callback URL loads in the browser with the authorization token.
- The entire URL is meant to be copied here as the program waits for it to be pasted in.
- Parse the query params here to get the authorization code.

## Step 4 - Call the Token Endpoint
- Call the token endpoint: https://accounts.spotify.com/api/token.
- Build the URL by setting required params with the following values. The benefit of doing this with the url.URL package is that it will encode the query.
	code -> this is the code from step 3
	grant_type -> grant type from the config
	redirect_uri -> redirect URI from the config
- Concatenate the clientID and clientSecret string as follows "clientID:clientSecret", and base64 encode it.
- Add headers to the request:
	- Content-Type -> content type from the config
	- Authorization -> base64 encoded client ID and client secret)

## Step 5 - Unmarshal the Token Endpoint to get the access token
- JSON decode the token response to get the access token!