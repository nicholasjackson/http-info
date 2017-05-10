package templates

const StatsTemplate = `
<html>
  <head>
    <title>http-info</title>
  <head>
  <body style="font-family: Verdana">
    <b>Architecture:</b> {{ .Arch }}<br/>
    <b>Operating System:</b> {{ .OS }}<br/>
    <b>Number of CPUs:</b> {{ .NumCPUs }}</br>
    <b>Network:</b></br>
    {{ range $index, $network := .Network }}
      &#160;&#160;<em>{{ $network.Name }}</em></br>
      {{ range $index, $address := $network.IPAdresses }}
        &#160;&#160;&#160;&#160;{{ $address }}</br>
      {{ end }}
      </br>
    {{ end }}</br>
    <b>AWS Details:</b></br>
    &#160;&#160;<em>PublicIP:</em>{{ .AWS.PublicIP }}</br>
    &#160;&#160;<em>PrivateIP:</em>{{ .AWS.PrivateIP }}</br>
  </body>
</html>
`
