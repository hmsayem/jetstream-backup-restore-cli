{{/*
Expand the name of the chart.
*/}}
{{- define "name" -}}
{{- default .Release.Name -}}
{{- end -}}


{{- define "nats.tlsConfig" -}}
tls {
{{- if .cert }}
    cert: {{ .secretPath }}/{{ .secret.name }}/{{ .cert }}
{{- end }}
{{- if .key }}
    key:  {{ .secretPath }}/{{ .secret.name }}/{{ .key }}
{{- end }}
{{- if .ca }}
    root: {{ .secretPath }}/{{ .secret.name }}/{{ .ca }}
{{- end }}
{{- if .insecure }}
    insecure: {{ .insecure }}
{{- end }}
{{- if .verify }}
    verify: {{ .verify }}
{{- end }}
{{- if .verifyAndMap }}
    verify_and_map: {{ .verifyAndMap }}
{{- end }}
{{- if .curvePreferences }}
    curve_preferences: {{ .curvePreferences }}
{{- end }}
{{- if .timeout }}
    timeout: {{ .timeout }}
{{- end }}
}
{{- end }}
