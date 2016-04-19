## TMAIL

A simple mail client (using mailgun api) for sending emails.
It needs a url, (to fetch html newsletter) and sends the content
to the desired destination.


```
export MAILGUN_API_KEY=asdf
export MAILGUN_DOMAIN=asdfadsf

tmail -url http://github.com/newsletter \
    -to j@tmpo.io -from j@tmpo.io \
    -subject "Test subject"

```
