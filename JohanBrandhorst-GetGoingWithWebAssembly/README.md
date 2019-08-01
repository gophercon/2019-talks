# Running this presentation

## Requirements

Go present:

```bash
$ go get golang.org/x/tools/present
```

## Running

In order to fit in all the speakers notes for each slide I've made some small
adjustments to the speaker notes CSS. To take advantage of this, make sure to
run the presentation locally, specifying a custom base path.

```bash
$ present -notes -base .
Open your web browser and visit http://127.0.0.1:3999
Notes are enabled, press 'N' from the browser to display them.
```

Navigate to http://127.0.0.1:3999 to view the presentation.
