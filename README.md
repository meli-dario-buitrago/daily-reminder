# daily-reminder
Daily reminder project provides a scheduller function to send notifications to a provided webhook


## How to use
1. Be sure to generate a webhook at slack application page 
2. Download this repository
3. Configure your properties at `./config/properties.yaml` 
   ```
   webhook-url: https://hooks.slack.com/services/.......
   jira-url: https://mercadolibre.atlassian.net/jira/software/projects/.....
   meet-url: https://meet.google.com/......
   days:
     Monday:
       presenter: {SLACK-USER-ID}
       alternate: {SLACK-USER-ID}
     Tuesday:
       presenter: {SLACK-USER-ID}
       alternate: {SLACK-USER-ID}
     Wednesday:
       presenter: {SLACK-USER-ID}
       alternate: {SLACK-USER-ID}
     Thursday:
       presenter: {SLACK-USER-ID}
       alternate: {SLACK-USER-ID}
     Friday:
       presenter: {SLACK-USER-ID}
       alternate: {SLACK-USER-ID}
   ```
   You can get the *SLACK-USER-ID* viewing their full profile and clicking at **More** option
   Each ID looks like this **U0E3GU34SHZ**
4. Configure your cron expresion at main.go L15, default is 9:25:00 AM every day
5. Run project 
   ```
   go run main.go
   ```
6. Relax and wait your reminders