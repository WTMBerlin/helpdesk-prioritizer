# Helpdesk Prioritizer

A tool that gives higher priority to support requests from nice
people using the [IBM Watson API](https://developer.ibm.com/watson/):

```
$ helpdesk-prioritizer 'Hi! My internet connection was not working this morning. Could you please find the time to look at it today? Would appreciate your help very much.'

Analytical (positive): high probability.
Conscientiousness (positive): high probability.
Agreeableness (positive): high probability.
Priority: 1.
```

This amazing tool is part of the ["Go for non-gophers" workshop](https://www.boosterconf.no/talks/846)
at the Booster conference in Bergen, Norway.

## Configuration

1. Get your IBM Watson API keys: go to the [Tone Analyzer page](https://www.ibm.com/watson/developercloud/tone-analyzer.html), click the large button in the top right corner that says "Start free in Bluemix," follow the instructions.

2. Set the environment variables `WATSON_USERNAME` and `WATSON_PASSWORD` using the API credentials you've got during Step 1.

3. You're all set!
