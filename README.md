# *How to deploy your Go chatbot server to Heroku*

* Install the Heroku toolbelt from here https://toolbelt.heroku.com to launch, stop and monitor instances. Sign up for free at https://www.heroku.com if you don't have an account yet.
* Fetch/clone/fork this repository
    * `go get github.com/uccmen/fbmessenger-boilerplate-go`
    * `cd $GOPATH/src/github.com/uccmen/fbmessenger-boilerplate-go`
* run `heroku apps:create [APP_NAME]`. Edit `[APP_NAME]` to your preferred app name of course ;)
### *Setup the Facebook App*
###### **special thanks to [jw84](https://github.com/jw84/messenger-bot-tutorial#setup-the-facebook-app) for providing this section

1. Create or configure a Facebook App or Page here https://developers.facebook.com/apps/

    ![Alt text](https://github.com/jw84/messenger-bot-tutorial/blob/master/demo/shot1.jpg)

2. In the app go to Messenger tab then click Setup Webhook. Here you will put in the URL of your Heroku server and a token. Make sure to check all the subscription fields.

    ![Alt text](https://github.com/jw84/messenger-bot-tutorial/blob/master/demo/shot3.jpg)

3. Get a Page Access Token and save this somewhere.

    ![Alt text](https://github.com/jw84/messenger-bot-tutorial/blob/master/demo/shot2.jpg)

4. Go back to Terminal and type in this command to trigger the Facebbook app to send messages. Remember to use the token you requested earlier.

    ```bash
    curl -X POST "https://graph.facebook.com/v2.6/me/subscribed_apps?access_token=<PAGE_ACCESS_TOKEN>"
    ```
5. Configure required envs listed below from Heroku dashboard UI for your app or from the command line:
    * `heroku config:set PORT=3000`
    * `heroku config:set HUB_VERIFY_TOKEN=my_secret_token`
6. `git push heroku master`
7. The app is now deployed. Check it out by visiting `https://[APP_NAME].herokuapp.com/health`
8. Go to the Facebook Page and click on Message to start chatting!

## *Required ENVs*
```
1. PORT="" //3000
2. HUB_VERIFY_TOKEN="" //your_own_secret_token
3. FB_PAGE_ACCESS_TOKEN="" //fb page access token saved earlier
4. FB_MESSENGER_URL="" //https://graph.facebook.com/v2.6/me/messages
5. ROLLBAR_TOKEN="" https://rollbar.com/why-rollbar/
6. RELEASE_STAGE="" //development
```