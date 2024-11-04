# ard-jellyfin

A program to generate XMLTV and M3U8 files for ARD's online live TV, using ARD's web APIs.

usage:<br>
install golang, run `go run ./cmd/ard-jellyfin`<br>
it will by default output into the folder `output`<br>
you can change it using the output flag:<br>
`go run ./cmd/ard-jellyfin -output output2`<br>
in jellyfin simply create a m3u tuner and xmltv epg using the files,<br>
by default the program will update it every hour<br>
you can set the timer using -timer {seconds}<br>
if your jellyfin cant access the folder<br>
you can run a builtin http server using the -listen flag:<br>
`go run ./cmd/ard-jellyfin -listen 0.0.0.0:1234`<br>
then just use http://the-machines-ip:1234/ard.m3u8 and http://the-machines-ip:1234/ard.xml