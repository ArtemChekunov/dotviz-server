# dotviz-server
WebBased DOT language visualization tool

# Description

It is web based tool for rendering DOT language https://en.wikipedia.org/wiki/DOT_(graph_description_language)
Inspired by plantuml-server.
Written on golang.

# How to RUN:

```
docker run -d -it -p 1234:1234 --name dotviz-server sc0rp1us/dotviz-server
# got to http://localhost:1234/
```

# How to build own assembly:

```
docker build --rm -t local/dotviz-server .
docker run -d -it -p 1234:1234 --name dotviz-server local/dotviz-server
```

# Screenshot

![alt tag](DotVizServer_screenshot.png)

